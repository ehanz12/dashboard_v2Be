package services

import (
	"context"
	"log"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"
	"google.golang.org/api/option"
)

var FCMClient *messaging.Client

func InitFirebase() {
	opt := option.WithCredentialsFile("firebase-service-account.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initializing Firebase: %v", err)
	}
	FCMClient, err = app.Messaging(context.Background())
	if err != nil {
		log.Fatalf("error initializing FCM client: %v", err)
	}
}

func SendPushNotification(
	token string,
	title string,
	body string,
) error {

	message := &messaging.Message{
		Token: token,
		Notification: &messaging.Notification{
			Title: title,
			Body:  body,
		},
	}

	_, err := FCMClient.Send(
		context.Background(),
		message,
	)

	return err
}
