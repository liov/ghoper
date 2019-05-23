package crawler

import (
	"hoper/utils/ulog"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
	"github.com/kataras/iris"
	"hoper/client/controller/common"
	"hoper/client/controller/upload"
	"hoper/initialize"
	"hoper/model"
	"hoper/model/e"
	"hoper/model/ov"
	"hoper/utils"
)

type crawlSelector struct {
	From                  int
	ListItemSelector      string
	ListItemTitleSelector string
	TitleSelector         string
	ContentSelector       string
}

func createCrawlSelector(from int) crawlSelector {
	selector := crawlSelector{
		From: from,
	}
	switch from {
	case model.ArticleFromJianShu:
		selector.ListItemSelector = ".note-list li"
		selector.ListItemTitleSelector = ".title"
		selector.TitleSelector = ".article .title"
		selector.ContentSelector = ".show-content"
	case model.ArticleFromZhihu:
		selector.ListItemSelector = ".PostListItem"
		selector.ListItemTitleSelector = ".PostListItem-info a"
		selector.TitleSelector = ".PostIndex-title"
		selector.ContentSelector = ".PostIndex-content"
	case model.ArticleFromHuxiu:
		selector.ListItemSelector = ".mod-art"
		selector.ListItemTitleSelector = ".mob-ctt h2 a"
		selector.TitleSelector = ".t-h1"
		selector.ContentSelector = ".article-content-wrap"
	case model.ArticleFromCustom:
		selector.ListItemSelector = ""
		selector.ListItemTitleSelector = ""
		selector.TitleSelector = ""
		selector.ContentSelector = ""
	case model.ArticleFromNULL:
		selector.ListItemSelector = ""
		selector.ListItemTitleSelector = ""
		selector.TitleSelector = ""
		selector.ContentSelector = ""
	}
	return selector
}

type sourceHTML struct {
	from          int
	sourceHTMLStr string
}

func createSourceHTML(from int) string {
	var htmlArr []string
	switch from {
	case model.ArticleFromJianShu:
		htmlArr = []string{
			"<div id=\"golang123-content-outter-footer\">",
			"<blockquote>",
			"<p>来源: <a href=\"https://www.jianshu.com/\" target=\"_blank\">简书</a><br>",
			"原文: <a href=\"{articleURL}\" target=\"_blank\">{title}</a></p>",
			"</blockquote>",
			"</div>",
		}
	case model.ArticleFromZhihu:
		htmlArr = []string{
			"<div id=\"golang123-content-outter-footer\">",
			"<blockquote>",
			"<p>来源: <a href=\"https://www.zhihu.com\" target=\"_blank\">知乎</a><br>",
			"原文: <a href=\"{articleURL}\" target=\"_blank\">{title}</a></p>",
			"</blockquote>",
			"</div>",
		}
	case model.ArticleFromHuxiu:
		htmlArr = []string{
			"<div id=\"golang123-content-outter-footer\">",
			"<blockquote>",
			"<p>来源: <a href=\"https://www.huxiu.com\" target=\"_blank\">虎嗅</a><br>",
			"原文: <a href=\"{articleURL}\" target=\"_blank\">{title}</a></p>",
			"</blockquote>",
			"</div>",
		}
	case model.ArticleFromCustom:
		htmlArr = []string{
			"<div id=\"golang123-content-outter-footer\">",
			"<blockquote>",
			"<p>来源: <a href=\"{siteURL}\" target=\"_blank\">{siteName}</a><br>",
			"原文: <a href=\"{articleURL}\" target=\"_blank\">{title}</a></p>",
			"</blockquote>",
			"</div>",
		}
	case model.ArticleFromNULL:
		htmlArr = []string{}
	}
	return strings.Join(htmlArr, "")
}

func createArticle(user model.User, category ov.Category, from int, data map[string]string) {
	var article model.Article
	article.Title = data["Title"]
	article.HTMLContent = data["Content"]
	article.ContentType = model.ContentTypeHTML
	article.UserID = user.ID
	article.Status = model.ArticleVerifying
	article.Categories = append(article.Categories, category)

	var crawlerArticle model.CrawlerArticle
	crawlerArticle.URL = data["URL"]
	crawlerArticle.Title = article.Title
	crawlerArticle.Content = article.HTMLContent
	crawlerArticle.From = from

	tx := initialize.DB.Begin()
	if err := tx.Create(&article).Error; err != nil {
		tx.Rollback()
		return
	}
	crawlerArticle.ID = article.ID
	if err := tx.Create(&crawlerArticle).Error; err != nil {
		tx.Rollback()
		return
	}
	tx.Commit()
}

