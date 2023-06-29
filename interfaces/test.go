package main

import (
	"fmt"
	"strings"
)

type VideoPlayer struct {
	File string
}

type AudioPlayer struct {
	File string
}

func (v VideoPlayer) Play() {
	fmt.Println("Playing video", v.File)
}

func (v VideoPlayer) Stop() {
	fmt.Println("Stop playing video", v.File)
}

func (a AudioPlayer) Play() {
	fmt.Println("Playing audio", a.File)
}

func (a AudioPlayer) Stop() {
	fmt.Println("Stop playing audio", a.File)
}

// Create the `DigitalPlayer` interface that implements the `Play()` and `Stop()` methods.
type DigitalPlayer interface {
	Play()
	Stop()
}

// nolint: gosimple // DO NOT remove this comment!
func main() {
	// DO NOT delete this code block; it takes as an input the name of the file to play or stop.
	var fileName string
	fmt.Scanln(&fileName)

	// Create the `player` variable of the `DigitalPlayer` type below:
	var player DigitalPlayer

	switch {
	case strings.HasSuffix(fileName, ".mp3"):
		// Make the `player` an `AudioPlayer` below:
		player = AudioPlayer{File: fileName}

	case strings.HasSuffix(fileName, ".mp4"):
		// Make the `player` a `VideoPlayer` below:
		player = VideoPlayer{File: fileName}

	default:
		fmt.Println("Unsupported file format")
		return
	}

	// Call the `Play()` and `Stop()` methods on the `player` below:
	player.Play()
	player.Stop()
}
