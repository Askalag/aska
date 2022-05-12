package main

import (
	"flag"
	"fmt"
	"github.com/Askalag/aska/microservices/history/pkg/server"
	"github.com/Askalag/aska/microservices/history/pkg/service"
	status_v1 "github.com/Askalag/aska/microservices/history/proto/status/v1"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {

	hsAddr := flag.String("hs_a", "localhost", "http history server address")
	hsPort := flag.String("hs_p", "9092", "http history port address")
	flag.Parse()

	srv := service.NewService()
	grpcSrv := grpc.NewServer()
	mSrv := server.NewServer(*srv)

	status_v1.RegisterStatusServiceServer(grpcSrv, mSrv.Hs)

	l, err := net.Listen("tcp", fmt.Sprintf("%s:%s", *hsAddr, *hsPort))
	if err != nil {
		log.Fatal(err.Error())
	}

	if err := grpcSrv.Serve(l); err != nil {
		log.Fatal(err.Error())
	}
}
