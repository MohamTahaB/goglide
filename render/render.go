package render

import (
	"image/color"
	"log"

	"github.com/MohamTahaB/goglide/model"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	deltaT           = 1.0 / 60
	w                = 640
	h                = 480
	number           = 100
	perceptionRadius = 70.0
	accMagnitude     = 100
	velMagnitude     = 150
)

type Render struct {
	m *model.Model
}

func InitiateRender() *Render {
	return &Render{
		m: model.NewModel(number, w, h, perceptionRadius),
	}
}

func (r *Render) Launch() {
	ebiten.SetWindowSize(w, h)
	ebiten.SetWindowTitle("GoGlide")
	if err := ebiten.RunGame(r); err != nil {
		log.Fatal(err)
	}
}

func (r *Render) Update() error {
	r.m.Update(deltaT, w, h)
	return nil
}

func (r *Render) Draw(screen *ebiten.Image) {
	for _, boid := range *r.m.GetBoids() {
		pos := boid.GetPosition()

		vector.DrawFilledCircle(screen, float32(pos.GetX()), float32(pos.GetY()), 5, color.White, false)
	}
}

func (r *Render) Layout(outsideWidth, outsideHeight int) (int, int) {
	return w, h
}
