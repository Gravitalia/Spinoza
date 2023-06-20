package main

import (
	"context"
	"log"
	"net"
	"os"

	"github.com/gravitalia/spinoza/helpers"
	"github.com/gravitalia/spinoza/proto"
	"github.com/gravitalia/spinoza/uploader"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

type server struct {
	proto.UnimplementedSpinozaServer
}

func (s *server) Upload(ctx context.Context, in *proto.UploadRequest) (*proto.UploadReply, error) {
	img, err := helpers.Compress(in.GetData(), in.GetWidth(), in.GetHeight())
	if err != nil {
		return &proto.UploadReply{Message: err.Error(), Error: true}, nil
	}

	id, err := uploader.UploadOnCloudinary(img)
	if err != nil {
		return &proto.UploadReply{Message: "Error while uploading file", Error: true}, nil
	}

	return &proto.UploadReply{Message: id, Error: false}, nil
}

func main() {
	// Load .env file
	godotenv.Load()

	lis, err := net.Listen("tcp", ":"+os.Getenv("PORT"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("Listening on port :%s\n", os.Getenv("PORT"))

	var opts []grpc.ServerOption
	maxMsgSize := 20 * 1024 * 1024
	opts = append(opts, grpc.MaxRecvMsgSize(maxMsgSize), grpc.MaxSendMsgSize(maxMsgSize))

	grpcServer := grpc.NewServer(opts...)
	proto.RegisterSpinozaServer(grpcServer, &server{})
	grpcServer.Serve(lis)
}
