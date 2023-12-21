package main

import (
	"fmt"
	"os"

	"github.com/at-wat/ebml-go"
	"github.com/at-wat/ebml-go/webm"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <path_to_webm_file>")
		os.Exit(1)
	}

	filePath := os.Args[1]

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	var doc webm.Segment
	if err := ebml.Unmarshal(file, &doc); err != nil {
		fmt.Printf("Error unmarshalling: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("WebM Information:\n")
	fmt.Printf("Duration: %v\n", doc.Info.Duration)
	fmt.Printf("Date UTC: %v\n", doc.Info.DateUTC)
	fmt.Printf("Timecode Scale: %v\n", doc.Info.TimecodeScale)
	fmt.Printf("MuxingApp: %v\n", doc.Info.MuxingApp)
	fmt.Printf("WritingApp: %v\n", doc.Info.WritingApp)

	fmt.Printf("\nTracks:\n")
	for _, track := range doc.Tracks.TrackEntry {
		fmt.Printf("Track Number: %d\n", track.TrackNumber)
		fmt.Printf("Track Type: %d\n", track.TrackType)
		fmt.Printf("Track Name: %s\n", track.Name)
		fmt.Printf("CodecID: %s\n", track.CodecID)
		fmt.Println()
	}
}
