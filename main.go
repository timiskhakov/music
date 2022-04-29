package main

import (
	"fmt"
	"github.com/faiface/beep"
	"github.com/faiface/beep/wav"
	"music/guitar"
	"music/karplusstrong"
	"os"
)

const sampleRate = 44100

func main() {
	kse := karplusstrong.NewExtended(sampleRate)
	g := guitar.NewGuitar(sampleRate, kse)

	f, _ := os.Create("out.wav")
	wav.Encode(f, beep.Seq(
		// 1 (2)
		g.Chord([]guitar.Note{
			{String: 5, Fret: 4},
			{String: 4, Fret: 6},
			{String: 3, Fret: 7},
		}, 2, 0),

		// 2 (2)
		g.Pluck(guitar.Note{String: 6, Fret: 3}, 0.5),
		g.Silence(1.5),

		// 3 (2)
		g.Pluck(guitar.Note{String: 6, Fret: 3}, 0.5),
		g.Silence(1.5),

		// 4 (2)
		g.Pluck(guitar.Note{String: 6, Fret: 3}, 0.125),
		g.Silence(0.125),
		g.Silence(0.125),
		g.Pluck(guitar.Note{String: 6, Fret: 3}, 0.125),
		g.Silence(0.125),
		g.Silence(0.125),
		g.Pluck(guitar.Note{String: 6, Fret: 4}, 0.125),
		g.Silence(0.125),
		g.Pluck(guitar.Note{String: 6, Fret: 3}, 0.250),
		g.Pluck(guitar.Note{String: 6, Fret: 1}, 0.250),
		g.Silence(0.5),

		// 5 (2)
		g.Silence(2),

		// 6 (2)
		g.Pluck(guitar.Note{String: 6, Fret: 3}, 0.125),
		g.Silence(0.125),
		g.Silence(0.125),
		g.Pluck(guitar.Note{String: 6, Fret: 3}, 0.125),
		g.Silence(0.125),
		g.Silence(0.125),
		g.Pluck(guitar.Note{String: 6, Fret: 4}, 0.125),
		g.Silence(0.125),
		g.Pluck(guitar.Note{String: 6, Fret: 3}, 0.250),
		g.Pluck(guitar.Note{String: 6, Fret: 1}, 0.250),
		g.Silence(0.5),

		// 7 (2)
		g.Pluck(guitar.Note{String: 4, Fret: 3}, 0.125),
		g.Silence(0.125),
		g.Silence(0.125),
		g.Pluck(guitar.Note{String: 4, Fret: 3}, 0.125),
		g.Silence(0.125),
		g.Silence(0.125),
		g.Pluck(guitar.Note{String: 4, Fret: 5}, 0.125),
		g.Silence(0.125),
		g.Pluck(guitar.Note{String: 4, Fret: 3}, 0.250),
		g.Pluck(guitar.Note{String: 4, Fret: 1}, 0.250),
		g.Silence(0.5),

		// 8 (2)
		g.Chord([]guitar.Note{{String: 5, Fret: 5}, {String: 3, Fret: 7}}, 0.125, 0.001),
		g.Silence(0.125),
		g.Silence(0.125),
		g.Chord([]guitar.Note{{String: 5, Fret: 5}, {String: 3, Fret: 7}}, 0.125, 0.001),
		g.Silence(0.125),
		g.Silence(0.125),
		g.Chord([]guitar.Note{{String: 5, Fret: 6}, {String: 3, Fret: 8}}, 0.125, 0.001),
		g.Silence(0.125),
		g.Chord([]guitar.Note{{String: 5, Fret: 5}, {String: 3, Fret: 7}}, 0.250, 0.001),
		g.Chord([]guitar.Note{{String: 5, Fret: 3}, {String: 3, Fret: 5}}, 0.250, 0.001),
		g.Silence(0.5),
	), beep.Format{SampleRate: sampleRate, NumChannels: 2, Precision: 3})

	fmt.Println("Done")
}
