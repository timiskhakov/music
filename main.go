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
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
}

func run() error {
	kse := karplusstrong.NewExtended(sampleRate)
	g := guitar.NewGuitar(sampleRate, kse)
	s := song.Portal(g)

	name := "portal.wav"
	f, err := os.Create(name)
	if err != nil {
		return err
	}

	fmt.Printf("Saving to %s...\n", name)
	if err = wav.Encode(f, s, beep.Format{SampleRate: sampleRate, NumChannels: 2, Precision: 3}); err != nil {
		return err
	}

	fmt.Println("Saved.")

	return nil
}
