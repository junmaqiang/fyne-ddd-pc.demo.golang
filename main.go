package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("MyPC")

	// 设置尺寸
	// w.Resize(fyne.NewSize(500, 400))
	// 边距填充
	w.SetPadded(false)
	// 是否禁止改变窗口尺寸
	w.SetFixedSize(true)

	hello := widget.NewLabel("Hello Fyne DDD")
	w.SetContent(container.NewVBox(
		hello,
		widget.NewButton("Hi!", func() {
			hello.SetText("Welcome :)")
		}),
	))

	w.ShowAndRun()
}
