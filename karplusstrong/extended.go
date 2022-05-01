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
)

type extended struct {
	sampleRate float64
	level      float64
}

func NewExtended(sampleRate int, level float64) *extended {
	return &extended{float64(sampleRate), level}
}

func (e *extended) Synthesize(frequency float64, duration float64) []float64 {
	// Create initial noise
	noise := make([]float64, int(e.sampleRate/frequency))
	for i := range noise {
		noise[i] = rand.Float64()*2 - 1
	}

	// Apply noise filters
	pickDirectionLowpass(noise)
	pickPositionComb(noise)

	// Create samples and add some noise in the beginning
	samples := make([]float64, int(e.sampleRate*duration))
	for i := 0; i < len(noise); i++ {
		samples[i] = noise[i]
	}

	// Apply single sample filters
	for i := len(noise); i < len(samples); i++ {
		samples[i] = firstOrderStringTuningAllpass(samples, i, len(noise))
	}

	// Apply all samples filters
	dynamicLevelLowpass(samples, math.Pi*frequency/e.sampleRate, e.level)

	return samples
}

func pickDirectionLowpass(noise []float64) {
	buffer := make([]float64, len(noise))
	buffer[0] = (1 - p) * noise[0]
	for i := 1; i < len(noise); i++ {
		buffer[i] = (1-p)*noise[i] + p*buffer[i-1]
	}
	noise = buffer
}

func pickPositionComb(noise []float64) {
	pick := int(beta*float64(len(noise)) + 1/2)
	if pick == 0 {
		pick = len(noise)
	}
	buffer := make([]float64, len(noise))
	for i := 0; i < len(noise); i++ {
		if i-pick < 0 {
			buffer[i] = noise[i]
		} else {
			buffer[i] = noise[i] - noise[i-pick]
		}
	}
	noise = buffer
}

func delayLine(samples []float64, n, N int) float64 {
	if n-N < 0 {
		return 0
	}
	return samples[n-N]
}

func oneZeroStringDamping(samples []float64, n, N int) float64 {
	return 0.996 * ((1-s)*delayLine(samples, n, N) + s*delayLine(samples, n-1, N))
}

func firstOrderStringTuningAllpass(samples []float64, n, N int) float64 {
	return c*(oneZeroStringDamping(samples, n, N)-samples[n-1]) + oneZeroStringDamping(samples, n-1, N)
}

func dynamicLevelLowpass(samples []float64, w float64, l float64) {
	buffer := make([]float64, len(samples))
	buffer[0] = w / (1 + w) * samples[0]
	for i := 1; i < len(samples); i++ {
		buffer[i] = w/(1+w)*(samples[i]+samples[i-1]) + (1-w)/(1+w)*buffer[i-1]
	}

	for i := 0; i < len(samples); i++ {
		samples[i] = (math.Pow(l, 4/3) * samples[i]) + (1-l)*buffer[i]
	}
}
