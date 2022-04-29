package karplusstrong

import "math/rand"

type Basic struct {
	sampleRate float64
}

func NewBasic(sampleRate int) *Basic {
	return &Basic{float64(sampleRate)}
}

func (b *Basic) Synthesize(frequency float64, duration float64) []float64 {
	noise := make([]float64, int(b.sampleRate/frequency))
	for i := range noise {
		noise[i] = rand.Float64()*2 - 1
	}

	samples := make([]float64, int(b.sampleRate*duration))
	for i := range noise {
		samples[i] = noise[i]
	}

	for i := len(noise) + 1; i < len(samples); i++ {
		samples[i] = (samples[i-len(noise)] + samples[i-len(noise)-1]) / 2
	}

	return samples
}
