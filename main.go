package main

import (
	"fmt"
	"time"

	"github.com/gopxl/beep"
	"github.com/gopxl/beep/speaker"
)

func main() {

	sr := beep.SampleRate(48000)

	speaker.Init(sr, 512)
	defer speaker.Close()

	streamer, err := (&GeneralSound{SampleRate: 48000, Frequency: 440, Gain: -.9, Phase: 0}).GenerateSound()
	if err != nil {
		panic(err)
	}
	limitedStreamer := beep.Take(sr.N(5*time.Second), streamer)

	first, second := beep.Dup(limitedStreamer)
	_, fourth := beep.Dup(first)

	//_ := sync.WaitGroup{}
	Save(fourth, sr, "out_fourth.wav")
	/*
		wg.Add(1)
		go func() {
			Save(third, sr, "out_third.wav")
			wg.Done()
		}()

		wg.Add(1)
		go func() {
			Save(first, sr, "out_first.wav")
			wg.Done()
		}()
	*/

	//wg.Wait()

	speaker.PlayAndWait(second)

	fmt.Println("The stream has ended")
}
