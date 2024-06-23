package main

import (
	myOpenAi "baia_service/openai"
	"fmt"
	"io/ioutil"

	"baia_service/utils"
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humachi"
	"github.com/danielgtaylor/huma/v2/humacli"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"github.com/sashabaranov/go-openai"
)

type Options struct {
	port int `help:"Port to listen on" short:"p" default:"8888"`
}

type GPTResponse struct {
	Body struct {
		Answer string `json:answer`
	}
}

type GPTRequest struct {
	Body struct {
		Question string `json:question`
	}
}

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
				Content: `Te llamas Mateo y eres un útil asistente de un restaurante diseñado para leer pedidos, compararlos con el menú "
						  y generar el pedido en formato JSON, asegúrate de que cada platillo de la orden del cliente
						  tenga los campos 'id', 'nombre_platillo', 'precio_por_cada_uno' y 'cantidad', debes devolver
						  un JSON con el siguiente formato: ` + string(jsonOrdersData) + ` si el usuario no ordena nada,
						  regresa el JSON vacío. Menu: ` + string(jsonMenuData) + `Se muy amigable, recuerda que nos puedes
						  ayudar a conseguir mas clientes si les caes bien, y no pongas tanto texto, se amable pero conciso
						  al mismo tiempo. Si te saludan en ingles, respondes todo en ingles`,
			},
		},
	}

	cli := humacli.New(func(hook humacli.Hooks, options *Options) {
		router := chi.NewMux()
		api := humachi.New(router, huma.DefaultConfig("My First API", "1.0.0"))

		hook.OnStart(func() {
			fmt.Printf("Starting server on port %d...\n", 8888)
			http.ListenAndServe(fmt.Sprintf(":%d", 8888), router)
		})

		huma.Register(api, huma.Operation{
			OperationID:   "ask-about-order",
			Method:        http.MethodPost,
			Path:          "/baia/",
			Summary:       "Answers about your order",
			Tags:          []string{"BAIA"},
			DefaultStatus: http.StatusCreated,
		}, func(ctx context.Context, input *GPTRequest) (*GPTResponse, error) {

			response := GPTResponse{}
			response.Body.Answer = utils.SendRequest(input.Body.Question)

			return &response, nil
		})
	})

	cli.Run()
}
