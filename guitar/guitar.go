package guitar

import (
	"fmt"
)

type Guitar struct {
	tunes     []float64
	processed int
}

func NewGuitar(sounds ...Sound) (*Guitar, error) {
	tunes := make([]float64, 0)
	for _, sound := range sounds {
		frequency, ok := notes[sound.Note]
		if !ok {
			return nil, fmt.Errorf("invalid note: %v", sound.Note)
		}

		tunes = append(tunes, extendedKarplusStrong(frequency, sound.Duration)...)
	}

	return &Guitar{tunes, 0}, nil
}

func (g *Guitar) Stream(samples [][2]float64) (int, bool) {
	if g.processed >= len(g.tunes) {
		return 0, false
	}

	if len(g.tunes)-g.processed < len(samples) {
		samples = samples[:len(g.tunes)-g.processed]
	}

	for i := range samples {
		samples[i][0] = g.tunes[g.processed+i]
		samples[i][1] = g.tunes[g.processed+i]
	}

	g.processed += len(samples)

	return len(samples), true
}

func (g *Guitar) Err() error {
	return nil
}
