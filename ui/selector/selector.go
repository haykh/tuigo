package selector

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/haykh/tuigo/ui/theme"
)

var (
	style        = lipgloss.NewStyle().MarginBottom(1)
	focusedStyle = style.Copy()
	noFocusStyle = style.Copy()

	noFocusItemStyle = lipgloss.NewStyle().Foreground(theme.ColorDimmed)

	itemStyle          = lipgloss.NewStyle()
	itemSelectedStyle  = itemStyle.Copy()
	itemFocusedStyle   = itemStyle.Copy()
	itemUnfocusedStyle = itemStyle.Copy().MarginLeft(2)

	markStyle   = lipgloss.NewStyle().Foreground(theme.ColorSuccess)
	cursorStyle = lipgloss.NewStyle()
)

var (
	checkMark  = "✓"
	selectMark = "+"
	cursor     = ">"
)

func itemFocusedRender(item string) string {
	return itemFocusedStyle.Render(
		fmt.Sprintf("%s %s", cursorStyle.Render(cursor), item),
	)
}

func itemUnfocusedRender(item string) string {
	return itemUnfocusedStyle.Render(item)
}

func View(
	multiselect bool,
	cursor int,
	items []string,
	selected, disabled map[string]struct{},
	focused bool,
) string {
	var itemViews []string
	var focusstyle, textstyle lipgloss.Style
	if focused {
		textstyle = lipgloss.NewStyle()
		focusstyle = focusedStyle
	} else {
		textstyle = noFocusItemStyle
		focusstyle = noFocusStyle
	}
	var mark string
	if multiselect {
		mark = selectMark
	} else {
		mark = checkMark
	}
	for i, item := range items {
		var itemView string
		lb := textstyle.Render("[")
		rb := textstyle.Render("]")
		it := textstyle.Render(item)
		cm := markStyle.Render(mark)
		if _, ok := disabled[item]; ok {
			ts := textstyle.Copy().Strikethrough(true)
			itemView = ts.Render("[ ] " + item)
		} else if _, ok := selected[item]; ok {
			itemView = itemSelectedStyle.Render(
				fmt.Sprintf("%s%s%s %s", lb, cm, rb, it),
			)
		} else {
			itemView = itemStyle.Render(
				fmt.Sprintf("%s %s %s", lb, rb, it),
			)
		}
		if (i == cursor) && focused {
			itemView = itemFocusedRender(itemView)
		} else {
			itemView = itemUnfocusedRender(itemView)
		}
		itemViews = append(itemViews, itemView)
	}
	return focusstyle.Render(lipgloss.JoinVertical(lipgloss.Left, itemViews...))
}
