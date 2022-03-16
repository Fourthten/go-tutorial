package main

import (
	"encoding/json"
	"log"

	"github.com/beevik/etree"
)

type Document struct {
	Title   string
	URL     string
	Content struct {
		Articles []struct {
			Title      string
			URL        string
			Categories []string
			Info       string
		}
	}
}

const jsonString = `{
	"Title": "My Website",
	"URL": "http://www.mywebsite.com",
	"Content": {
		"Articles": [{
			"Categories": ["News", "Sports"],
			"Title": "Article 1",
			"URL": "/article1/"
		}, {
			"Categories": ["News"],
			"Title": "Article 2",
			"URL": "/article2/"
		}, {
			"Categories": ["Sports"],
			"Info": "popular article",
			"Title": "Article 3",
			"URL": "/article3/"
		}]
	}
}`

func main() {
	data := Document{}
	err := json.Unmarshal([]byte(jsonString), &data)
	if err != nil {
		log.Fatal(err.Error())
	}
	doc := etree.NewDocument()
	doc.CreateProcInst("xml", `version="1.0" encoding="UTF-8"`)
	website := doc.CreateElement("website")
	website.CreateElement("title").SetText(data.Title)
	website.CreateElement("url").SetText(data.URL)
	content := website.CreateElement("contents")

	for _, each := range data.Content.Articles {
		article := content.CreateElement("article")
		article.CreateElement("title").SetText(each.Title)
		article.CreateElement("url").SetText(each.URL)
		for _, category := range each.Categories {
			article.CreateElement("category").SetText(category)
		}
		if each.Info != "" {
			article.CreateAttr("info", each.Info)
		}
	}
	doc.Indent(2)
	err = doc.WriteToFile("../excel/file/output.xml")
	if err != nil {
		log.Println(err.Error())
	}
}
