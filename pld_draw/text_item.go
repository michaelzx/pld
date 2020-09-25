package pld_draw

import (
	"github.com/fogleman/gg"
	"strings"
)

type TextItem struct {
	Content    string  // 内容
	LineWidth  float64 // 行宽
	LineHeight float64 // 行高:1\1.5,字体倍数
	FontPath   string  // 字体路径
	FontSize   float64 // 字体大小
	Color      string  // 字体颜色
	X          float64 // x轴坐标
	Y          float64 // y轴坐标
	En         bool
}

func (ti *TextItem) Draw(ggc *gg.Context) (endX, endY float64) {
	ggc.SetHexColor(ti.Color)
	err := ggc.LoadFontFace(ti.FontPath, ti.FontSize)
	if err != nil {
		panic("字体加载失败")
	}
	var lines []string
	if ti.En {
		lines = ti.getEnLines(ggc)
	} else {
		lines = ti.getCnLines(ggc)
	}
	endY = ti.Y + ti.FontSize
	lh := ti.FontSize * ti.LineHeight
	for _, line := range lines {
		lineW, _ := ggc.MeasureString(line)
		lineEndX := ti.X + lineW
		if lineEndX > endX {
			endX = lineEndX
		}
		ggc.DrawString(line, ti.X, endY)
		ggc.Push()
		endY += lh
	}
	return
}
func (ti *TextItem) getCnLines(ggc *gg.Context) []string {
	var lines []string
	words := []rune(ti.Content)
	currentLine := ""
	for i, word := range words {
		currentLine += string(word)
		if i == len(words)-1 {
			lines = append(lines, currentLine)
			currentLine = ""
		} else {
			currentLineNext := currentLine + string(words[i+1])
			lineW, _ := ggc.MeasureString(currentLineNext)
			if lineW > ti.LineWidth {
				lines = append(lines, currentLine)
				currentLine = ""
			}
		}
	}
	return lines
}

func (ti *TextItem) getEnLines(ggc *gg.Context) []string {
	words := strings.Split(ti.Content, " ")
	var lines []string
	currentLine := ""
	for i, word := range words {
		if currentLine == "" {
			currentLine += word
		} else {
			currentLine += " " + word
		}
		if i == len(words)-1 {
			lines = append(lines, currentLine)
			currentLine = ""
		} else {
			currentLineNext := currentLine + " " + words[i+1]
			lineW, _ := ggc.MeasureString(currentLineNext)
			if lineW > ti.LineWidth {
				lines = append(lines, currentLine)
				currentLine = ""
			}
		}
	}
	return lines
}
