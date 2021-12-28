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

	re                      *regexp.Regexp
	performArabicConversion bool
}

func NewFileRenamer(pathname, from, to string) *fileRenamer {
	return &fileRenamer{
		pathname: pathname,
		from:     from,
		to:       to,
	}
}

func (r *fileRenamer) Do() {
	if strings.Contains(r.to, "<arabic>") {
		r.performArabicConversion = true
	}

	re, err := regexp.Compile(r.from)
	if err != nil {
		log.Fatalf("Failed to compile regex from: %s. Error: %v", r.from, err)
	}
	r.re = re

	files, err := ioutil.ReadDir(r.pathname)
	if err != nil {
		log.Fatalf("Failed to list files from: %s. Error: %v", r.pathname, err)
	}
	for _, f := range files {
		fromName := filepath.Base(f.Name())
		if strings.HasPrefix(fromName, ".") {
			// skip hidden files
			continue
		}
		toName := r.convertFilename(fromName)
		if toName == "" || fromName == toName {
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
		toName := r.convertFilename(fromName)
		if toName == "" || fromName == toName {
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

func (r *fileRenamer) convertFilename(fromName string) string {
	to := r.to
	if r.performArabicConversion {
		// chinese number to arabic number
		groups := r.re.FindStringSubmatch(fromName)
		if len(groups) < 1 {
			return ""
		}
		to = strings.ReplaceAll(to, "<arabic>", ConvertChineseNumberToArabicNumber(groups[1]))
	}
	return r.re.ReplaceAllString(fromName, to)
}
