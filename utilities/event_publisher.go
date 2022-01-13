package utilities

import (
	"golang.org/x/net/context"
	"encoding/json"
	"cloud.google.com/go/pubsub"
	"log"
	"fmt"
)

const (
	projectID = "otp-service-337711"
	topicID = "verification"
)

type EventData struct {
	ID string `json:"id"`
    Name string `json:"name"`
    PhoneNumber string `json:"phone_number"`
	Otp string `json:"otp"`
}

func Publish(ctx context.Context, data EventData, otp string) (error) {
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		log.Print(err.Error())
		return err
	}

	topic := client.Topic(topicID)
	_, err = topic.Exists(ctx)
	if err != nil {
		log.Print(err.Error())
		return err
	}
	event, _ := json.Marshal(data)
	// publish message to Cloud Pub/Sub
	r := topic.Publish(ctx, &pubsub.Message{
		Data: event,
	})
	id, err := r.Get(ctx)
	if err != nil {
		log.Print(err.Error())
		return err
	}
	fmt.Printf("Published a message with a message ID: %s\n", id)
	return nil
}
