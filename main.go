package main

import (
	"log"
	"time"

	"github.com/Shopify/sarama"
)

func main() {
	broker := sarama.NewBroker("localhost:9092")
	config := sarama.NewConfig()
	config.Version = sarama.V1_0_0_0
	broker.Open(config)
	yes, err := broker.Connected()
	if err != nil {
		log.Print(err.Error())
	}
	log.Print(yes)
	name := "blah25s"

	topicDetail := &sarama.TopicDetail{}
	topicDetail.NumPartitions = int32(1)
	topicDetail.ReplicationFactor = int16(1)
	topicDetail.ConfigEntries = make(map[string]*string)

	topicDetails := make(map[string]*sarama.TopicDetail)
	topicDetails[name] = topicDetail

	request := sarama.CreateTopicsRequest{
		Timeout:      time.Second * 15,
		TopicDetails: topicDetails,
	}
	response, err := broker.CreateTopics(&request)
	if err != nil {
		log.Printf("%#v", &err)
	}
	t := response.TopicErrors
	log.Printf("length is %d", len(t))
	for key, val := range t {
		log.Printf("Key is %s", key)
		log.Printf("Value is %#v", val.Err.Error())
		log.Printf("Value3 is %#v", val.ErrMsg)
	}
	log.Printf("the response is %#v", response)
	broker.Close()
}
