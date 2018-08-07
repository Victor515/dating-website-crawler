package controller

import (
	"crawler/frontend/view"
	"gopkg.in/olivere/elastic.v5"
	"net/http"
	"strings"
	"strconv"
	"crawler/frontend/model"
	"context"
	"reflect"
	"crawler/engine"
	"regexp"
)

type SearchResultHandler struct {
	view view.SearchResultView
	client *elastic.Client
}

func CreateSearchResultHandler(template string) SearchResultHandler{
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil{
		panic(err)
	}

	return SearchResultHandler{
		view: view.CreateSearchResult(template),
		client: client,
	}
}

func (h SearchResultHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	queryString := strings.TrimSpace(req.FormValue("q"))

	from, err := strconv.Atoi(req.FormValue("from"))
	if err != nil{
		from = 0
	}

	//fmt.Fprintf(resp, "q=%s, from=%d", queryString, from)
	page, err := h.getSearchResult(queryString, from)
	if err != nil{
		http.Error(resp, err.Error(), http.StatusBadRequest)
	}

	err = h.view.Render(resp, page)

	if err != nil{
		http.Error(resp, err.Error(), http.StatusBadRequest)
	}
}

func (h SearchResultHandler) getSearchResult(queryString string, from int) (model.SearchResult, error){
	var result model.SearchResult
	result.Query = queryString
	resp, err := h.client.
		Search("dating_profile").
		Query(elastic.NewQueryStringQuery(rewriteQueryString(queryString))).
		From(from).
		Do(context.Background())

	if err != nil{
		return result, err
	}

	result.Hits = resp.TotalHits()
	result.Start = from
	result.Items = resp.Each(
		reflect.TypeOf(engine.Item{}),
	)
	result.PrevFrom = result.Start - len(result.Items)
	result.NextFrom = result.Start + len(result.Items)

	return result, nil
}

func rewriteQueryString(q string) string{
	re := regexp.MustCompile(`([A-Z][a-z]*):`)
	return  re.ReplaceAllString(q, "Payload.$1:")
}


