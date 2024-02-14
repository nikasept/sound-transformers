package main

import (
	"os"

	"github.com/gopxl/beep"
	"github.com/gopxl/beep/wav"
)

func Save(stream beep.Streamer, sr beep.SampleRate, filename string) error {

	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}

	err = wav.Encode(f, stream, beep.Format{SampleRate: 48000, NumChannels: 2, Precision: 3})

	if err != nil {
		return err
	}

	return nil
}
