package main

import (
	"ethereumLink/ethereumLink"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"
	pb "kunmel.com/sgx-multichain/multichain_proto"
	"log"
	"net"
)

type MultichainServiceServer struct {
	pb.MultichainServiceServer
}

func NewServer() *MultichainServiceServer {
	return &MultichainServiceServer{
	}
}
func main() {
	lis,err := net.Listen("tcp", "10.108.16.218:3333")
	if err != nil {
		log.Fatalln("cannot create a listener at the address")
	}
	creds, err := credentials.NewServerTLSFromFile("../key/test.pem", "../key/test.key")
	if err != nil {
		grpclog.Fatalf("Failed to generate credentials %v", err)
	}
	grpcServer := grpc.NewServer(grpc.Creds(creds))
	pb.RegisterMultichainServiceServer(grpcServer, NewServer())
	fmt.Println("connect created")
	log.Fatalln(grpcServer.Serve(lis))
}

func (s * MultichainServiceServer) NewEthTx (request *pb.NewEthRequest, stream pb.MultichainService_NewEthTxServer) error {
	statusCode := ethereumLink.EthereumLink(request.KeystorePath, request.ChainID, request.Pass, request.ContractAddress, request.FcnName, request.ContractArg)
	response := pb.CommitResponse{StatusCode: statusCode}
	err := stream.Send(&response)
	if err != nil {
		return err
	}
	return nil
}