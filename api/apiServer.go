package baiaAPI

import (
	firebaseService "baia_service/firebase"
	myOpenAi "baia_service/openai"
	"baia_service/utils"
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"cloud.google.com/go/firestore"
	"github.com/danielgtaylor/huma/v2"
)

type GPTResponse struct {
	Body struct {
		Answer string `json:answer`
	}
}

type UploadResponse struct {
	Body struct {
		Message string `json:"message"`
	}
}

type uploadFileRequest struct {
	Body struct {
		FileName string `json:"filename"`
	}
	RawBody multipart.Form
}

func RegisterEndPoints(api huma.API, fbClient *firestore.Client) {

	huma.Register(api, huma.Operation{
		OperationID:   "ask-about-order",
		Method:        http.MethodPost,
		Path:          "/baia/askGPT/text/{question}",
		Summary:       "Answers about your order",
		Tags:          []string{"BAIA"},
		DefaultStatus: http.StatusCreated,
	}, func(ctx context.Context, input *struct {
		Body *struct {
			Question string `json:"question" example:"Hola"`
			User     string `json:"senderID" example:"5212223201384@c.us"`
		}
	}) (*GPTResponse, error) {
		firebaseService.SaveUserMessage(input.Body.Question, input.Body.User, fbClient)

		response := GPTResponse{}
		answer := utils.SendRequest(input.Body.Question)
		response.Body.Answer = answer

		firebaseService.SaveBAIAMessage(answer, input.Body.User, fbClient)
		fmt.Println("********** MESSAGE **********")
		fmt.Println(input.Body.User)
		fmt.Println(input.Body.Question)
		return &response, nil
	})

	huma.Register(api, huma.Operation{
		OperationID:   "audio",
		Method:        http.MethodPost,
		Path:          "/baia/askGPT/audio/",
		Summary:       "Answers about your order sending the audio file",
		Tags:          []string{"BAIAudio"},
		DefaultStatus: http.StatusCreated,
	}, func(ctx context.Context, input *struct {
		RawBody multipart.Form
	}) (*GPTResponse, error) {
		if err := os.MkdirAll("audios", os.ModePerm); err != nil {
			return nil, huma.NewError(http.StatusInternalServerError, "Error creating 'audios' directory", err)
		}

		fileHeaders, ok := input.RawBody.File["audio"]
		if !ok || len(fileHeaders) == 0 {
			return nil, huma.NewError(http.StatusBadRequest, "No audio file uploaded")
		}

		file, err := fileHeaders[0].Open()
		if err != nil {
			return nil, huma.NewError(http.StatusBadRequest, "Error opening uploaded file", err)
		}
		defer file.Close()

		dst, err := os.Create(filepath.Join("audios/apiAudios", fileHeaders[0].Filename))
		if err != nil {
			return nil, huma.NewError(http.StatusInternalServerError, "Error creating file on server", err)
		}
		defer dst.Close()

		_, err = io.Copy(dst, file)
		if err != nil {
			return nil, huma.NewError(http.StatusInternalServerError, "Error saving file to server", err)
		}

		audioPath := "audios/apiAudios/" + fileHeaders[0].Filename
		translatedText := myOpenAi.Speech_to_text(audioPath)
		answerFromGPT := myOpenAi.AskGpt(translatedText)
		formatedAnswer := utils.FormatGPTResponse(answerFromGPT)
		response := GPTResponse{}
		response.Body.Answer = formatedAnswer

		return &response, nil
	})

}
