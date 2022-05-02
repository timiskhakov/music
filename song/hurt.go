package song

import (
	"github.com/faiface/beep"
	"github.com/timiskhakov/music/guitar"
)

func Hurt(g1, g2 *guitar.Guitar) beep.Streamer {
	return beep.Mix(hurtLead(g1), hurtRhythm(g2))
}

func hurtLead(g *guitar.Guitar) beep.Streamer {
	return beep.Seq(
		// Intro
		// 1
		g.Silence(0.5),
		g.Chord([]guitar.Note{{3, 2}, {2, 1}, {1, 0}}, 1.5, 0.01),
		// 2
		g.Silence(0.5),
		g.Pluck(guitar.Note{2, 1}, 1),
		g.Pluck(guitar.Note{1, 0}, 0.5),
		// 3
		g.Silence(0.5),
		g.Chord([]guitar.Note{{3, 2}, {2, 1}, {1, 0}}, 1.5, 0.01),
		// 4
		g.Silence(0.5),
		g.Pluck(guitar.Note{2, 1}, 1),
		g.Pluck(guitar.Note{1, 0}, 0.5),
		// 5
		g.Silence(0.5),
		g.Chord([]guitar.Note{{3, 2}, {2, 1}, {1, 0}}, 1, 0.01),
		g.Pluck(guitar.Note{1, 0}, 0.5),

		// Verse
		// 6
		g.Pluck(guitar.Note{1, 3}, 0.25),
		g.Pluck(guitar.Note{1, 0}, 0.5),
		g.Pluck(guitar.Note{2, 3}, 0.75),
		g.Pluck(guitar.Note{1, 0}, 0.25),
		g.Pluck(guitar.Note{2, 1}, 0.25),
		// 7
		g.Pluck(guitar.Note{1, 0}, 0.5),
		g.Pluck(guitar.Note{1, 0}, 1.25),
		g.Pluck(guitar.Note{3, 0}, 0.25),
		// 8
		g.Pluck(guitar.Note{2, 1}, 0.5),
		g.Pluck(guitar.Note{3, 2}, 0.25),
		g.Pluck(guitar.Note{2, 1}, 0.75),
		g.Pluck(guitar.Note{2, 3}, 0.5),
		// 9
		g.Pluck(guitar.Note{1, 0}, 0.5),
		g.Pluck(guitar.Note{1, 0}, 1.25),
		g.Pluck(guitar.Note{1, 0}, 0.25),
		// 10
		g.Pluck(guitar.Note{1, 3}, 0.25),
		g.Pluck(guitar.Note{1, 0}, 1),
		g.Pluck(guitar.Note{2, 3}, 0.25),
		g.Pluck(guitar.Note{2, 1}, 0.25),
		g.Pluck(guitar.Note{1, 0}, 0.25+0.5),
		// 11
		//g.Silence(0.5),
		g.Pluck(guitar.Note{1, 0}, 1.25),
		g.Pluck(guitar.Note{3, 0}, 0.25),
		// 12
		g.Pluck(guitar.Note{2, 1}, 0.5),
		g.Pluck(guitar.Note{3, 2}, 0.25),
		g.Pluck(guitar.Note{2, 3}, 0.75),
		g.Pluck(guitar.Note{1, 0}, 0.25),
		g.Pluck(guitar.Note{3, 0}, 0.25),
		// 13
		g.Pluck(guitar.Note{3, 2}, 0.5),
		g.Pluck(guitar.Note{1, 0}, 1.25),
		g.Pluck(guitar.Note{1, 0}, 0.25),
		// 14
		g.Pluck(guitar.Note{1, 3}, 0.25),
		g.Pluck(guitar.Note{1, 0}, 0.5),
		g.Pluck(guitar.Note{2, 3}, 0.75),
		g.Pluck(guitar.Note{1, 2}, 0.25),
		g.Pluck(guitar.Note{2, 1}, 0.25),
		// 15
		g.Pluck(guitar.Note{1, 0}, 0.5),
		g.Pluck(guitar.Note{1, 0}, 1.25),
		g.Pluck(guitar.Note{3, 0}, 0.25),
		// 16
		g.Pluck(guitar.Note{2, 1}, 0.5),
		g.Pluck(guitar.Note{3, 2}, 0.25),
		g.Pluck(guitar.Note{2, 1}, 0.75),
		g.Pluck(guitar.Note{1, 0}, 0.5),
		// 17
		g.Pluck(guitar.Note{1, 0}, 0.5),
		g.Pluck(guitar.Note{1, 0}, 1.25),
		g.Pluck(guitar.Note{1, 0}, 0.25),
		// 18
		g.Pluck(guitar.Note{1, 3}, 0.25),
		g.Pluck(guitar.Note{1, 0}, 1),
		g.Pluck(guitar.Note{2, 3}, 0.25),
		g.Pluck(guitar.Note{2, 1}, 0.25),
		g.Pluck(guitar.Note{1, 0}, 0.25+0.5),
		// 19
		//g.Silence(0.5)
		g.Pluck(guitar.Note{1, 0}, 1.25),
		g.Pluck(guitar.Note{3, 0}, 0.25),
		// 20
		g.Pluck(guitar.Note{2, 1}, 0.5),
		g.Pluck(guitar.Note{3, 2}, 0.25),
		g.Pluck(guitar.Note{2, 3}, 0.75),
		g.Pluck(guitar.Note{1, 2}, 0.25),
		g.Pluck(guitar.Note{3, 0}, 0.25),
		// 21
		g.Pluck(guitar.Note{3, 0}, 2),

		// Chorus
		// 22
		g.Pluck(guitar.Note{1, 3}, 0.25),
		g.Pluck(guitar.Note{1, 0}, 0.5),
		g.Pluck(guitar.Note{1, 0}, 0.5),
		g.Pluck(guitar.Note{2, 3}, 0.25),
		g.Pluck(guitar.Note{2, 1}, 0.25),
		g.Pluck(guitar.Note{2, 1}, 0.25+2),
		// 23
		//g.Silence(2),
		// 24
		g.Silence(0.5),
		g.Pluck(guitar.Note{2, 1}, 0.5),
		g.Pluck(guitar.Note{2, 3}, 0.25),
		g.Pluck(guitar.Note{2, 1}, 0.5),
		g.Pluck(guitar.Note{2, 3}, 0.25+2),
		// 25
		//g.Silence(2),
		// 26
		g.Pluck(guitar.Note{1, 3}, 0.25),
		g.Pluck(guitar.Note{1, 0}, 0.5),
		g.Pluck(guitar.Note{1, 0}, 0.5),
		g.Pluck(guitar.Note{2, 3}, 0.25),
		g.Pluck(guitar.Note{2, 1}, 0.25),
		g.Pluck(guitar.Note{2, 1}, 0.25+1),
		// 27
		//g.Silence(1),
		g.Pluck(guitar.Note{2, 1}, 0.5),
		g.Pluck(guitar.Note{3, 2}, 0.25),
		g.Pluck(guitar.Note{2, 1}, 0.25+1),
		// 28
		//g.Silence(1),
		g.Pluck(guitar.Note{2, 3}, 0.5),
		g.Pluck(guitar.Note{3, 0}, 0.25),
		g.Pluck(guitar.Note{3, 0}, 0.25+1),
		// 29
		//g.Silence(1),
		g.Silence(0.75),
		g.Pluck(guitar.Note{2, 1}, 0.25),
		// 30
		g.Pluck(guitar.Note{2, 3}, 0.5),
		g.Pluck(guitar.Note{2, 1}, 0.25),
		g.Pluck(guitar.Note{1, 0}, 0.5),
		g.Pluck(guitar.Note{2, 3}, 0.25),
		g.Pluck(guitar.Note{2, 1}, 0.25),
		g.Pluck(guitar.Note{2, 1}, 0.25+2),
		// 31
		//g.Silence(2),
		// 32
		g.Pluck(guitar.Note{2, 3}, 0.5),
		g.Pluck(guitar.Note{2, 3}, 0.5),
		g.Pluck(guitar.Note{2, 3}, 0.75),
		g.Pluck(guitar.Note{2, 3}, 0.25+2),
		// 33
		//g.Silence(2),
		// 34
		g.Pluck(guitar.Note{2, 3}, 0.5),
		g.Pluck(guitar.Note{2, 1}, 0.5),
		g.Pluck(guitar.Note{1, 0}, 0.25),
		g.Pluck(guitar.Note{2, 3}, 0.25),
		g.Pluck(guitar.Note{2, 1}, 0.25),
		g.Pluck(guitar.Note{2, 1}, 0.25+2),
		// 35
		//g.Silence(2),
		// 36
		g.Pluck(guitar.Note{2, 3}, 0.5),
		g.Pluck(guitar.Note{2, 1}, 0.25),
		g.Pluck(guitar.Note{1, 0}, 0.75),
		g.Pluck(guitar.Note{3, 0}, 0.25),
		g.Pluck(guitar.Note{3, 2}, 0.25+0.5),
		// 37
		//g.Silence(2)
	)
}

