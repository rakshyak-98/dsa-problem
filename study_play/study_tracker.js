const STORAGE_KEY = "studyTrackerState.v1";
const GEMINI_API_KEY_STORAGE = "studyTrackerGeminiApiKey.v1";
const GEMINI_MODEL = "gemini-3.5-flash";
const TOTAL_WEEKS = 12;
const DAY_IN_MS = 24 * 60 * 60 * 1000;
const FILE_HANDLE_DB_NAME = "studyTrackerFileHandles";
const FILE_HANDLE_STORE = "handles";
const SOLVED_PROBLEMS_HANDLE_KEY = "solvedProblemsMd";

const DAILY_SESSION_ITEMS = [
  { id: "understand", label: "Understand: restate today's ask in one sentence" },
  { id: "core5", label: "Core 5 reflex (twoSum, BS, dedupe, window, freq)" },
  { id: "reflex", label: "Specialty drill (node study_play/daily_drill.js)" },
  { id: "solve", label: "Primary: one problem from this week's plan" },
  { id: "lesson", label: "Log ask + pattern + one lesson" }
];

const TOPIC_ORDER = [
  "arrays",
  "strings",
  "simulation",
  "hashing",
  "two_pointers",
  "binary_search",
  "sliding_window",
  "trees",
  "graphs",
  "stack_queue",
  "dynamic_programming",
  "greedy",
  "matrix",
  "misc"
];

const STUDY_PLAN_FORM_OPTIONS = {
  topics: [
    { value: "", label: "Select topic from study plan" },
    { value: "arrays", label: "Arrays" },
    { value: "strings", label: "Strings" },
    { value: "simulation", label: "Simulation" },
    { value: "hashing", label: "Hashing" },
    { value: "two_pointers", label: "Two pointers" },
    { value: "sliding_window", label: "Sliding window" },
    { value: "binary_search", label: "Binary search" },
    { value: "trees", label: "Trees" },
    { value: "stack_queue", label: "Stack / queue" },
    { value: "matrix", label: "Matrix" },
    { value: "graphs", label: "Graphs" },
    { value: "dynamic_programming", label: "Dynamic programming" },
    { value: "greedy", label: "Greedy" },
    { value: "misc", label: "Misc" }
  ],
  difficulties: [
    { value: "easy", label: "easy" },
    { value: "medium", label: "medium" },
    { value: "hard", label: "hard" }
  ],
  statuses: [
    { value: "solved", label: "solved" },
    { value: "reviewed", label: "reviewed" },
    { value: "studied-editorial", label: "studied-editorial" }
  ],
  patterns: [
    { value: "", label: "Select pattern from study plan" },
    { value: "two pointers", label: "Two pointers" },
    { value: "prefix sum + hash map", label: "Prefix sum + hash map" },
    { value: "sliding window", label: "Sliding window" },
    { value: "binary search", label: "Binary search" },
    { value: "hash map / set", label: "Hash map / set" },
    { value: "bfs", label: "BFS" },
    { value: "dfs + visited", label: "DFS + visited" },
    { value: "dynamic programming", label: "Dynamic programming" },
    { value: "greedy", label: "Greedy" },
    { value: "monotonic stack", label: "Monotonic stack" }
  ],
  mistakes: [
    { value: "none", label: "none" },
    { value: "understanding", label: "understanding" },
    { value: "logic", label: "logic" },
    { value: "edge case", label: "edge case" },
    { value: "pattern miss", label: "pattern miss" },
    { value: "syntax", label: "syntax" },
    { value: "timeout", label: "timeout" },
    { value: "state definition", label: "state definition" },
    { value: "off-by-one", label: "off-by-one" }
  ]
};

