package main

import (
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

// Cam takes pictures
type Cam struct {
	FileFolder string
	FilePrefix string
}

// Capture ...
func (c *Cam) Capture() {
	// Grab current UTC time
	now := time.Now()

	// Check if the folder exists
	if _, err := os.Stat(c.FileFolder); os.IsNotExist(err) {
		os.MkdirAll(c.FileFolder, 0777)
	}

	// Build the filename
	// TODO: Raw format
	fileName := filepath.Join(c.FileFolder, c.FilePrefix+"-"+now.Format(fileDateFormat)+".jpg")

	// Execute the command against the OS

	// TODO: raspiyuv
	cmd := exec.Command("raspistill", "-vf", "-hf", "-a", "1024", "-a", "8", "-a", "github.com/brentpabst/wxcam | %F %r", "-o", fileName)
	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	err := cmd.Run()
	// Check for errors
	if err != nil {
		log.Fatalf("Failed to execute capture command: %s", err)
	}

	log.Info("Took a picture: " + fileName)
}
