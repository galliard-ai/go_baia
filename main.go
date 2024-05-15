package main

import (
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
	req = openai.ChatCompletionRequest{
		Model: openai.GPT4o,
		Messages: []openai.ChatCompletionMessage{
			{
				Role: openai.ChatMessageRoleSystem,
				Content: `Eres un útil asistente de un restaurante diseñado para leer pedidos, compararlos con el menú "
							  y generar el pedido en formato JSON, asegúrate de que cada platillo de la orden del cliente
							  tenga los campos 'id', 'nombre_platillo', 'precio_por_cada_uno' y 'cantidad', debes devolver
							  un JSON con el siguiente formato: ` + string(jsonOrdersData) + ` si el usuario no ordena nada,
							  regresa el JSON vacío. Cuando vayas a mandar el menu, asegurate de que en tu respuesta solo venga
							  el json y ninguna palabra mas ni menos, tampoco lo formates, es decir no pongas backticks ni nada, solo el json.
							  En el caso de que te hagan una pregunta si responde normalmente, pero si vas a mandar la orden, mandas la orden solita. Menu: ` + string(jsonMenuData),
			},
		},
	}

	listenMsgs()

}
