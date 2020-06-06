package mongoDB

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
)

type Mongodb struct {
	DBName     string
	Client     *mongo.Client
	Database   *mongo.Database
	Collection *mongo.Collection
	Context    context.Context
}

func NewMongoDB(name string) Mongodb {
	return Mongodb{DBName: name}
}

func (mdb *Mongodb) Connect(uri string) error {
	var err error
	mdb.Client, err = mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
		return err
	}
	mdb.Context, _ = context.WithTimeout(context.Background(), 10*time.Second)
	err = mdb.Client.Connect(mdb.Context)
	if err != nil {
		log.Fatal(err)
		return err
	}
	err = mdb.Client.Ping(mdb.Context, readpref.Primary())
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

func (mdb *Mongodb) InsertElement(i interface{}) error {
	_, err := mdb.Collection.InsertOne(mdb.Context, i)
	return err
}

func (mdb *Mongodb) GetLatest() (interface{}, error) {
	var i interface{}
	err := mdb.Collection.FindOne(context.TODO(), bson.D{}).Decode(&i)
	return i, err
}

func (mdb *Mongodb) GetLatestByKey(key string, value string, result interface{}) error {
	filter := bson.D{{key, value}}
	//var i interface{}
	err := mdb.Collection.FindOne(context.TODO(), filter).Decode(&result)
	return err
}

func (mdb *Mongodb) initDB(name string) {
	mdb.Database = mdb.Client.Database(name)
}

func (mdb *Mongodb) initCollection(name string) {
	mdb.Collection = mdb.Database.Collection(name)
}

func (mdb *Mongodb) InitDatabase(collectionName string) {
	mdb.initDB(mdb.DBName)
	mdb.initCollection(collectionName)
}
