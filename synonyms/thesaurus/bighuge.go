package thesaurus

import (
	"encoding/json"
	"errors"
	"net/http"
)

const apiUrl = "http://words.bighugelabs.com/api/2/"

type BigHuge struct {
	APIKey string
}

type synonyms struct {
	Noun *words `json:"noun"`
	Verb *words `json:"verb"`
}

type words struct {
	Syn []string `json:"syn"`
}

func (b *BigHuge) Synonyms(term string) ([]string, error) {
	var syns []string
	url := b.GetUrl(term)
	response, err := http.Get(url)
	if err != nil {
		errorMessage := "Bighuge: Failed when looking for synonyms for " + term + " " + err.Error()
		return syns, errors.New(errorMessage)
	}
	var data synonyms
	defer response.Body.Close()

	if err := json.NewDecoder(response.Body).Decode(&data); err != nil {
		return syns, err
	}

	if data.Noun != nil {
		syns = append(syns, data.Noun.Syn...)
	}

	if data.Verb != nil {
		syns = append(syns, data.Verb.Syn...)
	}
	return syns, nil

}

func (b *BigHuge) GetUrl(term string) string {
	return apiUrl + b.APIKey + "/" + term + "/json"
}
