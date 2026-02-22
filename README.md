# Waily

**Waily** is a lightweight, keyboard-driven CLI app for tracking your daily wins—what we call *Wails*. It's a CRUD application built around the idea of reflecting on small victories each day, implemented in Go with JSON storage.

---

## Features

- **Add Daily Wails** – Quickly add a new Wail for the current day.
- **View Wail Streams** – Browse all your past Wails by date.
- **Edit Wails** – Modify or delete any Wail within a stream.
- **Delete Streams** – Remove an entire day's Wails if needed.
- **Simple JSON Storage** – Each stream stored as a flat array of Wails.

---

## Prototype Workflow (V1)

1. **Main Menu**
   - Displays today's date and number of Wails added.
   - Shows notifications when Wails are added or edited.

2. **Add Daily Wail**
   - Prompt: `Daily Wail: *Enter Win*`
   - Press Enter to add to today’s stream.

3. **View Wail Streams**
   - List all dates with Wails.
   - Use cursor keys to select a stream.
   - `DEL` → Delete entire stream.
   - `INS` → Enter Edit Mode for a stream.

4. **Edit Mode**
   - Select individual Wails.
   - `DEL` → Delete selected Wail.
   - `INS` → Edit selected Wail.

5. **Edit Individual Wail**
   - Prompt: `Edit Wail: *Enter New Wail*`
   - Press Enter to update.

---

## Storage Format

Each Wail object contains:

```json
{
  "date": "2026-02-22",
  "timestamp": "15:30:12",
  "content": "Completed Go CLI tutorial"
}
