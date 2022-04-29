package main

import (
	"fmt"
	"github.com/faiface/beep"
	"github.com/faiface/beep/wav"
	"music/guitar"
	"music/karplusstrong"
	"music/song"
	"os"
)

const sampleRate = 44100

func main() {
	kse := karplusstrong.NewExtended(sampleRate)
	g := guitar.NewGuitar(sampleRate, kse)
	s := song.Portal(g)

	f, _ := os.Create("portal.wav")
	wav.Encode(f, s, beep.Format{SampleRate: sampleRate, NumChannels: 2, Precision: 3})

	fmt.Println("Done")
}
