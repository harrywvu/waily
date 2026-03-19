# Waily

**Waily** is a lightweight, keyboard-driven CLI app for tracking your daily wins — what we call *Wails*. Built in Go with SQLite storage, it groups entries by day into *Streams* and lets you add, view, edit, and delete them entirely from the keyboard.

---

## Features

- **Add Daily Wails** — Quickly log a win for today. Multiple wails per day are grouped into the same stream automatically.
- **View Streams** — Browse all past streams listed by date with their stream ID.
- **Open a Stream** — Inspect all wails inside any stream by ID.
- **Edit Wails** — Update the content of any individual wail.
- **Delete Wails** — Remove a single wail from a stream.
- **Delete Streams** — Wipe an entire day's stream at once.
- **Instant keyboard navigation** — Menu keys register immediately with no Enter required and no character echo.
- **Centered UI** — The interface auto-centers horizontally and vertically based on your terminal size.
- **Color-coded status** — Feedback messages are automatically green (success), red (error), or yellow (neutral).

---

## Storage

Waily uses a local SQLite database (`wails.db`) created automatically on first run.

**Schema:**

```sql
CREATE TABLE wails (
    id        INTEGER PRIMARY KEY AUTOINCREMENT,
    timestamp TEXT,   -- RFC3339Nano, e.g. "2026-02-22T15:30:12.000000000+08:00"
    date      TEXT,   -- "YYYY-MM-DD", e.g. "2026-02-22"
    content   TEXT,   -- "Finished a workout"
    stream_id INTEGER -- groups wails for the same day together
)
```

Each day gets one `stream_id`. If you add multiple wails on the same day, they all share that stream's ID.

---

## Navigation

All menu keys are **instant** — press once, no Enter needed, nothing echoes to the screen. ID inputs (stream ID, wail ID) are typed normally and confirmed with Enter, since they can be multi-digit.

### Main Menu

```
[Q] Quit
[A] Add Daily Wail    [V] View Streams
```

| Key | Action |
|-----|--------|
| `a` | Add a new wail to today's stream |
| `v` | View the stream list |
| `q` | Quit |

### Stream List (`v`)

Displays all streams grouped by date:

```
ID      DATE
──      ──────────────────
1       2026-02-14
2       2026-02-15
```

| Key | Action |
|-----|--------|
| `1` | Open a stream (enter Stream ID) |
| `2` | Delete a stream (enter Stream ID) |
| `0` | Back to main menu |

### Inside a Stream (`1`)

Displays all wails in the selected stream:

```
[2026-02-15 — STREAM 2]
────────────────────────────────────────────
ID      WAIL
──      ────────────────────────────
1       Finished a Workout
2       Code Habit: 1 Hour Go
```

| Key | Action |
|-----|--------|
| `1` | Edit a wail (enter Wail ID, then new content) |
| `2` | Delete a wail (enter Wail ID) |
| `0` | Back to main menu |

---

## Project Structure

```
.
├── main.go              # DB init, menu loop, routing
├── wails.go             # CRUD operations and stream logic
├── wails.db             # SQLite database (auto-created)
└── helpers/
    ├── print.go         # All terminal output, layout, ANSI, centering
    └── scanner.go       # Input: GetKeyPress (raw), GetUserInputString, GetUserInputInt
```

---

## Dependencies

```
golang.org/x/term   — raw terminal mode (instant keypresses + terminal sizing)
github.com/mattn/go-sqlite3 — SQLite driver
```

Install:

```bash
go get golang.org/x/term
go get github.com/mattn/go-sqlite3
```

> `go-sqlite3` requires CGO. Make sure you have a C compiler available (`gcc` on Linux/macOS, or [TDM-GCC](https://jmeubank.github.io/tdm-gcc/) on Windows).

---

## Running

```bash
go run .
```

Or build a binary:

```bash
go build -o waily .
./waily
```
