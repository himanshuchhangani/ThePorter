package main

import (
	"encoding/json"
	"flag"
	"log"
	"os"

	"github.com/go-redis/redis"
	"github.com/himanshuchhangani/theporter"
	"github.com/himanshuchhangani/theporter/dataprocessor"
	"github.com/himanshuchhangani/theporter/storagehandler"
)

var (
	confFile = flag.String("config", "/tmp/config.yml,", "ThePorter Config file")
	portFile = flag.String("portData", "/tmp/ports.json", "The Port Data to be Ported to the warehouse.")
)

func main() {
	log.Printf("ImPorting...")
	flag.Parse()
	config, err := theporter.ReadConfig(*confFile)
	if err != nil {
		log.Fatalf("error reading config file " + err.Error())
		return
	}

	client := redis.NewClient(&redis.Options{
		Addr:     config.DBConfig.Addr,
		Password: config.DBConfig.Password,
		DB:       config.DBConfig.DB,
	})

	pf, err := os.Open(*portFile)
	if err != nil {
		log.Fatalf("Error: opening port json file " + err.Error())
	}

	ds := storagehandler.DataStore{Client: client}
	de := json.NewDecoder(pf)

	fileProcessor := &dataprocessor.Processor{
		DecoderEngine: de,
		StorageEngine: &ds,
	}

	if err := fileProcessor.Process(); err != nil {
		log.Fatalf("Error processing the file " + err.Error())
	}

	log.Printf("ImPorted.")

}
