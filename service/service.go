package service

import "github.com/mp3tags/request-repository-proto"

type Service struct {
	RequestRepository request_repository.RequestRepositoryServiceClient
}
