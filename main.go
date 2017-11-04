package main

import (
	"log"
	"flag"
	"os"
	"io/ioutil"
)

var extractedFilePath string
var pubkFilePath string
var savePath string

func main() {
	flag.StringVar(&extractedFilePath, "extract", "", "Path to file to extract.")
	flag.StringVar(&pubkFilePath, "pubk", "", "Path to wc24pubk.mod")
	flag.StringVar(&savePath, "output", "", "Path to save decrypted file")
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

	if savePath == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}
	decrypted := Extract(fileToExtract, pubkFile)
	ioutil.WriteFile(savePath, decrypted, os.ModePerm)
	log.Print("Done! Saved to " + savePath)
}
