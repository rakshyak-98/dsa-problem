# DSA Study Tracker

Browser-based **12-week interview prep companion** — daily checklists, problem logging, weekly stats, and a Gemini-powered analyzer.

Part of the [dsa-problem](https://github.com/rakshyak-98/dsa-problem) practice workspace: plan + drills + tracker + visualizer in one place.

---

## Why this exists

Interview prep fails less from lack of resources and more from lack of rhythm:

- No clear “what am I doing today?”
- Solved problems vanish into browser history
- Weak topics stay weak because progress isn’t measured
- Lessons after a bug never get written down

This tracker is the front door for the study plan in this folder: turn the plan into daily action and keep a record of every problem.

---

## Features

### Daily Practice
- Session checklist (understand → reflex → one primary → log)
- Plan start date so the tracker knows which week you’re in
- Expandable **12-week roadmap** with checkable goals
- Progress saved in `localStorage`

### Problem Log
- Log name, topic, difficulty, status, pattern, mistake type, revisit date, one-sentence lesson
- Card-style history with tags
- Optional sync to `SOLVED_PROBLEMS.md` (Chrome/Edge over localhost)

### Weekly Stats
- Cards: total solved, this week, topics covered, revisits due
- **Topic × Week heat matrix** aligned to your plan start

### AI Analyzer
- Paste a problem → Gemini suggests topic tags
- Paste a solution → correctness / complexity-style review
- API key stored only in the browser

Also includes a **DSA Visualizer** at [`visualizer/index.html`](./visualizer/index.html).

---

## Stack

| Piece | Choice |
|--------|--------|
| UI | HTML + CSS |
| Logic | Vanilla JavaScript |
| Storage | Browser `localStorage` |
| AI | Google Gemini (optional) |
| Plan / drills | Markdown + JS in this folder |

No build step — open the HTML file and it runs.

---

## Quick start

```bash
git clone https://github.com/rakshyak-98/dsa-problem.git
cd dsa-problem
```

Open the tracker:

```bash
# macOS
open study_play/study_tracker.html

# Linux
xdg-open study_play/study_tracker.html

# Or serve the folder (better for markdown sync)
npx serve study_play
```

1. Set a **plan start date**
2. Check today’s session items
3. Log your first problem
4. (Optional) Add a Gemini API key under **AI Analyzer** from [Google AI Studio](https://aistudio.google.com/apikey)

Day-one practice flow: see [`START_HERE.md`](./START_HERE.md). Full roadmap: [`STUDY_PLAN.md`](./STUDY_PLAN.md).

---

## Folder map

```
study_play/
├── study_tracker.html      # tracker UI
├── study_tracker.js
├── START_HERE.md           # day-one instructions
├── STUDY_PLAN.md           # 12-week roadmap
├── DAILY_30MIN_DRILL.md
├── BLOG_POST.md            # short public write-up
├── drills/                 # reflex drills with tests
├── visualizer/             # algorithm visualizer
└── SOLVED_PROBLEMS.md      # optional synced log
```

---

## Design choices

- **Local-first** — no account, no backend
- **One lesson per problem** — short reflection, not empty checkmarks
- **Revisit dates** — log doubles as a spaced-review queue
- **Week-aware stats** — relative to *your* plan start
- **Calm UI** — dark theme, clear hierarchy, mobile-friendly tabs

---

## Roadmap ideas

- Export / import JSON for backup across browsers
- Deeper visualizer ↔ tracker links
- Optional cloud sync for multi-device practice
- Live demo URL

---

## Contributing

Issues, forks, and PRs are welcome. Short public write-up: [`BLOG_POST.md`](./BLOG_POST.md).

**License / usage:** explore freely; adapt the workflow to your own plan.
