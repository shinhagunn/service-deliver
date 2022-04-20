package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/smtp"
	"os"
	"strings"
	"time"

	"github.com/shinhagunn/service-deliver/config/collection"
	pb "github.com/shinhagunn/service-deliver/delivery"
	"github.com/shinhagunn/service-deliver/models"
	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/grpc"
)

type Service interface {
	Process()
}

type server struct {
	pb.UnimplementedDeliverServer
}

func (s *server) Deliver(ctx context.Context, in *pb.DeliverRequest) (*pb.DeliverResponse, error) {
	order := new(models.Order)
	err := collection.Order.FindByID(in.OrderId, order)

	if err != nil {
		panic(err)
	}

	order.Status = "Delivering"

	log.Println("Đã nhận được đơn hàng, ... đang xử lý")

	if err := collection.Order.Update(order); err != nil {
		log.Printf("Error update order: %s", order.ID)
		panic(err)
	}

	time.Sleep(time.Minute)

	order.Status = "Arrived"
	log.Println("Đơn hàng đã giao đến nơi")

	if err := collection.Order.Update(order); err != nil {
		log.Printf("Error update order: %s", order.ID)
		panic(err)
	}

	user := new(models.User)

	if err := collection.User.FindOne(context.Background(), bson.M{"_id": order.UserID}).Decode(&user); err != nil {
		log.Printf("Lỗi không tìm thấy user")
		panic(err)
	}

	subject := "Subject: Your order arrived ShinWatch\n\n"

	m := fmt.Sprintf("Người nhận: %v\n", user.UserProfile.Fullname)
	m += fmt.Sprintf("Địa chỉ: %v\n", user.UserProfile.Address)
	m += "Người gửi: Shin Watches\n\n"
	for i, v := range order.OrderProduct {
		m += fmt.Sprintf("Sản phẩm thứ %v:\nProduct: %v, \nPrice: %v \n\n", i, v.Name, v.Price)
	}

	body := m
	message := strings.Join([]string{subject, body}, " ")

	SendEmailArrived(message, user.Email)

	return &pb.DeliverResponse{Status: true}, nil
}

func SendEmailArrived(message string, toAddress string) (response bool, err error) {
	fromAddress := os.Getenv("EMAIL_FROM")
	fromEmailPassword := os.Getenv("EMAIL_PASSWORD")
	smtpServer := os.Getenv("SMTP_SERVER")
	smptPort := os.Getenv("SMTP_PORT")

	var auth = smtp.PlainAuth("", fromAddress, fromEmailPassword, smtpServer)
	if err := smtp.SendMail(smtpServer+":"+smptPort, auth, fromAddress, []string{toAddress}, []byte(message)); err == nil {
		return true, nil
	} else {
		return false, err
	}
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", os.Getenv("GRPC_PORT")))

	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()

	pb.RegisterDeliverServer(s, &server{})

	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
