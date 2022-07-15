package app

import (
	"context"
	"flag"
	"fmt"
	"ginweb/app/conf"
	"ginweb/app/server"
	"ginweb/app/service"
	"ginweb/common/logz"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func Run() {
	flag.Parse()
	conf.LoadConfig()
	logz.NewLogger(conf.Conf.Log, zap.AddCaller(), zap.AddCallerSkip(1))
	svc := service.New(conf.Conf)
	httpServer := server.NewHttp(conf.Conf, svc)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	eg := &errgroup.Group{}
	logz.Infof("server start listen on %s", conf.Conf.Server.Addr)

	eg.Go(func() error {
		if err := httpServer.ListenAndServe(); err != http.ErrServerClosed {
			logz.Fatal(err)
		}
		return nil
	})

	eg.Go(func() error {
		ch := make(chan os.Signal)
		signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
		select {
		case v := <-ch:
			fmt.Println()
			logz.Infof("捕获退出信号: %v", v)

		case <-ctx.Done():
		}

		httpServer.Shutdown(context.Background())
		logz.Warn("server stop")
		return nil
	})

	if err := eg.Wait(); err != nil {
		logz.Error(err)
	}
}
