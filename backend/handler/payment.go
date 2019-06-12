package handler

import (
	"context"
	"log"
	"net/http"
	"os"
	"payment-app/backend/db"
	"payment-app/backend/domain"
	gpay "payment-app/payment-service/proto"
	"strconv"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

// Charge : exec payment service charge
func Charge(c Context) {
	// load env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// receive parameters and body
	t := domain.Payment{}
	c.Bind(&t)
	identifer, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	// get item information from id
	res, err := db.SelectItem(int64(identifer))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	// create request to send to gRPC server
	greq := &gpay.PayRequest{
		Id:          int64(identifer),
		Token:       t.Token,
		Amount:      res.Amount,
		Name:        res.Name,
		Description: res.Description,
	}

	// Connect to the server by specifying the IP address(here localhost) and the port number(here 50051)
	conn, err := grpc.Dial(os.Getenv("GRPC_URL"), grpc.WithInsecure())
	if err != nil {
		c.JSON(http.StatusForbidden, err)
	}
	defer conn.Close()
	client := gpay.NewPayManagerClient(conn)

	// exec payment processing function of gRPC microservice
	gres, err := client.Charge(context.Background(), greq)
	if err != nil {
		c.JSON(http.StatusForbidden, err)
		return
	}
	c.JSON(http.StatusOK, gres)
}
