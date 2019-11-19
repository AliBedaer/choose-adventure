package cyoa

import (
	"encoding/json"
	"io"
)

// ReadJson : parse json file to Story

func ReadJson(r io.Reader) (Story, error) {

	d := json.NewDecoder(r)
	var story Story

	if err := d.Decode(&story); err != nil {
		return nil, err
	}

	return story,nil
}

// Story : story type
type Story map[string]Chapter


// Chapter : is struct
type Chapter struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []Option `json:"options"`
}

// Option : is struct
type Option struct {
	Text    string `json:"text"`
	Chapter string `json:"arc"`
}
