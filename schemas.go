package popugschema

import (
	_ "embed"

	"github.com/dangkaka/go-kafka-avro"
)

//go:embed user_schema
var userSchema string

var kafkaServers = []string{"localhost:9092"}
var schemaRegistryServers = []string{"http://localhost:8081"}
var topic = "users"

func NewUserProducer() (*kafka.AvroProducer, error) {
	return kafka.NewAvroProducer(kafkaServers, schemaRegistryServers)
}

func sendUserUpdate(producer *kafka.AvroProducer, schema string) {
	value := `{
		
	}`
}
