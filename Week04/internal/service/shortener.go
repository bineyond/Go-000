// solmyr

package service

import (
	"context"
	pb "Week04/api/url"
)

type ShortenService struct {
	pb.UnimplementedURLServiceServer
}

func NewShortenService() *ShortenService {
	return &ShortenService{}
}

func (s *ShortenService) Shorten(ctx context.Context, request *pb.ShortenRequest) (*pb.ShortenResponse, error) {
	return &pb.ShortenResponse{Shorten: request.GetOrigin()}, nil
}