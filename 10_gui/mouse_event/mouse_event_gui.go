package main

import (
	"image"
	"image/color"
	"log"

	"golang.org/x/exp/shiny/driver"
	"golang.org/x/exp/shiny/screen"
	"golang.org/x/mobile/event/key"
	"golang.org/x/mobile/event/lifecycle"
	"golang.org/x/mobile/event/mouse"
	"golang.org/x/mobile/event/paint"
)

var darkGray = color.RGBA{0x30, 0x30, 0x30, 0xff}

// Point 構造体
type Point struct {
	X, Y int
}

func main() {
	var points []Point

	driver.Main(func(s screen.Screen) {
		w, err := s.NewWindow(&screen.NewWindowOptions{
			Title:  "dots",
			Width:  400,
			Height: 300,
		})
		if err != nil {
			log.Fatal(err)
		}
		defer w.Release()

		for {
			e := w.NextEvent()

			switch e := e.(type) {
			case lifecycle.Event:
				if e.To == lifecycle.StageDead {
					return
				}

			case key.Event:
				if e.Code == key.CodeEscape {
					return
				}

			case mouse.Event:
				if e.Button == mouse.ButtonLeft && e.Direction == mouse.DirPress {
					p := Point{int(e.X), int(e.Y)}
					points = append(points, p)
					// 描画させるためにpaintイベントを送る
					w.Send(paint.Event{})
				}

			case paint.Event:
				for _, pos := range points {
					r := image.Rect(pos.X-5, pos.Y-5, pos.X+5, pos.Y+5)
					w.Fill(r, darkGray, screen.Src)
				}
				w.Publish()

			case error:
				log.Print(e)
			}
		}
	})
}
