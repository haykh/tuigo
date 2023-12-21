package text

import (
	"github.com/haykh/tuigo/ui/theme"
	"github.com/haykh/tuigo/utils"
)

var (
	style       = theme.ElementStyle.Copy()
	normalStyle = style.Copy()
	labelStyle  = style.Copy().Foreground(theme.ColorAccent).Padding(0, 1)
	dimmedStyle = style.Copy().Foreground(theme.ColorDimmed)
)

func View(focused bool, txt string, texttype utils.TextType) string {
	switch texttype {
	case utils.NormalText:
		return normalStyle.Render(txt)
	case utils.LabelText:
		return labelStyle.Render(txt)
	case utils.DimmedText:
		return dimmedStyle.Render(txt)
	default:
		return style.Render(txt)
	}
}
