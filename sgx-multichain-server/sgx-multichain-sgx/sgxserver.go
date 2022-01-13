package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"
	"io"
	pb "kunmel.com/sgx-multichain/multichain_proto"
	"log"
	"math/rand"
)

// actually not a server but a client
func main() {
	fabIP := "10.108.16.218:2222"
	fabPem := "/root/key/test.pem"
	ethIP := "10.108.16.218:3333"
	ethPem := "/root/key/test.pem"
	serverName := "www.kunmellh.com"
	var client pb.MultichainServiceClient
	var response int32
	for i:=0; i<10; i++ {
		if rand.Intn(10) < 5 {
			client = setupConnect(fabIP, fabPem, serverName)
			response = sendFab(client, "mycc", []string{"a","b","1"}, "invoke")
		} else {
			client = setupConnect(ethIP, ethPem, serverName)
			response = sendEth(client, "UTC--2021-09-15T05-28-00.066630039Z--6a2e96b16eec57ba9f21b1b5c5bdc476a2d3676f", "22b8", "123456", "0x120615426fbf165d6673befb179f7faaa9f08c6a", "GetHold", "")
		}
	}
	decodeRsponse(response)
}

func setupConnect(ip string, pemPath string, serverName string)  pb.MultichainServiceClient {
	creds, err := credentials.NewClientTLSFromFile(pemPath, serverName)
	if err != nil {
		grpclog.Fatalf("Failed to create TLS credentials %v", err)
	}
	conn, err := grpc.Dial(ip, grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalln("peer grpc client cannot dial ", ip)
	}
	fmt.Println("gRPC connection created")
	client := pb.NewMultichainServiceClient(conn)
	return client
}

func sendFab(client pb.MultichainServiceClient, ccName string, ccArg []string, fcnName string) int32 {
	serverStream, err := client.NewFabTx(context.Background(), &pb.NewFabRequest{
		CCName: ccName,
		CCArg: ccArg,
		FcnName: fcnName,
	})
	if err != nil {
		log.Fatalln(err)
	}
	var result *pb.CommitResponse
	var statuscode int32
	for {
		result, err = serverStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
		statuscode = result.StatusCode
	}
	return statuscode
}

func sendEth(client pb.MultichainServiceClient, keyPath string, chainID string, pass string, contractAddress string, fcnName string, contractArg string) int32 {
	serverStream, err := client.NewEthTx(context.Background(), &pb.NewEthRequest{
		KeystorePath: keyPath,
		ChainID: chainID,
		Pass: pass,
		ContractAddress: contractAddress,
		FcnName: fcnName,
		ContractArg: contractArg,
	})
	if err != nil {
		log.Fatalln(err)
	}
	var response *pb.CommitResponse
	var statuscode int32
	for {
		response, err = serverStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
		statuscode = response.StatusCode
	}
	return statuscode
}

func decodeRsponse(resp int32) {
	if resp == 200 {
		fmt.Println("send requst success")
	} else if resp == 404 {
		fmt.Println("send request failed")
	} else {
		fmt.Println(resp)
	}
}
