package data

import "io/ioutil"

type Article struct {
	Title string
	Body  []byte
}

func (a*Article) Save() error {
	filename := a.Title + ".txt"
	return ioutil.WriteFile(filename, a.Body, 0600)
}

func LoadArticle(title string) (*Article, error) {
	filename := title + ".txt"
	body, error := ioutil.ReadFile(filename)
	if error != nil {
		return nil, error
	}
	return &Article{Title:title, Body:body}, error
}
