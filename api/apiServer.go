package baiaAPI

import (
	"baia_service/utils"
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

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

func RegisterEndPoints(api huma.API) {

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
		}
	}) (*GPTResponse, error) {

		response := GPTResponse{}

		response.Body.Answer = utils.SendRequest(input.Body.Question)
		fmt.Println(input.Body.Question)
		return &response, nil
	})

	huma.Register(api, huma.Operation{
		OperationID:   "audio",
		Method:        http.MethodPost,
		Path:          "/baia/askGPT/audio/",
		Summary:       "Answers about your order sending the audio file",
		Tags:          []string{"BAIA"},
		DefaultStatus: http.StatusCreated,
	}, func(ctx context.Context, input *struct {
		RawBody multipart.Form
	}) (*UploadResponse, error) {
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

		dst, err := os.Create(filepath.Join("audios", fileHeaders[0].Filename))
		if err != nil {
			return nil, huma.NewError(http.StatusInternalServerError, "Error creating file on server", err)
		}
		defer dst.Close()

		_, err = io.Copy(dst, file)
		if err != nil {
			return nil, huma.NewError(http.StatusInternalServerError, "Error saving file to server", err)
		}
		response := UploadResponse{}
		response.Body.Message = "Succesfull"

		return &response, nil
	})

}
