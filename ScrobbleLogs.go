package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	session := os.Args[1]
	LastLines := strings.Split(os.Args[2], "\n")
	
	for index, each := range LastLines {
		var _ = index
		LastString := strings.Split(each, "]: ")[1]
		LastArray := strings.Split(LastString, "|")
		artist := LastArray[1]
		track := LastArray[2]
		timestamp := LastArray[4]
	
		fmt.Println(artist + " " + track + " " + timestamp)
		scrobble(session, artist, track, timestamp)
    }
}

func scrobble(session, artist, track, timestamp string) {
	cmd := exec.Command("/usr/bin/perl", "lfmCMD.pl", "method=track.scrobble", "artist=" + artist, "track=" + track, "timestamp=" + timestamp, "sk=" + session)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	var _ = err
	fmt.Printf("%s", out.String())
}
