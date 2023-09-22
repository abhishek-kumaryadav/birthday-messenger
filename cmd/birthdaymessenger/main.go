package main

import (
	"birthdaymessenger/pkg/fileutil"
	"birthdaymessenger/pkg/google"
	"birthdaymessenger/pkg/whatsapp"
	"fmt"
	"log"
	"time"

	"github.com/go-co-op/gocron"
)

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
		whatsapp.SendMessage(phoneNumber, whatsappTemplate, whatsappPhoneNumberId, whatsappAuthToken)
	}
}

func main() {
	props, err := fileutil.ReadPropertiesFile("birthday-messenger.properties")
	if err != nil {
		fmt.Println("Error while reading properties file")
	}

	s := gocron.NewScheduler(time.UTC)
	_, err = s.Every(props["job.frequency"]).Seconds().Do(func() {
		process(props["whatsapp.template"], props["whatsapp.phoneNumberId"], props["whatsapp.authToken"])
	})
	if err != nil {
		fmt.Println("Error executing job", err)
	}
	s.StartBlocking()
}
