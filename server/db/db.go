package db

import (
	"context"
	"log"
	"os"
	"schema"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DBHandler struct {
  dbConn *mongo.Client
}

func newDBHandler(dbConn *mongo.Client) *DBHandler{
  return &DBHandler{
    dbConn: dbConn,
  }
}

func getURI() (string, error) {
  err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading the dotenv file")
    return "", err
  }
  uri := os.Getenv("URI")
  return uri, nil
}

func GetDBConn() (*DBHandler, error) {
  uri, err := getURI()
  if err != nil {
    os.Getenv("URI")
  }
  serverAPI := options.ServerAPI(options.ServerAPIVersion1)
  opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)
  dbConn, err := mongo.Connect(context.TODO(), opts)
  if err != nil {
    log.Fatal(err.Error())
    return nil, err
  }
  return newDBHandler(dbConn), nil
}




func (d *DBHandler) FetchApplications() ([]schema.Application, error) {
  var results []schema.Application
  coll := d.dbConn.Database("jobber").Collection("applications")
  cursor, err := coll.Find(context.TODO(), bson.D{{}})
  if err != nil {
    log.Fatal(err.Error())
    return results, err
  }
  if err = cursor.All(context.TODO(), &results); err != nil {
    log.Fatal(err.Error())
    return results, err
  }
  return results, nil
}

func (d *DBHandler) AddApplication(application *schema.Application) error {
  coll := d.dbConn.Database("jobber").Collection("applications")
  _, err := coll.InsertOne(context.TODO(), application)
  if err != nil {
    log.Fatal(err.Error())
  }
  return err
}

func (d *DBHandler) Close() error {
  return d.dbConn.Disconnect(context.TODO())
}
