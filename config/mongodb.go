package config

import (
	"fmt"
	"log"
	"os"

	"github.com/kamva/mgm/v3"
	"github.com/shinhagunn/service-deliver/config/collection"
	"github.com/shinhagunn/service-deliver/models"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitMongoDB() {
	mgm.SetDefaultConfig(nil, os.Getenv("DATABASE_NAME"), options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s", os.Getenv("MONGO_HOST"), os.Getenv("MONGO_PORT"))))

	log.Printf("Connected to %s!", os.Getenv("DATABASE_NAME"))

	collection.Order = mgm.Coll(&models.Order{})
	collection.User = mgm.Coll(&models.User{})
}