const ROADMAP_WEEKS = [
  {
    weekNumber: 1,
    phase: "Phase 1",
    title: "Arrays and complexity",
    tasks: [
      "Mon: Reflex `01` + primary `arrays/easy/plus_one.js` (stretch: concatenation_of_array)",
      "Tue: Re-drill arrays + primary `find_closest_number_to_zero` (stretch: max_consecutive_ones)",
      "Wed: Reflex `02` + primary `hashing/easy/two_sum.js` (stretch: contains_duplicates)",
      "Thu: Re-drill hashing + primary `fair_candy_swap` (stretch: degree_of_an_array)",
      "Fri: Mixed + primary `majority_element` (stretch: summary_ranges)",
      "Sat: Review failed primaries and re-type drills blind",
      "Sun: Rest (optional light reflex)"
    ],
    goals: [
      "Write `reverseArray`, `maxInArray`, and `countFreq` from memory in under 3 minutes.",
      "Explain O(n) vs O(n^2) for every solution you wrote.",
      "Restate each solved problem in one plain sentence.",
      "Solve 6 or more primary easies without peeking."
    ]
  },
  {
    weekNumber: 2,
    phase: "Phase 1",
    title: "Strings and simulation",
    tasks: [
      "Mon: Primary `find_words_containing_character` (stretch: most_common_word)",
      "Tue: Primary `unique_morse_code_words` (stretch: find_resultant_array_after_removing_anagrams)",
      "Wed: Primary `baseball_game` (stretch: relative_ranks)",
      "Thu: Primary 1 unseen `strings/easy/` (stretch: 1 more)",
      "Fri: Timed mock with 2 easy problems in 45 minutes",
      "Sat: Re-drill `01` and `02`",
      "Sun: Rest"
    ],
    goals: [
      "Feel comfortable with `Map`, `Set`, and object frequency maps.",
      "Reach 12 or more primary easies in this plan.",
      "For simulation: list operations in order before coding."
    ]
  },
  {
    weekNumber: 3,
    phase: "Phase 2",
    title: "Two pointers",
    tasks: [
      "Mon: Reflex `03` + primary `remove_duplicates_from_sorted_array` (stretch: move_zeroes)",
      "Tue: Re-drill + primary `best_time_to_buy_sell_stock` (stretch: squares_of_a_sorted_array)",
      "Wed: Primary `container_with_most_water` (stretch only: 3sum)",
      "Thu: Primary `maximum_average_subarray_1` (stretch: min difference of k scores)",
      "Fri: Timed 1 easy or 1 medium from `two_pointers/`",
      "Sat: Re-type `templates/pattern_cheat_sheet.js` sections 1 to 3",
      "Sun: Rest"
    ],
    goals: []
  },
  {
    weekNumber: 4,
    phase: "Phase 2",
    title: "Binary search",
    tasks: [
      "Mon: Reflex `04` + primary `search_insertion_position` (stretch: smallest letter greater than target)",
      "Tue: Primary `longest_subsequence_with_limited_sum`",
      "Wed: Primary `find_the_duplicate_number`",
      "Thu: Primary `sort_colors` (stretch: longest_palindromic_substring)",
      "Fri: Mock interview with 1 medium in 45 minutes",
      "Sat: Review mediums and tag mistake type (include understanding)",
      "Sun: Rest"
    ],
    goals: [
      "Make `left`, `right`, `while (left < right)` muscle memory.",
      "Make `while (left <= right)` binary search muscle memory.",
      "Pass the translation test before every medium.",
      "Solve at least 4 medium two-pointer primaries."
    ]
  },
  {
    weekNumber: 5,
    phase: "Phase 3",
    title: "Hashing mastery and medium arrays",
    tasks: [
      "Mon: Rewrite `02` + primary `group_anagram` (stretch: top_k_ferquent_element)",
      "Tue: Primary `find_pivot_index` (stretch: subarray_sum_divisible_by_k)",
      "Wed: Primary `find_all_duplicates_in_an_array` (stretch: max_product_subarray)",
      "Thu: Primary `longest_consecutive_sequence`",
      "Fri: Primary `product_of_array_except_self`",
      "Sat: Re-drill files `01` to `04`",
      "Sun: Rest"
    ],
    goals: []
  },
  {
    weekNumber: 6,
    phase: "Phase 3",
    title: "Consolidation week",
    tasks: [
      "Mon: Re-solve 1 medium you previously got wrong (understand → blind)",
      "Tue: Re-solve 1 medium you previously got wrong",
      "Wed: Re-solve 1 medium you previously got wrong",
      "Thu: Re-solve 1 medium you previously got wrong",
      "Fri: Timed set of 2 medium problems in 90 minutes",
      "Sat: Write a one-page pattern journal",
      "Sun: Rest"
    ],
    goals: [
      "Reach 8 or more medium primaries solved.",
      "Explain prefix sum in one sentence with an example.",
      "Solve group anagrams without notes.",
      "Correctly say subarray vs subsequence for Week 5 primaries."
    ]
  },
  {
    weekNumber: 7,
    phase: "Phase 4",
    title: "Trees and stacks",
    tasks: [
      "Mon: `drills/05_trees_stacks_reflex.js` + `trees/easy/convert_sorted_array_to_binary_search_tree.js`",
      "Tue: BFS and DFS tree practice",
      "Wed: `stack_queue/easy/last_stone_weight.js`",
      "Thu: Study monotonic stack and try 1 problem",
      "Fri: Matrix and graph easy practice with flood fill",
      "Sat: `graphs/easy/island_permeter.js` and graph review",
      "Sun: Rest"
    ],
    goals: []
  },
  {
    weekNumber: 8,
    phase: "Phase 4",
    title: "Dynamic programming",
    tasks: [
      "Mon: `drills/06_dp_reflex.js` + `dynamic_programming/easy/fibonacci_number.js` + `dynamic_programming/easy/min_cost_climbing_staris.js`",
      "Tue: `dynamic_programming/easy/pascale_triangle_1.js` + `dynamic_programming/easy/pascale_triangle_2.js`",
      "Wed: Study DP on strings and add examples if needed",
      "Thu: Compare greedy vs DP with `greedy/easy/can_place_flower.js` + `greedy/easy/lemonade_change.js`",
      "Fri: Mixed set with 1 DP easy + 1 greedy easy",
      "Sat: Re-type DP templates from memory",
      "Sun: Rest"
    ],
    goals: []
  },
  {
    weekNumber: 9,
    phase: "Phase 4",
    title: "Graphs BFS and DFS",
    tasks: [
      "Mon: `drills/07_graphs_reflex.js` BFS grid practice",
      "Tue: `drills/07_graphs_reflex.js` DFS grid and visited set practice",
      "Wed: Repo graph easy practice",
      "Thu: Try 1 medium graph problem",
      "Fri: Mock interview with 1 medium tree or graph problem",
      "Sat: Full reflex review of all 7 drill files",
      "Sun: Rest"
    ],
    goals: []
  },
  {
    weekNumber: 10,
    phase: "Phase 5",
    title: "Medium consolidation and hard intro",
    tasks: [
      "Mon: 1 new medium, max 90 minutes (first 5 min = question literacy)",
      "Tue: Re-solve Monday's medium from scratch",
      "Wed: 1 new medium and classify the pattern",
      "Thu: 1 hard study problem, 30 minute attempt plus editorial rewrite",
      "Fri: Timed set of 2 medium problems in 90 minutes",
      "Sat: Random drill file + `templates/pattern_cheat_sheet.js`",
      "Sun: Rest"
    ],
    goals: []
  },
  {
    weekNumber: 11,
    phase: "Phase 5",
    title: "Medium consolidation and hard intro",
    tasks: [
      "Mon: 1 new medium, max 90 minutes (first 5 min = question literacy)",
      "Tue: Re-solve Monday's medium from scratch",
      "Wed: 1 new medium and classify the pattern",
      "Thu: 1 hard study problem, 30 minute attempt plus editorial rewrite",
      "Fri: Timed set of 2 medium problems in 90 minutes",
      "Sat: Random drill file + `templates/pattern_cheat_sheet.js`",
      "Sun: Rest"
    ],
    goals: []
  },
  {
    weekNumber: 12,
    phase: "Phase 5",
    title: "Medium consolidation and hard intro",
    tasks: [
      "Mon: 1 new medium, max 90 minutes (first 5 min = question literacy)",
      "Tue: Re-solve Monday's medium from scratch",
      "Wed: 1 new medium and classify the pattern",
      "Thu: 1 hard study problem, 30 minute attempt plus editorial rewrite",
      "Fri: Timed set of 2 medium problems in 90 minutes",
      "Sat: Random drill file + `templates/pattern_cheat_sheet.js`",
      "Sun: Rest"
    ],
    goals: [
      "Reach 20 or more medium problems total.",
      "Study 3 or more hard problems (attempt + editorial + rewrite).",
      "Pass translation test and list 2–3 approaches for a random medium in 5 minutes."
    ]
  }
];

const state = loadState();

let refs = {};

document.addEventListener("DOMContentLoaded", init);

function init() {
  refs = {
    tabButtons: Array.from(document.querySelectorAll(".tab-button")),
    tabPanels: Array.from(document.querySelectorAll(".tab-panel")),
    sessionDate: document.getElementById("session-date"),
    planStartDate: document.getElementById("plan-start-date"),
    dailySessionSummary: document.getElementById("daily-session-summary"),
    dailySessionChecklist: document.getElementById("daily-session-checklist"),
    dailyWeekEyebrow: document.getElementById("daily-week-eyebrow"),
    sessionProgressCount: document.getElementById("session-progress-count"),
    sessionProgressFill: document.getElementById("session-progress-fill"),
    roadmapSummary: document.getElementById("roadmap-summary"),
    weeklyRoadmap: document.getElementById("weekly-roadmap"),
    problemDate: document.getElementById("problem-date"),
    problemTopic: document.getElementById("problem-topic"),
    problemDifficulty: document.getElementById("problem-difficulty"),
    problemStatus: document.getElementById("problem-status"),
    problemPattern: document.getElementById("problem-pattern"),
    problemMistake: document.getElementById("problem-mistake"),
    problemForm: document.getElementById("problem-form"),
    resetForm: document.getElementById("reset-form"),
    formStatus: document.getElementById("form-status"),
    logStatus: document.getElementById("log-status"),
    syncMarkdown: document.getElementById("sync-markdown"),
    clearEntries: document.getElementById("clear-entries"),
    problemLogList: document.getElementById("problem-log-list"),
    logCount: document.getElementById("log-count"),
    statsCards: document.getElementById("stats-cards"),
    statsMatrix: document.getElementById("stats-matrix"),
    geminiApiKey: document.getElementById("gemini-api-key"),
    analyzerQuestion: document.getElementById("analyzer-question"),
    analyzerSolution: document.getElementById("analyzer-solution"),
    generateTags: document.getElementById("generate-tags"),
    analyzeSolution: document.getElementById("analyze-solution"),
    tagStatus: document.getElementById("tag-status"),
    generatedTags: document.getElementById("generated-tags"),
    analysisStatus: document.getElementById("analysis-status"),
    analysisMetrics: document.getElementById("analysis-metrics")
  };

  state.planStartDate = state.planStartDate || todayIso();
  state.activeSessionDate = state.activeSessionDate || todayIso();

  populateProblemFormOptions();
  refs.sessionDate.value = state.activeSessionDate;
  refs.planStartDate.value = state.planStartDate;
  refs.problemDate.value = todayIso();
  refs.geminiApiKey.value = loadGeminiApiKey();

  bindEvents();
  renderAll();
}

