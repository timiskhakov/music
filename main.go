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
	kse1 := karplusstrong.NewExtended(sampleRate)
	g1 := guitar.NewGuitar(sampleRate, kse1)

	kse2 := karplusstrong.NewExtended(sampleRate)
	g2 := guitar.NewGuitar(sampleRate, kse2)

	f, _ := os.Create("out.wav")
	wav.Encode(f, beep.Mix(
		// Guitar 1
		beep.Seq(
			// 1
			g1.Pluck(guitar.Note{1, 3}, 0.25),
			g1.Pluck(guitar.Note{1, 2}, 0.25),
			g1.Pluck(guitar.Note{1, 0}, 0.25),
			g1.Pluck(guitar.Note{1, 0}, 0.25),
			// 2
			g1.Pluck(guitar.Note{1, 2}, 2),
			// 3
			g1.Silence(0.75),
			g1.Pluck(guitar.Note{2, 3}, 0.25),
			g1.Pluck(guitar.Note{1, 3}, 0.25),
			g1.Pluck(guitar.Note{1, 2}, 0.25),
			g1.Pluck(guitar.Note{1, 0}, 0.25),
			g1.Pluck(guitar.Note{1, 0}, 0.25),
			// 4 (2)
			g1.Silence(.25),
			g1.Pluck(guitar.Note{1, 2}, .75),
			g1.Pluck(guitar.Note{2, 3}, .5),
			g1.Pluck(guitar.Note{1, 0}, .25),
			g1.Pluck(guitar.Note{3, 3}, .25),
			// 5
			g1.Silence(2),
			// 6
			g1.Pluck(guitar.Note{1, 0}, .5),
			g1.Pluck(guitar.Note{1, 2}, .25),
			g1.Pluck(guitar.Note{1, 3}, .75),
			g1.Pluck(guitar.Note{1, 0}, .25),
			g1.Pluck(guitar.Note{2, 2}, .25),
			// 7
			g1.Silence(0.5),
			g1.Pluck(guitar.Note{2, 3}, .25),
			g1.Pluck(guitar.Note{1, 0}, .75),
			g1.Pluck(guitar.Note{3, 3}, .25),
			g1.Pluck(guitar.Note{1, 0}, .25),
			// 8
			g1.Silence(0.25),
			g1.Pluck(guitar.Note{1, 2}, 1.75),
			// 9
			g1.Silence(1),
			g1.Pluck(guitar.Note{1, 3}, .25),
			g1.Pluck(guitar.Note{1, 2}, .25),
			g1.Pluck(guitar.Note{1, 0}, .25),
			g1.Pluck(guitar.Note{1, 0}, .25),
			// 10
			g1.Pluck(guitar.Note{1, 2}, 2),
			// 11
			g1.Silence(0.75),
			g1.Pluck(guitar.Note{2, 3}, 0.25),
			g1.Pluck(guitar.Note{1, 3}, 0.25),
			g1.Pluck(guitar.Note{1, 2}, 0.25),
			g1.Pluck(guitar.Note{1, 0}, 0.25),
			g1.Pluck(guitar.Note{1, 0}, 0.25),
			// 12
			g1.Silence(.25),
			g1.Pluck(guitar.Note{1, 2}, .75),
			g1.Pluck(guitar.Note{2, 3}, .5),
			g1.Pluck(guitar.Note{1, 0}, .25),
			g1.Pluck(guitar.Note{3, 3}, .25),
			// 13
			g1.Silence(1.75),
			g1.Pluck(guitar.Note{3, 3}, .25),
			// 14
			g1.Pluck(guitar.Note{1, 0}, .5),
			g1.Pluck(guitar.Note{1, 2}, .25),
			g1.Pluck(guitar.Note{1, 3}, .75),
			g1.Pluck(guitar.Note{1, 0}, .25),
			g1.Pluck(guitar.Note{2, 2}, .25),
			// 15
			g1.Silence(0.5),
			g1.Pluck(guitar.Note{2, 3}, .25),
			g1.Pluck(guitar.Note{1, 0}, .5),
			g1.Pluck(guitar.Note{3, 3}, .25),
			g1.Pluck(guitar.Note{2, 3}, .25),
			g1.Pluck(guitar.Note{1, 0}, .25),
			// 16
			g1.Pluck(guitar.Note{1, 1}, 0.25),
			g1.Pluck(guitar.Note{1, 0}, 0.25),
			g1.Pluck(guitar.Note{2, 3}, 0.25),
			g1.Pluck(guitar.Note{2, 1}, 0.25),
			g1.Silence(0.5),
			g1.Pluck(guitar.Note{3, 3}, 0.25),
			g1.Pluck(guitar.Note{3, 4}, 0.25),
		),
		// Guitar 2
		beep.Seq(
			// 1
			g2.Silence(1),
			// 2
			g2.Chord([]guitar.Note{{5, 0}, {4, 0}, {3, 0}, {4, 0}}, 1, 0.25),
			g2.Chord([]guitar.Note{{5, 2}, {4, 0}, {3, 0}, {4, 0}}, 1, 0.25),
			// 3
			g2.Chord([]guitar.Note{{5, 0}, {4, 0}, {3, 0}, {4, 0}}, 1, 0.25),
			g2.Chord([]guitar.Note{{5, 2}, {4, 0}, {3, 0}, {4, 0}}, 1, 0.25),
			// 4
			g2.Chord([]guitar.Note{{5, 0}, {4, 0}, {3, 0}, {4, 0}}, 1, 0.25),
			g2.Chord([]guitar.Note{{5, 2}, {4, 0}, {3, 0}, {4, 0}}, 1, 0.25),
			// 5
			g2.Chord([]guitar.Note{{5, 0}, {4, 0}, {3, 0}, {4, 0}}, 1, 0.25),
			g2.Chord([]guitar.Note{{5, 2}, {4, 0}, {3, 0}, {4, 0}}, 1, 0.25),
			// 6
			g2.Chord([]guitar.Note{{5, 2}, {4, 2}, {3, 1}, {4, 2}}, 1, 0.25),
			g2.Chord([]guitar.Note{{5, 2}, {4, 2}, {3, 1}, {4, 2}}, 1, 0.25),
			// 7
			g2.Chord([]guitar.Note{{6, 5}, {5, 4}, {4, 2}, {5, 4}}, 1, .25),
			g2.Chord([]guitar.Note{{6, 5}, {5, 4}, {4, 2}, {5, 4}}, 1, .25),
			// 8
			g2.Chord([]guitar.Note{{5, 0}, {4, 0}, {3, 0}, {4, 0}}, 1, 0.25),
			g2.Chord([]guitar.Note{{5, 2}, {4, 0}, {3, 0}, {4, 0}}, 1, 0.25),
			// 9
			g2.Chord([]guitar.Note{{5, 0}, {4, 0}, {3, 0}, {4, 0}}, 1, 0.25),
			g2.Chord([]guitar.Note{{5, 2}, {4, 0}, {3, 0}, {4, 0}}, 1, 0.25),
			// 10
			g2.Chord([]guitar.Note{{5, 0}, {4, 0}, {3, 0}, {4, 0}}, 1, 0.25),
			g2.Chord([]guitar.Note{{5, 2}, {4, 0}, {3, 0}, {4, 0}}, 1, 0.25),
			// 11
			g2.Chord([]guitar.Note{{5, 0}, {4, 0}, {3, 0}, {4, 0}}, 1, 0.25),
			g2.Chord([]guitar.Note{{5, 2}, {4, 0}, {3, 0}, {4, 0}}, 1, 0.25),
			// 12
			g2.Chord([]guitar.Note{{5, 0}, {4, 0}, {3, 0}, {4, 0}}, 1, 0.25),
			g2.Chord([]guitar.Note{{5, 2}, {4, 0}, {3, 0}, {4, 0}}, 1, 0.25),
			// 13
			g2.Chord([]guitar.Note{{5, 0}, {4, 0}, {3, 0}, {4, 0}}, 1, 0.25),
			g2.Chord([]guitar.Note{{5, 2}, {4, 0}, {3, 0}, {4, 0}}, 1, 0.25),
			// 14
			g2.Chord([]guitar.Note{{5, 2}, {4, 2}, {3, 1}, {4, 2}}, 1, 0.25),
			g2.Chord([]guitar.Note{{5, 2}, {4, 2}, {3, 1}, {4, 2}}, 1, 0.25),
			// 15
			g2.Chord([]guitar.Note{{6, 5}, {5, 4}, {4, 2}, {5, 4}}, 1, 0.25),
			g2.Chord([]guitar.Note{{6, 5}, {5, 4}, {4, 2}, {5, 4}}, 1, 0.25),
			// 16
			g2.Chord([]guitar.Note{{6, 6}, {5, 5}, {4, 3}, {3, 3}}, 1, 0.25),
			g2.Silence(1),
		),
	), beep.Format{SampleRate: sampleRate, NumChannels: 2, Precision: 3})

	fmt.Println("Done")
}
