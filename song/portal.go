package song

import (
	"github.com/faiface/beep"
	"github.com/timiskhakov/music/guitar"
)

func Portal(g *guitar.Guitar) beep.Streamer {
	return beep.Mix(portalLead(g), portalRhythm(g))
}

func portalLead(g *guitar.Guitar) beep.Streamer {
	return beep.Seq(
		// Verse 1
		// 1
		g.Pluck(guitar.Note{1, 3}, 0.25),
		g.Pluck(guitar.Note{1, 2}, 0.25),
		g.Pluck(guitar.Note{1, 0}, 0.25),
		g.Pluck(guitar.Note{1, 0}, 0.25),
		// 2
		g.Pluck(guitar.Note{1, 2}, 2+0.75),
		// 3
		//g1.Silence(0.75),
		g.Pluck(guitar.Note{2, 3}, 0.25),
		g.Pluck(guitar.Note{1, 3}, 0.25),
		g.Pluck(guitar.Note{1, 2}, 0.25),
		g.Pluck(guitar.Note{1, 0}, 0.25),
		g.Pluck(guitar.Note{1, 0}, 0.25+.25),
		// 4 (2)
		//g1.Silence(.25),
		g.Pluck(guitar.Note{1, 2}, .75),
		g.Pluck(guitar.Note{2, 3}, .5),
		g.Pluck(guitar.Note{1, 0}, .25),
		g.Pluck(guitar.Note{3, 3}, .25+2),
		// 5
		//g1.Silence(2),
		// 6
		g.Pluck(guitar.Note{1, 0}, .5),
		g.Pluck(guitar.Note{1, 2}, .25),
		g.Pluck(guitar.Note{1, 3}, .75),
		g.Pluck(guitar.Note{1, 0}, .25),
		g.Pluck(guitar.Note{2, 2}, .25+0.5),
		// 7
		//g1.Silence(0.5),
		g.Pluck(guitar.Note{2, 3}, .25),
		g.Pluck(guitar.Note{1, 0}, .75),
		g.Pluck(guitar.Note{3, 3}, .25),
		g.Pluck(guitar.Note{1, 0}, .25+.25),
		// 8
		//g1.Silence(0.25),
		g.Pluck(guitar.Note{1, 2}, 1.75+1),
		// 9
		//g1.Silence(1),
		g.Pluck(guitar.Note{1, 3}, .25),
		g.Pluck(guitar.Note{1, 2}, .25),
		g.Pluck(guitar.Note{1, 0}, .25),
		g.Pluck(guitar.Note{1, 0}, .25),
		// 10
		g.Pluck(guitar.Note{1, 2}, 2+.75),
		// 11
		//g1.Silence(0.75),
		g.Pluck(guitar.Note{2, 3}, 0.25),
		g.Pluck(guitar.Note{1, 3}, 0.25),
		g.Pluck(guitar.Note{1, 2}, 0.25),
		g.Pluck(guitar.Note{1, 0}, 0.25),
		g.Pluck(guitar.Note{1, 0}, 0.25+.25),
		// 12
		//g1.Silence(.25),
		g.Pluck(guitar.Note{1, 2}, .75),
		g.Pluck(guitar.Note{2, 3}, .5),
		g.Pluck(guitar.Note{1, 0}, .25),
		g.Pluck(guitar.Note{3, 3}, .25+1.75),
		// 13
		//g1.Silence(1.75),
		g.Pluck(guitar.Note{3, 3}, .25),
		// 14
		g.Pluck(guitar.Note{1, 0}, .5),
		g.Pluck(guitar.Note{1, 2}, .25),
		g.Pluck(guitar.Note{1, 3}, .75),
		g.Pluck(guitar.Note{1, 0}, .25),
		g.Pluck(guitar.Note{2, 2}, .25+0.5),
		// 15
		//g1.Silence(0.5),
		g.Pluck(guitar.Note{2, 3}, .25),
		g.Pluck(guitar.Note{1, 0}, .5),
		g.Pluck(guitar.Note{3, 3}, .25),
		g.Pluck(guitar.Note{2, 3}, .25),
		g.Pluck(guitar.Note{1, 0}, .25),
		// 16
		g.Pluck(guitar.Note{1, 1}, .25),
		g.Pluck(guitar.Note{1, 0}, .25),
		g.Pluck(guitar.Note{2, 3}, .25),
		g.Pluck(guitar.Note{2, 1}, .25+.5),
		//g1.Silence(0.5),
		g.Pluck(guitar.Note{3, 3}, .25),
		g.Pluck(guitar.Note{3, 4}, .25),

		// Chorus
		// 17
		g.Pluck(guitar.Note{2, 1}, .5),
		g.Pluck(guitar.Note{1, 1}, .5),
		g.Pluck(guitar.Note{1, 0}, .25),
		g.Pluck(guitar.Note{2, 3}, .25),
		g.Pluck(guitar.Note{2, 3}, .25),
		g.Pluck(guitar.Note{2, 1}, .25),
		// 18
		g.Pluck(guitar.Note{2, 3}, .25),
		g.Pluck(guitar.Note{2, 1}, .25),
		g.Pluck(guitar.Note{2, 1}, .25),
		g.Pluck(guitar.Note{3, 3}, .25),
		g.Pluck(guitar.Note{2, 1}, .5),
		g.Pluck(guitar.Note{3, 3}, .25),
		g.Pluck(guitar.Note{3, 4}, .25),
		// 19
		g.Pluck(guitar.Note{2, 1}, .5),
		g.Pluck(guitar.Note{1, 1}, .5),
		g.Pluck(guitar.Note{1, 3}, .25),
		g.Pluck(guitar.Note{1, 1}, .25),
		g.Pluck(guitar.Note{1, 0}, .25),
		g.Pluck(guitar.Note{1, 1}, .25),
		// 20
		g.Pluck(guitar.Note{1, 0}, .25),
		g.Pluck(guitar.Note{1, 1}, .25),
		g.Pluck(guitar.Note{1, 1}, .5),
		g.Pluck(guitar.Note{1, 1}, .5),
		g.Pluck(guitar.Note{1, 3}, .25),
		g.Pluck(guitar.Note{1, 5}, .25),
		// 21
		g.Pluck(guitar.Note{1, 6}, .25),
		g.Pluck(guitar.Note{1, 6}, .25),
		g.Pluck(guitar.Note{1, 5}, .5),
		g.Pluck(guitar.Note{1, 3}, .5),
		g.Pluck(guitar.Note{1, 1}, .25),
		g.Pluck(guitar.Note{1, 3}, .25),
		// 22
		g.Pluck(guitar.Note{1, 5}, .25),
		g.Pluck(guitar.Note{1, 5}, .25),
		g.Pluck(guitar.Note{1, 3}, .25),
		g.Pluck(guitar.Note{2, 6}, .25),
		g.Pluck(guitar.Note{2, 6}, .5),
		g.Pluck(guitar.Note{2, 3}, .25),
		g.Pluck(guitar.Note{2, 1}, .25),
		// 23
		g.Pluck(guitar.Note{2, 3}, .25),
		g.Pluck(guitar.Note{1, 1}, .25),
		g.Pluck(guitar.Note{1, 1}, .25),
		g.Pluck(guitar.Note{1, 0}, .5),
		g.Pluck(guitar.Note{1, 0}, .25),
		g.Pluck(guitar.Note{1, 2}, .25),
		g.Pluck(guitar.Note{1, 2}, .25+2),
		// 24
		//g1.Silence(2)
	)
}