function bindEvents() {
  refs.tabButtons.forEach((button) => {
    button.addEventListener("click", () => activateTab(button.dataset.tab));
  });

  refs.sessionDate.addEventListener("change", (event) => {
    state.activeSessionDate = event.target.value || todayIso();
    saveState();
    renderDailySession();
  });

  refs.planStartDate.addEventListener("change", (event) => {
    state.planStartDate = event.target.value || todayIso();
    saveState();
    renderRoadmap();
    renderStats();
  });

  refs.problemForm.addEventListener("submit", handleProblemSubmit);
  refs.resetForm.addEventListener("click", () => {
    refs.problemForm.reset();
    populateProblemFormOptions();
    refs.problemDate.value = todayIso();
    refs.formStatus.textContent = "";
  });

  refs.syncMarkdown.addEventListener("click", syncMarkdownToFile);
  refs.clearEntries.addEventListener("click", clearEntries);

  refs.geminiApiKey.addEventListener("change", () => {
    saveGeminiApiKey(refs.geminiApiKey.value);
  });
  refs.geminiApiKey.addEventListener("blur", () => {
    saveGeminiApiKey(refs.geminiApiKey.value);
  });
  refs.generateTags.addEventListener("click", handleGenerateTags);
  refs.analyzeSolution.addEventListener("click", handleAnalyzeSolution);

  refs.problemLogList.addEventListener("click", (event) => {
    const deleteButton = event.target.closest("button[data-entry-id]");
    if (!deleteButton) {
      return;
    }
    deleteEntry(deleteButton.dataset.entryId);
  });
}

function activateTab(tabId) {
  refs.tabButtons.forEach((button) => {
    button.classList.toggle("active", button.dataset.tab === tabId);
  });

  refs.tabPanels.forEach((panel) => {
    panel.classList.toggle("active", panel.id === tabId);
  });
}

function handleProblemSubmit(event) {
  event.preventDefault();

  const formData = new FormData(refs.problemForm);
  const entry = {
    id: createEntryId(),
    date: sanitizeText(formData.get("date")) || todayIso(),
    problemName: sanitizeText(formData.get("problemName")),
    topic: normalizeTopic(sanitizeText(formData.get("topic"))),
    difficulty: sanitizeText(formData.get("difficulty")) || "easy",
    status: sanitizeText(formData.get("status")) || "solved",
    problemLink: sanitizeText(formData.get("problemLink")),
    pattern: sanitizeText(formData.get("pattern")),
    mistakeType: sanitizeText(formData.get("mistakeType")) || "none",
    revisitDate: sanitizeText(formData.get("revisitDate")),
    lesson: sanitizeText(formData.get("lesson"))
  };

  if (!entry.problemName || !entry.topic) {
    setStatus(refs.formStatus, "Problem name and topic are required.", "warning");
    return;
  }

  state.problemEntries.push(entry);
  saveState();

  refs.problemForm.reset();
  refs.problemForm.elements.date.value = todayIso();
  setStatus(refs.formStatus, "Problem entry saved.", "success");

  renderProblemLog();
  renderStats();
}

function renderAll() {
  renderDailySession();
  renderRoadmap();
  renderProblemLog();
  renderStats();
}

function renderDailySession() {
  const sessionDate = state.activeSessionDate || todayIso();
  const sessionChecks = state.dailySessions?.[sessionDate] || {};
  const completedCount = DAILY_SESSION_ITEMS.filter((item) => sessionChecks[item.id]).length;
  const totalCount = DAILY_SESSION_ITEMS.length;

  refs.dailySessionChecklist.innerHTML = "";
  DAILY_SESSION_ITEMS.forEach((item) => {
    refs.dailySessionChecklist.appendChild(
      buildCheckboxItem({
        checked: Boolean(sessionChecks[item.id]),
        label: item.label,
        onChange: (checked) => {
          if (!state.dailySessions[sessionDate]) {
            state.dailySessions[sessionDate] = {};
          }
          state.dailySessions[sessionDate][item.id] = checked;
          saveState();
          renderDailySession();
        }
      })
    );
  });

  const currentWeek = getStudyWeekForDate(todayIso());
  if (refs.dailyWeekEyebrow) {
    refs.dailyWeekEyebrow.textContent = currentWeek
      ? `Week ${currentWeek} of 12`
      : "Week — of 12";
  }
  if (refs.sessionProgressCount) {
    refs.sessionProgressCount.textContent = `${completedCount}/${totalCount}`;
  }
  if (refs.sessionProgressFill) {
    refs.sessionProgressFill.style.width = `${totalCount ? (completedCount / totalCount) * 100 : 0}%`;
  }

  const currentWeekLabel = currentWeek ? `You are currently in Week ${currentWeek} of the roadmap.` : "Your 12-week roadmap has not started yet.";
  setStatus(
    refs.dailySessionSummary,
    `You completed ${completedCount} of ${totalCount} checklist steps for ${sessionDate}. ${currentWeekLabel}`
  );
}

function renderRoadmap() {
  refs.weeklyRoadmap.innerHTML = "";

  const currentWeek = getStudyWeekForDate(todayIso());
  let totalItems = 0;
  let completedItems = 0;

  ROADMAP_WEEKS.forEach((week) => {
    const details = document.createElement("details");
    const isCurrent = week.weekNumber === currentWeek;
    details.className = "week-card";
    details.open = isCurrent || week.weekNumber === 1;

    const itemCount = week.tasks.length + week.goals.length;
    const weekCompleted = [...week.tasks, ...week.goals].filter((_, index) => {
      const id = getRoadmapCheckId(week.weekNumber, index);
      return Boolean(state.roadmapChecks[id]);
    }).length;

    if (isCurrent) {
      details.classList.add("current");
    }
    if (itemCount > 0 && weekCompleted === itemCount) {
      details.classList.add("is-done");
    }

    totalItems += itemCount;
    completedItems += weekCompleted;

    const summary = document.createElement("summary");

    const badge = document.createElement("span");
    badge.className = "week-badge";
    badge.textContent = `W${week.weekNumber}`;

    const titleWrap = document.createElement("div");
    titleWrap.className = "week-title-wrap";

    const titleRow = document.createElement("div");
    titleRow.className = "week-title-row";

    const title = document.createElement("span");
    title.className = "week-title";
    title.textContent = week.title;
    titleRow.appendChild(title);

    if (isCurrent) {
      const currentLabel = document.createElement("span");
      currentLabel.className = "week-current-label";
      currentLabel.textContent = "current";
      titleRow.appendChild(currentLabel);
    }

    const meta = document.createElement("div");
    meta.className = "week-meta";
    meta.textContent = `${week.phase} study block`;

    titleWrap.appendChild(titleRow);
    titleWrap.appendChild(meta);

    const progress = document.createElement("span");
    progress.className = "week-progress";
    progress.textContent = `${weekCompleted}/${itemCount}`;

    const chevron = document.createElementNS("http://www.w3.org/2000/svg", "svg");
    chevron.setAttribute("class", "week-chevron");
    chevron.setAttribute("viewBox", "0 0 24 24");
    chevron.setAttribute("fill", "none");
    chevron.setAttribute("stroke", "currentColor");
    chevron.setAttribute("stroke-width", "2");
    chevron.setAttribute("aria-hidden", "true");
    const chevronPath = document.createElementNS("http://www.w3.org/2000/svg", "path");
    chevronPath.setAttribute("d", "M6 9l6 6 6-6");
    chevron.appendChild(chevronPath);

    summary.appendChild(badge);
    summary.appendChild(titleWrap);
    summary.appendChild(progress);
    summary.appendChild(chevron);
    details.appendChild(summary);

    const body = document.createElement("div");
    body.className = "week-body";

    const tasksBlock = document.createElement("div");
    tasksBlock.className = "week-block";
    const tasksHeading = document.createElement("h4");
    tasksHeading.textContent = "Checklist for this week";
    tasksBlock.appendChild(tasksHeading);

    week.tasks.forEach((task, index) => {
      const id = getRoadmapCheckId(week.weekNumber, index);
      tasksBlock.appendChild(
        buildCheckboxItem({
          checked: Boolean(state.roadmapChecks[id]),
          label: formatRoadmapChecklistItem(task),
          onChange: (checked) => {
            state.roadmapChecks[id] = checked;
            saveState();
            renderRoadmap();
          }
        })
      );
    });

    body.appendChild(tasksBlock);

    if (week.goals.length > 0) {
      const goalsBlock = document.createElement("div");
      goalsBlock.className = "week-block";
      const goalsHeading = document.createElement("h4");
      goalsHeading.textContent = "Success goals";
      goalsBlock.appendChild(goalsHeading);

      week.goals.forEach((goal, goalIndex) => {
        const id = getRoadmapCheckId(week.weekNumber, week.tasks.length + goalIndex);
        goalsBlock.appendChild(
          buildCheckboxItem({
            checked: Boolean(state.roadmapChecks[id]),
            label: goal,
            onChange: (checked) => {
              state.roadmapChecks[id] = checked;
              saveState();
              renderRoadmap();
            }
          })
        );
      });

      body.appendChild(goalsBlock);
    }

    details.appendChild(body);
    refs.weeklyRoadmap.appendChild(details);
  });

  const currentWeekText = currentWeek ? `Current focus: Week ${currentWeek}.` : "Your plan has not started yet.";
  setStatus(refs.roadmapSummary, `You have completed ${completedItems} of ${totalItems} roadmap checklist items. ${currentWeekText}`);
}

