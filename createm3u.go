package main

import "io/ioutil"
import "log"
import "strings"
import "os"

func main() {

	entries, err := ioutil.ReadDir(".")

	if err != nil {
		log.Fatal(err)
	}

	validFiles := make([]os.FileInfo, len(entries))

	for _, entry := range entries {
		if isMusicFile(entry.Name()) {
			log.Println(entry.Name() + " is a music file!")
			validFiles = append(validFiles, entry)
		} else {
			log.Println(entry.Name() + " is not a music file.")
		}
	}

	log.Println("No more files found!")
}

func isMusicFile(s string) bool {
	return strings.HasSuffix(s, ".mp3")
}
