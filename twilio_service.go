package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
)

func smsTwilio(message string) {
	from := os.Getenv("TWILIO_SMS_FROM")
	to := os.Getenv("TWILIO_SMS_TO")
	body := message
	var twilio_account_sid = os.Getenv("TWILIO_SID")
	var twilio_auth_token = os.Getenv("TWILIO_AUTH_TOKEN")

	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: twilio_account_sid,
		Password: twilio_auth_token,
	})

	params := &twilioApi.CreateMessageParams{
		To:   &to,
		From: &from,
		Body: &body,
	}

	resp, err := client.Api.CreateMessage(params)
	if err != nil {
		fmt.Println("Error sending SMS message: " + err.Error())
	} else {
		response, _ := json.Marshal(*resp)
		fmt.Println("Response: " + string(response))
	}

}

func whatsAppTwilio(message string) {
	var twilio_account_sid = os.Getenv("TWILIO_SID")
	var twilio_auth_token = os.Getenv("TWILIO_AUTH_TOKEN")

	twilioClient := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: twilio_account_sid,
		Password: twilio_auth_token,
	})
	from := os.Getenv("TWILIO_WA_FROM")
	to := os.Getenv("TWILIO_WA_TO")
	body := message

	params := &twilioApi.CreateMessageParams{}
	params.SetTo(to)
	params.SetFrom(from)
	params.SetBody(body)

	_, err := twilioClient.Api.CreateMessage(params)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Message sent successfully!")
	}
}
