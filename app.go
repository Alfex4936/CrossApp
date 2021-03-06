package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
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

// Notice ...
type Notice struct {
	ID       int64  `json:"id"`
	Category string `json:"category"`
	Title    string `json:"title"`
	Date     string `json:"date"`
	Link     string `json:"link"`
	Writer   string `json:"writer"`
}

/// Parse 공지 불러오는 함수 (url, length) -> Promise<Notice[]>
func (a *App) Parse(url string, length int) ([]Notice, error) {
	ajouHTML := url
	if url == "" { // As default, use main link
		ajouHTML = fmt.Sprintf("%v?mode=list&articleLimit=%v&article.offset=0", AjouLink, length)
	}

	notices := []Notice{}

	resp, err := soup.Get(ajouHTML)
	if err != nil {
		fmt.Println("[Parser] Check your HTML connection.", err)
		return notices, err
	}
	doc := soup.HTMLParse(resp)

	ids := doc.FindAll("td", "class", "b-num-box")
	// if len(ids) == 0 {
	// 	fmt.Println("[Parser] Check your parser.")
	// 	return notices, err
	// }

	top_notices := len(ids)

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
		top_notices--
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

	if top_notices == len(ids) {
		return notices, errors.New("[Parser] No notices")
	}

	return notices, nil
}

const NaverWeather = "https://m.search.naver.com/search.naver?sm=tab_hty.top&where=nexearch&query=%EB%82%A0%EC%94%A8+%EB%A7%A4%ED%83%843%EB%8F%99&oquery=%EB%82%A0%EC%94%A8"
const NaverWeatherIcon = "https://weather.naver.com/today/02117530?cpName=ACCUWEATHER"

// Weather 해외 기상은 일출부터 일몰 전이 낮, 일몰부터 일출 전이 밤
type Weather struct {
	MaxTemp       string `json:"max_temp"`        // 최고 온도
	MinTemp       string `json:"min_temp"`        // 최저 온도
	CurrentTemp   string `json:"current_temp"`    // 현재 온도
	CurrentStatus string `json:"current_stat"`    // 흐림, 맑음 ...
	RainDay       string `json:"rain_day"`        // 강수 확률 낮
	RainNight     string `json:"rain_night"`      // 강수 확률 밤
	FineDust      string `json:"fine_dust"`       // 미세먼지
	UltraDust     string `json:"ultra_dust"`      // 초미세먼지
	UV            string `json:"uv"`              // 자외선 지수
	Sunset        string `json:"sunset"`          // 일몰
	Icon          string `json:"icon"`            // 날씨 아이콘 (ico_animation_wt?)
	TmrRainDay    string `json:"tmrw_rain_day"`   // 내일 낮 강수 확률
	TmrRainNight  string `json:"tmrw_rain_night"` // 내일 낮 강수 확률
}

// GetWeather is a function that parses suwon's weather today
func (a *App) GetWeather() (Weather, error) { // Promise<resolve, reject>
	var weather Weather

	resp, err := soup.Get(NaverWeather)
	if err != nil {
		fmt.Println("Check your HTML connection.")
		return weather, err // nil
	}

	resp2, err := soup.Get(NaverWeatherIcon)
	if err != nil {
		fmt.Println("Check your HTML connection.")
		return weather, err // nil
	}
	doc := soup.HTMLParse(resp)
	doc2 := soup.HTMLParse(resp2)

	currentTemp := doc.Find("div", "class", "temperature_text").Find("strong").Text() + "도"

	// ! 해외 기상은 일출부터 일몰 전이 낮, 일몰부터 일출 전이 밤
	maxTemp := doc.Find("span", "class", "highest")
	dayTemp := maxTemp.Text() + "도"
	dayTemp = strings.Replace(dayTemp, "°", "", 1)

	minTemp := doc.Find("span", "class", "lowest")
	nightTemp := minTemp.Text() + "도"
	nightTemp = strings.Replace(nightTemp, "°", "", 1)

	currentStatElem := doc.Find("span", "class", "weather")
	currentStat := currentStatElem.Text()

	// [미세먼지, 초미세먼지, 자외선, 일몰 시간]
	statuses := doc.FindAll("li", "class", "item_today")
	fineDust := statuses[0].Find("a").Find("span").Text()
	ultraDust := statuses[1].Find("a").Find("span").Text()
	UV := statuses[2].Find("a").Find("span").Text()
	sunset := statuses[3].Find("a").Find("span").Text()

	// 강우량
	rainElems := doc.FindAll("span", "class", "rainfall")
	dayRain := rainElems[0].Text()
	nightRain := rainElems[1].Text()
	tmrwDayRain := rainElems[2].Text()
	tmrwNghtRain := rainElems[3].Text()

	// 날씨 아이콘
	i := doc2.Find("div", "class", "summary_img").Find("i").Attrs()
	img := i["data-ico"]

	if strings.Contains(i["class"], "night") {
		img += "_night"
	}

	// struct 값 변경
	weather.CurrentTemp = currentTemp
	weather.CurrentStatus = currentStat

	weather.MaxTemp = dayTemp // Assert that (day temp > night temp) in general
	weather.MinTemp = nightTemp

	weather.RainDay = dayRain
	weather.RainNight = nightRain
	weather.TmrRainDay = tmrwDayRain
	weather.TmrRainNight = tmrwNghtRain

	weather.FineDust = fineDust
	weather.UltraDust = ultraDust
	weather.UV = UV
	weather.Sunset = sunset
	weather.Icon = fmt.Sprintf("https://raw.githubusercontent.com/Alfex4936/KakaoChatBot-Golang/main/imgs/%s.png?raw=true", img)

	// fmt.Printf("%+v", weather)

	return weather, nil
}

