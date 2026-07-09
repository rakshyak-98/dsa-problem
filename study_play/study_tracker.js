const STORAGE_KEY = "studyTrackerState.v1";
const TOTAL_WEEKS = 12;
const DAY_IN_MS = 24 * 60 * 60 * 1000;
const FILE_HANDLE_DB_NAME = "studyTrackerFileHandles";
const FILE_HANDLE_STORE = "handles";
const SOLVED_PROBLEMS_HANDLE_KEY = "solvedProblemsMd";

const DAILY_SESSION_ITEMS = [
  { id: "trigger-scan", label: "Review the pattern triggers before you start coding" },
  { id: "restate", label: "Write down the input, output, constraints, and edge cases" },
  { id: "bruteforce", label: "Start with the brute-force idea before optimizing" },
  { id: "pattern-scan", label: "Choose the most likely pattern before writing code" },
  { id: "trace", label: "Trace one small example by hand" },
  { id: "tests", label: "Run tests or sample checks after coding" },
  { id: "postmortem", label: "Write a short lesson learned and a revisit date" }
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
      "Mon: `drills/01_arrays_reflex.js` + `arrays/easy/plus_one.js` + `arrays/easy/concatenation_of_array.js`",
      "Tue: Re-drill arrays + `arrays/easy/find_closest_number_to_zero.js` + `arrays/easy/max_consecutive_ones.js`",
      "Wed: `drills/02_hashing_reflex.js` + `hashing/easy/two_sum.js` + `hashing/easy/contains_duplicates.js`",
      "Thu: Re-drill hashing + `hashing/easy/fair_candy_swap.js` + `hashing/easy/degree_of_an_array.js`",
      "Fri: Mixed review + `arrays/easy/majority_element.js` + `hashing/easy/summary_ranges.js`",
      "Sat: Review failed problems and re-type drills blind",
      "Sun: Rest"
    ],
    goals: [
      "Write `reverseArray`, `maxInArray`, and `countFreq` from memory in under 3 minutes.",
      "Explain O(n) vs O(n^2) for every solution you wrote.",
      "Solve 8 or more easy problems without peeking."
    ]
  },
  {
    weekNumber: 2,
    phase: "Phase 1",
    title: "Strings and simulation",
    tasks: [
      "Mon: `strings/easy/find_words_containing_character.js` + `strings/easy/most_common_word.js`",
      "Tue: Hashing on strings + `strings/easy/unique_morse_code_words.js` + `hashing/easy/find_resultant_array_after_removing_anagrams.js`",
      "Wed: `simulation/easy/baseball_game.js` + `simulation/easy/relative_ranks.js`",
      "Thu: Solve 3 unseen `strings/easy/` problems",
      "Fri: Timed mock with 2 easy problems in 45 minutes",
      "Sat: Re-drill `01_arrays_reflex.js` and `02_hashing_reflex.js`",
      "Sun: Rest"
    ],
    goals: [
      "Feel comfortable with `Map`, `Set`, and object frequency maps.",
      "Reach 16 or more total easy problems solved in this repo."
    ]
  },
  {
    weekNumber: 3,
    phase: "Phase 2",
    title: "Two pointers",
    tasks: [
      "Mon: `drills/03_two_pointers_reflex.js` + `two_pointers/easy/remove_duplicates_from_sorted_array.js` + `two_pointers/easy/move_zeroes.js`",
      "Tue: Re-drill + `two_pointers/easy/best_time_to_buy_sell_stock.js` + `two_pointers/easy/squares_of_a_sorted_array.js`",
      "Wed: Medium intro with `two_pointers/medium/container_with_most_water.js` + `two_pointers/medium/3sum.js`",
      "Thu: Sliding window with `misc/easy/maximum_average_subarray_1.js` + `two_pointers/easy/minimum_difference_between_highest_and_lowest_of_k_score.js`",
      "Fri: Timed 1 easy + 1 medium from `two_pointers/`",
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
      "Mon: `drills/04_binary_search_reflex.js` + `binary_search/easy/search_insertion_position.js` + `binary_search/easy/find_smallest_letter_greater_than_target.js`",
      "Tue: Binary search on answer with `binary_search/easy/longest_subsequence_with_limited_sum.js`",
      "Wed: Mixed pointers and binary search with `two_pointers/medium/find_the_duplicate_number.js`",
      "Thu: Medium batch with `two_pointers/medium/sort_colors.js` + `two_pointers/medium/longest_palindromic_substring.js`",
      "Fri: Mock interview with 1 medium in 45 minutes",
      "Sat: Review all medium attempts and tag the mistake type",
      "Sun: Rest"
    ],
    goals: [
      "Make `left`, `right`, `while (left < right)` muscle memory.",
      "Make `while (left <= right)` binary search muscle memory.",
      "Solve at least 5 medium two-pointer problems."
    ]
  },
  {
    weekNumber: 5,
    phase: "Phase 3",
    title: "Hashing mastery and medium arrays",
    tasks: [
      "Mon: Full rewrite of `drills/02_hashing_reflex.js` + `hashing/medium/group_anagram.js` + `hashing/medium/top_k_ferquent_element.js`",
      "Tue: Prefix sum with `misc/medium/subarray_sum_divisible_by_k.js` + `misc/easy/find_pivot_index.js`",
      "Wed: Array medium with `arrays/medium/max_product_subarray.js` + `arrays/medium/find_all_duplicates_in_an_array.js`",
      "Thu: `hashing/medium/longest_consecutive_sequence.js`",
      "Fri: `misc/medium/product_of_array_except_self.js`",
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
      "Mon: Re-solve 2 medium problems you previously got wrong",
      "Tue: Re-solve 2 medium problems you previously got wrong",
      "Wed: Re-solve 2 medium problems you previously got wrong",
      "Thu: Re-solve 2 medium problems you previously got wrong",
      "Fri: Timed set of 3 medium problems in 2 hours",
      "Sat: Write a one-page pattern journal",
      "Sun: Rest"
    ],
    goals: [
      "Reach 10 or more medium problems solved.",
      "Explain prefix sum in one sentence with an example.",
      "Solve group anagrams and top-k without notes."
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
      "Mon: 1 new medium, max 90 minutes",
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
      "Mon: 1 new medium, max 90 minutes",
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
      "Mon: 1 new medium, max 90 minutes",
      "Tue: Re-solve Monday's medium from scratch",
      "Wed: 1 new medium and classify the pattern",
      "Thu: 1 hard study problem, 30 minute attempt plus editorial rewrite",
      "Fri: Timed set of 2 medium problems in 90 minutes",
      "Sat: Random drill file + `templates/pattern_cheat_sheet.js`",
      "Sun: Rest"
    ],
    goals: [
      "Reach 25 or more medium problems total.",
      "Study at least 3 hard problems.",
      "List 2 to 3 approaches for a random medium in 5 minutes."
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
    statsCards: document.getElementById("stats-cards"),
    statsMatrix: document.getElementById("stats-matrix")
  };

  state.planStartDate = state.planStartDate || todayIso();
  state.activeSessionDate = state.activeSessionDate || todayIso();

  populateProblemFormOptions();
  refs.sessionDate.value = state.activeSessionDate;
  refs.planStartDate.value = state.planStartDate;
  refs.problemDate.value = todayIso();

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
  const currentWeekLabel = currentWeek ? `You are currently in Week ${currentWeek} of the roadmap.` : "Your 12-week roadmap has not started yet.";
  setStatus(
    refs.dailySessionSummary,
    `You completed ${completedCount} of ${DAILY_SESSION_ITEMS.length} checklist steps for ${sessionDate}. ${currentWeekLabel}`
  );
}