function renderProblemLog() {
  refs.problemLogList.innerHTML = "";

  if (refs.logCount) {
    refs.logCount.textContent = String(state.problemEntries.length);
  }

  if (state.problemEntries.length === 0) {
    const emptyState = document.createElement("div");
    emptyState.className = "empty-card";
    emptyState.textContent = "No problems logged yet. Add one from the form.";
    refs.problemLogList.appendChild(emptyState);
    return;
  }

  getSortedEntries("desc").forEach((entry) => {
    refs.problemLogList.appendChild(buildLogEntry(entry));
  });
}

function renderStats() {
  const today = todayIso();
  const now = new Date(`${today}T00:00:00`).getTime();
  const weekAgo = now - 7 * DAY_IN_MS;
  const solvedEntries = state.problemEntries.filter((entry) => entry.status === "solved");
  const thisWeek = state.problemEntries.filter((entry) => {
    return new Date(`${entry.date}T00:00:00`).getTime() >= weekAgo;
  });
  const uniqueTopics = new Set(state.problemEntries.map((entry) => normalizeTopic(entry.topic)));
  const revisitDue = state.problemEntries.filter((entry) => entry.revisitDate && entry.revisitDate <= today);

  const cards = [
    {
      label: "Total solved",
      value: String(solvedEntries.length),
      tone: "success",
      icon: '<path d="M22 11.08V12a10 10 0 11-5.93-9.14"/><polyline points="22 4 12 14.01 9 11.01"/>'
    },
    {
      label: "This week",
      value: String(thisWeek.length),
      tone: "primary",
      icon: '<path d="M8.5 14.5A2.5 2.5 0 0011 12c0-1.38-.5-2-1.5-3-1-1.12-1-2.62 0-3.75a2.5 2.5 0 014.5 1.5c0 1.5-1 2.5-2 3.5s-1.5 2-1.5 3.5"/><path d="M12 18v.01"/>'
    },
    {
      label: "Topics covered",
      value: String(uniqueTopics.size),
      tone: "warning",
      icon: '<polygon points="12 2 2 7 12 12 22 7 12 2"/><polyline points="2 17 12 22 22 17"/><polyline points="2 12 12 17 22 12"/>'
    },
    {
      label: "Revisit due",
      value: String(revisitDue.length),
      tone: "danger",
      icon: '<polyline points="23 4 23 10 17 10"/><path d="M20.49 15a9 9 0 11-2.12-9.36L23 10"/>'
    }
  ];

  refs.statsCards.innerHTML = "";
  cards.forEach((card) => {
    const element = document.createElement("div");
    element.className = "stat-card";

    const iconWrap = document.createElement("div");
    iconWrap.className = `stat-icon ${card.tone}`;
    iconWrap.innerHTML = `<svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">${card.icon}</svg>`;

    const textWrap = document.createElement("div");
    const label = document.createElement("div");
    label.className = "stat-label";
    label.textContent = card.label;

    const value = document.createElement("div");
    value.className = "stat-value";
    value.textContent = card.value;

    textWrap.appendChild(label);
    textWrap.appendChild(value);
    element.appendChild(iconWrap);
    element.appendChild(textWrap);
    refs.statsCards.appendChild(element);
  });

  renderStatsMatrix(state.problemEntries);
}

function renderStatsMatrix(entries) {
  const topics = getMatrixTopics(entries);
  refs.statsMatrix.innerHTML = "";
  refs.statsMatrix.className = "heat-matrix";

  const headRow = document.createElement("tr");
  const topicHeader = document.createElement("th");
  topicHeader.textContent = "Topic";
  headRow.appendChild(topicHeader);

  for (let week = 1; week <= TOTAL_WEEKS; week += 1) {
    const header = document.createElement("th");
    header.textContent = `W${week}`;
    headRow.appendChild(header);
  }
  refs.statsMatrix.appendChild(headRow);

  topics.forEach((topic) => {
    const row = document.createElement("tr");
    const topicCell = document.createElement("td");
    topicCell.textContent = formatTopicLabel(topic);
    row.appendChild(topicCell);

    for (let week = 1; week <= TOTAL_WEEKS; week += 1) {
      const count = entries.filter((entry) => {
        return normalizeTopic(entry.topic) === topic && getStudyWeekForDate(entry.date) === week;
      }).length;

      const cell = document.createElement("td");
      const heat = document.createElement("div");
      heat.className = "heat-cell";
      if (count === 1) {
        heat.classList.add("c1");
      } else if (count > 1 && count <= 3) {
        heat.classList.add("c2");
      } else if (count > 3) {
        heat.classList.add("c3");
      }
      heat.textContent = count ? String(count) : "·";
      cell.appendChild(heat);
      row.appendChild(cell);
    }

    refs.statsMatrix.appendChild(row);
  });
}

function getMatrixTopics(entries) {
  const topics = new Set(TOPIC_ORDER);
  entries.forEach((entry) => topics.add(normalizeTopic(entry.topic)));
  return getOrderedTopics([...topics].map((topic) => ({ topic })));
}

function formatTopicLabel(topic) {
  const match = STUDY_PLAN_FORM_OPTIONS.topics.find((item) => item.value === topic);
  if (match && match.label) {
    return match.label;
  }
  return String(topic || "")
    .split("_")
    .filter(Boolean)
    .map((part) => part.charAt(0).toUpperCase() + part.slice(1))
    .join(" ");
}

