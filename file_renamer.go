package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	log "github.com/sirupsen/logrus"
)

type fileRenamer struct {
	pathname string
	from     string
	to       string
}

func NewFileRenamer(pathname, from, to string) *fileRenamer {
	return &fileRenamer{
		pathname: pathname,
		from:     from,
		to:       to,
	}
}

func (r *fileRenamer) Do() {
	files, err := ioutil.ReadDir(r.pathname)
	if err != nil {
		log.Fatalf("Failed to list files from: %s. Error: %v", r.pathname, err)
	}
	re, err := regexp.Compile(r.from)
	if err != nil {
		log.Fatalf("Failed to compile regex from: %s. Error: %v", r.from, err)
	}
	for _, f := range files {
		fromName := filepath.Base(f.Name())
		if strings.HasPrefix(fromName, ".") {
			// skip hidden files
			continue
		}
		toName := re.ReplaceAllString(fromName, r.to)
		if fromName == toName {
			// skip files with no change
			continue
		}
		log.Infof("[DRYRUN] %s => %s", fromName, toName)
	}

	fmt.Printf("\nAre you sure to proceed (y/N)? ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return
	}
	input := strings.TrimSpace(strings.ToLower(scanner.Text()))
	if input != "y" && input != "yes" {
		return
	}

	for _, f := range files {
		fromName := filepath.Base(f.Name())
		if strings.HasPrefix(fromName, ".") {
			// skip hidden files
			continue
		}
		toName := re.ReplaceAllString(fromName, r.to)
		if fromName == toName {
			// skip files with no change
			continue
		}

		err := os.Rename(filepath.Join(r.pathname, fromName), filepath.Join(r.pathname, toName))
		if err != nil {
			log.Fatalf("%s => %s FAIL --- %v", fromName, toName, err)
		}
		log.Infof("%s => %s SUCCESS", fromName, toName)
	}
}
