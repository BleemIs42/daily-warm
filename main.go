package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"
	"time"

	"github.com/barryyan/daily-warm/api"
	"github.com/barryyan/daily-warm/gomail"

	env "github.com/joho/godotenv"
	cron "github.com/robfig/cron/v3"
)

// User for receive email
type User struct {
	Email string `json:"email"`
	Local string `json:"local"`
}

func main() {
	loadConfig()

	nyc, _ := time.LoadLocation("Asia/Shanghai")
	cJob := cron.New(cron.WithLocation(nyc))

	cronCfg := os.Getenv("MAIL_CRON")
	if cronCfg != "" {
		cJob.AddFunc(cronCfg, func() {
			batchSendMail()
		})
		cJob.Start()
		select {}
	} else {
		batchSendMail()
	}
}

func loadConfig() {
	err := env.Load()
	if err != nil {
		log.Fatalf("Load .env file error: %s", err)
	}
}

func batchSendMail() {
	loadConfig()
	one := api.GetONE()
	english := api.GetEnglish()
	poem := api.GetPoem()

	users := getUsers("MAIL_TO")
	if len(users) == 0 {
		return
	}

	res := make(chan int)
	defer close(res)

	for _, user := range users {
		weather := api.GetWeather(user.Local)
		datas := map[string]interface{}{
			"one":     one,
			"weather": weather,
			"english": english,
			"poem":    poem,
		}
		html := generateHTML(HTML, datas)

		go func(email string) {
			sendMail(html, email)
			<-res
		}(user.Email)
		res <- 1
	}
}

func getUsers(envUser string) []User {
	var users []User
	userJSON := os.Getenv(envUser)
	err := json.Unmarshal([]byte(userJSON), &users)
	if err != nil {
		log.Fatalf("Parse users from %s error: %s", userJSON, err)
	}
	return users
}

func generateHTML(html string, datas map[string]interface{}) string {
	for key, data := range datas {
		rDataKey := reflect.TypeOf(data)
		rDataVal := reflect.ValueOf(data)
		fieldNum := rDataKey.NumField()
		for i := 0; i < fieldNum; i++ {
			fName := rDataKey.Field(i).Name
			rValue := rDataVal.Field(i)

			var fValue string
			switch rValue.Interface().(type) {
			case string:
				fValue = rValue.String()
			case []string:
				fValue = strings.Join(rValue.Interface().([]string), "<br>")
			}

			mark := fmt.Sprintf("{{%s.%s}}", key, fName)
			html = strings.ReplaceAll(html, mark, fValue)
		}
	}
	return html
}

func sendMail(content string, to string) {
	gomail.Config.Username = os.Getenv("MAIL_USERNAME")
	gomail.Config.Password = os.Getenv("MAIL_PASSWORD")
	gomail.Config.Host = os.Getenv("MAIL_HOST")
	gomail.Config.Port = os.Getenv("MAIL_PORT")
	gomail.Config.From = os.Getenv("MAIL_FROM")

	email := gomail.GoMail{
		To:      []string{to},
		Subject: os.Getenv("MAIL_SUBJECT"),
		Content: content,
	}

	err := email.Send()
	if err != nil {
		log.Printf("Send email fail, error: %s", err)
	} else {
		log.Printf("Send email %s success!", to)
	}
}
