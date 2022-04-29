package guitar

type sound struct {
	totalSamples []float64
	processed    int
}

type synthesizer interface {
	Synthesize(frequency float64, duration float64) []float64
}

func newSound(synth synthesizer, note Note, duration float64) *sound {
	return &sound{synth.Synthesize(noteFrequencies[note], duration), 0}
}

func (s *sound) Stream(samples [][2]float64) (int, bool) {
	if s.processed >= len(s.totalSamples) {
		return 0, false
	}

	if len(s.totalSamples)-s.processed < len(samples) {
		samples = samples[:len(s.totalSamples)-s.processed]
	}

	for i := range samples {
		samples[i][0] = s.totalSamples[s.processed+i]
		samples[i][1] = s.totalSamples[s.processed+i]
	}

	s.processed += len(samples)

	return len(samples), true
}

func (s *sound) Err() error {
	return nil
}
