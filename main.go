package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/accuknox/kmux"
	"github.com/accuknox/kmux/config"
	"github.com/spf13/viper"
)

var configFilePath *string

func main() {

	loadConfigFromFile()

	topicName := viper.GetString("app.topic")
	msgSize := viper.GetInt("app.msg.size")
	msgLength := viper.GetInt("app.msg.length")

	err := kmux.Init(&config.Options{
		LocalConfigFile: "kmux.yaml",
	})
	if err != nil {
		panic(err)
	}

	ss, err := kmux.NewStreamSink(topicName)
	if err != nil {
		panic(err)
	}
	err = ss.Connect()
	if err != nil {
		panic(err)
	}
	count := 0
	iteration := 1
	var timeVar time.Time
	for {
		randBytes := make([]byte, msgSize)
		rand.Read(randBytes)
		if count == 0 {
			timeVar = time.Now()
		}
		count++
		err = ss.Flush(randBytes)
		if err != nil {
			panic(err)
		}
		if count == msgLength {
			log.Println("time took : ", time.Since(timeVar))
			fmt.Printf("%v. Time: %v ; Size(Bytes): %v ; Time : %v\n", iteration, timeVar, msgSize, time.Since(timeVar))
			count = 0
			iteration++
		}
	}
}

func loadConfigFromFile() {

	//flag set with config file path
	configFilePath = flag.String("config-path", "config/", "config-file-path/")
	flag.Parse()

	viper.SetConfigName("app")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(*configFilePath)
	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Error reading config file: %s\n", err)
	}

}