func crawlContent(pageURL string, crawlSelector crawlSelector, siteInfo map[string]string, crawlExist bool) map[string]string {
	var crawlerArticle model.CrawlerArticle
	if err := initialize.DB.Where("url = ?", pageURL).Find(&crawlerArticle).Error; err == nil {
		if !crawlExist {
			// 当crawlExist为false时，已抓取过的文章就不再抓取
			return nil
		}
	}
	articleDOC, err := goquery.NewDocument(pageURL)
	if err != nil {
		return nil
	}
	title := articleDOC.Find(crawlSelector.TitleSelector).Text()
	if title == "" && crawlSelector.From != model.ArticleFromNULL {
		return nil
	}
	contentDOM := articleDOC.Find(crawlSelector.ContentSelector)
	imgs := contentDOM.Find("img")
	if imgs.Length() > 0 {
		imgs.Each(func(j int, img *goquery.Selection) {
			imgURL, exists := img.Attr("src")
			var ext string
			if !exists {
				if crawlSelector.From != model.ArticleFromJianShu {
					return
				}
				originalSrc, originalExists := img.Attr("data-original-src")
				if originalExists && originalSrc != "" {
					tempImgURL, tempErr := utils.RelativeURLToAbsoluteURL(originalSrc, pageURL)
					if tempErr != nil || tempImgURL == "" {
						return
					}
					imgURL = tempImgURL
					resp, err := http.Head(imgURL)
					if err != nil {
						ulog.Error(err)
						return
					}

					defer resp.Body.Close()

					contentType := resp.Header.Get("content-type")
					if contentType == "image/jpeg" {
						ext = ".jpg"
					} else if contentType == "image/gif" {
						ext = ".gif"
					} else if contentType == "image/png" {
						ext = ".png"
					}
				}
			}

			if imgURL == "" || crawlSelector.From == model.ArticleFromZhihu && strings.Index(imgURL, "data:image/svg+xml;utf8,") == 0 {
				actualsrc, actualsrcExists := img.Attr("data-actualsrc")
				if actualsrcExists && actualsrc != "" {
					imgURL = actualsrc
				}
			}
			var imgURLErr error
			imgURL, imgURLErr = utils.RelativeURLToAbsoluteURL(imgURL, pageURL)
			if imgURLErr != nil || imgURL == "" {
				return
			}
			urlData, urlErr := url.Parse(imgURL)
			if urlErr != nil {
				return
			}

			if ext == "" {
				index := strings.LastIndex(urlData.Path, ".")
				if index >= 0 {
					ext = urlData.Path[index:]
				}
			}

			resp, err := http.Get(imgURL)

			if err != nil {
				return
			}

			defer resp.Body.Close()

			imgUploadedInfo := upload.GenerateUploadedInfo(ext)
			if err := os.MkdirAll(initialize.Config.Server.UploadDir, 0777); err != nil {
				ulog.Error(err)
				return
			}
			out, outErr := os.OpenFile(imgUploadedInfo.UploadFilePath, os.O_WRONLY|os.O_CREATE, 0666)
			if outErr != nil {
				ulog.Error(err)
				return
			}

			defer out.Close()

			if _, err := io.Copy(out, resp.Body); err != nil {
				ulog.Error(err)
				return
			}
			img.SetAttr("src", initialize.Config.Server.UploadDir)
		})
	}
	contentDOM.Find("a").Each(func(j int, a *goquery.Selection) {
		oldHref, exists := a.Attr("href")
		if exists {
			href, err := utils.RelativeURLToAbsoluteURL(oldHref, pageURL)
			if err == nil {
				a.SetAttr("href", href)
			}
		}
	})
	articleHTML, htmlErr := contentDOM.Html()
	if htmlErr != nil {
		return nil
	}

	sourceHTML := createSourceHTML(crawlSelector.From)
	if crawlSelector.From == model.ArticleFromCustom {
		sourceHTML = strings.Replace(sourceHTML, "{siteURL}", siteInfo["siteURL"], -1)
		sourceHTML = strings.Replace(sourceHTML, "{siteName}", siteInfo["siteName"], -1)
	}
	sourceHTML = strings.Replace(sourceHTML, "{title}", title, -1)
	sourceHTML = strings.Replace(sourceHTML, "{articleURL}", pageURL, -1)
	articleHTML += sourceHTML
	articleHTML = "<div id=\"golang123-content-outter\">" + articleHTML + "</div>"
	return map[string]string{
		"Title":   title,
		"Content": articleHTML,
		"URL":     pageURL,
	}
}

