package selector

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/haykh/tuigo/ui/theme"
)

var (
	style        = theme.ElementStyle.Copy().MarginLeft(0)
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
	focused bool,
	multiselect bool,
	cursor int,
	items []string,
	selected, disabled map[string]struct{},
	viewLimit int,
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
	if viewLimit == len(items)+1 {
		viewLimit = -1
	}
	if viewLimit > 0 {
		if cursor >= viewLimit && cursor < len(items)-1 {
			itemViews = append(itemViews, noFocusItemStyle.Render("^v ..."))
		} else if cursor < viewLimit {
			itemViews = append(itemViews, noFocusItemStyle.Render("vv ..."))
		} else if cursor >= len(items)-1 {
			itemViews = append(itemViews, noFocusItemStyle.Render("^^ ..."))
		}
	}
	for i, item := range items {
		var itemView string
		if viewLimit > 0 {
			if cursor < viewLimit && i >= viewLimit {
				break
			}
			if cursor >= viewLimit && i == cursor+1 {
				break
			}
			if i <= cursor-viewLimit {
				continue
			}
		}
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
