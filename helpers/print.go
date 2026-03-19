package helpers

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"golang.org/x/term"
)

// ── ANSI ──────────────────────────────────────────────────────────────────────

const (
	Reset = "\033[0m"
	Bold  = "\033[1m"
	Dim   = "\033[2m"

	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Cyan   = "\033[36m"
)

// contentWidth is the fixed width of the UI block (divider length).
// All content is assumed to fit within this width.
const contentWidth = 46

// ── Terminal sizing ───────────────────────────────────────────────────────────

type layout struct {
	hPad string // horizontal prefix applied to every line
	vPad int    // number of blank lines printed before the block
}

// contentLines is the total number of lines the main menu block occupies.
// Count: 1 (shortcuts) + 5 (title art + blank) + 1 (date) + 1 (blank) +
//
//	1 (divider) + 1 (options) + 1 (divider) + 2 (status + blank, worst case) +
//	2 (blanks) + 1 (prompt) = ~16 lines. Use 18 for comfortable breathing room.
const contentLines = 18

func getLayout() layout {
	cols, rows, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		// Non-TTY or unsupported — no padding.
		return layout{hPad: "", vPad: 0}
	}

	hPadCount := (cols - contentWidth) / 2
	if hPadCount < 0 {
		hPadCount = 0
	}

	vPadCount := (rows - contentLines) / 2
	if vPadCount < 0 {
		vPadCount = 0
	}

	return layout{
		hPad: strings.Repeat(" ", hPadCount),
		vPad: vPadCount,
	}
}

// p prefixes a line with the horizontal padding.
func p(pad, s string) string {
	return pad + s
}

// ── Low-level primitives ──────────────────────────────────────────────────────

func ClearTerminal() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func PrintNewLine() { fmt.Print("\n") }

func printDivider(pad string) {
	fmt.Println(p(pad, Dim+strings.Repeat("─", contentWidth)+Reset))
}

// ── Header blocks ─────────────────────────────────────────────────────────────

func PrintTitle(pad string) {
	lines := []string{
		"▗▖ ▗▖ ▗▄▖ ▗▄▄▄▖▗▖ ▗▖  ▗▖",
		"▐▌ ▐▌▐▌ ▐▌  █  ▐▌  ▝▚▞▘ ",
		"▐▌ ▐▌▐▛▀▜▌  █  ▐▌   ▐▌  ",
		"▐▙█▟▌▐▌ ▐▌▗▄█▄▖▐▙▄▄▖▐▌  ",
	}
	fmt.Println()
	for _, line := range lines {
		fmt.Println(p(pad, Cyan+line+Reset))
	}
	today := time.Now()
	fmt.Println(p(pad, Dim+today.Format("Monday, January 2, 2006")+Reset))
}

func PrintShortcuts(pad string) {
	fmt.Println(p(pad, Dim+"[Q] Quit"+Reset))
}

func PrintOptions(pad string) {
	PrintNewLine()
	printDivider(pad)
	fmt.Println(p(pad, Bold+Cyan+"[A]"+Reset+" Add Daily Wail    "+Bold+Cyan+"[V]"+Reset+" View Streams"))
	printDivider(pad)
}

// PrintStatus prints the status line with automatic color coding.
func PrintStatus(pad, s string) {
	if s == "" {
		return
	}
	color := Yellow
	lower := strings.ToLower(s)
	switch {
	case strings.Contains(lower, "success") ||
		strings.Contains(lower, "added") ||
		strings.Contains(lower, "edited") ||
		strings.Contains(lower, "deleted"):
		color = Green
	case strings.Contains(lower, "error") ||
		strings.Contains(lower, "does not exist") ||
		strings.Contains(lower, "cannot") ||
		strings.Contains(lower, "empty") ||
		strings.Contains(lower, "invalid") ||
		strings.Contains(lower, "wrong"):
		color = Red
	}
	fmt.Println()
	fmt.Println(p(pad, color+Bold+s+Reset))
}

// PrintHeader is the main screen renderer.
func PrintHeader(statusMessage string) {
	l := getLayout()

	ClearTerminal()

	// Vertical padding — push block to vertical center.
	fmt.Print(strings.Repeat("\n", l.vPad))

	PrintShortcuts(l.hPad)
	PrintTitle(l.hPad)
	PrintStatus(l.hPad, statusMessage)
	PrintOptions(l.hPad)
	PrintNewLine()
	fmt.Print(p(l.hPad, Bold+"› "+Reset))
}

// ── Section headers ───────────────────────────────────────────────────────────

func PrintSection(label string) {
	l := getLayout()
	PrintNewLine()
	fmt.Println(p(l.hPad, Bold+Cyan+label+Reset))
	printDivider(l.hPad)
}

// ── Stream list view ──────────────────────────────────────────────────────────

func PrintMasterStreamFilteredByStreamID(ids []int, dates []string) {
	l := getLayout()
	PrintSection("STREAMS")
	if len(ids) == 0 {
		fmt.Println(p(l.hPad, Dim+"No streams yet."+Reset))
		return
	}
	fmt.Println(p(l.hPad, Dim+Bold+fmt.Sprintf("%-6s  %s", "ID", "DATE")+Reset))
	fmt.Println(p(l.hPad, Dim+fmt.Sprintf("%-6s  %s", "──", "──────────────────")+Reset))
	for i, id := range ids {
		fmt.Println(p(l.hPad, fmt.Sprintf("%-6d  %s", id, dates[i])))
	}
}

// ── Single stream view ────────────────────────────────────────────────────────

func PrintStream(streamID int, date string, wailIDs []int, wails []string) {
	l := getLayout()
	label := fmt.Sprintf("[%s — STREAM %d]", strings.ToUpper(date), streamID)
	PrintSection(label)
	if len(wailIDs) == 0 {
		fmt.Println(p(l.hPad, Dim+"No wails in this stream."+Reset))
		return
	}
	fmt.Println(p(l.hPad, Dim+Bold+fmt.Sprintf("%-6s  %s", "ID", "WAIL")+Reset))
	fmt.Println(p(l.hPad, Dim+fmt.Sprintf("%-6s  %s", "──", "────────────────────────────")+Reset))
	for i, id := range wailIDs {
		fmt.Println(p(l.hPad, fmt.Sprintf("%-6d  %s", id, wails[i])))
	}
}

// ── Action prompts ────────────────────────────────────────────────────────────

func PrintActionBar(options string) {
	l := getLayout()
	fmt.Println()
	fmt.Println(p(l.hPad, Dim+options+Reset))
	fmt.Print(p(l.hPad, Bold+"› "+Reset))
}

func PrintInlinePrompt(label string) {
	l := getLayout()
	fmt.Println()
	fmt.Print(p(l.hPad, Cyan+"›"+Reset+" "+label+": "))
}