package karplusstrong

import (
	"math"
	"math/rand"
)

const (
	p    = 0.9
	beta = 0.1
	s    = 0.5
	c    = 0.1
	l    = 0.1
)

type extended struct {
	sampleRate float64
}

func NewExtended(sampleRate int) *extended {
	return &extended{float64(sampleRate)}
}

func (e *extended) Synthesize(frequency float64, duration float64) []float64 {
	N := int(e.sampleRate / frequency)

	// Create noise
	noise := make([]float64, N)
	for i := range noise {
		noise[i] = rand.Float64()*2 - 1
	}

	// Pick-Direction Lowpass Filter
	buffer := make([]float64, N)
	buffer[0] = (1 - p) * noise[0]
	for i := 1; i < N; i++ {
		buffer[i] = (1-p)*noise[i] + p*buffer[i-1]
	}
	noise = buffer

	// Pick-Position Comb Filter
	pick := int(beta*float64(N) + 1/2)
	if pick == 0 {
		pick = N
	}
	buffer = make([]float64, N)
	for i := 0; i < N; i++ {
		if i-pick < 0 {
			buffer[i] = noise[i]
		} else {
			buffer[i] = noise[i] - noise[i-pick]
		}
	}
	noise = buffer

	// Create samples and add some noise in the beginning
	samples := make([]float64, int(e.sampleRate*duration))
	for i := 0; i < N; i++ {
		samples[i] = noise[i]
	}

	for i := N; i < len(samples); i++ {
		samples[i] = firstOrderStringTuningAllpassFilter(samples, i, N)
	}

	// Dynamic-level lowpass filter
	w_tilde := math.Pi * frequency / e.sampleRate
	buffer = make([]float64, len(samples))
	buffer[0] = w_tilde / (1 + w_tilde) * samples[0]
	for i := 1; i < len(samples); i++ {
		buffer[i] = w_tilde/(1+w_tilde)*(samples[i]+samples[i-1]) + (1-w_tilde)/(1+w_tilde)*buffer[i-1]
	}

	for i := 0; i < len(samples); i++ {
		samples[i] = (math.Pow(l, 4/3) * samples[i]) + (1-l)*buffer[i]
	}

	return samples
}

func delayLine(samples []float64, n, N int) float64 {
	if n-N < 0 {
		return 0
	}
	return samples[n-N]
}

func stringDamplingFilter(samples []float64, n, N int) float64 {
	return 0.996 * ((1-s)*delayLine(samples, n, N) + s*delayLine(samples, n-1, N))
}

func firstOrderStringTuningAllpassFilter(samples []float64, n, N int) float64 {
	return c*(stringDamplingFilter(samples, n, N)-samples[n-1]) + stringDamplingFilter(samples, n-1, N)
}