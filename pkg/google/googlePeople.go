package google

import (
	"birthdaymessenger/pkg/fileutil"
	"context"
	"fmt"
	"log"
	"os"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
	"google.golang.org/api/people/v1"
)

func GetPhoneNumbers(name string) (string, error) {
	props, err := fileutil.ReadPropertiesFile("birthday-messenger.properties")

	ctx := context.Background()
	b, err := os.ReadFile(props["gcp.credentials.path"])
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, people.ContactsReadonlyScope, calendar.CalendarReadonlyScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	client := getClient(config)

	srv, err := people.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to create people Client %v", err)
	}

	r, err := srv.People.SearchContacts().Query(name).ReadMask("names,phoneNumbers,birthdays").Do()
	if err != nil {
		log.Fatalf("error retrieving contacts: %v", err)
	}
	for _, c := range r.Results {
		names := c.Person.Names
		if len(names) > 0 {
			personName := names[0].DisplayName
			if personName == name {
				return c.Person.PhoneNumbers[0].Value, err
			}
		}
	}

	return "", fmt.Errorf("Names did not match the birthday person's name")

}
