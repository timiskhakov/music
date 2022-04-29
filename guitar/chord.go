package guitar

import "github.com/faiface/beep"

func Chord(synthesizer Synthesizer, notes []Note, sampleRate int, duration, delay float64) beep.Streamer {
	streamers := make([]beep.Streamer, len(notes))
	for i, note := range notes {
		silence := beep.Silence(int(float64(sampleRate) * delay * float64(i)))
		sound := NewSound(synthesizer, note, duration-delay*float64(i))
		streamers[i] = beep.Seq(silence, sound)
	}

	return beep.Mix(streamers...)
}
