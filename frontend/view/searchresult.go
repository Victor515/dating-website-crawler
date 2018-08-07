package view

import (
	"html/template"
	"io"
	"crawler/frontend/model"
)

type SearchResultView struct {
	template *template.Template
}

func CreateSearchResult(filename string) SearchResultView{
	return SearchResultView{
		template: template.Must(template.ParseFiles(filename)),
	}
}

// render the page
func (s SearchResultView) Render(w io.Writer, data model.SearchResult) error{
	return s.template.Execute(w, data)
}
