package guitar

import (
	"fmt"
	"math/rand"
)

const SampleRate = 44100

var notes = map[Note]float64{
	Note{1, 0}: 329.63,
	Note{2, 0}: 246.94,
	Note{3, 0}: 196.00,
	Note{4, 0}: 146.83,
	Note{5, 0}: 110.00,
	Note{5, 0}: 82.41,
}

type Guitar struct {
	sound     []float64
	processed int
}

func NewGuitar(note Note, duration int) (*Guitar, error) {
	f, ok := notes[note]
	if !ok {
		return nil, fmt.Errorf("invalid note: %v", note)
	}

	return &Guitar{pluck(f, duration), 0}, nil
}

func (g *Guitar) Stream(samples [][2]float64) (int, bool) {
	if g.processed >= len(g.sound) {
		return 0, false
	}

	if len(g.sound)-g.processed < len(samples) {
		samples = samples[:len(g.sound)-g.processed]
	}

	for i := range samples {
		samples[i][0] = g.sound[g.processed+i]
		samples[i][1] = g.sound[g.processed+i]
	}

	g.processed += len(samples)

	return len(samples), true
}

func (g *Guitar) Err() error {
	return nil
}

func pluck(frequency float64, duration int) []float64 {
	noise := make([]float64, int(SampleRate/frequency))
	for i := range noise {
		noise[i] = rand.Float64()*2 - 1
	}

	samples := make([]float64, SampleRate*duration)
	for i := range noise {
		samples[i] = noise[i]
	}

	for i := len(noise) + 1; i < len(samples); i++ {
		samples[i] = (samples[i-len(noise)] + samples[i-len(noise)-1]) / 2
	}

	return samples
}