func portalRhythm(g *guitar.Guitar) beep.Streamer {
	return beep.Seq(
		// Verse 1
		// 1
		g.Silence(1),
		// 2
		g.Chord([]guitar.Note{{5, 0}, {4, 0}, {3, 0}, {4, 0}}, 1, 0.25),
		g.Chord([]guitar.Note{{5, 2}, {4, 0}, {3, 0}, {4, 0}}, 1, 0.25),
		// 3
		g.Chord([]guitar.Note{{5, 0}, {4, 0}, {3, 0}, {4, 0}}, 1, 0.25),
		g.Chord([]guitar.Note{{5, 2}, {4, 0}, {3, 0}, {4, 0}}, 1, 0.25),
		// 4
		g.Chord([]guitar.Note{{5, 0}, {4, 0}, {3, 0}, {4, 0}}, 1, 0.25),
		g.Chord([]guitar.Note{{5, 2}, {4, 0}, {3, 0}, {4, 0}}, 1, 0.25),
		// 5
		g.Chord([]guitar.Note{{5, 0}, {4, 0}, {3, 0}, {4, 0}}, 1, 0.25),
		g.Chord([]guitar.Note{{5, 2}, {4, 0}, {3, 0}, {4, 0}}, 1, 0.25),
		// 6
		g.Chord([]guitar.Note{{5, 2}, {4, 2}, {3, 1}, {4, 2}}, 1, 0.25),
		g.Chord([]guitar.Note{{5, 2}, {4, 2}, {3, 1}, {4, 2}}, 1, 0.25),
		// 7
		g.Chord([]guitar.Note{{6, 5}, {5, 4}, {4, 2}, {5, 4}}, 1, .25),
		g.Chord([]guitar.Note{{6, 5}, {5, 4}, {4, 2}, {5, 4}}, 1, .25),
		// 8
		g.Chord([]guitar.Note{{5, 0}, {4, 0}, {3, 0}, {4, 0}}, 1, 0.25),
		g.Chord([]guitar.Note{{5, 2}, {4, 0}, {3, 0}, {4, 0}}, 1, 0.25),
		// 9
		g.Chord([]guitar.Note{{5, 0}, {4, 0}, {3, 0}, {4, 0}}, 1, 0.25),
		g.Chord([]guitar.Note{{5, 2}, {4, 0}, {3, 0}, {4, 0}}, 1, 0.25),
		// 10
		g.Chord([]guitar.Note{{5, 0}, {4, 0}, {3, 0}, {4, 0}}, 1, 0.25),
		g.Chord([]guitar.Note{{5, 2}, {4, 0}, {3, 0}, {4, 0}}, 1, 0.25),
		// 11
		g.Chord([]guitar.Note{{5, 0}, {4, 0}, {3, 0}, {4, 0}}, 1, 0.25),
		g.Chord([]guitar.Note{{5, 2}, {4, 0}, {3, 0}, {4, 0}}, 1, 0.25),
		// 12
		g.Chord([]guitar.Note{{5, 0}, {4, 0}, {3, 0}, {4, 0}}, 1, 0.25),
		g.Chord([]guitar.Note{{5, 2}, {4, 0}, {3, 0}, {4, 0}}, 1, 0.25),
		// 13
		g.Chord([]guitar.Note{{5, 0}, {4, 0}, {3, 0}, {4, 0}}, 1, 0.25),
		g.Chord([]guitar.Note{{5, 2}, {4, 0}, {3, 0}, {4, 0}}, 1, 0.25),
		// 14
		g.Chord([]guitar.Note{{5, 2}, {4, 2}, {3, 1}, {4, 2}}, 1, 0.25),
		g.Chord([]guitar.Note{{5, 2}, {4, 2}, {3, 1}, {4, 2}}, 1, 0.25),
		// 15
		g.Chord([]guitar.Note{{6, 5}, {5, 4}, {4, 2}, {5, 4}}, 1, 0.25),
		g.Chord([]guitar.Note{{6, 5}, {5, 4}, {4, 2}, {5, 4}}, 1, 0.25),
		// 16
		g.Chord([]guitar.Note{{6, 6}, {5, 5}, {4, 3}, {3, 3}}, 1, 0.25),
		g.Silence(1),

		// Chorus
		// 17
		g.Pluck(guitar.Note{5, 3}, .5),
		g.Pluck(guitar.Note{5, 3}, .5),
		g.Pluck(guitar.Note{5, 3}, .5),
		g.Pluck(guitar.Note{5, 3}, .5),
		// 18
		g.Pluck(guitar.Note{5, 1}, .5),
		g.Pluck(guitar.Note{5, 1}, .5),
		g.Pluck(guitar.Note{5, 3}, .5),
		g.Pluck(guitar.Note{5, 3}, .5),
		// 19
		g.Pluck(guitar.Note{5, 3}, .5),
		g.Pluck(guitar.Note{5, 3}, .5),
		g.Pluck(guitar.Note{5, 3}, .5),
		g.Pluck(guitar.Note{5, 3}, .5),
		// 20
		g.Pluck(guitar.Note{5, 1}, .5),
		g.Pluck(guitar.Note{5, 1}, .5),
		g.Pluck(guitar.Note{5, 3}, .5),
		g.Pluck(guitar.Note{5, 3}, .5),
		// 21
		g.Pluck(guitar.Note{4, 0}, .5),
		g.Pluck(guitar.Note{4, 0}, .5),
		g.Pluck(guitar.Note{5, 3}, .5),
		g.Pluck(guitar.Note{5, 3}, .5),
		// 22
		g.Pluck(guitar.Note{5, 3}, .5),
		g.Pluck(guitar.Note{5, 3}, .5),
		g.Pluck(guitar.Note{4, 0}, .5),
		g.Pluck(guitar.Note{4, 0}, .5),
		// 23
		g.Pluck(guitar.Note{5, 1}, .5),
		g.Pluck(guitar.Note{5, 1}, .5),
		g.Pluck(guitar.Note{5, 0}, .5),
		g.Pluck(guitar.Note{5, 0}, .5),
		// 24
		g.Chord([]guitar.Note{{5, 0}, {4, 0}, {3, 0}, {4, 0}}, 1, 0.25),
		g.Chord([]guitar.Note{{5, 2}, {4, 0}, {3, 0}, {4, 0}}, 1, 0.25),
	)
}
