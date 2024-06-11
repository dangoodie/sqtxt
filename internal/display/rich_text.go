package display

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
	"image/color"
)

type RichText struct {
	widget.BaseWidget
	Chars      []RichChar
	LineHeight float32
	CharWidth  float32
}

type RichChar struct {
	Char    rune
	FgColor color.Color
	BgColor color.Color
}

func NewRichText() *RichText {
	r := &RichText{
		Chars:      []RichChar{},
		LineHeight: 20,
		CharWidth:  10,
	}
	r.ExtendBaseWidget(r)
	return r
}

func (r *RichText) CreateRenderer() fyne.WidgetRenderer {
	objects := []fyne.CanvasObject{}
	for i, ch := range r.Chars {
		// Background rectangle
		bgRect := canvas.NewRectangle(ch.BgColor)
		bgRect.Resize(fyne.NewSize(r.CharWidth, r.LineHeight))

		x := float32(i) * r.CharWidth
		y := float32(0) // TODO: Calculate y position for multiline text

		bgRect.Move(fyne.NewPos(x, y))
		objects = append(objects, bgRect)

		// Foreground text
		text := canvas.NewText(string(ch.Char), ch.FgColor)
		text.TextSize = r.LineHeight
		text.Move(fyne.NewPos(x, y))
		objects = append(objects, text)
	}
	return &RichTextRenderer{
		objects:  objects,
		richText: r,
	}
}

func (r *RichText) SetChars(chars []RichChar) {
	r.Chars = chars
	r.Refresh()
}

type RichTextRenderer struct {
	objects  []fyne.CanvasObject
	richText *RichText
}

func (r *RichTextRenderer) Layout(size fyne.Size) {
	for i, obj := range r.objects {
		x := float32(i) * r.richText.CharWidth
		y := float32(0) // TODO: Calculate y position for multiline text
		obj.Move(fyne.NewPos(x, y))
	}
}

func (r *RichTextRenderer) MinSize() fyne.Size {
	return fyne.NewSize(float32(len(r.richText.Chars))*r.richText.CharWidth, r.richText.LineHeight)
}

func (r *RichTextRenderer) Refresh() {
	canvas.Refresh(r.richText)
}

func (r *RichTextRenderer) BackgroundColor() color.Color {
	return color.White
}

func (r *RichTextRenderer) Objects() []fyne.CanvasObject {
	return r.objects
}

func (r *RichTextRenderer) Destroy() {}
