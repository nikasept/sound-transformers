package main

import (
	"github.com/gopxl/beep"
	"github.com/gopxl/beep/effects"
	"github.com/gopxl/beep/generators"
)

type GeneralSound struct {
	SampleRate int
	Frequency  float64
	Gain       float64
	Phase      float64 // dunno what the fuck this is
}

func (sound *GeneralSound) GenerateSound() (beep.Streamer, error) {
	streamer, err := generators.SineTone(beep.SampleRate(sound.SampleRate), sound.Frequency)
	effect := effects.Gain{Streamer: streamer, Gain: sound.Gain}

	return &effect, err
}
