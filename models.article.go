package main

import(
	"errors"
)

type article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

// For this demo, we're storing the article list in memory
// In a real application, this list will most likely be fetched
// from a database or from static files
var articleList = []article{
	article{ID: 1, Title: "Covid Crisis", Content: "This is very terrifying Situation"},
	article{ID: 2, Title: "Vaccination shortage", Content: "There is a dire shortage of vccines."},
}

// Return a list of all the articles
func getAllArticles() []article {
	return articleList
}

func getArticleByID(id int) (*article, error) {
	for _, a := range articleList {
	  if a.ID == id {
		return &a, nil
	  }
	}
	return nil, errors.New("Article not found")
  }