func hurtRhythm(g *guitar.Guitar) beep.Streamer {
	return beep.Seq(
		// Intro
		// 1
		g.Chord([]guitar.Note{{5, 0}, {4, 2}}, 2, 0.25),
		// 2
		g.Chord([]guitar.Note{{5, 3}, {4, 2}, {3, 0}}, 0.75, 0.25),
		g.Chord([]guitar.Note{{5, 3}, {4, 0}, {3, 2}}, 0.75, 0.25),
		g.Chord([]guitar.Note{{2, 3}, {3, 2}}, 0.5, 0.25),
		// 3
		g.Chord([]guitar.Note{{5, 0}, {4, 2}}, 2, 0.25),
		// 4
		g.Chord([]guitar.Note{{5, 3}, {4, 2}, {3, 0}}, 0.75, 0.25),
		g.Chord([]guitar.Note{{5, 3}, {4, 0}, {3, 2}}, 0.75, 0.25),
		g.Chord([]guitar.Note{{2, 3}, {3, 2}}, 0.5, 0.25),
		// 5
		g.Chord([]guitar.Note{{5, 0}, {4, 2}}, 2, 0.25),

		// Verse
		// 6
		g.Chord([]guitar.Note{{5, 3}, {3, 0}}, 1, 0.5),
		g.Chord([]guitar.Note{{4, 0}, {3, 2}, {2, 3}}, 1, 0.25),
		// 7
		g.Chord([]guitar.Note{{5, 0}, {4, 2}, {3, 2}}, 0.75, 0.25),
		g.Chord([]guitar.Note{{5, 0}, {4, 2}, {3, 2}}, 0.75, 0.25),
		g.Pluck(guitar.Note{5, 0}, 0.5),
		// 8
		g.Chord([]guitar.Note{{5, 3}, {4, 2}}, 1, 0.25),
		g.Chord([]guitar.Note{{4, 0}, {3, 2}, {2, 3}, {3, 2}}, 1, 0.25),
		// 9
		g.Chord([]guitar.Note{{5, 0}, {4, 2}, {3, 2}}, 0.75, 0.25),
		g.Chord([]guitar.Note{{5, 0}, {4, 2}, {3, 2}}, 0.75, 0.25),
		g.Pluck(guitar.Note{5, 0}, 0.5),
		// 10
		g.Pluck(guitar.Note{5, 3}, 0.5),
		g.Chord([]guitar.Note{{3, 0}, {2, 1}}, 0.25, 0),
		g.Pluck(guitar.Note{5, 3}, 0.25),
		g.Pluck(guitar.Note{4, 0}, 1),
		// 11
		g.Chord([]guitar.Note{{5, 0}, {4, 2}, {3, 2}}, 0.75, 0.25),
		g.Chord([]guitar.Note{{5, 0}, {4, 2}, {3, 2}}, 0.75, 0.25),
		g.Pluck(guitar.Note{5, 0}, 0.5),
		// 12
		g.Chord([]guitar.Note{{5, 3}, {4, 2}}, 1, 0.25),
		g.Chord([]guitar.Note{{4, 0}, {3, 2}, {2, 3}}, 1, 0.25),
		// 13
		g.Chord([]guitar.Note{{5, 0}, {4, 2}, {3, 2}}, 0.75, 0.25),
		g.Chord([]guitar.Note{{5, 0}, {4, 2}, {3, 2}}, 0.75, 0.25),
		g.Pluck(guitar.Note{5, 0}, 0.5),
		// 14
		g.Pluck(guitar.Note{5, 0}, 0.5),
		g.Pluck(guitar.Note{3, 0}, 0.5),
		g.Chord([]guitar.Note{{4, 0}, {3, 2}, {2, 3}}, 1, 0.25),
		// 15
		g.Chord([]guitar.Note{{5, 0}, {4, 2}, {3, 2}}, 0.75, 0.25),
		g.Chord([]guitar.Note{{5, 0}, {4, 2}, {3, 2}}, 0.75, 0.25),
		g.Pluck(guitar.Note{5, 0}, 0.5),
		// 16
		g.Chord([]guitar.Note{{5, 3}, {4, 2}}, 1, 0.25),
		g.Chord([]guitar.Note{{4, 0}, {3, 2}, {2, 3}, {3, 2}}, 1, 0.25),
		// 17
		g.Chord([]guitar.Note{{5, 0}, {4, 2}, {3, 2}}, 0.75, 0.25),
		g.Chord([]guitar.Note{{5, 0}, {4, 2}, {3, 2}}, 0.75, 0.25),
		g.Pluck(guitar.Note{5, 0}, 0.5),
		// 18
		g.Chord([]guitar.Note{{5, 3}, {2, 1}}, 0.75, 0.5),
		g.Chord([]guitar.Note{{5, 3}, {4, 0}}, 1.25, 0.25),
		// 19
		g.Chord([]guitar.Note{{5, 0}, {4, 2}, {3, 2}}, 0.75, 0.25),
		g.Chord([]guitar.Note{{5, 0}, {4, 2}, {3, 2}}, 0.75, 0.25),
		g.Pluck(guitar.Note{5, 0}, 0.5),
		// 20
		g.Chord([]guitar.Note{{5, 3}, {4, 2}}, 1, 0.25),
		g.Chord([]guitar.Note{{4, 0}, {3, 2}, {2, 3}}, 1, 0.25),
		// 21
		g.Chord([]guitar.Note{{6, 3}, {4, 0}, {3, 0}}, 0.75, 0.25),
		g.Chord([]guitar.Note{{5, 2}, {4, 0}, {3, 0}}, 0.75, 0.25),
		g.Chord([]guitar.Note{{5, 2}, {4, 0}}, 0.5, 0.25),

		// Chorus
		// 22
		g.Chord([]guitar.Note{{5, 0}, {4, 2}, {3, 2}}, 0.75, 0.25),
		g.Chord([]guitar.Note{{5, 0}, {4, 2}, {3, 2}}, 0.75, 0.25),
		g.Chord([]guitar.Note{{5, 0}, {4, 2}}, 0.5, 0.25),
		// 23
		g.Chord([]guitar.Note{{6, 1}, {5, 3}, {4, 3}}, 0.75, 0.25),
		g.Chord([]guitar.Note{{6, 1}, {5, 3}, {4, 3}}, 0.75, 0.25),
		g.Chord([]guitar.Note{{6, 1}, {5, 3}}, 0.5, 0.25),
		// 24
		g.Chord([]guitar.Note{{6, 3}, {4, 2}, {3, 0}}, 0.75, 0.25),
		g.Chord([]guitar.Note{{6, 3}, {4, 2}, {3, 0}}, 0.75, 0.25),
		g.Chord([]guitar.Note{{6, 3}, {4, 2}}, 0.5, 0.25),
		// 25
		g.Chord([]guitar.Note{{6, 3}, {4, 0}, {3, 0}}, 0.75, 0.25),
		g.Chord([]guitar.Note{{6, 3}, {4, 0}, {3, 0}}, 0.75, 0.25),
		g.Chord([]guitar.Note{{6, 3}, {4, 0}}, 0.5, 0.25),
		// 26
		g.Chord([]guitar.Note{{5, 0}, {4, 2}, {3, 2}}, 0.75, 0.25),
		g.Chord([]guitar.Note{{5, 0}, {4, 2}, {3, 2}}, 0.75, 0.25),
		g.Chord([]guitar.Note{{5, 0}, {4, 2}}, 0.5, 0.25),
		// 27
		g.Chord([]guitar.Note{{6, 1}, {5, 3}, {4, 3}}, 0.75, 0.25),
		g.Chord([]guitar.Note{{6, 1}, {5, 3}, {4, 3}}, 0.75, 0.25),
		g.Chord([]guitar.Note{{6, 1}, {5, 3}}, 0.5, 0.25),
		// 28
		g.Chord([]guitar.Note{{6, 3}, {4, 2}, {3, 0}}, 0.75, 0.25),
		g.Chord([]guitar.Note{{6, 3}, {4, 2}, {3, 0}}, 0.75, 0.25),
		g.Chord([]guitar.Note{{6, 3}, {4, 2}}, 0.5, 0.25),
		// 29
		g.Chord([]guitar.Note{{6, 3}, {4, 0}, {3, 0}}, 0.75, 0.25),
		g.Chord([]guitar.Note{{6, 3}, {4, 0}, {3, 0}}, 0.75, 0.25),
		g.Chord([]guitar.Note{{6, 3}, {4, 0}}, 0.5, 0.25),
		// 30
		g.Chord([]guitar.Note{{5, 0}, {4, 2}, {3, 2}}, 0.75, 0.25),
		g.Chord([]guitar.Note{{5, 0}, {4, 2}, {3, 2}}, 0.75, 0.25),
		g.Chord([]guitar.Note{{5, 0}, {4, 2}}, 0.5, 0.25),
		// 31
		g.Chord([]guitar.Note{{6, 1}, {5, 3}, {4, 3}}, 0.75, 0.25),
		g.Chord([]guitar.Note{{6, 1}, {5, 3}, {4, 3}}, 0.75, 0.25),
		g.Chord([]guitar.Note{{6, 1}, {5, 3}}, 0.5, 0.25),
		// 32
		g.Chord([]guitar.Note{{6, 3}, {4, 2}, {3, 0}}, 0.75, 0.25),
		g.Chord([]guitar.Note{{6, 3}, {4, 2}, {3, 0}}, 0.75, 0.25),
		g.Chord([]guitar.Note{{6, 3}, {4, 2}}, 0.5, 0.25),
		// 33
		g.Chord([]guitar.Note{{6, 3}, {5, 2}, {4, 0}}, 0.75, 0.25),
		g.Chord([]guitar.Note{{6, 3}, {5, 2}, {4, 0}}, 0.75, 0.25),
		g.Chord([]guitar.Note{{6, 3}, {4, 2}}, 0.5, 0.25),
		// 34
		g.Chord([]guitar.Note{{5, 0}, {4, 2}, {3, 2}}, 0.75, 0.25),
		g.Chord([]guitar.Note{{5, 0}, {4, 2}, {3, 2}}, 0.75, 0.25),
		g.Chord([]guitar.Note{{5, 0}, {4, 2}}, 0.5, 0.25),
		// 35
		g.Chord([]guitar.Note{{6, 1}, {5, 3}, {4, 3}}, 0.75, 0.25),
		g.Chord([]guitar.Note{{6, 1}, {5, 3}, {4, 3}}, 0.75, 0.25),
		g.Chord([]guitar.Note{{6, 1}, {5, 3}}, 0.5, 0.25),
		// 36
		g.Chord([]guitar.Note{{6, 3}, {4, 2}, {3, 0}}, 0.75, 0.25),
		g.Chord([]guitar.Note{{6, 3}, {4, 2}, {3, 0}}, 0.75, 0.25),
		g.Pluck(guitar.Note{6, 3}, 0.5),
		// 37
		g.Chord([]guitar.Note{{5, 0}, {4, 2}}, 2, 0.25),
	)
}
