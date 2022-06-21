package server

import (
	"context"
	"errors"
	"google.golang.org/protobuf/types/known/timestamppb"
	pb "grpc-simple/proto"
	"log"
	"time"
)

type Server struct {
	pb.ProductServer
}

func (s *Server) Post(ctx context.Context, in *pb.GetByRequest) (*pb.GetByResponse, error) {
	return nil, nil
}

func (s *Server) GetBy(ctx context.Context, in *pb.GetByRequest) (*pb.GetByResponse, error) {
	log.Printf("GetBy %s request", in)

	var (
		p   *product
		err error
	)

	switch {
	case in.GetId() != "":
		p, err = GetProduct("id", in.GetId())
		CheckErr(err, "GetBy")

	case in.GetName() != "":
		p, err = GetProduct("name", in.GetName())
		CheckErr(err, "GetByName")

	case in.GetCreatedAt() != 0:
		ts := time.Unix(in.GetCreatedAt(), 0).Format("2006-01-02 15:04:05.000000")
		//fmt.Println("ts:", ts)
		p, err = GetProduct("created_at", ts)
		CheckErr(err, "GetByCreatedAt")
	}

	pbts := timestamppb.New(p.CreatedAt).AsTime().Unix()
	if pbts > -62135596800 {
		return &pb.GetByResponse{
			Id:          p.Id,
			Name:        p.Name,
			CreatedAt:   pbts,
			Price:       p.Price,
			Description: p.Description,
		}, nil
	}

	return nil, errors.New("what you mean?")
}

//func (s *Server) GetAllById(ctx context.Context, _ *emptypb.Empty) ([]*pb.GetAllByIdResponse, error) {
//	log.Printf("GetAllById")
//
//	p, err := GetAllProducts("id", order.asc)
//	CheckErr(err, "GetAllById")
//
//	var arr []*pb.GetAllByIdResponse
//	for _, i := range p {
//		arr = append(arr,
//			&pb.GetAllByIdResponse{
//				Id:          i.Id,
//				Name:        i.Name,
//				CreatedAt:   timestamppb.New(i.CreatedAt).AsTime().Unix(),
//				Price:       i.Price,
//				Description: i.Description,
//			})
//	}
//
//	return arr, nil
//}
