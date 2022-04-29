package guitar

type Sound struct {
	totalSamples []float64
	processed    int
}

type Synthesizer interface {
	Synthesize(frequency float64, duration float64) []float64
}

func NewSound(synthesizer Synthesizer, note Note, duration float64) *Sound {
	return &Sound{synthesizer.Synthesize(noteFrequencies[note], duration), 0}
}

func (s *Sound) Stream(samples [][2]float64) (int, bool) {
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

func (s *Sound) Err() error {
	return nil
}
