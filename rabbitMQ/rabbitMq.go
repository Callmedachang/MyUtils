package rabbitMQ

import (
	"fmt"
	"log"
	"bytes"
	"github.com/streadway/amqp"
)

var Conn *amqp.Connection
var Channel *amqp.Channel
var count = 0

const (
	QueueName = "wrcredit"
	Exchange  = "rb_xjkc"
	mqurl     = "amqp://zcmlc_xjkc:zcmlc2017@10.139.96.248:5672"
)

func init() {
	if Channel == nil {
		mqConnect()
	}
}

//func main() {
//	go func() {
//		for {
//			push()
//			time.Sleep(1 * time.	Second)
//		}
//	}()
//	receive()
//	fmt.Println("end")
//	close()
//}

func failOnErr(err error, msg string) {
	if err != nil {
		log.Fatalf("%s:%s", msg, err)
		panic(fmt.Sprintf("%s:%s", msg, err))
	}
}

func mqConnect() {
	var err error
	Conn, err = amqp.Dial(mqurl)
	failOnErr(err, "failed to connect tp rabbitmq")

	Channel, err = Conn.Channel()
	failOnErr(err, "failed to open a channel")
}

func close() {
	Channel.Close()
	Conn.Close()
}

//连接rabbitmq server
func push() {

	if Channel == nil {
		mqConnect()
	}
	msgContent := "hello world!"

	Channel.Publish(Exchange, QueueName, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(msgContent),
	})
}

func receive() {
	if Channel == nil {
		mqConnect()
	}

	msgs, err := Channel.Consume(QueueName, "", true, false, false, false, nil)
	failOnErr(err, "")

	forever := make(chan bool)

	go func() {
		//fmt.Println(*msgs)
		for d := range msgs {
			s := BytesToString(&(d.Body))
			count++
			fmt.Printf("receve msg is :%s -- %d\n", *s, count)
		}
	}()

	fmt.Printf(" [*] Waiting for messages. To exit press CTRL+C\n")
	<-forever
}

func BytesToString(b *[]byte) *string {
	s := bytes.NewBuffer(*b)
	r := s.String()
	return &r
}
