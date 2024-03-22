package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"log"
)

func main() {
	var windowWidth, windowHeight int32 = 800, 600

	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		log.Fatalf("Erro ao inicializar SDL: %s", err)
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow("Pong em Go com SDL2", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, 800, 600, sdl.WINDOW_SHOWN)
	if err != nil {
		log.Fatalf("Erro ao criar janela: %s", err)
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED|sdl.RENDERER_PRESENTVSYNC)
	if err != nil {
		log.Fatalf("Erro ao criar renderer: %s", err)
	}
	defer renderer.Destroy()

	positionY := (windowHeight - 150) / 2

	leftRectangleX := windowWidth / 6
	rightRectangleX := (windowWidth * 5 / 6) - 20
	var ballVelX, ballVelY int32 = 5, 5

	leftRectangle := sdl.Rect{X: leftRectangleX, Y: positionY, W: 20, H: 150}
	rightRectangle := sdl.Rect{X: rightRectangleX, Y: positionY, W: 20, H: 150}
	ball := sdl.Rect{X: windowWidth / 2, Y: windowHeight / 2, W: 15, H: 15}

	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch e := event.(type) {
			case *sdl.QuitEvent:
				return
			case *sdl.KeyboardEvent:
				switch e.Keysym.Sym {
				case sdl.K_UP:
					leftRectangle.Y -= 10
					if leftRectangle.Y < 0 {
						leftRectangle.Y = 0
					}
				case sdl.K_DOWN:
					leftRectangle.Y += 10
					if leftRectangle.Y > windowHeight-leftRectangle.H {
						leftRectangle.Y = windowHeight - leftRectangle.H
					}
				}
			}
		}

		ball.X += ballVelX
		ball.Y += ballVelY

		if ball.Y <= 0 || ball.Y+ball.H >= windowHeight {
			ballVelY = -ballVelY
		}
		if ball.X <= 0 || ball.X+ball.W >= windowWidth {
			ballVelX = -ballVelX
		}

		renderer.SetDrawColor(0, 0, 0, 255)
		renderer.Clear()

		renderer.SetDrawColor(255, 255, 255, 255)

		renderer.FillRect(&leftRectangle)
		renderer.FillRect(&rightRectangle)
		renderer.FillRect(&ball)

		renderer.Present()
	}
}
