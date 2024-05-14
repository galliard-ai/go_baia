package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	router := gin.Default()

	router.POST("/", func(context *gin.Context) {
		// Aquí es donde procesarías el mensaje recibido de WhatsApp
		// Puedes acceder al cuerpo del mensaje en context.Request.Body
		body, err := io.ReadAll(context.Request.Body)
		if err != nil {
			context.String(http.StatusInternalServerError, err.Error())
			return
		}
		values, err := url.ParseQuery(string(body))
		if err != nil {
			context.String(http.StatusInternalServerError, err.Error())
			return
		}

		// Extraer el valor del cuerpo del mensaje
		messageBody := values.Get("Body")
		answer := AskGpt(messageBody)
		fmt.Println("MENSAJE: " + messageBody)
		// Procesar el cuerpo del mensaje
		fmt.Println("Mensaje de WhatsApp recibido:", string(body))
		// Responder a la solicitud de ngrok
		context.String(http.StatusOK, answer)
	})

	// Iniciar el servidor en el puerto 3000
	if err := router.Run(":8000"); err != nil {
		log.Fatal(err)
	}
	//fmt.Println(AskGpt("Con quién hablo?"))
	//whatsAppTwilio("HOLA armandoOOO")

}
