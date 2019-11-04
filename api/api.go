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

	limitDesc := ([]rune)(wrap.Find(".wea_about b").Text())
	limit := ""
	if len(limitDesc) <= 4 {
		limit = string(limitDesc)
	} else {
		limit = string(limitDesc[4:])
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
	imgURL, _ := wrap.Find(".fp-one-imagen").Attr("src")
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
	imgURL, _ := wrap.Find(".himg").Attr("src")
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

// GetWallpaper from bing
func GetWallpaper() Wallpaper {
	url := "https://cn.bing.com"
	doc := FetchHTML(url)
	imgURL, _ := doc.Find("#bgLink").Attr("href")
	imgURL = url + imgURL
	title, titleOk := doc.Find("#sh_cp").Attr("title")
	if titleOk {
		title = strings.Split(title, " ")[0]
	}
	return Wallpaper{
		Title:  title,
		ImgURL: imgURL,
	}
}

// GetTrivia data
func GetTrivia() Trivia {
	url := "http://www.lengdou.net/random"
	doc := FetchHTML(url)
	wrap := doc.Find(".container .media .media-body")
	imgURL, _ := wrap.Find(".topic-img img").Attr("src")
	return Trivia{
		ImgURL:      imgURL,
		Description: strings.Split(wrap.Find(".topic-content").Text(), "#")[0],
	}
}
