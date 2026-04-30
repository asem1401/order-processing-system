package main

import (
	"context"
	"log"
	"net"

	"payment-service/internal/messaging"
	"payment-service/internal/usecase"
	pb "payment-service/proto/payment"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedPaymentServiceServer
	uc *usecase.PaymentUsecase
}

func (s *server) ProcessPayment(ctx context.Context, req *pb.PaymentRequest) (*pb.PaymentResponse, error) {
	status := s.uc.Process(req.OrderId, req.Amount, req.Email)

	return &pb.PaymentResponse{
		Status: status,
	}, nil
}

func main() {
	pub := messaging.NewPublisher()
	uc := usecase.NewPaymentUsecase(pub)

	lis, _ := net.Listen("tcp", ":50051")
	s := grpc.NewServer()

	pb.RegisterPaymentServiceServer(s, &server{uc: uc})

	log.Println("Payment running")
	s.Serve(lis)
}