func crawlList(listURL string, user model.User, category ov.Category, crawlSelector crawlSelector, siteInfo map[string]string, crawlExist bool, wg *sync.WaitGroup) {
	defer wg.Done()

	if _, err := url.Parse(listURL); err != nil {
		return
	}

	doc, docErr := goquery.NewDocument(listURL)
	if docErr != nil {
		ulog.Error(docErr)
		return
	}

	var articleURLArr []string
	doc.Find(crawlSelector.ListItemSelector).Each(func(i int, s *goquery.Selection) {
		articleLink := s.Find(crawlSelector.ListItemTitleSelector)
		href, exists := articleLink.Attr("href")
		href = strings.TrimSpace(href)
		if exists {
			url, err := utils.RelativeURLToAbsoluteURL(href, listURL)
			if err == nil {
				articleURLArr = append(articleURLArr, url)
			}
		}
	})

	for i := 0; i < len(articleURLArr); i++ {
		articleMap := crawlContent(articleURLArr[i], crawlSelector, siteInfo, crawlExist)
		if articleMap != nil {
			createArticle(user, category, crawlSelector.From, articleMap)
		}
	}
}

// Crawl 抓取文章
func Crawl(c iris.Context) {

	type JSONData struct {
		URLS       []string `json:"urls"`
		From       int      `json:"from"`
		CategoryID int      `json:"categoryID"`
		Scope      string   `json:"scope"`
		CrawlExist bool     `json:"crawlExist"`
	}
	var jsonData JSONData
	if err := c.ReadJSON(&jsonData); err != nil {
		common.Response(c, "参数无效")
		return
	}

	if jsonData.From != model.ArticleFromJianShu && jsonData.From != model.ArticleFromZhihu &&
		jsonData.From != model.ArticleFromHuxiu {
		common.Response(c, "无效的from")
		return
	}
	if jsonData.Scope != model.CrawlerScopePage && jsonData.Scope != model.CrawlerScopeList {
		common.Response(c, "无效的scope")
		return
	}

	user := c.GetViewData()["user"].(model.User)

	if user.Name != initialize.Config.Server.CrawlerName {
		common.Response(c, "您没有权限执行此操作, 请使用爬虫账号")
		return
	}

	var category ov.Category
	if err := initialize.DB.First(&category, jsonData.CategoryID).Error; err != nil {
		ulog.Error(err)
		common.Response(c, "错误的categoryID")
		return
	}

	crawlSelector := createCrawlSelector(jsonData.From)

	if jsonData.Scope == model.CrawlerScopeList {
		var wg sync.WaitGroup
		for i := 0; i < len(jsonData.URLS); i++ {
			wg.Add(1)
			go crawlList(jsonData.URLS[i], user, category, crawlSelector, nil, jsonData.CrawlExist, &wg)
		}
		wg.Wait()
	} else if jsonData.Scope == model.CrawlerScopePage {
		for i := 0; i < len(jsonData.URLS); i++ {
			data := crawlContent(jsonData.URLS[i], crawlSelector, nil, jsonData.CrawlExist)
			if data != nil {
				createArticle(user, category, jsonData.From, data)
			}
		}
	}

	common.Response(c, "抓取完成")
}

