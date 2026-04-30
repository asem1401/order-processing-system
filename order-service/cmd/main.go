package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"

	grpcclient "order-service/internal/grpc"
	"order-service/internal/transport/http"
	"order-service/internal/usecase"

	pb "order-service/proto/payment"
)

func main() {

	
	conn, err := grpc.Dial("payment:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal("failed to connect to payment:", err)
	}

	paymentClient := pb.NewPaymentServiceClient(conn)

	client := grpcclient.NewClient(paymentClient)


	uc := usecase.NewOrderUsecase(client)


	handler := http.NewHandler(uc)

	r := gin.Default()

	r.POST("/orders", handler.CreateOrder)

	log.Println("Order running :8080")
	r.Run(":8080")
}