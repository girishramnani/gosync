package main

import (
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"time"

	"github.com/girishramnani/gosync/gosync"
)

func main() {
	config := gosync.ParseConfigFromCli()
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)

	uploader := gosync.NewS3BucketUploader(config.Bucket)
	walker := gosync.NewWalker(uploader)

	ticker := time.NewTicker(time.Duration(config.PoolInteval) * time.Second)

	// first run should be immediate
	log.Println("[INFO] Syncing ", config.Directory)
	err := filepath.Walk(config.Directory, walker.Walk)
	if err != nil {
		log.Println("[ERROR] ", err)
	}
	for {
		select {
		case <-ticker.C:
			log.Println("[INFO] Syncing ", config.Directory)
			err = filepath.Walk(config.Directory, walker.Walk)
			if err != nil {
				log.Println("[ERROR] ", err)
			}
			// do stuff
		case <-signalChan:
			ticker.Stop()
			log.Println("[INFO] Exiting")
			os.Exit(0)
			break
		}
	}
}
