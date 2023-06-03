package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/gravitalia/spinoza/helpers"
	"github.com/gravitalia/spinoza/proto"
	"github.com/gravitalia/spinoza/uploader"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

type server struct {
	proto.UnimplementedSpinozaServiceServer
}

func (s *server) Upload(ctx context.Context, in *proto.UploadRequest) (*proto.UploadReply, error) {
	start := time.Now().UnixMilli()
	img, err := helpers.Compress(in.GetData(), in.GetWidth(), in.GetHeight())
	if err != nil {
		return &proto.UploadReply{Message: err.Error(), Error: true}, nil
	}
	fmt.Println(time.Now().UnixMilli() - start)

	id, err := uploader.UploadOnCloudinary(img)
	if err != nil {
		return &proto.UploadReply{Message: "Error while uploading file", Error: true}, nil
	}

	return &proto.UploadReply{Message: id, Error: false}, nil
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	lis, err := net.Listen("tcp", ":"+os.Getenv("PORT"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("Listening on port :%s\n", os.Getenv("PORT"))

	var opts []grpc.ServerOption
	maxMsgSize := 40 * 1024 * 1024
	opts = append(opts, grpc.MaxRecvMsgSize(maxMsgSize), grpc.MaxSendMsgSize(maxMsgSize))

	grpcServer := grpc.NewServer(opts...)
	proto.RegisterSpinozaServiceServer(grpcServer, &server{})
	grpcServer.Serve(lis)
}
