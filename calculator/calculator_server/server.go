package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"math"
	"net"

	"gRPC_course/calculator/calculatorpb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {}

func (*server) Sum(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
	fmt.Printf("Received Sum RPC: %v", req)
	firstNumber := req.FirstNumber
	secondNumber := req.SecondNumber

	sum := firstNumber + secondNumber
	res := &calculatorpb.SumResponse{
		SumResult: sum,
	}
	return res, nil
}

func (*server) PrimeNumberDecomposition(req *calculatorpb.PrimeNumberDecompositionRequest,
	stream calculatorpb.CalculatorService_PrimeNumberDecompositionServer) error {
	fmt.Printf("Received PrimeNumberDecomposition RPC: %v", req)
	number := req.GetNumber()
	divisor := int64(2)
	for number > 1 {
		if number % divisor == 0 {
			err := stream.Send(&calculatorpb.PrimeNumberDecompositionResponse{
				PrimeFactor: divisor,
			})
			if err != nil {
				log.Fatalf("Can not send a response: %v", err)
			}
			number = number / divisor
		} else {
			divisor++
			fmt.Printf("Divisor has increased to %v\n", divisor)
		}
	}
	return nil
}

func (*server) ComputeAverage(stream calculatorpb.CalculatorService_ComputeAverageServer) error {
	fmt.Printf("Received ComputeAverage RPC stream: %v", stream)
	count := 0
	sum := int32(0)
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&calculatorpb.ComputeAverageResponse{Average: float64(sum) / float64(count)})
		}
		if err != nil {
			log.Fatalf("while reading client stream: %v", err)
		}
		sum += req.GetNumber()
		count++
	}
}

func (*server) FindMaximum(stream calculatorpb.CalculatorService_FindMaximumServer) error {
	fmt.Printf("Received FindMaximum RPC stream: %v", stream)
	currMax := int32(0)

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Can not get request from the client: %v", err)
			return err
		}
		number := req.GetNumber()
		if number > currMax {
			currMax = number
			sendErr := stream.Send(&calculatorpb.FindMaximumResponse{Max: currMax})
			if sendErr != nil {
				log.Fatalf("Can not send message to the client: %v", sendErr)
			}
		}
	}
}


func (*server) SquareRoot(ctx context.Context, req *calculatorpb.SquareRootRequest) (
	*calculatorpb.SquareRootResponse, error) {
	fmt.Printf("Received SquareRoot RPC: %v\n", req)
	number := req.GetNumber()

	if number < 0 {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Received a negative number: %v", number),
		)
	}
	return &calculatorpb.SquareRootResponse{NumberRoot: math.Sqrt(float64(number))}, nil
}

func main() {
	fmt.Println("Calculator server")

	lis, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	calculatorpb.RegisterCalculatorServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
