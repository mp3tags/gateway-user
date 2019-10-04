package service

import (
	"context"
	"github.com/go-openapi/runtime/middleware"
	"github.com/golang/protobuf/ptypes"
	"github.com/mp3tags/request-repository-proto"
	"io/ioutil"
	"log"
	"mp3tags/gateway-user/models"
	"mp3tags/gateway-user/restapi/operations"
	"time"
)

func UploadFile(s *Service) operations.PostAPIV1UploadHandlerFunc {
	return func(params operations.PostAPIV1UploadParams) middleware.Responder {
		result := new(operations.PostAPIV1UploadOK)

		response, err := s.uploadFile(params)
		if err != nil {
			log.Println("UploadFile error: " + err.Error())
		}

		return result.WithPayload(response)
	}
}

func (s *Service) uploadFile(params operations.PostAPIV1UploadParams) (*models.UploadResponse, error) {
	data, err := ioutil.ReadAll(params.File)
	if err != nil {
		return nil, err
	}

	createdAt, err := ptypes.TimestampProto(time.Now())
	if err != nil {
		return nil, err
	}

	// save request
	createRequest := request_repository.CreateRequestParams{
		CreatedAt: createdAt,
		UserUuid:  "",
		UserIp:    "",
		Url:       params.HTTPRequest.URL.RequestURI(),
		Data:      string(data),
	}

	_, err = s.RequestRepository.CreateRequest(context.Background(), &createRequest)
	if err != nil {
		return nil, err
	}

	// save file

	// read file
	return nil, nil
}
