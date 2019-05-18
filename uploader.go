package main

import (
	"os"
	"path/filepath"
	"sort"

	"io/ioutil"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/pkg/errors"
)

// Uploader uploads files to S3
type Uploader struct {
	S3Region   string
	AwsID      string
	AwsSecret  string
	FileFolder string
}

// Upload ...
func (u *Uploader) Upload() {
	// Build AWS Session Object
	s, err := session.NewSession(&aws.Config{
		Region:      aws.String(u.S3Region),
		Credentials: credentials.NewStaticCredentials(u.AwsID, u.AwsSecret, ""),
	})

	// Check for errors connecting to AWS, stop if so
	if err != nil {
		log.Fatalf("Failed to connect to AWS: %s", err)
	}

	// Load list of files to upload
	files, err := ioutil.ReadDir(u.FileFolder)

	// Check for error getting files
	if err != nil {
		log.Fatalf("Failed to list files to upload: %s", err)
	}

	// Check for no files
	if len(files) == 0 {
		log.Info("No files to upload")
	}
	log.Infof("Found %d files", len(files))

	// Sort the List of Files
	sort.Slice(files, func(i, j int) bool { return files[i].ModTime().Before(files[j].ModTime()) })

	// Build the uploader Object
	uploader := s3manager.NewUploader(s)

	// Loop through files and upload them
	for _, f := range files {
		err = u.uploadFile(uploader, f)
		if err != nil {
			log.Errorf("Failed to upload file: %s", err)
			break
		}
	}
}

func (u *Uploader) uploadFile(uploader *s3manager.Uploader, fileInfo os.FileInfo) error {
	// Open the file
	file, err := os.Open(filepath.Join(u.FileFolder, fileInfo.Name()))

	// Check for read errors
	if err != nil {
		return errors.Wrap(err, "reading file")
	}
	defer file.Close()

	// Build upload params
	params := &s3manager.UploadInput{
		Bucket: aws.String(*s3Bucket),
		Key:    aws.String(file.Name()),
		Body:   file,
	}

	// Upload
	_, err = uploader.Upload(params)
	if err != nil {
		return errors.Wrap(err, "uploading file")
	}
	log.Info("Uploaded file: %s", fileInfo.Name())

	// Delete files
	err = os.Remove(filepath.Join(u.FileFolder, file.Name()))
	if err != nil {
		return errors.Wrap(err, "removing file")
	}
	return nil
}
