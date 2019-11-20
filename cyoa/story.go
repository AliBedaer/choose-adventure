package cyoa

import (
	"encoding/json"
	"html/template"
	"io"
	"net/http"
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

// ParseTemplate : Parsing the package html template
func ParseTemplate(tpl string, story Story)  {
	tmpl := template.Must(template.ParseFiles(tpl))
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		data := story
		err := tmpl.Execute(writer,data)

		if err != nil{
			panic(err)
		}
	})

	err := http.ListenAndServe(":8000",nil)

	if err != nil{
		panic(err)
	}
}