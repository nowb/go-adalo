package main

import (
	"fmt"
	"log"

	"github.com/nowb/go-adalo"
)

func main() {
	client, _ := adalo.NewClient("s7f98jo32jlseznml3n")

	if output, err := client.SendNotification(&adalo.SendNotificationInput{
		AppID:                 "2bc38ea8-2eb1-4db0-974c-e9e031f2c0e0",
		AudienceEmail:         "user@example.com",
		NotificationTitleText: "Hello There",
		NotificationBodyText:  "This is just a test...",
	}); err != nil {
		log.Fatal(fmt.Errorf("sending notification failed: %w", err).Error())
	} else {
		log.Printf("successfully sent notification, output: %v\n", output)
	}
}
