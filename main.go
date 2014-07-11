package main

import (
	"fmt"
	"log"
	"path/filepath"
	"strings"
)

func Unix2win(unix int64) int64 {
	return (unix + 11644473600) * 10000000
}

func Win2unix(win int64) int64 {
	return (win / 10000000) - 11644473600
}

func main() {
	// [1] Read the tsv file containing firstPlayed info
	tracksFirstPlayed, err := ReadTSV("ps.tsv")
	if err != nil {
		log.Fatal(err)
	}

	// [2] Read the scrobbles to get the lastPlayed info
	first2last := map[int64]int64{}
	first2track := map[int64]string{}
	paths, err := filepath.Glob("scrobbles/*.json")
	if err != nil {
		log.Fatal(err)
	}
	for _, path := range paths {
		fmt.Printf("Parsing: %s\n", path)
		lastfmScrobblesJSON, err := ReadLastfmScrobblesJSON(path)
		if err != nil {
			log.Fatal(err)
		}
		for _, scrobble := range lastfmScrobblesJSON {
			trackid := strings.ToLower(scrobble.UncorrectedTrack.Artist.Name + "\t" + scrobble.Album.Name + "\t" + scrobble.UncorrectedTrack.Name)
			if firstPlayed, ok := tracksFirstPlayed[trackid]; ok {
				lastPlayed := int64(scrobble.Timestamp.UnixTimestamp)
				if lastPlayed > first2last[firstPlayed] {
					first2last[firstPlayed] = lastPlayed
				}
				first2track[firstPlayed] = trackid
			}
		}
	}

	// [3] Update the lastPlayed info in the xml file
	UpdateXML("ps.xml", "ps_updated.xml", first2last, first2track)
}