function renderRoadmap() {
  refs.weeklyRoadmap.innerHTML = "";

  const currentWeek = getStudyWeekForDate(todayIso());
  let totalItems = 0;
  let completedItems = 0;

  ROADMAP_WEEKS.forEach((week) => {
    const details = document.createElement("details");
    details.className = "week-card";
    details.open = week.weekNumber === currentWeek || week.weekNumber === 1;

    const summary = document.createElement("summary");

    const titleWrap = document.createElement("div");
    const title = document.createElement("span");
    title.textContent = `Week ${week.weekNumber}: ${week.title}`;
    const meta = document.createElement("div");
    meta.className = "week-meta";
    meta.textContent = `${week.phase} study block`;
    titleWrap.appendChild(title);
    titleWrap.appendChild(meta);

    const itemCount = week.tasks.length + week.goals.length;
    const weekCompleted = [...week.tasks, ...week.goals].filter((_, index) => {
      const id = getRoadmapCheckId(week.weekNumber, index);
      return Boolean(state.roadmapChecks[id]);
    }).length;

    totalItems += itemCount;
    completedItems += weekCompleted;

    const progress = document.createElement("span");
    progress.className = "week-meta";
    progress.textContent = `${weekCompleted}/${itemCount} completed`;

    summary.appendChild(titleWrap);
    summary.appendChild(progress);
    details.appendChild(summary);

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

    details.appendChild(tasksBlock);

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

      details.appendChild(goalsBlock);
    }

    refs.weeklyRoadmap.appendChild(details);
  });

  const currentWeekText = currentWeek ? `Current focus: Week ${currentWeek}.` : "Your plan has not started yet.";
  setStatus(refs.roadmapSummary, `You have completed ${completedItems} of ${totalItems} roadmap checklist items. ${currentWeekText}`);
}

function renderProblemLog() {
  refs.problemLogList.innerHTML = "";

  if (state.problemEntries.length === 0) {
    const emptyState = document.createElement("div");
    emptyState.className = "muted";
    emptyState.textContent = "No problems logged yet.";
    refs.problemLogList.appendChild(emptyState);
    return;
  }

  getSortedEntries("desc").forEach((entry) => {
    refs.problemLogList.appendChild(buildLogEntry(entry));
  });
}

