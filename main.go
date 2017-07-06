package main

import (
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/alecthomas/kingpin"
	"github.com/coreos/go-systemd/daemon"
	"github.com/robfig/cron"
)

var (
	log              = logrus.New()
	cronPhotographer = kingpin.Flag("cron-photo", "CRON schedule for taking pictures.").Envar("GOPICAM_PHOTO_CRON").Default("0/24 * * * * *").String()
	cronUploader     = kingpin.Flag("cron-upload", "CRON schedule for uploading pictures.").Envar("GOPICAM_UPLOAD_CRON").Default("45 * * * * *").String()
	fileFolder       = kingpin.Flag("file-folder", "Path where pictures should be stored.").Envar("GOPICAM_PHOTO_FOLDER").Default("pictures").String()
	filePrefix       = kingpin.Flag("file-prefix", "Prefix for picture filenames.").Envar("GOPICAM_PHOTO_PREFIX").Default("wxcam").String()
	awsID            = kingpin.Flag("aws-id", "AWS security ID").Envar("AWS_ACCESS_KEY_ID").String()
	awsSecret        = kingpin.Flag("aws-secret", "AWS security secret.").Envar("AWS_SECRET_ACCESS_KEY").String()
	s3Region         = kingpin.Flag("s3-region", "AWS S3 region to upload pictures to.").Envar("AWS_REGION").Default("us-east-1").String()
	s3Bucket         = kingpin.Flag("s3-bucket", "AWS S3 bucket to upload pictures to.").Envar("GOPICAM_S3_BUCKET").Default("wxcam").String()
)

const fileDateFormat = time.RFC3339

func main() {
	log.Info("Application started")

	kingpin.Parse()

	c := cron.New()

	c.AddFunc(*cronPhotographer, cam)
	c.AddFunc(*cronUploader, uploader)
	c.AddFunc("@every 15s", func() { daemon.SdNotify(false, "WATCHDOG=1") })

	var wg sync.WaitGroup
	terminate(c, &wg)

	c.Start()
	wg.Add(1)

	daemon.SdNotify(false, "READY=1")

	wg.Wait()

	log.Info("Application stopped")
}

func terminate(c *cron.Cron, wg *sync.WaitGroup) {
	ch := make(chan os.Signal, 2)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
	go func() {
		s := <-ch
		log.Info("Terminating: ", s)
		c.Stop()
		wg.Done()
	}()
}
