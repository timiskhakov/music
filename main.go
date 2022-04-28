package main

import (
	"fmt"
	"github.com/faiface/beep"
	"github.com/faiface/beep/wav"
	"music/guitar"
	"os"
)

func main() {
	file := "out.wav"
	f, _ := os.Create(file)

	g, _ := guitar.NewGuitar(guitar.Note{String: 1}, 1)

	wav.Encode(f, g, beep.Format{SampleRate: guitar.SampleRate, NumChannels: 2, Precision: 1})

	fmt.Println("Done")
}