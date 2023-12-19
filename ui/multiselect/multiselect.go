package multiselect

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

var (
	style              = lipgloss.NewStyle()
	focusedStyle       = style.Copy().Border(lipgloss.RoundedBorder(), false, false, false, true)
	noFocusStyle       = style.Copy().PaddingLeft(1)
	itemStyle          = style.Copy()
	itemSelectedStyle  = style.Copy()
	itemDisabledStyle  = style.Copy().Strikethrough(true)
	itemFocusedStyle   = style.Copy()
	itemUnfocusedStyle = style.Copy().MarginLeft(2)

	checkMarkStyle = lipgloss.NewStyle()
	cursorStyle    = lipgloss.NewStyle()
)

var (
	checkMark = "✓"
	cursor    = ">"
)

func itemDisabledRender(item string) string {
	return itemDisabledStyle.Render(
		fmt.Sprintf("[ ] %s", item),
	)
}

func itemSelectedRender(item string) string {
	return itemSelectedStyle.Render(
		fmt.Sprintf("[%s] %s", checkMarkStyle.Render(checkMark), item),
	)
}

func itemRender(item string) string {
	return itemStyle.Render(
		fmt.Sprintf("[ ] %s", item),
	)
}

func itemFocusedRender(item string) string {
	return itemFocusedStyle.Render(
		fmt.Sprintf("%s %s", cursorStyle.Render(cursor), item),
	)
}

func itemUnfocusedRender(item string) string {
	return itemUnfocusedStyle.Render(item)
}

func View(
	cursor int,
	items []string,
	selected, disabled map[string]struct{},
	focused bool,
) string {
	var itemViews []string
	for i, item := range items {
		var itemView string
		if _, ok := disabled[item]; ok {
			itemView = itemDisabledRender(item)
		} else if _, ok := selected[item]; ok {
			itemView = itemSelectedRender(item)
		} else {
			itemView = itemRender(item)
		}
		if (i == cursor) && focused {
			itemView = itemFocusedRender(itemView)
		} else {
			itemView = itemUnfocusedRender(itemView)
		}
		itemViews = append(itemViews, itemView)
	}
	var focusstyle lipgloss.Style
	if focused {
		focusstyle = focusedStyle
	} else {
		focusstyle = noFocusStyle
	}
	return focusstyle.Render(lipgloss.JoinVertical(lipgloss.Left, itemViews...))
}
