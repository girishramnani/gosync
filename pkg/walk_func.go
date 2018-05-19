package pkg

import (
	"log"
	"os"
)

type Walker struct {
	uploader Uploader
}

func NewWalker(uploader Uploader) *Walker {
	return &Walker{
		uploader: uploader,
	}
}

func (wlk *Walker) Walk(path string, info os.FileInfo, err error) error {
	if err != nil {
		log.Printf("[INFO] Error while walking path %q: %s\n", path, err.Error())
		return err
	}
	if info.IsDir() {
		return nil
	}

	log.Printf("[INFO] uploading file: %q\n", path)
	err2 := wlk.uploader.Upload(path)
	if err2 != nil {
		// in most cases if upload fails we do not want to move forward and let the error get logged
		return err2
	}

	// delete the file after upload gets successful
	err2 = os.Remove(path)
	if err2 != nil {
		log.Printf("[ERROR] error uploading the file: %q, err - %s", path, err2.Error())
	}
	return nil
}
