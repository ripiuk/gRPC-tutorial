package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"gRPC_course/calculator/calculatorpb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func main() {
	fmt.Println("Calculator client")
	cc, err := grpc.Dial("127.0.0.1:8000", grpc.WithInsecure())  // because we do not have SSL certificate
	if err != nil {
		log.Fatalf("Could not conect: %v", err)
	}
	defer cc.Close()

	c := calculatorpb.NewCalculatorServiceClient(cc)
	// fmt.Printf("Created client: %f", c)

	doUnary(c)

	doServerStreaming(c)

	doClientStreaming(c)

	doBiDirectionalStreaming(c)

	doErrorUnary(c)
}

func doUnary(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting a Sum Unary RPC...")
	req := &calculatorpb.SumRequest{
		FirstNumber: 2,
		SecondNumber: 2,
	}
	res, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Sum RPC: %v", err)
	}
	log.Printf("Response from Sum: %v", res.SumResult)
}

func doServerStreaming(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting a Prime Decomposition Server Streaming RPC...")
	req := &calculatorpb.PrimeNumberDecompositionRequest{
		Number: 12390392840,
	}
	stream, err := c.PrimeNumberDecomposition(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling PrimeNumberDecomposition RPC: %v", err)
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error receieving msg: %v", err)
		}
		fmt.Println(res.GetPrimeFactor())
	}
}

func doClientStreaming(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting a Average number Client Streaming RPC...")
	numbers := []int32{1, 2, 3, 4}
	stream, err := c.ComputeAverage(context.Background())
	if err != nil {
		log.Fatalf("Error while calling ComputeAverage: %v", err)
	}
	for _, number := range numbers {
		fmt.Printf("Sending req: %v\n", number)
		err := stream.Send(&calculatorpb.ComputeAverageRequest{Number: number,})
		if err != nil {
			log.Fatalf("Error sending request: %v", err)
		}
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error receiving response from ComputeAverage: %v", err)
	}
	fmt.Printf("ComputeAverage response: %v", res)
}

func doBiDirectionalStreaming(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting a Max number BiDi Streaming RPC...")

	stream, err := c.FindMaximum(context.Background())
	if err != nil {
		log.Fatalf("Could not create sream: %v", err)
	}

	numbers := []int32{1,5,3,6,2,20}
	waitch := make(chan struct{})

	go func() {
		for _, number := range numbers {
			fmt.Printf("Sending message: %v\n", number)
			err := stream.Send(&calculatorpb.FindMaximumRequest{Number: number})
			if err != nil {
				log.Fatalf("Error sending request: %v", err)
			}
		}
		err := stream.CloseSend()
		if err != nil {
			log.Fatalf("Could not close the sream: %v", err)
		}
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error while receieving data from server: %v", err)
			}
			fmt.Printf("Received: %v\n", res.GetMax())
		}
		close(waitch)
	}()

	<- waitch
}

func doErrorUnary(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting a SquareRoot Unary RPC...")

	// correct call
	doErrorCall(c, 10)

	// error call
	doErrorCall(c, -2)
}

func doErrorCall(c calculatorpb.CalculatorServiceClient, n int32) {
	res, err := c.SquareRoot(context.Background(), &calculatorpb.SquareRootRequest{Number: n})
	if err != nil {
		respErr, ok := status.FromError(err)
		if ok {
			// actual error from gRPC (user error)
			fmt.Printf("Error msg from server: %v\n", respErr.Message())
			fmt.Println(respErr.Code())
			if respErr.Code() == codes.InvalidArgument {
				fmt.Println("We probably sent a negative number!")
				return
			}
		} else {
			log.Fatalf("Error calling SquareRoot: %v", err)
			return
		}
	}
	fmt.Printf("Result of SquareRoot of %v: %v\n", n, res.GetNumberRoot())
}