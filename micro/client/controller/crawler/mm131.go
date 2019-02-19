package crawler

import (
	"fmt"
	"github.com/gocolly/colly"
	"hoper/utils"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"

	"strconv"
)

//Go中，双引号是用来表示字符串string，其实质是一个byte类型的数组，
// 单引号表示rune类型。
// 还有一个反引号，用来创建原生的字符串字面量，它可以由多行组成，但不支持任何转义序列

/*选择器	例子	例子描述	CSS
.class	.intro	选择 class="intro" 的所有元素。	1
#id	#firstname	选择 id="firstname" 的所有元素。	1
*	*	选择所有元素。	2
element	p	选择所有 <p> 元素。	1
element,element	div,p	选择所有 <div> 元素和所有 <p> 元素。	1
element element	div p	选择 <div> 元素内部的所有 <p> 元素。	1
element>element	div>p	选择父元素为 <div> 元素的所有 <p> 元素。	2
element+element	div+p	选择紧接在 <div> 元素之后的所有 <p> 元素。	2
[attribute]	[target]	选择带有 target 属性所有元素。	2
[attribute=value]	[target=_blank]	选择 target="_blank" 的所有元素。	2
[attribute~=value]	[title~=flower]	选择 title 属性包含单词 "flower" 的所有元素。	2
[attribute|=value]	[lang|=en]	选择 lang 属性值以 "en" 开头的所有元素。	2
:link	a:link	选择所有未被访问的链接。	1
:visited	a:visited	选择所有已被访问的链接。	1
:active	a:active	选择活动链接。	1
:hover	a:hover	选择鼠标指针位于其上的链接。	1
:focus	input:focus	选择获得焦点的 input 元素。	2
:first-letter	p:first-letter	选择每个 <p> 元素的首字母。	1
:first-line	p:first-line	选择每个 <p> 元素的首行。	1
:first-child	p:first-child	选择属于父元素的第一个子元素的每个 <p> 元素。	2
:before	p:before	在每个 <p> 元素的内容之前插入内容。	2
:after	p:after	在每个 <p> 元素的内容之后插入内容。	2
:lang(language)	p:lang(it)	选择带有以 "it" 开头的 lang 属性值的每个 <p> 元素。	2
element1~element2	p~ul	选择前面有 <p> 元素的每个 <ul> 元素。	3
[attribute^=value]	a[src^="https"]	选择其 src 属性值以 "https" 开头的每个 <a> 元素。	3
[attribute$=value]	a[src$=".pdf"]	选择其 src 属性以 ".pdf" 结尾的所有 <a> 元素。	3
[attribute*=value]	a[src*="abc"]	选择其 src 属性中包含 "abc" 子串的每个 <a> 元素。	3
:first-of-type	p:first-of-type	选择属于其父元素的首个 <p> 元素的每个 <p> 元素。	3
:last-of-type	p:last-of-type	选择属于其父元素的最后 <p> 元素的每个 <p> 元素。	3
:only-of-type	p:only-of-type	选择属于其父元素唯一的 <p> 元素的每个 <p> 元素。	3
:only-child	p:only-child	选择属于其父元素的唯一子元素的每个 <p> 元素。	3
:nth-child(n)	p:nth-child(2)	选择属于其父元素的第二个子元素的每个 <p> 元素。	3
:nth-last-child(n)	p:nth-last-child(2)	同上，从最后一个子元素开始计数。	3
:nth-of-type(n)	p:nth-of-type(2)	选择属于其父元素第二个 <p> 元素的每个 <p> 元素。	3
:nth-last-of-type(n)	p:nth-last-of-type(2)	同上，但是从最后一个子元素开始计数。	3
:last-child	p:last-child	选择属于其父元素最后一个子元素每个 <p> 元素。	3
:root	:root	选择文档的根元素。	3
:empty	p:empty	选择没有子元素的每个 <p> 元素（包括文本节点）。	3
:target	#news:target	选择当前活动的 #news 元素。	3
:enabled	input:enabled	选择每个启用的 <input> 元素。	3
:disabled	input:disabled	选择每个禁用的 <input> 元素	3
:checked	input:checked	选择每个被选中的 <input> 元素。	3
:not(selector)	:not(p)	选择非 <p> 元素的每个元素。	3
::selection	::selection	选择被用户选取的元素部分。	3*/
//写起来是真没python方便，用了框架都不如
func MM131() {

	fmt.Println("mm131")
	c1 := colly.NewCollector(
		colly.DetectCharset(),
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.132 Safari/537.36"),
		colly.AllowURLRevisit(),
	)

	c2 := colly.NewCollector(
		colly.DetectCharset(),
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.132 Safari/537.36"),
		colly.AllowURLRevisit(),
	)

	c2.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c2.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	c2.OnResponse(func(r *colly.Response) {
		fmt.Println("Visited", r.Request.URL)
	})

	aVisited := false
	max := 0
	c1.OnHTML("a[target=_blank]", func(e *colly.HTMLElement) {
		if aVisited {
			return
		}

		link := e.Attr("href")
		aVisited = true
		// Print link
		pattern := regexp.MustCompile(`\d{4,}`).FindAllString(link, -1)
		if len(pattern) > 0 {
			max, _ = strconv.Atoi(pattern[0])
		}

		for max != 0 {
			fmt.Println(max)
			page := strconv.Itoa(max)
			url := "http://www.mm131.com/xinggan/" + page + ".html"
			c2.Visit(url)
		}
	})

	c1.Visit("http://www.mm131.com/xinggan/")

	var dir string
	// On every a element which has href attribute call callback
	c2.OnHTML("h5", func(e *colly.HTMLElement) {

		dir = "E:/pic/" + e.Text

	})

	c2.OnHTML("span.page-ch", func(e *colly.HTMLElement) {

		pattern := regexp.MustCompile(`\d+`).FindAllString(e.Text, -1)

		if len(pattern) > 0 {
			dir = dir + pattern[0] + "P"
			fmt.Println(dir)
			if utils.CheckNotExist(dir) == true {
				if err := utils.Mkdir(dir); err != nil {
					fmt.Println(err)
				}
			} else {
				max = 0
				return
			}

			num, _ := strconv.Atoi(pattern[0])
			for i := 1; i <= num; i++ {
				getImg(strconv.Itoa(max), strconv.Itoa(i), dir)
			}
			max--
		}

	})

	c2.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished", r.Request.URL)
	})

}
func getImg(page string, num string, dir string) {
	path := "/" + num + ".jpg"
	client := http.Client{}
	//提交请求
	reqest, _ := http.NewRequest("GET", "http://img1.mm131.me/pic/"+page+path, nil)

	//增加header选项
	reqest.Header.Add("Cookie", "UM_distinctid=160c072721f36a-049309acceadc2-e323462-144000-160c0727220f67; CNZZDATA3866066=cnzz_eid%3D1829424698-1494676185-%26ntime%3D1494676185; bdshare_firstime=1515057214243; Hm_lvt_9a737a8572f89206db6e9c301695b55a=1515057214,1515074260,1515159455; Hm_lpvt_9a737a8572f89206db6e9c301695b55a=1515159455")
	reqest.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.132 Safari/537.36")
	reqest.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8")
	reqest.Header.Add("Accept-Encoding", "gzip, deflate")
	reqest.Header.Add("Connection", "keep-alive")
	reqest.Header.Add("Host", "img1.mm131.me")
	reqest.Header.Add("Referer", "http://www.mm131.com/")

	res, _ := client.Do(reqest)
	defer res.Body.Close()
	img, _ := ioutil.ReadAll(res.Body)
	if utils.CheckNotExist(dir+path) == true {
		fh, _ := os.Create(dir + path)
		defer fh.Close()
		fh.Write(img)
	}
}
