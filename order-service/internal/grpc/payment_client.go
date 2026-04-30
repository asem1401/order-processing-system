package grpc

import (
	"context"

	pb "order-service/proto/payment"
)

type Client struct {
	client pb.PaymentServiceClient
}

func NewClient(c pb.PaymentServiceClient) *Client {
	return &Client{client: c}
}


func (c *Client) ProcessPayment(orderID string, amount int64, email string) (string, error) {
	res, err := c.client.ProcessPayment(context.Background(), &pb.PaymentRequest{
		OrderId: orderID,
		Amount:  amount,
		Email:   email,
	})

	if err != nil {
		return "", err
	}

	return res.Status, nil
}