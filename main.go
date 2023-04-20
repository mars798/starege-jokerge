package main

import (
	"bytes"
	"image"
	"image/color"
	"log"
	"strconv"

	"math/rand"

	"github.com/mars798/starege-jokerge/assets/encodedimages"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var (
	windowTitle = "Starege-Jokerge 0:0"
	img         [3]*ebiten.Image
	cells       = [3][3]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}
	x, y        int
	jMove       = false
	win         = 0
	score       = [2]int{0, 0}
)

func init() {
	imgDecode, _, err := image.Decode(bytes.NewReader(encodedimages.ImgEmpty))
	if err != nil {
		log.Fatal(err)
	}
	img[0] = ebiten.NewImageFromImage(imgDecode)

	imgDecode, _, err = image.Decode(bytes.NewReader(encodedimages.ImgStarege))
	if err != nil {
		log.Fatal(err)
	}
	img[1] = ebiten.NewImageFromImage(imgDecode)

	imgDecode, _, err = image.Decode(bytes.NewReader(encodedimages.ImgJokerge))
	if err != nil {
		log.Fatal(err)
	}
	img[2] = ebiten.NewImageFromImage(imgDecode)
}

type Game struct{}

func (g *Game) Update() error {
	x, y := ebiten.CursorPosition()
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		x /= 100
		y /= 100
		if cells[y][x] == 0 {
			cells[y][x] = 1
		out:
			for _, v := range cells {
				for _, w := range v {
					if w == 0 {
						jMove = true
						break out
					}
				}
			}
		}
	}
	for jMove {
		jY := rand.Intn(3)
		jX := rand.Intn(3)
		if cells[jY][jX] == 0 {
			cells[jY][jX] = 2
			jMove = false
		}
	}

	//Проверка на победу Starege
	for i, v := range cells {
		scoreX := 0
		for _, w := range v {
			if w == 1 {
				scoreX++
			}
		}
		if scoreX == 3 {
			win = 1
		}
		scoreX = 0
		if cells[0][i] == 1 && cells[1][i] == 1 && cells[2][i] == 1 {
			win = 1
		}
	}
	if cells[0][0] == 1 && cells[1][1] == 1 && cells[2][2] == 1 {
		win = 1
	}
	if cells[0][2] == 1 && cells[1][1] == 1 && cells[2][0] == 1 {
		win = 1
	}

	//Проверка на победу Jokerge
	if win == 0 {
		for i, v := range cells {
			scoreX := 0
			for _, w := range v {
				if w == 2 {
					scoreX++
				}
			}
			if scoreX == 3 {
				win = 2
			}
			scoreX = 0
			if cells[0][i] == 2 && cells[1][i] == 2 && cells[2][i] == 2 {
				win = 2
			}
		}
		if cells[0][0] == 2 && cells[1][1] == 2 && cells[2][2] == 2 {
			win = 2
		}
		if cells[0][2] == 2 && cells[1][1] == 2 && cells[2][0] == 2 {
			win = 2
		}
	}

	// Проверка на ничью
	if win == 0 {
		sum := 0
		for _, v := range cells {
			for _, w := range v {
				if w != 0 {
					sum++
				}
			}
		}
		if sum == 9 {
			win = 3
		}
	}

	if win == 1 {
		score[0] += 1
		str := "Starege-Jokerge " + strconv.Itoa(score[0]) + ":" + strconv.Itoa(score[1])
		ebiten.SetWindowTitle(str)
	}
	if win == 2 {
		score[1] += 1
		str := "Starege-Jokerge " + strconv.Itoa(score[0]) + ":" + strconv.Itoa(score[1])
		ebiten.SetWindowTitle(str)
	}
	if win != 0 {
		for i, v := range cells {
			for j := range v {
				cells[i][j] = 0
			}
		}
		win = 0
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	w0, h0 := ebiten.WindowSize()
	w := float64(w0)
	h := float64(h0)
	ebitenutil.DrawRect(screen, 0, 0, w, h, color.White)
	ebitenutil.DrawLine(screen, 0, 100, 300, 100, color.Black)
	ebitenutil.DrawLine(screen, 0, 200, 300, 200, color.Black)
	ebitenutil.DrawLine(screen, 100, 0, 100, 300, color.Black)
	ebitenutil.DrawLine(screen, 200, 0, 200, 300, color.Black)
	for i, v := range cells {
		for j, w := range v {
			op := &ebiten.DrawImageOptions{}
			x = 100 * j
			y = 100 * i
			switch w {
			case 1:
				op.GeoM.Scale(0.333, 0.333)
				op.GeoM.Translate(float64(x), float64(y))
				screen.DrawImage(img[1], op)
			case 2:
				op.GeoM.Scale(0.249, 0.249)
				op.GeoM.Translate(float64(x), float64(y))
				screen.DrawImage(img[2], op)
			}
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 300, 300
}

func main() {
	ebiten.SetWindowSize(300, 300)
	ebiten.SetWindowTitle(windowTitle)
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
