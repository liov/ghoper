package client

import (
	"context"
	"fmt"
	"github.com/kataras/iris"
	"github.com/micro/go-micro"
	"github.com/satori/go.uuid"
	"hoper/client/controller/common/logging"
	"hoper/client/controller/cron"
	"hoper/client/controller/hwebsocket"
	"hoper/client/router"
	"hoper/initialize"
	"hoper/protobuf"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Client() {
	defer initialize.DB.Close()
	//crawler.M131()
	//go crawler.MM131()

	cron.New().Start()
	defer cron.New().Stop()

	defer logging.F.Close()

	go hwebsocket.Start()

	//go hcache.Start()

	irisRouter := router.IrisRouter()
Loop:
	for {
		signal.Notify(router.Ch,
			// kill -SIGINT XXXX 或 Ctrl+c
			os.Interrupt,
			syscall.SIGINT, // register that too, it should be ok
			// os.Kill等同于syscall.Kill
			os.Kill,
			syscall.SIGKILL, // register that too, it should be ok
			// kill -SIGTERM XXXX
			syscall.SIGTERM,
		)
		select {
		case <-router.Ch:
			break Loop
		default:
			// listen and serve on http://0.0.0.0:8000.
			if err := irisRouter.Run(iris.TLS(initialize.Config.Server.HttpPort, "../config/tls/pem.pem", "../config/tls/key.key"),
				iris.WithConfiguration(iris.YAML("../config/iris.yml"))); err != nil && err != http.ErrServerClosed {
				log.Printf("Listen: %s\n", err)
			}
		}

	}
	/*	opts := groupcache.HTTPPoolOptions{BasePath: hcache.BasePath}
		peers := groupcache.NewHTTPPoolOpts("", &opts)
		peers.Set("http://localhost:8333", "http://localhost:8222")

		val, err := hcache.GetFromPeer("helloworld", "wjs1", peers)*/

}

func pub() {
	service := micro.NewService(
		micro.Name("go.micro.cli.pubsub"),
	)
	// parse command line
	service.Init()

	// create publisher
	pub1 := micro.NewPublisher("example.topic.pubsub.1", service.Client())
	pub2 := micro.NewPublisher("example.topic.pubsub.2", service.Client())

	// pub to topic 1
	go sendEv("example.topic.pubsub.1", pub1)
	// pub to topic 2
	go sendEv("example.topic.pubsub.2", pub2)

	// block forever
	select {}
}

func sendEv(topic string, p micro.Publisher) {
	t := time.NewTicker(time.Second)

	for range t.C {
		// create new event
		ev := &protobuf.Event{
			Id:        uuid.NewV4().String(),
			Timestamp: time.Now().Unix(),
			Message:   fmt.Sprintf("Messaging you all day on %s", topic),
		}

		log.Fatalf("publishing %+v\n", ev)

		// publish an event
		if err := p.Publish(context.Background(), ev); err != nil {
			log.Fatalf("error publishing: %v", err)
		}
	}
}
