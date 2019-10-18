# Daily-warm
> 每天定时发邮件给你关心的人

## Usage

```bash
# 1. Config `.env` file in root dir, e.g.
MAIL_USERNAME = user@qq.com
MAIL_PASSWORD = ***********
MAIL_SUBJECT = 每日一暖, 温情一生
MAIL_CRON = "0/1 * * * *"
MAIL_FROM = 天心<790956404@qq.com>
MAIL_TO = [{"email": "user<user@qq.com>", "local": "shaanxi/xian"}]

# 2. Run
sh daily-warm.sh
```

## Build

```bash
go mod download
go run *.go
```

## Package

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

```go
// import "github.com/barryyan/daily-warm/api"
package api 

func CreateClient() *http.Client
func FetchDom(url string) *query.Document

type One struct{ 
  Date   string
  ImgURL string
  Word   string
}
func GetONE() One

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
