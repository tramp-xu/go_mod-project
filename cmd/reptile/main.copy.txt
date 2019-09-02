package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"net/http"
	"regexp"
)
// determineEncoding 拿到页面的 encoding
func determineEncoding(r io.Reader) encoding.Encoding  {
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		panic(err)
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}

// printCityList 获取所有城市的 a 链接

func printCityList (contents []byte)  {
	// ^ not
	re := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`)
	matches := re.FindAllSubmatch(contents, -1)
	fmt.Printf("%s", matches)
	for _, m := range matches {
		fmt.Printf("City: %s , Url: %s\n", m[2], m[1])
	}
	fmt.Printf("Matches found: %d\n", len(matches))

}

func main() {
	resp, err := http.Get("http://www.zhenai.com/zhenghun")
	if err != nil {
		panic(err)
	}
	// http.Get 方法后必须调用 close
	defer resp.Body.Close()
	// StatusCode: http 的返回状态码
	if resp.StatusCode != http.StatusOK {
		fmt.Println(http.StatusOK)
		return
	}
	if resp.StatusCode == http.StatusOK {
		e := determineEncoding(resp.Body)
		// 转换成使用 UTF-8 的 reader
		utf8Reader := transform.NewReader(resp.Body, e.NewDecoder())
		all, err := ioutil.ReadAll(utf8Reader)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s/n", all)
		printCityList(all)
	}
}
