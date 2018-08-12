package engine

import "crawler/config"

type ParserFunc func(contents []byte, url string) ParserResult

type Parser interface {
	Parse(contents []byte, url string) ParserResult
	Serialize() (name string, args interface{})
}

type Request struct {
	Url string
	Parser Parser
}

type ParserResult struct {
	Requests []Request
	Items    []Item
}

type Item struct {
	Url 	string
	Type    string
	Id  	string
	Payload interface{}
}

type NilParser struct {
}

func (NilParser) Parse(contents []byte, url string) ParserResult {
	return ParserResult{}
}

func (NilParser) Serialize() (name string, args interface{}) {
	return config.NilParser, nil
}

type FuncParser struct {
	Parser ParserFunc
	Name string // func name
}

func (f *FuncParser) Parse(contents []byte, url string) ParserResult {
	return f.Parser(contents, url)
}

func (f *FuncParser) Serialize() (name string, args interface{}) {
	return f.Name, nil
}

// factory function
func NewFuncParser(p ParserFunc, name string) *FuncParser{
	return &FuncParser{
		Parser: p,
		Name:name,
	}
}

