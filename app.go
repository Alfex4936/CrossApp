package main

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/anaskhan96/soup"
)

const AjouLink = "https://www.ajou.ac.kr/kr/ajou/notice.do"

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (b *App) shutdown(ctx context.Context) {

}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return name
}

// Notice ...
type Notice struct {
	ID       int64  `json:"id"`
	Category string `json:"category"`
	Title    string `json:"title"`
	Date     string `json:"date"`
	Link     string `json:"link"`
	Writer   string `json:"writer"`
}

func (a *App) Parse(url string, length int) []Notice {
	ajouHTML := url
	if url == "" { // As default, use main link
		ajouHTML = fmt.Sprintf("%v?mode=list&articleLimit=%v&article.offset=0", AjouLink, length)
	}

	notices := []Notice{}

	resp, err := soup.Get(ajouHTML)
	if err != nil {
		fmt.Println("[Parser] Check your HTML connection.", err)
		return notices
	}
	doc := soup.HTMLParse(resp)

	ids := doc.FindAll("td", "class", "b-num-box")
	if len(ids) == 0 {
		fmt.Println("[Parser] Check your parser.")
		return notices
	}

	titles := doc.FindAll("div", "class", "b-title-box")
	dates := doc.FindAll("span", "class", "b-date")
	categories := doc.FindAll("span", "class", "b-cate")
	//links := doc.FindAll("div", "class", "b-title-box")
	writers := doc.FindAll("span", "class", "b-writer")
	for i := 0; i < len(ids); i++ {
		id, err := strconv.ParseInt(strings.TrimSpace(ids[i].Text()), 10, 64)
		if err != nil {
			continue // 최상위 공지
		}
		title := strings.TrimSpace(titles[i].Find("a").Text())
		link := titles[i].Find("a").Attrs()["href"]
		category := strings.TrimSpace(categories[i].Text())
		date := strings.TrimSpace(dates[i].Text())
		writer := writers[i].Text()

		duplicate := "[" + writer + "]"
		if strings.Contains(title, duplicate) {
			title = strings.TrimSpace(strings.Replace(title, duplicate, "", 1))
		}

		notice := Notice{ID: id, Category: category, Title: title, Date: date, Link: AjouLink + link, Writer: writer}
		notices = append(notices, notice)
	}

	// fmt.Printf("%v", notices)

	return notices
}
