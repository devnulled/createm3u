package main

import "io/ioutil"
import "log"
import "strings"
import "os"
import "bufio"
import "fmt"

func main() {
	const musicPath = "."
	const playlistName = "playlist.m3u"
	const shouldCleanPlaylists = true

	if shouldCleanPlaylists {
		cleanPlaylists(musicPath)
	}

	musicFiles := getMusicFilesForPath(musicPath)
	writePlaylist(musicPath, playlistName, musicFiles)
}

func getMusicFilesForPath(path string) []string {

	entries, err := ioutil.ReadDir(path)

	if err != nil {
		log.Fatal("Unable to open path to find music files", err)
	}

	validFiles := make([]string, len(entries))

	for _, entry := range entries {
		if isMusicFile(entry) {
			log.Println(entry.Name() + " is a music file!")
			validFiles = append(validFiles, entry.Name())
		} else {
			log.Println(entry.Name() + " is not a music file.")
		}
	}

	log.Println("No more files found!")

	return validFiles
}

func writePlaylist(path string, filename string, musicFiles []string) {
	if len(musicFiles) <= 0 {
		return
	}

	fullPath := []string{path, "/", filename}
	file, err := os.Create(strings.Join(fullPath, ""))

	if err != nil {
		log.Fatal("Unable to write playlist", err)
	}

	defer file.Close()

	w := bufio.NewWriter(file)

	for _, musicFile := range musicFiles {
		if musicFile != "" {
			fmt.Fprintln(w, musicFile)
		}
	}

	w.Flush()

	log.Println("Playlist created!")
}

func cleanPlaylists(path string) {
	entries, err := ioutil.ReadDir(path)

	if err != nil {
		log.Fatal("Unable to open path to find playlists to clean", err)
	}

	for _, entry := range entries {
		if isPlaylistFile(entry) {
			log.Println(entry.Name() + " is a playlist file!")
			os.Remove(path + entry.Name())
		} else {
			log.Println(entry.Name() + " is not a playlist file.")
		}
	}
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

func isPlaylistFile(f os.FileInfo) bool {

	if f.IsDir() {
		return false
	}

	if f.Size() < 10 {
		return false
	}

	return strings.HasSuffix(f.Name(), ".m3u")
}
