/**
 * Main application — wires UI, code execution, and playback.
 */
(function () {
  let structureType = "array";
  let arrayView = "bars";
  let engine = new VizEngine("array");
  let editor = null;
  let consoleLines = [];
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

  function initEditor() {
    editor = CodeMirror.fromTextArea($("#code-editor"), {
      mode: "javascript",
      theme: "dracula",
      lineNumbers: true,
      tabSize: 2,
      lineWrapping: true,
    });
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
  }

  function setStructure(type) {
    structureType = type;
    engine = new VizEngine(type);
    engine.onStepChange = onStepChange;

    document.querySelectorAll(".structure-tabs .tab-button").forEach((btn) => {
      btn.classList.toggle("active", btn.dataset.structure === type);
    });

    showInputSection();
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

    const data = getInputData();
    const code = editor.getValue();

    engine.onStepChange = onStepChange;
    const viz = engine.createVizAPI();

    const helpers = {
      viz,
      log: logToConsole,
      buildTreeFromLevelOrder,
      parseArrayInput,
      parseTreeInput,
      parseGraphInput,
    };

    try {
      const wrappedCode = `
        ${code}
        return await run(viz, data);
      `;
      const fn = new AsyncFunction(
        "viz", "data", "log", "buildTreeFromLevelOrder",
        "parseArrayInput", "parseStringInput", "parseTreeInput", "parseGraphInput",
        wrappedCode
      );
      const result = await fn(
        viz,
        data,
        logToConsole,
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
    btn.addEventListener("click", () => setStructure(btn.dataset.structure));
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
  setStructure("array");
})();
