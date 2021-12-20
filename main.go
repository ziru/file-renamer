package main

import (
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	pathname = kingpin.Arg("pathname", "Specifies the folder pathname to list files from").Required().String()
	from     = kingpin.Flag("from", "Regex to match the input filename").String()
	to       = kingpin.Flag("to", "Substitute the match with").String()
)

func main() {
	kingpin.Parse()
	renamer := NewFileRenamer(*pathname, *from, *to)
	renamer.Do()
}
