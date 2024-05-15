package main

import (
	"fmt"

	"io/ioutil"

	"github.com/joho/godotenv"
	"github.com/sashabaranov/go-openai"
)

func main() {

	// texto := `Lo siento, pero en nuestro menú no tenemos "caldo de poblano". ¿Podrías verificar el nombre del platillo que deseas? Tenemos una "Crema de Poblano".

	// De acuerdo a tu pedido preliminar, aquí está tu orden en formato JSON con los elementos disponibles en el menú:

	// json
	// {
	// 	"orden": [
	// 		{
	// 			"id": 5,
	// 			"nombre_platillo": "Crema de Poblano",
	// 			"precio_por_cada_uno": 18.00,
	// 			"cantidad": 1
	// 		},
	// 		{
	// 			"id": 14,
	// 			"nombre_platillo": "Horchata",
	// 			"precio_por_cada_uno": 12.00,
	// 			"cantidad": 2
	// 		}
	// 	]
	// }

	// Por favor, revisa y confirma si deseas estos platillos u otros.`
	// fmt.Println(strings.Replace(formatGPTResponse(texto), "}", "", -1))

	//************************************************************************
	godotenv.Load()

	jsonMenuData, err := ioutil.ReadFile("jsons/menu.json")
	if err != nil {
		fmt.Println("Error at parsing menu json")
	}

	jsonOrdersData, err := ioutil.ReadFile("jsons/orders/order.json")
	if err != nil {
		fmt.Println("Error at parsing order json")
	}
	req = openai.ChatCompletionRequest{
		Model: openai.GPT4o,
		Messages: []openai.ChatCompletionMessage{
			{
				Role: openai.ChatMessageRoleSystem,
				Content: `Eres un útil asistente de un restaurante diseñado para leer pedidos, compararlos con el menú "
						  y generar el pedido en formato JSON, asegúrate de que cada platillo de la orden del cliente
						  tenga los campos 'id', 'nombre_platillo', 'precio_por_cada_uno' y 'cantidad', debes devolver
						  un JSON con el siguiente formato: ` + string(jsonOrdersData) + ` si el usuario no ordena nada,
						  regresa el JSON vacío. Menu: ` + string(jsonMenuData) + `Se muy amigable, recuerda que nos puedes
						  ayudar a conseguir mas clientes si les caes bien`,
			},
		},
	}

	listenMsgs()

}
