package main

import (
	myOpenAi "baia_service/openai"
	twilioService "baia_service/twilio"
	"fmt"
	"io/ioutil"

	"github.com/joho/godotenv"
	"github.com/sashabaranov/go-openai"
)

func main() {
	godotenv.Load()

	jsonMenuData, err := ioutil.ReadFile("jsons/menu.json")
	if err != nil {
		fmt.Println("Error at parsing menu json")
	}

	jsonOrdersData, err := ioutil.ReadFile("jsons/orders/order.json")
	if err != nil {
		fmt.Println("Error at parsing order json")
	}
	myOpenAi.Req = openai.ChatCompletionRequest{
		Model: openai.GPT4o,
		Messages: []openai.ChatCompletionMessage{
			{
				Role: openai.ChatMessageRoleSystem,
				Content: `Eres un útil asistente de un restaurante diseñado para leer pedidos, compararlos con el menú "
						  y generar el pedido en formato JSON, asegúrate de que cada platillo de la orden del cliente
						  tenga los campos 'id', 'nombre_platillo', 'precio_por_cada_uno' y 'cantidad', debes devolver
						  un JSON con el siguiente formato: ` + string(jsonOrdersData) + ` si el usuario no ordena nada,
						  regresa el JSON vacío. Menu: ` + string(jsonMenuData) + `Se muy amigable, recuerda que nos puedes
						  ayudar a conseguir mas clientes si les caes bien, y no pongas tanto texto, se amable pero conciso
						  al mismo tiempo.`,
			},
		},
	}

	twilioService.ListenMsgs()

}
