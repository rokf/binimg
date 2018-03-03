package main

import (
	"flag"
	"github.com/gdamore/tcell"
	"github.com/nfnt/resize"
	"image"
	"image/color"
	"image/png"
	"os"
)

func main() {
	// flags
	scale := flag.Uint("scale", 1, "output image scale factor")
	filename := flag.String("file", "image.png", "name of the produced file")
	red := flag.Uint("r", 0, "red [0-255]")
	green := flag.Uint("g", 0, "green [0-255]")
	blue := flag.Uint("b", 0, "blue [0-255]")
	alpha := flag.Uint("a", 255, "alpha [0-255]")

	flag.Parse()

	// screen initialisation
	s, _ := tcell.NewScreen()
	s.Init()

	s.EnableMouse()

	s.Clear()

	// main loop
loop:
	for {
		ev := s.PollEvent()

		switch ev := ev.(type) {
		case *tcell.EventKey:
			switch ev.Key() {
			case tcell.KeyEsc:
				break loop
			case tcell.KeyCtrlC:
				s.Clear()
			case tcell.KeyCtrlS:
				w, h := s.Size()
				img := image.NewNRGBA(image.Rect(0, 0, w, h))
				for y := 0; y < h; y++ {
					for x := 0; x < w; x++ {
						if rn, _, _, _ := s.GetContent(x, y); rn != ' ' {
							img.Set(x, y, color.NRGBA{
								R: uint8(*red),
								G: uint8(*green),
								B: uint8(*blue),
								A: uint8(*alpha),
							})
						}
					}
				}
				f, _ := os.Create(*filename)

				if *scale != 1 {
					resized := resize.Resize(uint(uint(w)*(*scale)), 0, img, resize.NearestNeighbor)
					png.Encode(f, resized)
				} else {
					png.Encode(f, img)
				}

				f.Close()
			}

		case *tcell.EventMouse:
			x, y := ev.Position()
			if ev.Buttons()&tcell.Button1 != 0 {
				s.SetContent(x, y, tcell.RuneBlock, nil, tcell.StyleDefault)
			} else if ev.Buttons()&tcell.Button3 != 0 {
				s.SetContent(x, y, ' ', nil, tcell.StyleDefault)
			}
		case *tcell.EventResize:
			s.Sync()
		}

		s.Show()
	}

	s.Fini()
}
