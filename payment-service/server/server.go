package main

import (
	"os"
	"net"
	"log"

	gpay "payment-app/payment-service/proto"

	"github.com/joho/godotenv"
	payjp "github.com/payjp/payjp-go/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

// server is used to implement sa
type server struct{}

func (c *server) Charge(ctx context.Context, req *gpay.PayRequest) (*gpay.PayResponse, error) {
	// load env setting
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// init API
	pay := payjp.New(os.Getenv("PAYJP_TEST_SECRET_KEY"), nil)

	// Make a payment. Payment amount for the first argument.
	// Method and setting of payment for the second argument.
	charge, err := pay.Charge.Create(int(req.Amount), payjp.Charge{
		Currency:    "jpy",
		CardToken:   req.Token, // Specify card information, customer ID, or card token. Token this time.
		Capture:     true,
		Description: req.Name + ":" + req.Description, // Set summary text.
	})
	if err != nil {
		return nil, err
	}

	// Generate Response from the paid result
	res := &gpay.PayResponse {
		Paid: 		charge.Paid,
		Captured: 	charge.Captured,
		Amount: 	int64(charge.Amount),
	}
	return res, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	gpay.RegisterPayManagerServer(s, &server{})

	// Register reflection service on gRPC server.
	reflection.Register(s)
	log.Printf("gRPC Sever started: localhost%s\n", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}