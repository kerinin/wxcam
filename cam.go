package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"time"
)

func cam() {
	log.Info("Taking a picture")

	// Check if the folder exists
	if _, err := os.Stat(*fileFolder); os.IsNotExist(err) {
		os.MkdirAll(*fileFolder, 0777)
	}

	// Grab current UTC time
	now := time.Now()

	// Build the filename
	fileName := *fileFolder + "/" + *filePrefix + "-" + now.Format(fileDateFormat) + ".jpg"

	// Execute the command against the OS

	cmd := exec.Command("raspistill", "-vf", "-hf", "-a", "1024", "-a", "8", "-a", "github.com/brentpabst/wxcam | %F %r", "-o", fileName)
	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	err := cmd.Run()
	// Check for errors
	if err != nil {
		log.Fatal(fmt.Sprint(err) + ": " + stderr.String())
	}

	log.Info("Took a picture: " + fileName)
}
