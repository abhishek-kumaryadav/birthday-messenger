package google

import (
	"birthdaymessenger/models"
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
	"google.golang.org/api/people/v1"
)

func GetBirthDays() ([]*models.Person, error) {
	var persons []*models.Person
	ctx := context.Background()
	b, err := os.ReadFile("credentials.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, calendar.CalendarReadonlyScope, people.ContactsReadonlyScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	client := getClient(config)

	srv, err := calendar.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to retrieve Calendar client: %v", err)
	}

	endTime := Bod(time.Now().AddDate(0, 0, 1)).Format(time.RFC3339)
	startTime := Bod(time.Now()).Format(time.RFC3339)
	events, err := srv.Events.List("addressbook#contacts@group.v.calendar.google.com").
		TimeMax(endTime).TimeMin(startTime).TimeZone("IST").Do()

	fmt.Println(events.Summary)
	// events, err := srv.Events.List("primary").TimeMin(t).MaxResults(1000).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve next ten of the user's events: %v", err)
	}

	if err != nil {
		log.Fatalf("Unable to retrieve next ten of the user's events: %v", err)
	}
	fmt.Println("Upcoming events:")
	if len(events.Items) == 0 {
		fmt.Println("No upcoming events found.")
	} else {
		for _, item := range events.Items {
			startDate := item.Start.Date
			person := models.Person{Name: strings.TrimSuffix(item.Summary, "'s birthday"), DateOfBirth: startDate}
			persons = append(persons, &person)
		}
	}
	return persons, err
}