type PeopleJSON struct {
	MsgCode     string `json:"msgCode"`
	PhoneNumber []struct {
		BussNm    string `json:"bussNm"`           // 업무명: "XXX학과(공학인증)"
		DeptCd    string `json:"deptCd"`           // "DS01234657"
		DeptNm    string `json:"deptNm"`           // 부서명: "정보통신대학교학팀(팔달관 777-1)"
		Email     string `json:"email"`            // 이메일: "example@ajou.ac.kr"
		Name      string `json:"korNm"`            // 이름(신분): "이름1(직원)" | "이름2(교원)"
		MdfLineNo int64  `json:"mdfLineNo,string"` // "289"
		TelNo     string `json:"telNo"`            // 전화번호: 031-219-"1234"
		UserNo    int64  `json:"userNo,string"`    // "201900000"
	} `json:"phoneNumber"`
}

type People struct {
	BussNm    string `json:"bussNm"`           // 업무명: "XXX학과(공학인증)"
	DeptCd    string `json:"deptCd"`           // "DS01234657"
	DeptNm    string `json:"deptNm"`           // 부서명: "정보통신대학교학팀(팔달관 777-1)"
	Email     string `json:"email"`            // 이메일: "example@ajou.ac.kr"
	Name      string `json:"korNm"`            // 이름(신분): "이름1(직원)" | "이름2(교원)"
	MdfLineNo int64  `json:"mdfLineNo,string"` // "289"
	TelNo     string `json:"telNo"`            // 전화번호: 031-219-"1234"
	UserNo    int64  `json:"userNo,string"`    // "201900000"
}

var AjouPeople = "https://mportal.ajou.ac.kr/system/phone/selectList.ajax"

// GetPeople searches a given name in AjouPeople link
func (a *App) GetPeople(keyword string) ([]People, error) {
	var people PeopleJSON
	result := []People{}

	// POST
	jsonValue, _ := json.Marshal(map[string]string{"keyword": keyword})
	// Disable SSL authentication and post jsonValue
	// http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	client := &http.Client{}
	req, err := http.NewRequest("POST", AjouPeople, bytes.NewBuffer(jsonValue))
	if err != nil {
		fmt.Println(err)
		return result, err
	}

	req.Header.Add("User-Agent", `Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36`)
	req.Header.Add("Content-Type", `application/json;charset=UTF-8`)
	resp, err := client.Do(req)

	// resp, err := http.Post((AjouPeople), "application/json;charset=UTF-8", bytes.NewBuffer(jsonValue))
	if err != nil {
		fmt.Println(err)
		return result, err
	}
	defer resp.Body.Close()

	// Response 체크.
	respBody, err := io.ReadAll(resp.Body)
	// fmt.Printf("%s\n", respBody)
	json.Unmarshal(respBody, &people)
	if err != nil {
		fmt.Println(err)
		return result, err
	}

	// fmt.Println(people.PhoneNumber[0].Name)
	for _, person := range people.PhoneNumber {
		result = append(result, person)
	}

	if len(result) > 10 {
		result = result[:10]
	}

	return result, nil
}
