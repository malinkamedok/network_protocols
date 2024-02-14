package main

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"strconv"
	"time"
)

var gotMessage bool

var callbackFirstMessage mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("TOPIC: %s\n", msg.Topic())
	fmt.Printf("MSG: %s\n", msg.Payload())

	if token := client.Unsubscribe(msg.Topic()); token.Wait() && token.Error() != nil {
		fmt.Println("Failed to unsubscribe: ", token.Error())
	}

	gotMessage = true
	fmt.Println("Got the message. Unsubscribed.")
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected to broker successfully!")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connect lost: %v\n", err)
}

func sub(client mqtt.Client, topic string) {
	token := client.Subscribe(topic, 0, callbackFirstMessage)
	token.Wait()
	fmt.Printf("Subscribed to topic %s\n", topic)
	return
}

func main() {

	options := mqtt.NewClientOptions()
	options.AddBroker("tcp://broker.hivemq.com:1883")
	options.OnConnect = connectHandler
	options.OnConnectionLost = connectLostHandler

	client := mqtt.NewClient(options)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		return
	}

	topic1 := "ITMO/Student16/Value1"
	topic2 := "ITMO/Student16/Value2"

	sub(client, topic1)
	sub(client, topic2)

	time.Sleep(2 * time.Second)

	fmt.Println("Messages from the two first topics were read")
	fmt.Println("Reading from all /Value3 topics")

	for i := 0; i < 30; i++ {
		gotMessage = false
		topic := "ITMO/Student" + strconv.Itoa(i) + "/Value3"
		sub(client, topic)
		for !gotMessage {
		}
	}

	fmt.Println("All messages have been read. Lab is done!")

}
