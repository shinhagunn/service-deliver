package services

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/smtp"
	"os"
	"strings"
	"time"

	"github.com/segmentio/kafka-go"
	"github.com/shinhagunn/service-deliver/config/collection"
	"github.com/shinhagunn/service-deliver/models"
	"go.mongodb.org/mongo-driver/bson"
)

type Deliver struct{}

func NewDeliver() *Deliver {
	return &Deliver{}
}

func (Deliver) Process() {
	brokers := strings.Split(os.Getenv("KAFKA_BROKERS"), ",")

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: brokers,
		Topic:   "deliver",
		GroupID: "deliver-order",
	})

	log.Println("Service deliver running ...")

	for {
		msg, err := r.ReadMessage(context.Background())
		if err != nil {
			panic("could not read message " + err.Error())
		}
		userData := msg.Value

		order := new(models.Order)

		err = json.Unmarshal(userData, &order)
		if err != nil {
			panic("could not parse userData " + err.Error())
		}

		go Delivering(order)
	}
}

func Delivering(order *models.Order) {
	order.Status = "Delivering"
	log.Println("Đã nhận được đơn hàng, ... đang xử lý")

	if err := collection.Order.Update(order); err != nil {
		log.Printf("Error update order: %s", order.ID)
		return
	}

	time.Sleep(time.Minute)

	order.Status = "Arrived"
	log.Println("Đơn hàng đã giao đến nơi")

	if err := collection.Order.Update(order); err != nil {
		log.Printf("Error update order: %s", order.ID)
		return
	}

	user := new(models.User)

	if err := collection.User.FindOne(context.Background(), bson.M{"_id": order.UserID}).Decode(&user); err != nil {
		log.Printf("Lỗi không tìm thấy user")
		return
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
