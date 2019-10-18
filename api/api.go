package api

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"strings"

	query "github.com/PuerkitoBio/goquery"
)

// One site info
type One struct {
	Date   string
	ImgURL string
	Word   string
}

// Weather data from https://tianqi.moji.com/weather/china
type Weather struct {
	City     string
	Temp     string
	Weather  string
	Air      string
	Humidity string
	Wind     string
	Limit    string
	Note     string
}

// CreateClient a http
func CreateClient() *http.Client {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	return &http.Client{Transport: tr}
}

// FetchDom from url
func FetchDom(url string) *query.Document {
	client := CreateClient()
	resp, err := client.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", resp.StatusCode, resp.Status)
	}
	doc, err := query.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return doc
}

// GetWeather data
func GetWeather(local string) Weather {
	url := "https://tianqi.moji.com/weather/china/" + local
	doc := FetchDom(url)
	wrap := doc.Find(".wea_info .left")
	humidityDesc := strings.Split(wrap.Find(".wea_about span").Text(), " ")
	humidity := "未知"
	if len(humidityDesc) >= 2 {
		humidity = humidityDesc[1]
	}

	limit := wrap.Find(".wea_about b").Text()
	if limit != "" {
		limit = strings.ReplaceAll(limit, "尾号限行", "")
	} else {
		limit = "无"
	}
	return Weather{
		City:     doc.Find("#search .search_default em").Text(),
		Temp:     wrap.Find(".wea_weather em").Text() + "°",
		Weather:  wrap.Find(".wea_weather b").Text(),
		Air:      wrap.Find(" .wea_alert em").Text(),
		Humidity: humidity,
		Wind:     wrap.Find(".wea_about em").Text(),
		Limit:    limit,
		Note:     strings.ReplaceAll(wrap.Find(".wea_tips em").Text(), "。", ""),
	}
}

// GetONE data
func GetONE() One {
	url := "http://wufazhuce.com/"
	doc := FetchDom(url)
	wrap := doc.Find(".fp-one .carousel .item.active")
	day := wrap.Find(".dom").Text()
	monthYear := wrap.Find(".may").Text()
	imgURL, ok := wrap.Find(".fp-one-imagen").Attr("src")
	if !ok {
		imgURL = ""
	}
	return One{
		ImgURL: imgURL,
		Date:   fmt.Sprintf("%s %s", day, monthYear),
		Word:   wrap.Find(".fp-one-cita a").Text(),
	}
}
