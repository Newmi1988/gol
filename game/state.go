package game

import (
	"bytes"
	"math/rand"
)

type Life struct {
	currState, nextState *Field
}

func GameOfLife(width, height int, perc float32) *Life {
	currState := NewField(width, height)
	// set a percentage of the field to random trues
	var p int = int(float32(width) * float32(height) * perc)
	for i := 0; i < p; i++ {
		currState.Set(rand.Intn(width), rand.Intn(height), true)
	}
	return &Life{
		currState: currState,
		nextState: NewField(width, height),
	}
}

func (life *Life) Step() {
	for y := 0; y < life.currState.heigth; y++ {
		for x := 0; x < life.currState.width; x++ {
			life.nextState.Set(x, y, life.currState.Next(x, y))
		}
	}
	life.currState, life.nextState = life.nextState, life.currState
}

func (life *Life) String() string {
	var buf bytes.Buffer
	for y := 0; y < life.currState.heigth; y++ {
		for x := 0; x < life.currState.width; x++ {
			b := byte(' ')
			if life.currState.IsAlive(x, y) {
				b = '*'
			}
			buf.WriteByte(b)
		}
		buf.WriteByte('\n')
	}
	return buf.String()
}
