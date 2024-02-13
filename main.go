package main

import (
	"github.com/gopxl/beep"
	"github.com/gopxl/beep/speaker"
)

func main() {
	speaker.Init(beep.SampleRate(48000), 512)
	streamer, err := (&GeneralSound{SampleRate: 48000, Frequency: 27, Gain: -.9, Phase: 0}).GenerateSound()

	if err != nil {
		panic(err)
	}

	speaker.PlayAndWait(streamer)
	defer speaker.Close()
}
