package main

import (
	"bufio"
	"fmt"
	"go.etcd.io/bbolt/rpc"
	pb "go.etcd.io/bbolt/rpc/service/proto"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

func main() {
	var port string
	filePath := os.Args[1]
	f, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	// read the last line
	for scanner.Scan() {
		port = scanner.Text()
	}
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		panic(err)
	}
	fmt.Printf("Listen on port %s", port)
	s := grpc.NewServer()
	pb.RegisterStorageServiceServer(s, &rpc.Server{})
	s.Serve(lis)
}