package fetcher

import (
	"net/http"
	"fmt"
	"golang.org/x/text/transform"
	"io/ioutil"
	"golang.org/x/text/encoding"
	"bufio"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding/unicode"
	"log"
)

// Given a url string, return contents as a byte slice and an error
func Fetch(url string) ([]byte, error){
	resp, err := http.Get(url)

	if err != nil{
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Errorf("wrong status code: %d", resp.StatusCode)
	}

	bufferReader:= bufio.NewReader(resp.Body)
	e := determineEncoding(bufferReader)
	utf8Reader := transform.NewReader(bufferReader, e.NewDecoder())
	return ioutil.ReadAll(utf8Reader)
}

func determineEncoding(reader *bufio.Reader) encoding.Encoding{
	// get the first 1024 bytes
	bytes, err := reader.Peek(1024)
	if err != nil{
		log.Printf("fetcher error: %v", err)
		return unicode.UTF8 // default encoding
	}

	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}

