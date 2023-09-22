package main

import (
	"birthdaymessenger/models"
	"birthdaymessenger/pkg/fileutil"
	"birthdaymessenger/pkg/google"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/go-co-op/gocron"
)

func sendMessage(phoneNumber string, whatsappTemplate string, whatsappPhoneNumberId string, whatsappAuthToken string) {
	language := models.Language{Code: "en_US"}
	template := models.Template{Name: whatsappTemplate, Language: language}
	data := models.Payload{MessagingProduct: "whatsapp", To: "+91" + strings.TrimSpace(phoneNumber), Type: "template", Template: template}
	payloadBytes, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error creating payload", err)
	}

	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", "https://graph.facebook.com/v17.0/"+whatsappPhoneNumberId+"/messages", body)
	if err != nil {
		fmt.Println("Error creating request", err)
	}
	req.Header.Set("Authorization", whatsappAuthToken)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error sending message", err)
	}
	defer resp.Body.Close()
}

func process(whatsappTemplate string, whatsappPhoneNumberId string, whatsappAuthToken string) {
	birthdays, err := google.GetBirthDays()
	if err != nil {
		log.Fatalf("Unable to get birthdays: %v", err)
	}
	for _, item := range birthdays {
		phoneNumber, err := google.GetPhoneNumbers(item.Name)
		if err != nil {
			log.Fatalf("Unable to get phone numbers: %v", err)
		}
		fmt.Println(phoneNumber)
		sendMessage(phoneNumber, whatsappTemplate, whatsappPhoneNumberId, whatsappAuthToken)
	}
}

func main() {
	props, err := fileutil.ReadPropertiesFile("birthday-messenger.properties")
	if err != nil {
		fmt.Println("Error while reading properties file")
	}

	if props["whatsapp.phoneNumberId"] == "" || props["whatsapp.template"] == "" || props["whatsapp.authToken"] == "" {
		fmt.Println("Error properties not loaded correctly")
	}

	s := gocron.NewScheduler(time.UTC)
	_, err = s.Every(5).Seconds().Do(func() {
		process(props["whatsapp.template"], props["whatsapp.phoneNumberId"], props["whatsapp.authToken"])
	})
	if err != nil {
		fmt.Println("Error executing job", err)
	}
	s.StartBlocking()
}
