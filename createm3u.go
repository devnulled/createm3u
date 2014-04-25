package main

import "io/ioutil"
import "log"
import "strings"
import "os"

func main() {
	getMusicFilesForPath(".")
}

func isMusicFile(f os.FileInfo) bool {
	if f.IsDir() {
		return false
	}

	if f.Size() < 10 {
		return false
	}

	return strings.HasSuffix(f.Name(), ".mp3")
}

func getMusicFilesForPath(path string) []os.FileInfo {

	entries, err := ioutil.ReadDir(path)

	if err != nil {
		log.Fatal(err)
	}

	validFiles := make([]os.FileInfo, len(entries))

	for _, entry := range entries {
		if isMusicFile(entry) {
			log.Println(entry.Name() + " is a music file!")
			validFiles = append(validFiles, entry)
		} else {
			log.Println(entry.Name() + " is not a music file.")
		}
	}

	log.Println("No more files found!")

	return validFiles
}
