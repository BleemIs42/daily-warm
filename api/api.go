package api

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	query "github.com/PuerkitoBio/goquery"
)

// Weather site data
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

// One site info
type One struct {
	Date     string
	ImgURL   string
	Sentence string
}

// English info
type English struct {
	ImgURL   string
	Sentence string
}

// Poem info
type Poem struct {
	Title   string   `json:"title"`
	Dynasty string   `json:"dynasty"`
	Author  string   `json:"author"`
	Content []string `json:"content"`
}

// PoemRes response data
type PoemRes struct {
	Status string `json:"status"`
	Data   struct {
		Origin Poem `json:"origin"`
	} `json:"data"`
}

// CreateClient a http
func CreateClient() *http.Client {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	return &http.Client{Transport: tr}
}

// Fetch from url
func Fetch(url string) io.Reader {
	client := CreateClient()
	resp, err := client.Get(url)
	if err != nil {
		log.Fatalf("Fetch url from %s error: %s", url, err)
	}
	if resp.StatusCode != 200 {
		log.Fatalf("Status code error: %d %s", resp.StatusCode, resp.Status)
	}
	return resp.Body
}

// FetchHTML from url
func FetchHTML(url string) *query.Document {
	doc, err := query.NewDocument(url)
	if err != nil {
		log.Fatalf("Fetch html from %s error: %s", url, err)
	}
	return doc
}

// GetWeather data
func GetWeather(local string) Weather {
	url := "https://tianqi.moji.com/weather/china/" + local
	doc := FetchHTML(url)
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
	doc := FetchHTML(url)
	wrap := doc.Find(".fp-one .carousel .item.active")
	day := wrap.Find(".dom").Text()
	monthYear := wrap.Find(".may").Text()
	imgURL, ok := wrap.Find(".fp-one-imagen").Attr("src")
	if !ok {
		imgURL = ""
	}
	return One{
		ImgURL:   imgURL,
		Date:     fmt.Sprintf("%s %s", day, monthYear),
		Sentence: wrap.Find(".fp-one-cita a").Text(),
	}
}

// GetEnglish data
func GetEnglish() English {
	url := "http://dict.eudic.net/home/dailysentence"
	doc := FetchHTML(url)
	wrap := doc.Find(".containter .head-img")
	imgURL, ok := wrap.Find(".himg").Attr("src")
	if !ok {
		imgURL = ""
	}
	return English{
		ImgURL:   imgURL,
		Sentence: wrap.Find(".sentence .sect_en").Text(),
	}
}

// GetPoem data
func GetPoem() Poem {
	url := "https://v2.jinrishici.com/one.json"

	buf := new(bytes.Buffer)
	res := Fetch(url)
	buf.ReadFrom(res)
	resByte := buf.Bytes()

	var resJSON PoemRes
	err := json.Unmarshal(resByte, &resJSON)
	if err != nil {
		log.Fatalf("Fetch json from %s error: %s", url, err)
	}

	status := resJSON.Status
	if status != "success" {
		log.Fatalf("Get poem status %s, res: %s", status, resJSON)
	}
	return resJSON.Data.Origin
}
