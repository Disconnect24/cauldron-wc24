package main

import (
	"log"
	"flag"
	"os"
	"io/ioutil"
)

var extractedFilePath string
var pubkFilePath string

func main() {
	flag.StringVar(&extractedFilePath, "extract", "", "Path to file to extract.")
	flag.StringVar(&pubkFilePath, "pubk", "", "Path to wc24pubk.mod")
	flag.Parse()
	log.Print("Reading " + extractedFilePath + "...")
	if extractedFilePath == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	fileToExtract, err := ioutil.ReadFile(extractedFilePath)
	if err != nil {
		panic(err)
	}

	pubkFile, err := ioutil.ReadFile(pubkFilePath)
	if err != nil {
		// We don't actually have to worry about this as it's an optional argument.
	}

	Extract(fileToExtract, pubkFile)
}
