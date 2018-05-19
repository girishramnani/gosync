package pkg

import (
	"fmt"
	"os"

	"log"
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
		fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
		return err
	}
	if info.IsDir() {
		return nil
	}

	fmt.Printf("[INFO] uploading file: %q\n", path)
	err2 := wlk.uploader.Upload(path)
	if err2 != nil {
		log.Printf("[ERROR] error uploading the file: %q, err -  %s", path, err2.Error())
		return nil
	}

	err2 = os.Remove(path)
	if err2 != nil {
		log.Printf("[ERROR] error uploading the file: %q, err - %s", path, err2.Error())
	}
	return nil
}
