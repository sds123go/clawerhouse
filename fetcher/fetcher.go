package fetcher

import (
	"bufio"
	"io/ioutil"
	"net/http"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

// determineEncoding 确定编码方式
func determineEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024)
	if err != nil {
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
func Fetch(url string) ([]byte, error) {
	var data []byte
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		bodyReader := bufio.NewReader(resp.Body)
		e := determineEncoding(bodyReader)
		utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())
		data, err = ioutil.ReadAll(utf8Reader)
		if err != nil {
			return nil, err
		}
	}
	return data, err
}
