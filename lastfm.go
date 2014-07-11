package main

import (
	"encoding/json"
	"os"
)

type LastfmTimestamp struct {
	Iso           string  `json:"iso"`
	UnixTimestamp float64 `json:"unixtimestamp"`
}

type LastfmArtist struct {
	Name string `json:"name"`
	Mbid string `json:"mbid"`
}

type LastfmThing struct {
	Name   string       `json:"name"`
	Mbid   string       `json:"mbid"`
	Artist LastfmArtist `json:"artist"`
}

type LastfmScrobble struct {
	Timestamp        LastfmTimestamp `json:"timestamp"`
	Track            LastfmThing     `json:"track"`
	UncorrectedTrack LastfmThing     `json:"uncorrectedTrack"`
	Album            LastfmThing     `json:"album"`
	Application      string          `json:"application"`
}

func ReadLastfmScrobblesJSON(path string) ([]LastfmScrobble, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	dec := json.NewDecoder(file)
	var scrobbles []LastfmScrobble
	err = dec.Decode(&scrobbles)
	if err != nil {
		return nil, err
	}
	return scrobbles, nil
}
