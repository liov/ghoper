package hnsq

import (
	"encoding/json"
	"fmt"
	"github.com/nsqio/go-nsq"
	"hoper/utils/reflectinvoke"
	"log"
)

// 消费者
type Consumer struct{}

// 主函数

//处理消息
func (*Consumer) HandleMessage(msg *nsq.Message) error {
	fmt.Println("receive", msg.NSQDAddress, "message:", string(msg.Body))
	return nil
}

//初始化消费者
func NewConsumer(topic string, channel string, handle nsq.HandlerFunc) {
	cfg := nsq.NewConfig()
	//cfg.LookupdPollInterval = time.Second          //设置重连时间
	c, err := nsq.NewConsumer(topic, channel, cfg) // 新建一个消费者
	if err != nil {
		panic(err)
	}
	//c.SetLogger(nil, 0)       //屏蔽系统日志
	c.AddHandler(handle) // 添加消费者接口

	//建立NSQLookupd连接
	if err := c.ConnectToNSQLookupd(Addr4161); err != nil {
		log.Println("consumer 新建失败")
	}

	//建立多个nsqd连接
	// if err := c.ConnectToNSQDs([]string{"127.0.0.1:4150", "127.0.0.1:4152"}); err != nil {
	//  panic(err)
	// }

	// 建立一个nsqd连接
	/*	if err := c.ConnectToNSQD(address); err != nil {
		 panic(err)
		}
		<-c.StopChan*/
}

func handleStringMessage(message *nsq.Message) error {

	fmt.Printf("handleStringMessage get a message  %v\n\n", string(message.Body))
	return nil
}

var reflectinvoker *reflectinvoke.Reflectinvoker

func handleJsonMessage(message *nsq.Message) error {

	resultJson := reflectinvoker.InvokeByJson([]byte(message.Body))
	result := reflectinvoke.Response{}
	err := json.Unmarshal(resultJson, &result)
	if err != nil {
		return err
	}
	info := "handleJsonMessage get a result\n"
	info += "raw:\n" + string(resultJson) + "\n"
	info += "function: " + result.FuncName + " \n"
	info += fmt.Sprintf("result: %v\n", result.Data)
	info += fmt.Sprintf("error: %d,%s\n\n", result.ErrorCode,
		reflectinvoke.ErrorMsg(result.ErrorCode))

	fmt.Println(info)

	return nil
}
