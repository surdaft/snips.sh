package styles

import "github.com/charmbracelet/lipgloss"

var (
	Colors = struct {
		Primary lipgloss.TerminalColor
		Green   lipgloss.TerminalColor
		Red     lipgloss.TerminalColor
		Cyan    lipgloss.TerminalColor
		Yellow  lipgloss.TerminalColor
		Blue    lipgloss.TerminalColor
		Pink    lipgloss.TerminalColor
		Purple  lipgloss.TerminalColor
		White   lipgloss.TerminalColor
		Muted   lipgloss.TerminalColor
		Black   lipgloss.TerminalColor
	}{
		Primary: lipgloss.AdaptiveColor{Dark: "#0ac5b2", Light: "#12a899"},
		Green:   lipgloss.AdaptiveColor{Dark: "#63c174", Light: "#347540"},
		Red:     lipgloss.AdaptiveColor{Dark: "#ff6368", Light: "#aa3337"},
		Yellow:  lipgloss.AdaptiveColor{Dark: "#f1a10d", Light: "#f1a10d"},
		Blue:    lipgloss.AdaptiveColor{Dark: "#52a9ff", Light: "#2667a8"},
		Pink:    lipgloss.AdaptiveColor{Dark: "#f76191", Light: "#ba4368"},
		Purple:  lipgloss.AdaptiveColor{Dark: "#bf7af0", Light: "#6b3393"},
		White:   lipgloss.AdaptiveColor{Dark: "7", Light: "#000000"},
		Muted:   lipgloss.AdaptiveColor{Dark: "8", Light: "#404040"},
		Black:   lipgloss.AdaptiveColor{Dark: "16", Light: "#FFFFFF"},
	}
)

func C(c lipgloss.TerminalColor, s string) string {
	return lipgloss.NewStyle().Foreground(c).Render(s)
}

func B(s string) string {
	return lipgloss.NewStyle().Bold(true).Render(s)
}

func BC(c lipgloss.TerminalColor, s string) string {
	return lipgloss.NewStyle().Foreground(c).Bold(true).Render(s)
}

func U(s string) string {
	return lipgloss.NewStyle().Underline(true).Render(s)
}

func UC(c lipgloss.TerminalColor, s string) string {
	return lipgloss.NewStyle().Foreground(c).Underline(true).Render(s)
}
