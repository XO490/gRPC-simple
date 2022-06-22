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

func (s *Server) GetBy(_ context.Context, in *pb.GetByRequest) (*pb.GetByResponse, error) {
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

	return nil, errors.New("not found")
}

func (s *Server) GetAllById(_ context.Context, _ *pb.GetByRequest) (*pb.GetAllByIdResponse, error) {
	log.Printf("GetAllById")

	p, err := GetAllProducts("id", order.asc)
	CheckErr(err, "GetAllById")

	items := &pb.GetAllByIdResponse{}
	for _, i := range p {
		items.Result = append(items.Result,
			&pb.GetByResponse{
				Id:          i.Id,
				Name:        i.Name,
				CreatedAt:   timestamppb.New(i.CreatedAt).AsTime().Unix(),
				Price:       i.Price,
				Description: i.Description,
			})
	}

	return items, nil
}

func (s *Server) GetAllByIdDesc(_ context.Context, _ *pb.GetByRequest) (*pb.GetAllByIdResponse, error) {
	log.Printf("GetAllByIdDesc")

	p, err := GetAllProducts("id", order.desc)
	CheckErr(err, "GetAllByIdDesc")

	items := &pb.GetAllByIdResponse{}
	for _, i := range p {
		items.Result = append(items.Result,
			&pb.GetByResponse{
				Id:          i.Id,
				Name:        i.Name,
				CreatedAt:   timestamppb.New(i.CreatedAt).AsTime().Unix(),
				Price:       i.Price,
				Description: i.Description,
			})
	}

	return items, nil
}

func (s *Server) GetAllByName(_ context.Context, _ *pb.GetByRequest) (*pb.GetAllByIdResponse, error) {
	log.Printf("GetAllByName")

	p, err := GetAllProducts("name", order.asc)
	CheckErr(err, "GetAllByName")

	items := &pb.GetAllByIdResponse{}
	for _, i := range p {
		items.Result = append(items.Result,
			&pb.GetByResponse{
				Id:          i.Id,
				Name:        i.Name,
				CreatedAt:   timestamppb.New(i.CreatedAt).AsTime().Unix(),
				Price:       i.Price,
				Description: i.Description,
			})
	}

	return items, nil
}

func (s *Server) GetAllByNameDesc(_ context.Context, _ *pb.GetByRequest) (*pb.GetAllByIdResponse, error) {
	log.Printf("GetAllByNameDesc")

	p, err := GetAllProducts("name", order.desc)
	CheckErr(err, "GetAllByNameDesc")

	items := &pb.GetAllByIdResponse{}
	for _, i := range p {
		items.Result = append(items.Result,
			&pb.GetByResponse{
				Id:          i.Id,
				Name:        i.Name,
				CreatedAt:   timestamppb.New(i.CreatedAt).AsTime().Unix(),
				Price:       i.Price,
				Description: i.Description,
			})
	}

	return items, nil
}

func (s *Server) GetAllByCreated(_ context.Context, _ *pb.GetByRequest) (*pb.GetAllByIdResponse, error) {
	log.Printf("GetAllByCreated")

	p, err := GetAllProducts("created_at", order.asc)
	CheckErr(err, "GetAllByCreated")

	items := &pb.GetAllByIdResponse{}
	for _, i := range p {
		items.Result = append(items.Result,
			&pb.GetByResponse{
				Id:          i.Id,
				Name:        i.Name,
				CreatedAt:   timestamppb.New(i.CreatedAt).AsTime().Unix(),
				Price:       i.Price,
				Description: i.Description,
			})
	}

	return items, nil
}

func (s *Server) GetAllByCreatedDesc(_ context.Context, _ *pb.GetByRequest) (*pb.GetAllByIdResponse, error) {
	log.Printf("GetAllByCreatedDesc")

	p, err := GetAllProducts("created_at", order.desc)
	CheckErr(err, "GetAllByCreatedDesc")

	items := &pb.GetAllByIdResponse{}
	for _, i := range p {
		items.Result = append(items.Result,
			&pb.GetByResponse{
				Id:          i.Id,
				Name:        i.Name,
				CreatedAt:   timestamppb.New(i.CreatedAt).AsTime().Unix(),
				Price:       i.Price,
				Description: i.Description,
			})
	}

	return items, nil
}
