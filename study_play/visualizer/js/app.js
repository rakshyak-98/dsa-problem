/**
 * Main application — wires UI, code execution, and playback.
 */
(function () {
  let structureType = "array";
  let arrayView = "bars";
  let engine = new VizEngine("array");
  let editor = null;
  let consoleLines = [];
  let lastAnalysis = null;
  let userPickedStructure = false;
  let parseDebounce = null;
  const AsyncFunction = Object.getPrototypeOf(async function () {}).constructor;

  const $ = (sel) => document.querySelector(sel);
  const canvas = $("#viz-canvas");
  const stepMessage = $("#step-message");
  const stepCounter = $("#step-counter");
  const variablesPanel = $("#variables-panel");
  const consolePanel = $("#console-panel");
  const errorPanel = $("#error-panel");
  const exampleSelect = $("#example-select");
  const playBtn = $("#btn-play");
  const parserTypeEl = $("#parser-type");
  const parserPatternsEl = $("#parser-patterns");
  const parserInlineEl = $("#parser-inline");
  const autoVizToggle = $("#auto-viz-toggle");

  function initEditor() {
    editor = CodeMirror.fromTextArea($("#code-editor"), {
      mode: "javascript",
      theme: "dracula",
      lineNumbers: true,
      tabSize: 2,
      lineWrapping: true,
    });
    editor.on("change", () => {
      clearTimeout(parseDebounce);
      parseDebounce = setTimeout(() => analyzeEditorCode(), 400);
    });
  }

  function updateParserUI(analysis, instrumented) {
    if (!analysis) return;
    lastAnalysis = analysis;

    const type = analysis.structureType || "array";
    const pct = Math.round((analysis.confidence || 0) * 100);
    parserTypeEl.textContent = `${type} (${pct}%)`;
    parserPatternsEl.textContent = analysis.patterns?.length
      ? analysis.patterns.filter((p) => p !== "manual-viz").join(" · ")
      : "";

    let inline = `Detected: <strong>${type}</strong>`;
    if (analysis.hasManualViz) inline += " · manual viz";
    else if (instrumented) inline += " · auto-instrumented";
    else if (autoVizToggle.checked) inline += " · will auto-instrument on run";
    parserInlineEl.innerHTML = inline;

    if (!userPickedStructure && analysis.confidence >= 0.5 && type !== structureType) {
      applyDetectedStructure(type, false);
    }
  }

  function analyzeEditorCode() {
    if (!editor || typeof CodeParser === "undefined") return;
    const code = editor.getValue();
    const analysis = CodeParser.analyze(code);
    const prepared = CodeParser.prepare(code, { autoInstrument: false });
    updateParserUI(analysis, prepared.instrumented);
  }

  function applyDetectedStructure(type, userInitiated) {
    if (userInitiated) userPickedStructure = true;
    structureType = type;
    engine = new VizEngine(type);
    engine.onStepChange = onStepChange;

    document.querySelectorAll(".structure-tabs .tab-button").forEach((btn) => {
      btn.classList.toggle("active", btn.dataset.structure === type);
    });

    showInputSection();
  }

  function parseArrayInput() {
    const raw = $("#array-input").value.trim();
    if (!raw) return [];
    return raw.split(",").map((s) => {
      const t = s.trim();
      if (t === "null") return null;
      const n = Number(t);
      return Number.isNaN(n) ? t : n;
    });
  }

  function parseStringInput() {
    return $("#string-input").value;
  }

  function parseTreeInput() {
    const raw = $("#tree-input").value.trim();
    if (!raw) return [];
    return raw.split(",").map((s) => {
      const t = s.trim().toLowerCase();
      if (t === "null" || t === "") return null;
      const n = Number(t);
      return Number.isNaN(n) ? t : n;
    });
  }

  function parseGraphInput() {
    const raw = $("#graph-input").value.trim();
    if (!raw) return [];
    return raw
      .split("\n")
      .map((line) => line.trim())
      .filter(Boolean)
      .map((line) =>
        line.split(",").map((s) => {
          const t = s.trim();
          const n = Number(t);
          return Number.isNaN(n) ? t : n;
        })
      );
  }

  function getInputData() {
    switch (structureType) {
      case "array":
        return parseArrayInput();
      case "string":
        return parseStringInput();
      case "tree":
        return parseTreeInput();
      case "graph":
        return parseGraphInput();
      default:
        return [];
    }
  }

  function showInputSection() {
    $("#input-array").classList.toggle("hidden", structureType !== "array");
    $("#input-string").classList.toggle("hidden", structureType !== "string");
    $("#input-tree").classList.toggle("hidden", structureType !== "tree");
    $("#input-graph").classList.toggle("hidden", structureType !== "graph");
    $("#array-view-toggle").classList.toggle("hidden", structureType !== "array");
  }

  function populateExamples() {
    const list = EXAMPLES[structureType] || [];
    exampleSelect.innerHTML = list
      .map((ex, i) => `<option value="${i}">${ex.name}</option>`)
      .join("");

    if (list.length) loadExample(0);
  }

  function loadExample(index) {
    const list = EXAMPLES[structureType] || [];
    const ex = list[Number(index)];
    if (!ex) return;

    if (structureType === "array") $("#array-input").value = ex.input;
    else if (structureType === "string") $("#string-input").value = ex.input;
    else if (structureType === "tree") $("#tree-input").value = ex.input;
    else $("#graph-input").value = ex.input;

    editor.setValue(ex.code);
    userPickedStructure = false;
    setTimeout(analyzeEditorCode, 0);
  }

  function setStructure(type, userInitiated = true) {
    if (userInitiated) userPickedStructure = true;
    applyDetectedStructure(type, userInitiated);
    populateExamples();
    buildVisualization();
  }

  function buildVisualization() {
    engine.reset();
    const data = getInputData();
    errorPanel.classList.add("hidden");

    if (structureType === "array") {
      engine.setArray(data.map((v) => (typeof v === "number" ? v : Number(v) || v)));
      engine.setArrayView(arrayView);
    } else if (structureType === "string") {
      engine.setString(String(data));
    } else if (structureType === "tree") {
      const root = buildTreeFromLevelOrder(data);
      engine.setTree(root);
    } else if (structureType === "graph") {
      engine.setGraph(data);
    }

    engine.recordStep({ message: "Initial state" });
    engine.goToStep(0);
    updatePlaybackUI();
  }

  function onStepChange(snapshot, index, total) {
    Renderers.render(canvas, structureType, snapshot);
    stepMessage.textContent = snapshot.message || "—";
    stepCounter.textContent = `Step ${index + 1} / ${total}`;
    variablesPanel.textContent = JSON.stringify(snapshot.variables || {}, null, 2);

    if (snapshot.line != null && editor) {
      editor.setCursor(snapshot.line - 1, 0);
      editor.addLineClass(snapshot.line - 1, "background", "cm-active-line-viz");
      setTimeout(() => {
        for (let i = 0; i < editor.lineCount(); i++) {
          editor.removeLineClass(i, "background", "cm-active-line-viz");
        }
      }, 300);
    }
  }

  function updatePlaybackUI() {
    const total = engine.steps.length;
    const idx = engine.currentStep;
    stepCounter.textContent = total ? `Step ${idx + 1} / ${total}` : "Step 0 / 0";
    playBtn.textContent = engine.isPlaying ? "⏸" : "▶";
    playBtn.title = engine.isPlaying ? "Pause" : "Play";
  }

  function logToConsole(...args) {
    const line = args.map((a) => (typeof a === "object" ? JSON.stringify(a) : String(a))).join(" ");
    consoleLines.push(line);
    consolePanel.textContent = consoleLines.join("\n");
    consolePanel.scrollTop = consolePanel.scrollHeight;
  }

  async function runCode() {
    engine.stopPlayback();
    engine.reset();
    consoleLines = [];
    consolePanel.textContent = "";
    errorPanel.classList.add("hidden");

    const rawCode = editor.getValue();
    const autoInstrument = autoVizToggle.checked;
    const prepared = CodeParser.prepare(rawCode, { autoInstrument });
    const code = prepared.code;
    const analysis = prepared.analysis;

    if (prepared.instrumented) {
      logToConsole("Auto-instrumented:", prepared.instrumentReason, `(${analysis.structureType}, ${analysis.patterns.join(", ")})`);
    } else if (!analysis.hasManualViz && autoInstrument) {
      logToConsole("Auto-viz:", prepared.instrumentReason || "using detected patterns");
    }

    updateParserUI(analysis, prepared.instrumented);

    if (analysis.structureType && analysis.structureType !== structureType && analysis.confidence >= 0.5) {
      applyDetectedStructure(analysis.structureType, false);
    }

    engine.onStepChange = onStepChange;
    const viz = engine.createVizAPI();
    const autoVizConfig = {
      structureType: analysis.structureType || structureType,
      primaryVar: analysis.primaryVar || "arr",
    };

    const data = getInputData();

    try {
      const wrappedCode = `
        const __autoViz = createAutoViz(viz, ${JSON.stringify(autoVizConfig)});
        ${code}
        return await run(viz, data);
      `;
      const fn = new AsyncFunction(
        "viz", "data", "log", "createAutoViz", "buildTreeFromLevelOrder",
        "parseArrayInput", "parseStringInput", "parseTreeInput", "parseGraphInput",
        wrappedCode
      );
      const result = await fn(
        viz,
        data,
        logToConsole,
        createAutoViz,
        buildTreeFromLevelOrder,
        parseArrayInput,
        parseStringInput,
        parseTreeInput,
        parseGraphInput
      );

      if (result !== undefined) {
        logToConsole("Result:", result);
      }

      if (engine.steps.length === 0) {
        engine.recordStep({ message: "Execution finished (no steps recorded)" });
      }
      engine.goToStep(0);
      Renderers.render(canvas, structureType, engine.getCurrentSnapshot());
      stepMessage.textContent = engine.getCurrentSnapshot().message || "Run complete — use playback controls.";

      // Auto-start playback after run so visualization animates without extra click
      if (engine.steps.length > 1) {
        engine.play(getSpeed(), updatePlaybackUI);
      }
    } catch (err) {
      errorPanel.textContent = err.stack || err.message;
      errorPanel.classList.remove("hidden");
      logToConsole("Error:", err.message);
    }

    updatePlaybackUI();
  }

  function getSpeed() {
    return Number($("#speed-range").value);
  }

  document.querySelectorAll(".structure-tabs .tab-button").forEach((btn) => {
    btn.addEventListener("click", () => setStructure(btn.dataset.structure, true));
  });

  document.querySelectorAll("#array-view-toggle .view-btn").forEach((btn) => {
    btn.addEventListener("click", () => {
      arrayView = btn.dataset.view;
      document.querySelectorAll("#array-view-toggle .view-btn").forEach((b) => {
        b.classList.toggle("active", b.dataset.view === arrayView);
      });
      if (structureType === "array") {
        engine.setArrayView(arrayView);
        if (engine.steps.length) {
          Renderers.render(canvas, structureType, engine.getCurrentSnapshot());
        } else {
          buildVisualization();
        }
      }
    });
  });

  exampleSelect.addEventListener("change", () => loadExample(exampleSelect.value));

  $("#btn-build").addEventListener("click", buildVisualization);
  $("#btn-run").addEventListener("click", runCode);

  $("#btn-reset").addEventListener("click", () => {
    engine.stopPlayback();
    if (engine.steps.length) engine.goToStep(0);
    updatePlaybackUI();
  });

  $("#btn-prev").addEventListener("click", () => {
    engine.stopPlayback();
    engine.prevStep();
    updatePlaybackUI();
  });

  $("#btn-next").addEventListener("click", () => {
    engine.stopPlayback();
    engine.nextStep();
    updatePlaybackUI();
  });

  playBtn.addEventListener("click", () => {
    if (!engine.steps.length) return;
    if (engine.steps.length <= 1) {
      stepMessage.textContent = "Run code first to record algorithm steps, then press Play.";
      return;
    }
    const playing = engine.togglePlayback(getSpeed(), updatePlaybackUI);
    if (playing && engine.currentStep < 0) engine.goToStep(0);
    updatePlaybackUI();
  });

  $("#btn-format").addEventListener("click", () => {
    try {
      const val = editor.getValue();
      editor.setValue(val);
    } catch (_) {
      /* noop */
    }
  });

  // Inject active line highlight style
  const style = document.createElement("style");
  style.textContent = ".cm-active-line-viz { background: rgba(56, 189, 248, 0.2) !important; }";
  document.head.appendChild(style);

  initEditor();
  setStructure("array", false);
  analyzeEditorCode();
})();
