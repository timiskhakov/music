package guitar

import "github.com/faiface/beep"

type Note struct {
	String int
	Fret   int
}

type Guitar struct {
	sampleRate  float64
	synthesizer synthesizer
}

func NewGuitar(sampleRate int, synthesizer synthesizer) *Guitar {
	return &Guitar{float64(sampleRate), synthesizer}
}

func (g *Guitar) Pluck(note Note, duration float64) beep.Streamer {
	return newSound(g.synthesizer, note, duration)
}

func (g *Guitar) Chord(notes []Note, duration, delay float64) beep.Streamer {
	streamers := make([]beep.Streamer, len(notes))
	for i, note := range notes {
		silence := beep.Silence(int(g.sampleRate * delay * float64(i)))
		sound := newSound(g.synthesizer, note, duration-delay*float64(i))
		streamers[i] = beep.Seq(silence, sound)
	}

	return beep.Mix(streamers...)
}

func (g *Guitar) Silence(duration float64) beep.Streamer {
	return beep.Silence(int(g.sampleRate * duration))
}