function buildCheckboxItem({ checked, label, onChange }) {
  const wrapper = document.createElement("label");
  wrapper.className = checked ? "check-item is-checked" : "check-item";

  const input = document.createElement("input");
  input.type = "checkbox";
  input.checked = checked;
  input.addEventListener("change", () => onChange(input.checked));

  const text = document.createElement("span");
  text.textContent = label;

  wrapper.appendChild(input);
  wrapper.appendChild(text);

  return wrapper;
}

function buildLogEntry(entry) {
  const card = document.createElement("article");
  card.className = "log-entry";

  const top = document.createElement("div");
  top.className = "log-entry-top";

  const titleWrap = document.createElement("div");
  titleWrap.style.minWidth = "0";
  titleWrap.style.flex = "1";

  const titleRow = document.createElement("div");
  titleRow.className = "log-entry-title";

  const title = document.createElement("h4");
  title.textContent = entry.problemName;
  titleRow.appendChild(title);

  if (entry.problemLink) {
    const link = document.createElement("a");
    link.className = "ext-link";
    link.href = entry.problemLink;
    link.target = "_blank";
    link.rel = "noopener noreferrer";
    link.setAttribute("aria-label", "Open problem");
    link.innerHTML = '<svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M18 13v6a2 2 0 01-2 2H5a2 2 0 01-2-2V8a2 2 0 012-2h6"/><polyline points="15 3 21 3 21 9"/><line x1="10" y1="14" x2="21" y2="3"/></svg>';
    titleRow.appendChild(link);
  }

  const date = document.createElement("div");
  date.className = "log-date";
  date.textContent = entry.date;

  titleWrap.appendChild(titleRow);
  titleWrap.appendChild(date);

  const deleteButton = document.createElement("button");
  deleteButton.type = "button";
  deleteButton.className = "icon-btn";
  deleteButton.dataset.entryId = entry.id;
  deleteButton.setAttribute("aria-label", "Delete");
  deleteButton.innerHTML = '<svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><polyline points="3 6 5 6 21 6"/><path d="M19 6l-1 14a2 2 0 01-2 2H8a2 2 0 01-2-2L5 6"/><path d="M10 11v6M14 11v6"/><path d="M9 6V4a1 1 0 011-1h4a1 1 0 011 1v2"/></svg>';

  top.appendChild(titleWrap);
  top.appendChild(deleteButton);

  const tags = document.createElement("div");
  tags.className = "tag-row";

  const difficulty = document.createElement("span");
  difficulty.className = `tag diff-${String(entry.difficulty || "").toLowerCase()}`;
  difficulty.textContent = entry.difficulty || "n/a";
  tags.appendChild(difficulty);

  const status = document.createElement("span");
  status.className = `tag status-${String(entry.status || "").toLowerCase()}`;
  status.textContent = entry.status || "n/a";
  tags.appendChild(status);

  const topic = document.createElement("span");
  topic.className = "tag";
  topic.textContent = formatTopicLabel(entry.topic);
  tags.appendChild(topic);

  card.appendChild(top);
  card.appendChild(tags);

  if (entry.pattern || entry.mistakeType || entry.revisitDate) {
    const meta = document.createElement("div");
    meta.className = "log-meta";

    if (entry.pattern) {
      const pattern = document.createElement("p");
      pattern.innerHTML = `<span class="label">Pattern: </span>${escapeHtml(entry.pattern)}`;
      meta.appendChild(pattern);
    }

    if (entry.mistakeType && entry.mistakeType !== "none") {
      const mistake = document.createElement("p");
      mistake.innerHTML = `<span class="label">Mistake: </span>${escapeHtml(entry.mistakeType)}`;
      meta.appendChild(mistake);
    }

    if (entry.revisitDate) {
      const revisit = document.createElement("p");
      revisit.innerHTML = `<span class="label">Revisit: </span><span class="log-date" style="display:inline">${escapeHtml(entry.revisitDate)}</span>`;
      meta.appendChild(revisit);
    }

    card.appendChild(meta);
  }

  if (entry.lesson) {
    const lesson = document.createElement("p");
    lesson.className = "log-lesson";
    lesson.textContent = `"${entry.lesson}"`;
    card.appendChild(lesson);
  }

  return card;
}

function generateMarkdown() {
  const entries = getSortedEntries("asc");
  const solvedCount = state.problemEntries.filter((entry) => entry.status === "solved").length;
  const topics = getOrderedTopics(state.problemEntries);

  const lines = [
    "# Solved Problems Log",
    "",
    "> Generated from `study_play/study_tracker.html`.",
    "",
    "## Summary",
    `- Total entries: ${state.problemEntries.length}`,
    `- Solved entries: ${solvedCount}`,
    `- Topics tracked: ${topics.length}`,
    ""
  ];

  if (entries.length === 0) {
    lines.push("## Entries", "", "No problems logged yet.");
    return lines.join("\n");
  }

  lines.push("## Entries", "");

  entries.forEach((entry) => {
    lines.push(`### ${entry.date} - ${entry.problemName}`);
    lines.push(`- Status: ${entry.status}`);
    lines.push(`- Topic: ${entry.topic}`);
    lines.push(`- Difficulty: ${entry.difficulty}`);
    lines.push(`- Problem link: ${entry.problemLink || "N/A"}`);
    lines.push(`- Pattern: ${entry.pattern || "N/A"}`);
    lines.push(`- Mistake type: ${entry.mistakeType || "N/A"}`);
    lines.push(`- Lesson: ${entry.lesson || "N/A"}`);
    lines.push(`- Revisit: ${entry.revisitDate || "N/A"}`);
    lines.push("");
  });

  return lines.join("\n");
}

async function syncMarkdownToFile() {
  if (!supportsFileSync()) {
    downloadMarkdownFile();
    setStatus(refs.logStatus, getSyncFallbackMessage(), "warning");
    return;
  }

  try {
    const handle = await getSolvedProblemsHandle();
    const writable = await handle.createWritable();
    await writable.write(generateMarkdown());
    await writable.close();
    setStatus(refs.logStatus, "Synced to SOLVED_PROBLEMS.md.", "success");
  } catch (error) {
    if (error && error.name === "AbortError") {
      setStatus(refs.logStatus, "Sync cancelled.", "warning");
      return;
    }

    setStatus(refs.logStatus, getSyncErrorMessage(error), "warning");
  }
}

function downloadMarkdownFile() {
  const markdown = generateMarkdown();
  const blob = new Blob([markdown], { type: "text/markdown;charset=utf-8" });
  const url = URL.createObjectURL(blob);
  const link = document.createElement("a");
  link.href = url;
  link.download = "SOLVED_PROBLEMS.md";
  document.body.appendChild(link);
  link.click();
  document.body.removeChild(link);
  URL.revokeObjectURL(url);
}

function getSyncFallbackMessage() {
  if (window.location.protocol === "file:") {
    return "Downloaded SOLVED_PROBLEMS.md. Replace study_play/SOLVED_PROBLEMS.md in your repo. For one-click sync, run `python3 -m http.server 8080` in study_play and open http://localhost:8080/study_tracker.html in Chrome or Edge.";
  }

  return "Downloaded SOLVED_PROBLEMS.md. Replace study_play/SOLVED_PROBLEMS.md in your repo. Direct sync requires Chrome or Edge.";
}

function clearEntries() {
  if (!window.confirm("Clear all saved problem entries from browser storage?")) {
    return;
  }

  state.problemEntries = [];
  saveState();
  renderProblemLog();
  renderStats();
  setStatus(refs.logStatus, "All saved entries were cleared.", "warning");
}

function deleteEntry(entryId) {
  state.problemEntries = state.problemEntries.filter((entry) => entry.id !== entryId);
  saveState();
  renderProblemLog();
  renderStats();
  setStatus(refs.logStatus, "Entry deleted.", "success");
}

