package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"log"
)

func main() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		log.Fatalf("Erro ao inicializar SDL: %s", err)
	}
	defer sdl.Quit() window, err := sdl.CreateWindow("Pong em Go com SDL2", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, 800, 600, sdl.WINDOW_SHOWN|sdl.WINDOW_RESIZABLE) if err != nil {
		log.Fatalf("Erro ao criar janela: %s", err)
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED|sdl.RENDERER_PRESENTVSYNC)
	if err != nil {
		log.Fatalf("Erro ao criar renderer: %s", err)
	}
	defer renderer.Destroy()

	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch e := event.(type) {
			case *sdl.QuitEvent:
				return
			case *sdl.WindowEvent:
				if e.Event == sdl.WINDOWEVENT_RESIZED {
					renderer.SetViewport(&sdl.Rect{X: 0, Y: 0, W: e.Data1, H: e.Data2})
				}
			}
		}

		renderer.SetDrawColor(255, 0, 0, 255)
		renderer.Clear()
		renderer.Present()
	}
}
