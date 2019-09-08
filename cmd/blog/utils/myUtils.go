package utils

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
	"github.com/sourcegraph/syntaxhighlight"
	"html/template"
	"time"
)

func MD5(str string) string {
	md5str := fmt.Sprintf("%x", md5.Sum([]byte(str)))
	return md5str
}

func SwitchTimeStampToData(tm int64) string {
	date := time.Unix(tm, 0).Format("YYYY-MM-DD HH:mm:ss")
	return date
}

func SwitchMarkdownToHtml(content string) template.HTML {
	unsafe  := blackfriday.MarkdownCommon([]byte(content))
	markdown := bluemonday.UGCPolicy().SanitizeBytes(unsafe)

	doc, _ := goquery.NewDocumentFromReader(bytes.NewReader(markdown))
	/**
	对document进程查询，选择器和css的语法一样
	第一个参数：i是查询到的第几个元素
	第二个参数：selection就是查询到的元素
	*/
	fmt.Println(doc)
	doc.Find("code").Each(func(i int, selection *goquery.Selection) {
		light, _ := syntaxhighlight.AsHTML([]byte(selection.Text()))
		selection.SetHtml(string(light))
		//fmt.Println(selection.Html())
		//fmt.Println("light: ", string(light))
		//fmt.Println("\n\n\n")
	})
	htmlString, _ := doc.Html()
	return template.HTML(string(htmlString))
}