function getStudyWeekForDate(dateString) {
  if (!state.planStartDate || !dateString) {
    return null;
  }

  const startDate = new Date(`${state.planStartDate}T00:00:00`);
  const targetDate = new Date(`${dateString}T00:00:00`);

  if (Number.isNaN(startDate.getTime()) || Number.isNaN(targetDate.getTime()) || targetDate < startDate) {
    return null;
  }

  const diffDays = Math.floor((targetDate.getTime() - startDate.getTime()) / DAY_IN_MS);
  const week = Math.floor(diffDays / 7) + 1;

  if (week < 1 || week > TOTAL_WEEKS) {
    return null;
  }

  return week;
}

function getOrderedTopics(entries) {
  const topics = new Set(entries.map((entry) => normalizeTopic(entry.topic)));
  const ordered = [];

  TOPIC_ORDER.forEach((topic) => {
    if (topics.has(topic)) {
      ordered.push(topic);
      topics.delete(topic);
    }
  });

  return [...ordered, ...Array.from(topics).sort()];
}

function getSortedEntries(order) {
  const factor = order === "desc" ? -1 : 1;
  return [...state.problemEntries].sort((a, b) => {
    const dateCompare = a.date.localeCompare(b.date) * factor;
    if (dateCompare !== 0) {
      return dateCompare;
    }
    return a.problemName.localeCompare(b.problemName) * factor;
  });
}

function getRoadmapCheckId(weekNumber, index) {
  return `week-${weekNumber}-item-${index}`;
}

