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
		log.Fatal("Error loading .env file, ", err)
	}
}

func batchSendMail() {
	loadConfig()
	one := api.GetONE()

	users := getUsers("MAIL_TO")
	if len(users) == 0 {
		return
	}

	res := make(chan int)
	defer close(res)

	for _, user := range users {
		weather := api.GetWeather(user.Local)
		datas := map[string]interface{}{"one": one, "weather": weather}
		html := generateHTML(HTML, datas)
		go func(email string) {
			sendMail(html, email)
			<-res
		}(user.Email)
		res <- 1
	}
}

func getUsers(envKey string) []User {
	var users []User
	err := json.Unmarshal([]byte(os.Getenv(envKey)), &users)
	if err != nil {
		log.Fatal(err)
	}
	return users
}

func generateHTML(html string, datas map[string]interface{}) string {
	for key, data := range datas {
		rKey := reflect.TypeOf(data)
		rVal := reflect.ValueOf(data)
		fieldNum := rKey.NumField()
		for i := 0; i < fieldNum; i++ {
			field := rKey.Field(i).Name
			value := rVal.Field(i).String()
			mark := fmt.Sprintf("{{%s.%s}}", key, field)
			html = strings.ReplaceAll(html, mark, value)
		}
	}
	return html
}

func sendMail(content string, to string) {
	gomail.Config.Username = os.Getenv("MAIL_USERNAME")
	gomail.Config.Password = os.Getenv("MAIL_PASSWORD")
	gomail.Config.From = os.Getenv("MAIL_FROM")
	gomail.Config.Host = os.Getenv("MAIL_HOST")
	gomail.Config.Port = os.Getenv("MAIL_PORT")

	email := gomail.GoMail{
		To:      []string{to},
		Subject: os.Getenv("MAIL_SUBJECT"),
		Content: content,
	}

	err := email.Send()
	if err != nil {
		fmt.Println(err, "\nSend email fail!")
	} else {
		fmt.Println("Send email " + to + " success!")
	}
}
