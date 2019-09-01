package main

import (
	"flag"
	pb "locust-grpc/protos"
	"log"
	"strconv"
	"time"

	"github.com/myzhan/boomer"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var verbose bool
var serverAddress string
var conn *grpc.ClientConn

func NewConn() *grpc.ClientConn {
	conn, err := grpc.Dial(serverAddress, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	return conn
}

func sendTransaction() (*pb.ResponseStatus, error) {
	c := pb.NewOteClient(conn)
	ctx := context.Background()
	body := &pb.SendTransactionRequest{}

	return c.SendTransaction(ctx, body)
}

func worker() {
	startTime := time.Now()
	response, err := sendTransaction()
	elapsed := time.Since(startTime)

	if err != nil {
		if verbose {
			log.Printf("%v\n", err)
		}
		boomer.RecordFailure("grpc", "error", 0.0, err.Error())
	} else {
		boomer.RecordSuccess("grpc", strconv.Itoa(int(response.Status)),
			elapsed.Nanoseconds()/int64(time.Millisecond), 10)
	}
}

func main() {
	defer func() {
		_ = conn.Close()
	}()

	flag.StringVar(&serverAddress, "url", "localhost:8080", "URL")
	flag.BoolVar(&verbose, "verbose", false, "Print debug log")
	flag.Parse()

	conn = NewConn()

	task := &boomer.Task{
		Name:   "worker",
		Weight: 10,
		Fn:     worker,
	}

	boomer.Run(task)
}
