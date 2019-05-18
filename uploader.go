package main

import (
	"os"
	"path/filepath"
	"sort"

	"io/ioutil"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/pkg/errors"
)

// Uploader uploads files to S3
type Uploader struct {
	s3Config   *aws.Config
	fileFolder string
	sem        chan struct{}
}

// NewUploader returns a new uploader
func NewUploader(s3Config *aws.Config, fileFolder string) *Uploader {
	return &Uploader{
		s3Config:   s3Config,
		fileFolder: fileFolder,
		sem:        make(chan struct{}, 1),
	}
}

// Upload uploads any files in the target folder to S3, deleting them after successful uploads
func (u *Uploader) Upload() {
	// Lock the semaphore to prevent concurrent upload attempts
	u.sem <- struct{}{}
	defer func() { <-u.sem }()

	// Build AWS Session Object
	s, err := session.NewSession(u.s3Config)

	// Check for errors connecting to AWS, stop if so
	if err != nil {
		log.Fatalf("Failed to connect to AWS: %s", err)
	}

	// Load list of files to upload
	files, err := ioutil.ReadDir(u.fileFolder)

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
	file, err := os.Open(filepath.Join(u.fileFolder, fileInfo.Name()))

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
	log.Infof("Uploaded file: %s", fileInfo.Name())

	// Delete files
	err = os.Remove(filepath.Join(u.fileFolder, file.Name()))
	if err != nil {
		return errors.Wrap(err, "removing file")
	}
	return nil
}
