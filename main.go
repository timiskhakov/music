package main

import (
	"fmt"
	"github.com/faiface/beep"
	"github.com/faiface/beep/wav"
	"github.com/timiskhakov/music/guitar"
	"github.com/timiskhakov/music/karplusstrong"
	"github.com/timiskhakov/music/song"
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
	kse1 := karplusstrong.NewExtended(sampleRate, 0.1)
	g1 := guitar.NewGuitar(sampleRate, kse1)
	kse2 := karplusstrong.NewExtended(sampleRate, 0.05)
	g2 := guitar.NewGuitar(sampleRate, kse2)
	//s := song.Portal(g)
	s := song.Hurt(g1, g2)

	name := "hurt.wav"
	f, err := os.Create(name)
	if err != nil {
		return err
	}

	fmt.Printf("Saving to %s...\n", name)
	if err = wav.Encode(f, s, beep.Format{SampleRate: sampleRate, NumChannels: 2, Precision: 2}); err != nil {
		return err
	}

	fmt.Println("Saved.")

	return nil
}
