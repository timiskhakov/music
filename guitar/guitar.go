package guitar

import "github.com/faiface/beep"

type guitar struct {
	sampleRate  float64
	synthesizer synthesizer
}

func NewGuitar(sampleRate int, synthesizer synthesizer) *guitar {
	return &guitar{float64(sampleRate), synthesizer}
}

func (g *guitar) Pluck(note Note, duration float64) beep.Streamer {
	return newSound(g.synthesizer, note, duration)
}

func (g *guitar) Chord(notes []Note, duration, delay float64) beep.Streamer {
	streamers := make([]beep.Streamer, len(notes))
	for i, note := range notes {
		silence := beep.Silence(int(g.sampleRate * delay * float64(i)))
		sound := newSound(g.synthesizer, note, duration-delay*float64(i))
		streamers[i] = beep.Seq(silence, sound)
	}

	return beep.Mix(streamers...)
}

func (g *guitar) Silence(duration float64) beep.Streamer {
	return beep.Silence(int(g.sampleRate * duration))
}
