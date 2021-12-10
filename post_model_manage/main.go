package main

import (
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	post_model_manage "post_model_manage/idl"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)


func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	reflection.Register(s)
	post_model_manage.RegisterPostModelManageServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
