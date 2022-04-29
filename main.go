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

	chord := guitar.Chord(kse, []guitar.Note{
		{String: 6},
		{String: 3},
		{String: 2},
		{String: 1},
	}, sampleRate, 4, 0.01)

	f, _ := os.Create("out.wav")
	wav.Encode(f, chord, beep.Format{SampleRate: sampleRate, NumChannels: 2, Precision: 3})

	fmt.Println("Done")
}
