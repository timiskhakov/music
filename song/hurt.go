package song

import (
	"github.com/faiface/beep"
	"github.com/timiskhakov/music/guitar"
)

func Hurt(g *guitar.Guitar) beep.Streamer {
	return beep.Mix(hurtLead(g), hurtRhythm(g))
}

func hurtLead(g *guitar.Guitar) beep.Streamer {
	return beep.Seq(
		// Intro
		// 1
		g.Silence(0.5),
		g.Chord([]guitar.Note{{3, 2}, {2, 1}, {1, 0}}, 1.5, 0.5),
	)
}

func hurtRhythm(g *guitar.Guitar) beep.Streamer {
	return beep.Seq(
		// Intro
		// 1
		g.Chord([]guitar.Note{{5, 0}, {4, 2}}, 2, 0.5),
	)
}
