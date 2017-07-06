package main

import (
	"os"
	"sort"

	"io/ioutil"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func uploader() {
	log.Info("Connecting to AWS S3")

	// Build AWS Session Object
	s, err := session.NewSession(&aws.Config{
		Region:      aws.String(*s3Region),
		Credentials: credentials.NewStaticCredentials(*awsID, *awsSecret, ""),
	})

	// Check for errors connecting to AWS, stop if so
	if err != nil {
		log.Panic(err)
	}

	// Load list of files to upload
	files, err := ioutil.ReadDir(*fileFolder)

	// Check for error getting files
	if err != nil {
		log.Panic(err)
	}

	// Check for no files
	if len(files) == 0 {
		log.Info("No files to upload")
	}
	log.Info("Found files: ", len(files))

	// Sort the List of Files
	sort.Slice(files, func(i, j int) bool { return files[i].ModTime().Before(files[j].ModTime()) })

	// Build the uploader Object
	uploader := s3manager.NewUploader(s)

	// Loop through files and upload them
	for i := 0; i < len(files); i++ {
		log.Info("Uploading file: ", files[i].Name())

		// Open the file
		file, err := os.Open(*fileFolder + "/" + files[i].Name())

		// Check for read errors
		if err != nil {
			log.Error("Failed to read file. ", err)
			break
		}
		defer file.Close()

		// Build upload params
		params := &s3manager.UploadInput{
			Bucket: aws.String(*s3Bucket),
			Key:    aws.String("wxcam.jpg"),
			Body:   file,
		}

		// Upload
		_, err = uploader.Upload(params)
		if err != nil {
			log.Error("Failed to upload file. ", err)
			break
		} else {
			log.Info("File uploaded")
		}

		// Delete files
		err = os.Remove(*fileFolder + "/" + files[i].Name())
		if err != nil {
			log.Error("Failed to remove file. ", err)
			break
		}
	}
}
