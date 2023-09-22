package whatsapp

import (
	"birthdaymessenger/models"
	"birthdaymessenger/models/properties"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func SendMessage(phoneNumber string) {
	language := models.Language{Code: "en_US"}
	template := models.Template{Name: properties.GetWhatsappTemplate(), Language: language}
	data := models.Payload{MessagingProduct: "whatsapp", To: "+91" + strings.TrimSpace(phoneNumber), Type: "template", Template: template}
	payloadBytes, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error creating payload", err)
	}

	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", "https://graph.facebook.com/v17.0/"+properties.GetWhatsappPhoneNumberId()+"/messages", body)
	if err != nil {
		fmt.Println("Error creating request", err)
	}

	req.Header.Set("Authorization", properties.GetWhatsappAuthToken())
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error sending message", err)
	}
	defer resp.Body.Close()
}
