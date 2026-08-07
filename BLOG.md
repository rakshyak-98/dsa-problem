# Building a Local-First Interview Prep Gym: Intent, Features, and Skill Growth

**Repo:** [github.com/rakshyak-98/dsa-problem](https://github.com/rakshyak-98/dsa-problem)

Interview prep usually fails from missing rhythm, not missing problems. I had study plans, LeetCode lists, and scattered notes — but no single system that turned *today* into a concrete action and kept a record of what I actually learned.
So I built **dsa-problem**: a local-first workspace that treats interview prep like gym training — daily drills, spaced repetition, honest logging, and tools that stay in your repo instead of on a SaaS dashboard.
This post explains why it exists, what it does, and how the design is meant to improve skill over time.

---

## The intent

### Problem I was solving

Most prep stacks look like this:

- A **problem list** (170+ files sorted by topic)
- A **study plan** (week-by-week roadmap)
- A **tracker** (spreadsheet or Notion)

What was missing was the **bridge** between plan and practice:

1. **Rhythm** — same daily shape, not reinventing the session each morning
2. **Reflex** — patterns written from memory until basics are automatic
3. **Understanding** — slowing down before pattern-matching the wrong problem
4. **Reading** — training the skill of *reading* solutions, not just writing them
5. **Backend depth** — verbal + code drills mapped to a real resume, not generic trivia
6. **Evidence** — a log of mistakes, lessons, and revisit dates

### Design principles

| Principle | What it means in practice |
|-----------|---------------------------|
| **Local-first** | Go stdlib + static HTML. No backend, no accounts, no API keys required |
| **Consistency > intensity** | Minimum tier (20 min) counts. Missed days are forward-only — no catch-up guilt |
| **Understand first** | Restate the ask in one sentence before any code |
| **Blind before peek** | Solutions and answer keys exist — but only after an honest attempt |
| **Same ritual daily** | Core drills every session, weekday specialty on rotation |

The goal is not to “finish LeetCode.” It is to make medium problems *readable* — where you can open a prompt, pass a translation test, and list two or three approaches in five minutes without panic.

---

## Features

### 1. Unified daily command

One entry point from the repo root:

```bash
go run .                              # full DSA day: read + write
go run . -- --track backend           # backend interview track
go run . -- --track read              # reading drills only
go run . -- --track write             # writing drills only
go run . -- --drill core              # Core 5 + Core Read 3
go run . -- --drill reflex            # today's specialty only
go run . -- --run core                # validate core answers
go run . -- --solution reflex         # show solution paths (after attempt)
```

The runner prints a **DAILY** header with today’s read file, write file, Core 5 functions, and the exact commands to run. No guessing which drill is “due.”

### 2. Write reflex drills (DSA coding muscle)

**Core 5** — five functions every session, target under 8 minutes from memory:

- `twoSum` — hash map pairing
- `binarySearch` — sorted search template
- `removeDuplicates` — in-place two-pointer on sorted array
- `maxSumSubarrayK` — fixed-size sliding window
- `frequencyMap` — counting with a map

**Weekday specialty drills** rotate through ten pattern files:

| Drill | Pattern family |
|-------|----------------|
| `01_arrays_reflex` | iteration, reversal, max/min |
| `02_hashing_reflex` | maps, sets, frequency |
| `03_two_pointers_reflex` | opposite ends, fast/slow |
| `04_binary_search_reflex` | search space, rotated arrays |
| `05_trees_stacks_reflex` | BFS/DFS, stack operations |
| `06_dp_reflex` | 1D DP, state definition |
| `07_graphs_reflex` | grid BFS/DFS, visited sets |
| `08_heap_reflex` | priority queue patterns |
| `09_backtrack_reflex` | choice trees, pruning |
| `10_math_reflex` | counting, modular arithmetic |

Each drill is a runnable Go file with `TODO: REFLEX` stubs and inline `PASS:` asserts. Implement blind, run, fix, repeat.

### 3. Code reading drills

Writing code is half the interview. Reading code — editorials, teammates’ PRs, your own old solutions — is the other half.

The **6-pass read method** trains structured scanning:

1. **Signature** — input, output, side effects
2. **Skeleton** — loops, recursion, early returns
3. **State** — what each variable *means*
4. **Trace** — hand-execute one example
5. **Pattern** — name the template (window, BS, DFS, DP…)
6. **Ask + bound** — one-sentence problem + complexity

**Core Read 3** runs daily; weekday specialties include find-the-bug, name-the-pattern, complexity-at-a-glance, and compare-variants.

### 4. Backend interview pack

For system design + backend rounds, a separate track maps resume bullets to timed drills:

- **Explain drills** — verbal concept cards (`TODO: EXPLAIN`)
- **Write drills** — Go reflex implementations (JWT middleware, worker pools, circuit breakers)
- **Scenario drills** — STAR stories and mock prompts
- **8 resume blocks** — REST/JWT, SQL, distributed resilience, WebRTC, workflows, DevOps/AWS, compliance, Go systems

```bash
go run ./bin/study_backend -- --cram     # hour-by-hour cram schedule
go run ./bin/study_backend -- --run      # validate written answers
```

### 5. Study tracker (browser)

A single-page app at `drills/tracker/study_tracker.html`:

- **Daily checklist** tied to the 12-week plan
- **Problem log** — topic, difficulty, pattern, mistake type, one-line lesson, revisit date
- **Weekly heat map** — visual consistency over time
- **Optional Gemini analyzer** — paste a problem for pattern hints (API key stays in your browser)

Open the file directly — no server required. Data lives in `localStorage`.

### 6. Algorithm visualizer

`reference/visualizer/` — write plain algorithm code; the parser auto-detects array, string, tree, or graph structures and visualizes steps. Useful when a trace on paper is not enough.

### 7. Problem catalog + jargon glossary

- **~170 problems** in `reference/problems/` organized by topic and difficulty
- **`doc/DSA_JARGON.md`** — plain-English one-liner definitions for every term you hit in drills (subarray vs subsequence, invariant, monotonic stack, etc.)

### 8. Solutions layer (honest-attempt only)

After you try:

- `drills/solutions/*.md` — triggers, bugs, pattern notes
- `drills/solutions/reflex/` — runnable Go per reflex drill
- `drills/read/answers/` — reading drill answer keys

The workflow is deliberate: **attempt → run tests → peek one snippet → re-type from memory**.

### Repository layout

```
dsa-problem/
├── doc/           guides, study plans, jargon glossary
├── drills/        practice files (write, read, backend, tracker)
├── bin/           CLI helpers (study_play, study_code, study_backend)
├── reference/     problem catalog + visualizer
├── main.go        unified daily runner
└── setup.sh       one-command bootstrap
```

---

## How this improves skill

The system is built around a **5-phase, 12-week roadmap** — but skill growth comes from repeatable mechanics, not calendar dates.

### Phase progression

```
Phase 1  Weeks 1–2   Foundations — arrays, hashing, question literacy
Phase 2  Weeks 3–4   Two pointers, sliding window, binary search
Phase 3  Weeks 5–6   Hashing mastery, prefix sums, medium arrays
Phase 4  Weeks 7–9   Trees, graphs, stacks, DP intro
Phase 5  Weeks 10–12 Medium consolidation, hard introduction
```

Each phase has checklists: “can write `reverseArray` in under 3 minutes,” “can pass the translation test every time,” “solved 8+ medium primaries.”

### The daily ritual (same shape every day)

```
1. Understand  →  restate the question in your own words
2. Reflex      →  Core 5 + today's specialty drill
3. Primary     →  ONE problem from the weekly table
4. Log         →  one sentence lesson + revisit date
```

**Energy tiers** protect the habit:

| Tier | Time | What counts |
|------|------|-------------|
| Minimum | ~20–30 min | Core 5 + one-line log |
| Standard | 45–60 min | Core 5 + specialty + one primary problem |
| Stretch | 75–90 min | Standard + optional second problem or blind re-solve |

Missed a day? Resume today. No catch-up pile.

### Skill layers the system trains

| Layer | Mechanism | Skill outcome |
|-------|-----------|---------------|
| **Reflex** | Core 5 + weekday drills from memory | Working memory free for reasoning, not syntax recall |
| **Literacy** | Question checklist + translation test | Fewer “solved the wrong problem” failures |
| **Pattern recognition** | Trigger scan table + reading drills | See sorted → two pointers; subarray sum → prefix + map |
| **Reading** | 6-pass method on snippets | Editorials become study material, not copy-paste |
| **Post-mortem** | Mistake taxonomy in tracker | Same failure type does not repeat three times |
| **Spaced repetition** | Revisit dates (+1, +3, +7, +21 days) | Problems stick past the first solve |
| **Verbal** | Backend explain + STAR drills | Resume bullets become 90-second stories with metrics |

### The problem-solving framework (medium/hard)

When stuck, the plan forces a process instead of panic:

1. **Understand** (3–5 min) — literacy checklist, translation test
2. **Restate** — input, output, constraints, edges
3. **Brute force** — name the naive solution and its complexity
4. **Pattern scan** — walk the trigger table
5. **Sketch** — trace on paper
6. **Code** — use templates from `pattern_cheat_sheet.go`
7. **Post-mortem** — pattern used, mistake type, one-line lesson

### What “improvement” looks like concretely

Early weeks:

- Hash map syntax is slow; two-pointer invariants are fuzzy
- Misread “return index” vs “return value”
- Medium problems feel like a wall after 10 minutes

After consistent Standard-tier weeks:

- Core 5 completes in under 8 minutes blind
- Open a medium, pass translation test in under 2 minutes
- Mistake log shifts from `pattern miss` → `edge case` → occasional `timeout`
- Reading an editorial: Pass 4 (trace) succeeds before peeking at code
- Backend block: lead with EXPLAIN plan → index → pool, not a rambling architecture tour

The tracker’s heat map and lesson log make this visible — not just “I did 50 problems,” but “I stopped repeating the same misunderstanding.”

---

## Who this is for

- **You have a problem list but no rhythm** — the unified `go run .` and daily checklist fix that
- **You pattern-match before understanding** — question literacy is built into every primary problem
- **You only write code, never read it** — the read track is a separate skill path
- **You are prepping backend/system design** — resume-mapped blocks, not generic flashcards
- **You want local tools** — no subscription, data stays on your machine

## Quick start

```bash
git clone https://github.com/rakshyak-98/dsa-problem.git
cd dsa-problem
./setup.sh
go run .                    # see today's plan
open drills/tracker/study_tracker.html
```

Read the full roadmap: [`doc/write/STUDY_PLAN.md`](doc/write/STUDY_PLAN.md)  
Drills overview: [`doc/drills.md`](doc/drills.md)  
Plain-English glossary: [`doc/DSA_JARGON.md`](doc/DSA_JARGON.md)

---

## Closing thought

Consistency beats intensity. The Minimum tier still counts.

Log one problem. Write one lesson. Show up with the same shape tomorrow.

That is the whole system.
