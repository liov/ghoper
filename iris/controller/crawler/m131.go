package crawler

import (
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html/charset"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"service/utils"
	"strconv"
)

func M131() {

	doc := getDoc("http://www.mm131.com/xinggan/", "text/html; charset=gb2312")
	max := 0
	link := doc.Find("a[target=_blank]").Nodes[0].Attr[1].Val
	pattern := regexp.MustCompile(`\d{4,}`).FindAllString(link, -1)
	if len(pattern) > 0 {
		max, _ = strconv.Atoi(pattern[0])
	}

	for max != 0 {
		page := strconv.Itoa(max)
		url := "http://www.mm131.com/xinggan/" + page + ".html"
		doc := getDoc(url, "text/html; charset=gb2312")
		s1 := doc.Find("h5").Text()
		dir := "E:/pic/" + s1
		s := doc.Find("span.page-ch").Nodes[0].FirstChild.Data
		pattern := regexp.MustCompile(`\d+`).FindAllString(s, -1)
		if len(pattern) > 0 {
			dir = dir + pattern[0] + "P"
			fmt.Println(dir)
			if utils.CheckNotExist(dir) == true {
				if err := utils.Mkdir(dir); err != nil {
					fmt.Println(err)
				}
			} else {
				return
			}
			num, _ := strconv.Atoi(pattern[0])
			for i := 1; i <= num; i++ {
				getImg(strconv.Itoa(max), strconv.Itoa(i), dir)
			}
			max--
		}
	}

}

func getDoc(url string, contentType string) *goquery.Document {

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	rb, _ := ioutil.ReadAll(res.Body)
	r, _ := charset.NewReader(bytes.NewReader(rb), contentType)
	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		log.Fatal(err)
	}
	return doc
}
