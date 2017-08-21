package nsq

import "github.com/nsqio/go-nsq"
var Procuctor *nsq.Producer

func init() {
	Procuctor, _ = nsq.NewProducer("127.0.0.1:4150", nsq.NewConfig())

}
func PublishMessage(message string) {
	if err := Procuctor.Publish("spider", []byte(message)); err != nil {           // 发布消息
		panic(err)
	}
}
