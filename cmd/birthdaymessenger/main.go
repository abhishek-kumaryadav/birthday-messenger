package main

import (
	"birthdaymessenger/models/properties"
	"birthdaymessenger/pkg/google"
	"birthdaymessenger/pkg/whatsapp"
	"fmt"
	"log"
	"time"

	"github.com/go-co-op/gocron"
)

func process() {
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
		whatsapp.SendMessage(phoneNumber)
	}
}

func main() {

	s := gocron.NewScheduler(time.UTC)
	fmt.Println(properties.GetJobFrequency())
	_, err := s.Every(properties.GetJobFrequency() + "s").Do(func() { process() })
	if err != nil {
		fmt.Println("Error executing job", err)
	}
	s.StartBlocking()
}
