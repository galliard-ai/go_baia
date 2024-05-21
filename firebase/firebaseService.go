package firebaseService

import (
	"context"
	"fmt"
	"log"

	firebase "firebase.google.com/go/v4"
)

func SendTest() {
	ctx := context.Background()
	config := &firebase.Config{
		ProjectID: "baia-afb63",
	}
	app, err := firebase.NewApp(ctx, config)
	if err != nil {
		log.Fatalln(err)
	}
	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	ref := client.Collection("todooooos").Doc("cccola")
	result, err := ref.Set(ctx, map[string]interface{}{
		"title": "hola",
	})

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(result)
}
