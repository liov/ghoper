package main

import (
	"github.com/kataras/iris"
	"go.uber.org/zap"
	"hoper/controller/hwebsocket"
	"hoper/initialize"
	"hoper/router"
	"hoper/utils/ulog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	if log, ok := ulog.Log.(*zap.SugaredLogger); ok {
		defer log.Sync()
	}

	defer initialize.DB.Close()
	defer initialize.BoltDB.Close()

	/*	cron.New().Start()
		defer cron.New().Stop()*/

	defer ulog.LogFile.Close()

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
			// listen and serve on https://0.0.0.0:8000.
			//if err := irisRouter.Run(iris.TLS(initialize.Config.Server.HttpPort, "../../config/tls/cert.pem", "../../config/tls/cert.key"),
			if err := irisRouter.Run(iris.Addr(initialize.Config.Server.HttpPort)); err != nil && err != http.ErrServerClosed {
				ulog.Error(err)
			}
		}

	}
	/*	opts := groupcache.HTTPPoolOptions{BasePath: hcache.BasePath}
		peers := groupcache.NewHTTPPoolOpts("", &opts)
		peers.Set("http://localhost:8333", "http://localhost:8222")

		val, err := hcache.GetFromPeer("helloworld", "wjs1", peers)*/

}
