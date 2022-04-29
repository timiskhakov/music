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

	g1, _ := guitar.NewGuitar(guitar.Sound{Note: guitar.Note{String: 1}, Duration: 4})
	g2, _ := guitar.NewGuitar(guitar.Sound{Note: guitar.Note{String: 6}, Duration: 1})

	mixer := beep.Mix(g1, g2)

	wav.Encode(f, mixer, beep.Format{SampleRate: guitar.SampleRate, NumChannels: 2, Precision: 3})

	fmt.Println("Done")
}
