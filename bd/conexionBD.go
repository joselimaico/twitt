package bd

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

// MongoCN es el objeto de conexión a la BD
var MongoCN = ConectarBD()
var clientOptions = options.Client().ApplyURI("mongodb+srv://jfl_96:0PwPY36gGKY2wunb@cluster-px3sbnfz.g9x7l.mongodb.net/?retryWrites=true&w=majority")

// ConectarBD se conecta a la base de datos
func ConectarBD() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("Conexión Exitosa con la BD")
	return client
}

// CheckConnection verifica la conexion a la base de datos
func CheckConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
