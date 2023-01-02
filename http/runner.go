package http

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/viper"
)

const (
	ConfigName = "config"
	ConfigPath = "./config"
)

type Component interface {
	Instantiate(configMap interface{}) (err error)
	Start() (err error)
	Stop() (err error)
}

type DaemonRunner interface {
	Run(component Component, configMap interface{})
}

type daemonRunner struct{}

func NewDaemonRunner() DaemonRunner {
	return daemonRunner{}
}

func (d daemonRunner) Run(component Component, configMap interface{}) {
	closed := make(chan struct{})

	signalChan := make(chan os.Signal)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		select {
		case sig := <-signalChan:
			log.Println(fmt.Sprintf("[Daemon Runner] receiving signal [%s], exiting service", sig))
			err := component.Stop()
			if err != nil {
				log.Println(fmt.Errorf("[Daemon Runner] error stopping service. \ncaused by: %v", err))
			}
			close(closed)
		}
	}()

	if err := LoadConfig(&configMap); err != nil {
		log.Fatal(fmt.Sprintf("error loading config. cause: %v", err.Error()))
	}

	if err := component.Instantiate(configMap); err != nil {
		log.Fatal(fmt.Sprintf("error starting component. cause: %v", err.Error()))
	}

	log.Println("[Daemon Runner] starting service")

	if err := component.Start(); err != nil {
		log.Fatal(fmt.Errorf("error starting service. cause: %v", err.Error()))
		err = component.Stop()
		if err != nil {
			log.Fatal(fmt.Errorf("error stopping service. cause: %v", err.Error()))
		}
		close(closed)
	}

	<-closed
	log.Println("service stopped")
}

func LoadConfig(configMap interface{}) error {
	viper.SetConfigName(ConfigName)
	viper.AddConfigPath(ConfigPath)

	err := viper.ReadInConfig()
	if err != nil {
		return fmt.Errorf("load config error. cause: %v", err.Error())
	}

	err = viper.Unmarshal(&configMap)
	if err != nil {
		return fmt.Errorf("[WKND Conf] load config error. Cause: %v", err)
	}

	return nil
}