function renderStats() {
  const solvedEntries = state.problemEntries.filter((entry) => entry.status === "solved");
  const currentWeek = getStudyWeekForDate(todayIso());
  const currentWeekSolved = solvedEntries.filter((entry) => getStudyWeekForDate(entry.date) === currentWeek).length;
  const uniqueTopics = new Set(state.problemEntries.map((entry) => normalizeTopic(entry.topic)));

  const cards = [
    { label: "Total entries", value: String(state.problemEntries.length) },
    { label: "Solved entries", value: String(solvedEntries.length) },
    { label: "Topics tracked", value: String(uniqueTopics.size) },
    { label: "Current week", value: currentWeek ? `Week ${currentWeek}` : "Not started" },
    { label: "Solved this week", value: String(currentWeekSolved) }
  ];

  refs.statsCards.innerHTML = "";
  cards.forEach((card) => {
    const element = document.createElement("div");
    element.className = "stat-card";

    const label = document.createElement("div");
    label.className = "stat-label";
    label.textContent = card.label;

    const value = document.createElement("div");
    value.className = "stat-value";
    value.textContent = card.value;

    element.appendChild(label);
    element.appendChild(value);
    refs.statsCards.appendChild(element);
  });

  renderStatsMatrix(solvedEntries);
}

function renderStatsMatrix(solvedEntries) {
  const topics = getOrderedTopics(solvedEntries);
  refs.statsMatrix.innerHTML = "";

  const headRow = document.createElement("tr");
  const topicHeader = document.createElement("th");
  topicHeader.textContent = "Topic";
  headRow.appendChild(topicHeader);

  for (let week = 1; week <= TOTAL_WEEKS; week += 1) {
    const header = document.createElement("th");
    header.textContent = `W${week}`;
    headRow.appendChild(header);
  }

  const totalHeader = document.createElement("th");
  totalHeader.textContent = "Total";
  headRow.appendChild(totalHeader);
  refs.statsMatrix.appendChild(headRow);

  if (topics.length === 0) {
    const emptyRow = document.createElement("tr");
    const emptyCell = document.createElement("td");
    emptyCell.colSpan = TOTAL_WEEKS + 2;
    emptyCell.textContent = "No solved problems yet. Add entries in the Problem Log tab.";
    emptyRow.appendChild(emptyCell);
    refs.statsMatrix.appendChild(emptyRow);
    return;
  }

  topics.forEach((topic) => {
    const row = document.createElement("tr");
    const topicCell = document.createElement("td");
    topicCell.textContent = topic;
    row.appendChild(topicCell);

    let total = 0;

    for (let week = 1; week <= TOTAL_WEEKS; week += 1) {
      const count = solvedEntries.filter((entry) => {
        return normalizeTopic(entry.topic) === topic && getStudyWeekForDate(entry.date) === week;
      }).length;
      total += count;

      const cell = document.createElement("td");
      cell.textContent = count ? String(count) : "-";
      row.appendChild(cell);
    }

    const totalCell = document.createElement("td");
    totalCell.textContent = String(total);
    row.appendChild(totalCell);

    refs.statsMatrix.appendChild(row);
  });
}

function buildCheckboxItem({ checked, label, onChange }) {
  const wrapper = document.createElement("label");
  wrapper.className = "check-item";

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

  const header = document.createElement("div");
  header.className = "summary";

  const titleWrap = document.createElement("div");
  const title = document.createElement("h4");
  title.textContent = `${entry.date} - ${entry.problemName}`;
  titleWrap.appendChild(title);

  if (entry.problemLink) {
    const link = document.createElement("a");
    link.href = entry.problemLink;
    link.target = "_blank";
    link.rel = "noopener noreferrer";
    link.textContent = entry.problemLink;
    titleWrap.appendChild(link);
  }

  const deleteButton = document.createElement("button");
  deleteButton.type = "button";
  deleteButton.className = "danger";
  deleteButton.dataset.entryId = entry.id;
  deleteButton.textContent = "Delete";

  header.appendChild(titleWrap);
  header.appendChild(deleteButton);

  const tags = document.createElement("div");
  tags.className = "tag-row";
  [
    entry.topic,
    entry.difficulty,
    entry.status,
    entry.pattern || "pattern not set",
    entry.mistakeType
  ].forEach((value) => {
    const tag = document.createElement("span");
    tag.className = "tag";
    tag.textContent = value;
    tags.appendChild(tag);
  });

  const lesson = document.createElement("p");
  lesson.textContent = entry.lesson || "No lesson recorded yet.";

  const revisit = document.createElement("div");
  revisit.className = "muted";
  revisit.textContent = entry.revisitDate ? `Revisit: ${entry.revisitDate}` : "Revisit date not set";

  card.appendChild(header);
  card.appendChild(tags);
  card.appendChild(lesson);
  card.appendChild(revisit);

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
  element.textContent = message;
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
