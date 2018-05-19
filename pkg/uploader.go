package pkg

// Uploader is the interface implemented by any syncing service that needs to be implemented
type Uploader interface {
	Upload(filePath string) error
}
