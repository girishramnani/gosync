package gosync

import (
	"flag"
	"log"
	"os"
)

type Config struct {
	Directory   string
	Bucket      string
	PoolInteval int
}

// ParseConfigFromCli gets the config from cli and returns the config
func ParseConfigFromCli() *Config {
	dir := flag.String("dir", "", "directory to watch (default cwd)")
	interval := flag.Int("interval", 30, "The time interval between sync runs in seconds")
	bucket := flag.String("bucket", "", "destination bucket for the files")

	flag.Parse()

	if *dir == "" {
		directory, err := os.Getwd()
		if err != nil {
			// shouldn't continue if cannot get the CWD
			log.Fatal("Cannot get the current working directory: ", err.Error())
		}
		*dir = directory
	}

	if *bucket == "" {
		flag.Usage()
		os.Exit(1)
	}

	return &Config{
		Bucket:      *bucket,
		Directory:   *dir,
		PoolInteval: *interval,
	}
}