function formatRoadmapChecklistItem(task) {
  let formatted = task.replace(/^Mon:/, "Monday:")
    .replace(/^Tue:/, "Tuesday:")
    .replace(/^Wed:/, "Wednesday:")
    .replace(/^Thu:/, "Thursday:")
    .replace(/^Fri:/, "Friday:")
    .replace(/^Sat:/, "Saturday:")
    .replace(/^Sun:/, "Sunday:");

  formatted = formatted.replace(/`([^`]+)`/g, (_, reference) => humanizeStudyReference(reference));
  formatted = formatted.replace(/\s+\+\s+/g, " and ");

  return formatted;
}

function humanizeStudyReference(reference) {
  const normalized = String(reference || "").trim().replace(/\\/g, "/");

  if (!normalized) {
    return "";
  }

  if (normalized === "templates/pattern_cheat_sheet.js") {
    return "the pattern cheat sheet";
  }

  if (normalized.startsWith("drills/")) {
    const fileName = normalized.split("/").pop() || normalized;
    const drillName = fileName.replace(/^\d+_/, "").replace(/_reflex\.js$/, "");
    return `${toReadableTitle(drillName)} reflex drill`;
  }

  if (normalized.endsWith("/")) {
    const parts = normalized.split("/").filter(Boolean);
    return `${toReadableTitle(parts.join(" "))} problem set`;
  }

  const fileName = normalized.split("/").pop() || normalized;
  return toReadableTitle(fileName.replace(/\.js$/, ""));
}

function toReadableTitle(value) {
  return String(value || "")
    .replace(/_/g, " ")
    .replace(/\bdp\b/gi, "DP")
    .replace(/\bbfs\b/gi, "BFS")
    .replace(/\bdfs\b/gi, "DFS")
    .replace(/\bii\b/gi, "II")
    .replace(/\biii\b/gi, "III")
    .replace(/\biv\b/gi, "IV")
    .replace(/\bvi\b/gi, "VI")
    .replace(/\bvii\b/gi, "VII")
    .replace(/\b([a-z])/g, (match) => match.toUpperCase());
}

function createEntryId() {
  return `entry-${Date.now()}-${Math.random().toString(16).slice(2, 8)}`;
}

function normalizeTopic(topic) {
  return (topic || "misc").trim().toLowerCase().replace(/\s+/g, "_");
}

function populateProblemFormOptions() {
  populateSelect(refs.problemTopic, STUDY_PLAN_FORM_OPTIONS.topics, "");
  populateSelect(refs.problemDifficulty, STUDY_PLAN_FORM_OPTIONS.difficulties, "easy");
  populateSelect(refs.problemStatus, STUDY_PLAN_FORM_OPTIONS.statuses, "solved");
  populateSelect(refs.problemPattern, STUDY_PLAN_FORM_OPTIONS.patterns, "");
  populateSelect(refs.problemMistake, STUDY_PLAN_FORM_OPTIONS.mistakes, "none");
}

function populateSelect(selectElement, options, selectedValue) {
  if (!selectElement) {
    return;
  }

  selectElement.innerHTML = "";

  options.forEach((optionConfig) => {
    const option = document.createElement("option");
    option.value = optionConfig.value;
    option.textContent = optionConfig.label;
    selectElement.appendChild(option);
  });

  if (selectedValue !== undefined) {
    selectElement.value = selectedValue;
  }
}

function sanitizeText(value) {
  return String(value || "").trim();
}

function todayIso() {
  return new Date().toISOString().slice(0, 10);
}

function setStatus(element, message, tone) {
  if (!element) {
    return;
  }
  element.textContent = message || "";
  element.className = "status";
  if (tone) {
    element.classList.add(tone);
  }
}

function loadState() {
  const fallback = {
    planStartDate: todayIso(),
    activeSessionDate: todayIso(),
    dailySessions: {},
    roadmapChecks: {},
    problemEntries: []
  };

  try {
    const raw = localStorage.getItem(STORAGE_KEY);
    if (!raw) {
      return fallback;
    }

    const parsed = JSON.parse(raw);
    return {
      planStartDate: parsed.planStartDate || fallback.planStartDate,
      activeSessionDate: parsed.activeSessionDate || fallback.activeSessionDate,
      dailySessions: parsed.dailySessions || {},
      roadmapChecks: parsed.roadmapChecks || {},
      problemEntries: Array.isArray(parsed.problemEntries) ? parsed.problemEntries : []
    };
  } catch (error) {
    return fallback;
  }
}

function saveState() {
  localStorage.setItem(STORAGE_KEY, JSON.stringify(state));
}

function supportsFileSync() {
  return typeof window.showOpenFilePicker === "function" && typeof indexedDB !== "undefined";
}

async function getSolvedProblemsHandle() {
  const savedHandle = await loadStoredFileHandle();
  if (savedHandle && await verifyFileHandlePermission(savedHandle)) {
    return savedHandle;
  }

  const [selectedHandle] = await window.showOpenFilePicker({
    multiple: false,
    types: [
      {
        description: "Markdown files",
        accept: {
          "text/markdown": [".md"],
          "text/plain": [".md"]
        }
      }
    ]
  });

  if (!selectedHandle || selectedHandle.name !== "SOLVED_PROBLEMS.md") {
    throw new Error("Please select the existing SOLVED_PROBLEMS.md file.");
  }

  const granted = await verifyFileHandlePermission(selectedHandle);
  if (!granted) {
    throw new Error("Write permission was not granted for SOLVED_PROBLEMS.md.");
  }

  await storeFileHandle(selectedHandle);
  return selectedHandle;
}

async function verifyFileHandlePermission(handle) {
  const options = { mode: "readwrite" };

  if (await handle.queryPermission(options) === "granted") {
    return true;
  }

  return (await handle.requestPermission(options)) === "granted";
}

function openFileHandleDb() {
  return new Promise((resolve, reject) => {
    const request = indexedDB.open(FILE_HANDLE_DB_NAME, 1);

    request.onupgradeneeded = () => {
      const db = request.result;
      if (!db.objectStoreNames.contains(FILE_HANDLE_STORE)) {
        db.createObjectStore(FILE_HANDLE_STORE);
      }
    };

    request.onsuccess = () => resolve(request.result);
    request.onerror = () => reject(request.error);
  });
}

async function storeFileHandle(handle) {
  const db = await openFileHandleDb();

  return new Promise((resolve, reject) => {
    const transaction = db.transaction(FILE_HANDLE_STORE, "readwrite");
    const store = transaction.objectStore(FILE_HANDLE_STORE);
    const request = store.put(handle, SOLVED_PROBLEMS_HANDLE_KEY);

    request.onsuccess = () => resolve();
    request.onerror = () => reject(request.error);
    transaction.oncomplete = () => db.close();
    transaction.onerror = () => reject(transaction.error);
  });
}

async function loadStoredFileHandle() {
  const db = await openFileHandleDb();

  return new Promise((resolve, reject) => {
    const transaction = db.transaction(FILE_HANDLE_STORE, "readonly");
    const store = transaction.objectStore(FILE_HANDLE_STORE);
    const request = store.get(SOLVED_PROBLEMS_HANDLE_KEY);

    request.onsuccess = () => resolve(request.result || null);
    request.onerror = () => reject(request.error);
    transaction.oncomplete = () => db.close();
    transaction.onerror = () => reject(transaction.error);
  });
}

function getSyncErrorMessage(error) {
  if (!error) {
    return "Sync failed.";
  }

  if (typeof error.message === "string" && error.message.trim()) {
    return error.message;
  }

  return "Sync failed.";
}

function loadGeminiApiKey() {
  try {
    return localStorage.getItem(GEMINI_API_KEY_STORAGE) || "";
  } catch (error) {
    return "";
  }
}

function saveGeminiApiKey(value) {
  const key = sanitizeText(value);
  try {
    if (key) {
      localStorage.setItem(GEMINI_API_KEY_STORAGE, key);
    } else {
      localStorage.removeItem(GEMINI_API_KEY_STORAGE);
    }
  } catch (error) {
    // Ignore storage failures; the in-memory input value still works for this session.
  }
  return key;
}

function getGeminiApiKey() {
  const fromInput = sanitizeText(refs.geminiApiKey && refs.geminiApiKey.value);
  if (fromInput) {
    saveGeminiApiKey(fromInput);
    return fromInput;
  }
  return loadGeminiApiKey();
}

function setButtonBusy(button, busy, busyLabel) {
  if (!button) {
    return;
  }
  if (busy) {
    if (!button.dataset.idleLabel) {
      button.dataset.idleLabel = button.textContent;
    }
    button.disabled = true;
    button.textContent = busyLabel || "Working...";
    return;
  }
  button.disabled = false;
  button.textContent = button.dataset.idleLabel || button.textContent;
}

async function handleGenerateTags() {
  const question = sanitizeText(refs.analyzerQuestion.value);
  const apiKey = getGeminiApiKey();

  if (!apiKey) {
    setStatus(refs.tagStatus, "Add your Gemini API key first.", "warning");
    refs.geminiApiKey.focus();
    return;
  }

  if (!question) {
    setStatus(refs.tagStatus, "Paste a question before generating tags.", "warning");
    refs.analyzerQuestion.focus();
    return;
  }

  setButtonBusy(refs.generateTags, true, "Generating...");
  setStatus(refs.tagStatus, "Asking Gemini for related topics...", "");

  try {
    const topicCatalog = STUDY_PLAN_FORM_OPTIONS.topics
      .filter((topic) => topic.value)
      .map((topic) => topic.label)
      .join(", ");

    const prompt = [
      "You are a DSA coach. Read the coding interview question and identify related topics and patterns.",
      "Prefer tags from this study-plan catalog when they fit: " + topicCatalog + ".",
      "You may also add precise pattern tags such as prefix sum, monotonic stack, BFS, DFS, union-find, etc.",
      "Return ONLY valid JSON with this shape:",
      '{"tags":["tag1","tag2"],"primaryTopic":"main topic","rationale":"one short sentence"}',
      "Use 3 to 8 concise lowercase tags. No markdown fences.",
      "",
      "QUESTION:",
      question
    ].join("\n");

    const data = await callGeminiJson(apiKey, prompt);
    const tags = normalizeTagList(data.tags);

    if (!tags.length) {
      throw new Error("Gemini returned no usable tags.");
    }

    renderGeneratedTags(tags, data.primaryTopic, data.rationale);
    setStatus(refs.tagStatus, "Tags generated.", "success");
  } catch (error) {
    setStatus(refs.tagStatus, getAnalyzerErrorMessage(error), "error");
  } finally {
    setButtonBusy(refs.generateTags, false);
  }
}

async function handleAnalyzeSolution() {
  const question = sanitizeText(refs.analyzerQuestion.value);
  const solution = sanitizeText(refs.analyzerSolution.value);
  const apiKey = getGeminiApiKey();

  if (!apiKey) {
    setStatus(refs.analysisStatus, "Add your Gemini API key first.", "warning");
    refs.geminiApiKey.focus();
    return;
  }

  if (!question) {
    setStatus(refs.analysisStatus, "Paste the question above so complexity analysis has context.", "warning");
    refs.analyzerQuestion.focus();
    return;
  }

  if (!solution) {
    setStatus(refs.analysisStatus, "Paste your solution before analyzing.", "warning");
    refs.analyzerSolution.focus();
    return;
  }

  setButtonBusy(refs.analyzeSolution, true, "Analyzing...");
  setStatus(refs.analysisStatus, "Asking Gemini to review your solution...", "");

  try {
    const prompt = [
      "You are a strict but helpful DSA interviewer reviewing a candidate solution.",
      "Use the question and the candidate solution below.",
      "Find what is wrong or incomplete in the answer (logic bugs, missing edge cases, incorrect approach, unclear reasoning, wrong complexity claims, etc.).",
      "If the solution looks correct, say so clearly and still note any weak spots or missing proofs.",
      "Estimate the time and space complexity of THIS submitted solution.",
      "Also list ALL realistic time complexities with which this problem can be solved,",
      "from brute force through better approaches to the best known practical interview target.",
      "Do NOT only return the optimal complexity. Include multiple achievable complexities with a short approach label for each.",
      "Mark which listed complexity is optimal and which one matches the submitted solution when possible.",
      "Return ONLY valid JSON with this exact shape:",
      "{",
      '  "verdict":"correct|partially_correct|incorrect",',
      '  "whatIsWrong":["issue 1","issue 2"],',
      '  "whatIsCorrect":["strength 1"],',
      '  "submittedTimeComplexity":"O(...)",',
      '  "submittedSpaceComplexity":"O(...)",',
      '  "possibleTimeComplexities":[',
      '    {"complexity":"O(n^2)","approach":"brute force","isOptimal":false,"matchesSubmitted":true},',
      '    {"complexity":"O(n)","approach":"hash map","isOptimal":true,"matchesSubmitted":false}',
      "  ],",
      '  "targetComplexitiesNote":"short note about which complexities are acceptable interview targets",',
      '  "fixSuggestions":["suggestion 1"]',
      "}",
      "No markdown fences. Keep lists concise.",
      "",
      "QUESTION:",
      question,
      "",
      "CANDIDATE SOLUTION:",
      solution
    ].join("\n");

    const data = await callGeminiJson(apiKey, prompt);
    renderAnalysisMetrics(data);
    setStatus(refs.analysisStatus, "Analysis complete.", "success");
  } catch (error) {
    setStatus(refs.analysisStatus, getAnalyzerErrorMessage(error), "error");
  } finally {
    setButtonBusy(refs.analyzeSolution, false);
  }
}

async function callGeminiJson(apiKey, prompt) {
  const url = `https://generativelanguage.googleapis.com/v1beta/models/${GEMINI_MODEL}:generateContent?key=${encodeURIComponent(apiKey)}`;

  const response = await fetch(url, {
    method: "POST",
    headers: {
      "Content-Type": "application/json"
    },
    body: JSON.stringify({
      contents: [
        {
          role: "user",
          parts: [{ text: prompt }]
        }
      ],
      generationConfig: {
        temperature: 0.2,
        responseMimeType: "application/json"
      }
    })
  });

  const payload = await response.json().catch(() => ({}));

  if (!response.ok) {
    const apiMessage = payload && payload.error && payload.error.message
      ? payload.error.message
      : `Gemini request failed (${response.status}).`;
    throw new Error(apiMessage);
  }

  const text = extractGeminiText(payload);
  return parseGeminiJson(text);
}

