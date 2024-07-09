package firebaseService

import (
	"context"
	"fmt"
	"log"
	"strings"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/iterator"
)

func InitFirebase() (*firestore.Client, error) {
	ctx := context.Background()
	config := &firebase.Config{
		ProjectID: "baia-1df5a",
	}
	app, err := firebase.NewApp(ctx, config)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	return client, nil
}

func getUserMessageNumber(client *firestore.Client, senderID string, isFromUser bool) int {
	ctx := context.Background()
	var messageType string
	if isFromUser {
		messageType = "message"
	} else {
		messageType = "response"
	}
	collectionRef := client.Collection("El Sabor de Tlaxcala").Doc("conversations").Collection(senderID).Doc("messages").Collection("Conversation")
	iter := collectionRef.Documents(ctx)
	messageCount := 0
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
		docID := doc.Ref.ID
		if strings.Contains(docID, messageType) {
			messageCount++

		}
	}
	fmt.Printf("%v number %v", messageType, messageCount+1)
	return messageCount + 1
}
func SaveBAIAMessage(message string, senderID string, client *firestore.Client) {
	ctx := context.Background()
	messageNumber := getUserMessageNumber(client, senderID, false)
	ref1 := client.Collection("El Sabor de Tlaxcala").Doc("conversations").Collection(senderID).Doc("messages")
	ref := ref1.Collection("Conversation").Doc(fmt.Sprintf("#%v response", messageNumber))
	result, err := ref.Set(ctx, map[string]interface{}{
		"Message":       message,
		"isUserMessage": false,
	})
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(result)
}

func SaveUserMessage(message string, senderID string, client *firestore.Client) {
	ctx := context.Background()
	messageNumber := getUserMessageNumber(client, senderID, true)
	ref1 := client.Collection("El Sabor de Tlaxcala").Doc("conversations").Collection(senderID).Doc("messages")
	ref := ref1.Collection("Conversation").Doc(fmt.Sprintf("#%v message", messageNumber))
	result, err := ref.Set(ctx, map[string]interface{}{
		"Message":       message,
		"isUserMessage": true,
	})
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(result)
}
