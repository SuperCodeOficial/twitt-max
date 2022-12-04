package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoCN ES EL OBJETO DE CONEXION A LA BD
var MongoCN = ConectarBD()
var clientOptions = options.Client().ApplyURI("mongodb+srv://max:SQrKMRmxe.V6S5U@cluster0.fg2eol9.mongodb.net/?retryWrites=true&w=majority")

// ConectarBD() ESTA ES LA FUNCIÓN QUE ME PERMITE CONECTAR LA BD
func ConectarBD() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("Conexión Exitosa con la BD")
	return client
}

// ChequeoConnection() ES LA FUNC QUE HACE PING A LA BD
func ChequeoConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