function extractGeminiText(payload) {
  const parts = payload
    && payload.candidates
    && payload.candidates[0]
    && payload.candidates[0].content
    && payload.candidates[0].content.parts;

  if (!Array.isArray(parts) || !parts.length) {
    const blockReason = payload
      && payload.promptFeedback
      && payload.promptFeedback.blockReason;
    throw new Error(blockReason ? `Gemini blocked the request: ${blockReason}` : "Gemini returned an empty response.");
  }

  return parts.map((part) => part.text || "").join("\n").trim();
}

function parseGeminiJson(text) {
  const cleaned = String(text || "")
    .replace(/^```(?:json)?\s*/i, "")
    .replace(/\s*```$/i, "")
    .trim();

  try {
    return JSON.parse(cleaned);
  } catch (error) {
    const start = cleaned.indexOf("{");
    const end = cleaned.lastIndexOf("}");
    if (start >= 0 && end > start) {
      return JSON.parse(cleaned.slice(start, end + 1));
    }
    throw new Error("Could not parse Gemini JSON response.");
  }
}

function normalizeTagList(tags) {
  if (!Array.isArray(tags)) {
    return [];
  }

  const seen = new Set();
  const result = [];

  tags.forEach((tag) => {
    const normalized = sanitizeText(tag).toLowerCase();
    if (!normalized || seen.has(normalized)) {
      return;
    }
    seen.add(normalized);
    result.push(normalized);
  });

  return result;
}

function renderGeneratedTags(tags, primaryTopic, rationale) {
  refs.generatedTags.classList.remove("empty-hint");
  refs.generatedTags.style.display = "flex";
  refs.generatedTags.innerHTML = "";

  const hint = document.createElement("p");
  hint.className = "hint";
  hint.style.flexBasis = "100%";
  hint.textContent = "Suggested topics";
  refs.generatedTags.appendChild(hint);

  tags.forEach((tag) => {
    const span = document.createElement("span");
    span.className = "tag accent";
    span.textContent = tag;
    refs.generatedTags.appendChild(span);
  });

  if (sanitizeText(primaryTopic)) {
    const primary = document.createElement("span");
    primary.className = "tag diff-easy";
    primary.textContent = `primary: ${sanitizeText(primaryTopic)}`;
    refs.generatedTags.appendChild(primary);
  }

  if (sanitizeText(rationale)) {
    const note = document.createElement("div");
    note.className = "muted";
    note.style.flexBasis = "100%";
    note.style.marginTop = "4px";
    note.textContent = sanitizeText(rationale);
    refs.generatedTags.appendChild(note);
  }
}

function renderAnalysisMetrics(data) {
  const tbody = refs.analysisMetrics.querySelector("tbody") || refs.analysisMetrics;
  tbody.innerHTML = "";

  const verdict = sanitizeText(data.verdict) || "unknown";
  const wrongItems = toStringList(data.whatIsWrong);
  const correctItems = toStringList(data.whatIsCorrect);
  const suggestions = toStringList(data.fixSuggestions);
  const complexities = Array.isArray(data.possibleTimeComplexities) ? data.possibleTimeComplexities : [];

  appendMetricRow(tbody, "Verdict", formatVerdict(verdict));
  appendMetricRow(tbody, "What is wrong", renderBulletList(wrongItems.length ? wrongItems : ["Nothing major flagged."]));
  appendMetricRow(tbody, "What looks correct", renderBulletList(correctItems.length ? correctItems : ["No strengths listed."]));
  appendMetricRow(
    tbody,
    "Your solution complexity",
    escapeHtml(`Time ${sanitizeText(data.submittedTimeComplexity) || "n/a"} · Space ${sanitizeText(data.submittedSpaceComplexity) || "n/a"}`)
  );
  appendMetricRow(tbody, "All possible time complexities", renderComplexityList(complexities));
  appendMetricRow(
    tbody,
    "Target note",
    escapeHtml(sanitizeText(data.targetComplexitiesNote) || "Aim for the best complexity you can explain cleanly in an interview.")
  );
  appendMetricRow(tbody, "Fix suggestions", renderBulletList(suggestions.length ? suggestions : ["No extra suggestions."]));
}

function appendMetricRow(tbody, label, valueHtml) {
  const row = document.createElement("tr");
  const th = document.createElement("th");
  const td = document.createElement("td");
  th.scope = "row";
  th.textContent = label;
  td.innerHTML = valueHtml;
  row.appendChild(th);
  row.appendChild(td);
  tbody.appendChild(row);
}

function renderBulletList(items) {
  return `<ul class="metrics-list">${items.map((item) => `<li>${escapeHtml(item)}</li>`).join("")}</ul>`;
}

function renderComplexityList(items) {
  if (!items.length) {
    return `<span class="muted">No complexity options returned.</span>`;
  }

  return items.map((item) => {
    const complexity = sanitizeText(item && item.complexity) || "O(?)";
    const approach = sanitizeText(item && item.approach);
    const classes = ["complexity-chip"];
    if (item && item.isOptimal) {
      classes.push("optimal");
    }
    if (item && item.matchesSubmitted) {
      classes.push("yours");
    }
    const badges = [];
    if (item && item.isOptimal) {
      badges.push("optimal");
    }
    if (item && item.matchesSubmitted) {
      badges.push("yours");
    }
    const badgeText = badges.length ? ` [${badges.join(", ")}]` : "";
    const label = approach ? `${complexity} — ${approach}${badgeText}` : `${complexity}${badgeText}`;
    return `<span class="${classes.join(" ")}">${escapeHtml(label)}</span>`;
  }).join("");
}

function formatVerdict(verdict) {
  const normalized = String(verdict || "").toLowerCase().replace(/\s+/g, "_");
  const labels = {
    correct: "Correct",
    partially_correct: "Partially correct",
    incorrect: "Incorrect",
    unknown: "Unknown"
  };
  return escapeHtml(labels[normalized] || sanitizeText(verdict) || "Unknown");
}

function toStringList(value) {
  if (!Array.isArray(value)) {
    return [];
  }
  return value.map((item) => sanitizeText(item)).filter(Boolean);
}

function escapeHtml(value) {
  return String(value)
    .replace(/&/g, "&amp;")
    .replace(/</g, "&lt;")
    .replace(/>/g, "&gt;")
    .replace(/"/g, "&quot;")
    .replace(/'/g, "&#39;");
}

function getAnalyzerErrorMessage(error) {
  if (!error) {
    return "Analyzer request failed.";
  }
  if (typeof error.message === "string" && error.message.trim()) {
    return error.message;
  }
  return "Analyzer request failed.";
}
