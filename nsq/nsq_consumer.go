package nsq

import (
	"github.com/nsqio/go-nsq"
)
func Consumer(handler nsq.Handler,channel chan int) {
	c, err := nsq.NewConsumer("spider", "test-channel", nsq.NewConfig())   // 新建一个消费者
	if err != nil {
		panic(err)
	}
	c.AddHandler(handler)                                           // 添加消息处理
	if err := c.ConnectToNSQD("127.0.0.1:4150"); err != nil {            // 建立连接
		panic(err)
	}
	channel<-1
}