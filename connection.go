package petback

import (
	"os"
	"context"

	"github.com/aiteung/atdb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetConnectionMongo(MongoString, dbname string) *mongo.Database {
	MongoInfo := atdb.DBInfo{
		DBString: os.Getenv(MongoString),
		DBName:   dbname,
	}
	conn := atdb.MongoConnect(MongoInfo)
	return conn
}

func GetAllData(MongoConnect *mongo.Database, colname string) []GeoJson {
	data := atdb.GetAllDoc[[]GeoJson](MongoConnect, colname)
	return data
}

func InsertDatajson(MongoConn *mongo.Database, colname string, coordinate []float64, name, volume, tipe string) (InsertedID interface{}) {
	req := new(LonLatProperties)
	req.Type = tipe
	req.Coordinates = coordinate
	req.Name = name
	req.Volume = volume

	ins := atdb.InsertOneDoc(MongoConn, colname, req)
	return ins
}

func UpdateDatajson(MongoConn *mongo.Database, colname, name, newVolume, newTipe string) error {
    filter := bson.M{"name": name}


    update := bson.M{
        "$set": bson.M{
            "volume": newVolume,
            "tipe":   newTipe,
        },
    }

    _, err := MongoConn.Collection(colname).UpdateOne(context.TODO(), filter, update)
    if err != nil {
        return err
    }

    return nil
}

func DeleteDatajson(MongoConn *mongo.Database, colname string, name string) (*mongo.DeleteResult, error) {
    filter := bson.M{"name": name}
    del, err := MongoConn.Collection(colname).DeleteOne(context.TODO(), filter)
    if err != nil {
        return nil, err
    }
    return del, nil
}