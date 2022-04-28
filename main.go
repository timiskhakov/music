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

	g, _ := guitar.NewGuitar(guitar.Note{String: 6}, 4)

	wav.Encode(f, g, beep.Format{SampleRate: guitar.SampleRate, NumChannels: 2, Precision: 3})

	fmt.Println("Done")
}