// CustomCrawl 自定义抓取
func CustomCrawl(c iris.Context) {

	type JSONData struct {
		URLS                  []string `json:"urls"`
		From                  int      `json:"from"`
		CategoryID            int      `json:"categoryID"`
		Scope                 string   `json:"scope"`
		CrawlExist            bool     `json:"crawlExist"`
		ListItemSelector      string   `json:"listItemSelector"`
		ListItemTitleSelector string   `json:"listItemTitleSelector"`
		TitleSelector         string   `json:"titleSelector"`
		ContentSelector       string   `json:"contentSelector"`
		SiteURL               string   `json:"siteURL" binding:"required,url"`
		SiteName              string   `json:"siteName" binding:"required"`
	}
	var jsonData JSONData
	if err := c.ReadJSON(&jsonData); err != nil {
		common.Response(c, "参数无效")
		return
	}

	if jsonData.From != model.ArticleFromCustom {
		common.Response(c, "无效的from")
		return
	}
	if jsonData.Scope != model.CrawlerScopePage && jsonData.Scope != model.CrawlerScopeList {
		common.Response(c, "无效的scope")
		return
	}

	user := c.GetViewData()["user"].(model.User)

	if user.Name != initialize.Config.Server.CrawlerName {
		common.Response(c, "您没有权限执行此操作, 请使用爬虫账号")
		return
	}

	var category ov.Category
	if err := initialize.DB.First(&category, jsonData.CategoryID).Error; err != nil {
		ulog.Error(err)
		common.Response(c, "错误的categoryID")
		return
	}

	crawlSelector := createCrawlSelector(model.ArticleFromCustom)
	crawlSelector.ListItemSelector = jsonData.ListItemSelector
	crawlSelector.ListItemTitleSelector = jsonData.ListItemTitleSelector
	crawlSelector.TitleSelector = jsonData.TitleSelector
	crawlSelector.ContentSelector = jsonData.ContentSelector

	siteInfo := map[string]string{
		"siteURL":  jsonData.SiteURL,
		"siteName": jsonData.SiteName,
	}
	if jsonData.Scope == model.CrawlerScopeList {
		var wg sync.WaitGroup
		for i := 0; i < len(jsonData.URLS); i++ {
			wg.Add(1)
			go crawlList(jsonData.URLS[i], user, category, crawlSelector, siteInfo, jsonData.CrawlExist, &wg)
		}
		wg.Wait()
	} else if jsonData.Scope == model.CrawlerScopePage {
		for i := 0; i < len(jsonData.URLS); i++ {
			data := crawlContent(jsonData.URLS[i], crawlSelector, siteInfo, jsonData.CrawlExist)
			if data != nil {
				createArticle(user, category, crawlSelector.From, data)
			}
		}
	}

	common.Response(c, "抓取完成")
}

// CrawlNotSaveContent 抓取的内容直接返回，而不保存到数据库
func CrawlNotSaveContent(c iris.Context) {

	type JSONData struct {
		URL             string `json:"url"`
		TitleSelector   string `json:"titleSelector"`
		ContentSelector string `json:"contentSelector"`
	}
	var jsonData JSONData
	if err := c.ReadJSON(&jsonData); err != nil {
		common.Response(c, "参数无效")
		return
	}

	crawlSelector := createCrawlSelector(model.ArticleFromNULL)
	crawlSelector.TitleSelector = jsonData.TitleSelector
	crawlSelector.ContentSelector = jsonData.ContentSelector

	data := crawlContent(jsonData.URL, crawlSelector, nil, true)

	common.Response(c, iris.Map{
		"content": data["Content"],
	}, "success")
}

// CrawlAccount 获取爬虫账号
func CrawlAccount(c iris.Context) {

	var users []model.User
	if err := initialize.DB.Where("name = ?", initialize.Config.Server.CrawlerName).Find(&users).Error; err != nil {
		ulog.Error(err)
		common.Response(c, "error")
		return
	}
	common.Response(c, common.H{
		"errNo": e.SUCCESS,
		"msg":   "success",
		"data":  users,
	})
}

// CreateAccount 创建爬虫账号
func CreateAccount(c iris.Context) {

	var users []model.User
	if err := initialize.DB.Where("name = ?", initialize.Config.Server.CrawlerName).Find(&users).Error; err != nil {
		ulog.Error(err)
		common.Response(c, "error")
		return
	}
	if len(users) <= 0 {
		var user model.User
		user.Name = initialize.Config.Server.CrawlerName
		user.Role = model.UserRoleCrawler
		user.AvatarURL = "/images/avatar/spider.png"
		user.Status = model.UserStatusActived
		if err := initialize.DB.Save(&user).Error; err != nil {
			ulog.Error(err)
			common.Response(c, "error")
			return
		}
		common.Response(c, common.H{
			"errNo": e.SUCCESS,
			"msg":   "success",
			"data":  []model.User{user},
		})
		return
	}
	common.Response(c, "爬虫账号已存在")
}
