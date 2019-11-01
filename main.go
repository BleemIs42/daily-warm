package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"
	"sync"
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

func isDev() bool {
	return os.Getenv("MAIL_MODE") == "dev"
}

func main() {
	loadConfig()

	if isDev() {
		batchSendMail()
		return
	}

	nyc, _ := time.LoadLocation("Asia/Shanghai")
	cJob := cron.New(cron.WithLocation(nyc))

	cronCfg := os.Getenv("MAIL_CRON")
	if cronCfg == "" {
		batchSendMail()
	} else {
		cJob.AddFunc(cronCfg, func() {
			batchSendMail()
		})
		cJob.Start()
		select {}
	}
}

func loadConfig() {
	err := env.Load()
	if err != nil {
		log.Fatalf("Load .env file error: %s", err)
	}
}

// TODO: refactor
func getParts() map[string]interface{} {
	wrapMap := map[string]func() interface{}{
		"one":       func() interface{} { return api.GetONE() },
		"english":   func() interface{} { return api.GetEnglish() },
		"poem":      func() interface{} { return api.GetPoem() },
		"wallpaper": func() interface{} { return api.GetWallpaper() },
		"trivia":    func() interface{} { return api.GetTrivia() },
	}

	wg := sync.WaitGroup{}
	parts := map[string]interface{}{}
	for name, getPart := range wrapMap {
		wg.Add(1)
		go func(key string, fn func() interface{}) {
			defer wg.Done()
			parts[key] = fn()
		}(name, getPart)
	}
	wg.Wait()
	return parts
}

func batchSendMail() {
	loadConfig()

	users := getUsers("MAIL_TO")
	if len(users) == 0 {
		return
	}

	parts := getParts()

	if isDev() {
		weather := api.GetWeather(users[0].Local)
		parts["weather"] = weather
		html := generateHTML(HTML, parts)
		fmt.Println(html)
		return
	}

	wg := sync.WaitGroup{}
	lock := sync.Mutex{}
	for _, user := range users {
		wg.Add(1)
		go func(user User) {
			defer wg.Done()
			weather := api.GetWeather(user.Local)
			lock.Lock()
			parts["weather"] = weather
			html := generateHTML(HTML, parts)
			lock.Unlock()
			sendMail(html, user.Email)
		}(user)
	}
	wg.Wait()
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
