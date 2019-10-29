# Daily-warm
> 每天定时发邮件给你关心的人, 内容包含天气, one 的一句话, 一句英语, 一首古诗

## Usage

```bash
# 1. Config `.env` file in root dir, e.g.
MAIL_USERNAME = user@qq.com
MAIL_PASSWORD = ***********
MAIL_HOST = smtp.qq.com
MAIL_PORT = 25
# 每分钟发一次
MAIL_CRON = "0/1 * * * *"
MAIL_SUBJECT = 每日一暖, 温情一生
MAIL_FROM = user<user@qq.com>
MAIL_TO = [{"email": "user<user@qq.com>", "local": "shaanxi/xian"}]

# 2.Build
go mod download
go build -o dwm.out *.go

# 3. Run
./dwm.out
```

## Screenshot

<img width="400" src="https://github.com/BarryYan/daily-warm/blob/master/screenshot.jpg?raw=true">

## Package
[github.com/barryyan/daily-warm/gomail](https://godoc.org/github.com/BarryYan/daily-warm/gomail)

```go
// import "github.com/barryyan/daily-warm/gomail"
package gomail 

type Configuration struct{
  Host     string
  Port     uint16
  Username string
  Password string
  From     string
}
var Config = Configuration{}

type GoMail struct{
  From    string
  To      []string
  Cc      []string
  Bcc     []string
  Subject string
  Content string
}
func (gm *GoMail) Send() error {}
```

[github.com/barryyan/daily-warm/api](https://godoc.org/github.com/BarryYan/daily-warm/api)
```go
// import "github.com/barryyan/daily-warm/api"
package api 

func CreateClient() *http.Client
func Fetch(url string) io.Reader
func FetchHTML(url string) *query.Document
type English struct{
  ImgURL   string
  Sentence string
}
func GetEnglish() English
type One struct{
  Date     string
  ImgURL   string
  Sentence string
}
func GetONE() One
type Poem struct{
  Title   string   `json:"title"`
  Dynasty string   `json:"dynasty"`
  Author  string   `json:"author"`
  Content []string `json:"content"`
}
func GetPoem() Poem
type PoemRes struct{
  Status string `json:"status"`
  Data   struct {
    Origin Poem `json:"origin"`
  } `json:"data"`
}
type Weather struct{
  City     string
  Temp     string
  Weather  string
  Air      string
  Humidity string
  Wind     string
  Limit    string
  Note     string
}
func GetWeather(local string) Weather
```
