package model

import "encoding/json"

type BookIndex struct {
	ID         int     `json:"Id"`
	Name       string  `json:"Name"`
	Author     string  `json:"Author"`
	HomeImg    string  `json:"HomeImg"`
	Abstract   string  `json:"Abstract"`
	PayPrice   float64 `json:"PayPrice"`
	CreateDate string  `json:"CreateDate"`
	Detail     []BookDetail
}

type BookDetail struct {
	Name          string `json:"Name"`
	Title         string `json:"Title"`
	HomeImg       string `json:"HomeImg"`
	AudioAbstract string `json:"AudioAbstract"`
	FileSize      int    `json:"FileSize"`
	FileDuration  int    `json:"FileDuration"`
	CreateDate    string `json:"CreateDate"`
	FilePath      string `json:"FilePath"`
}

// FromJsonObj
func FromJsonObj(o interface{}) (BookIndex, error) {
	var bookIndex BookIndex

	s, err := json.Marshal(o)
	if err != nil {
		return bookIndex, err
	}

	err = json.Unmarshal(s, &bookIndex)

	return bookIndex, err

}
