package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

var (
	Config Configuration
)

type App struct {
	AppName   string `json:"appName"`
	Port      string `json:"port"`
	SecretJWT string `json:"secretJWT"`
}

type Db struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Port     string `json:"port"`
	Host     string `json:"host"`
	DbName   string `json:"name"`
}

type S3 struct {
	Key    string `json:"s3Key"`
	Secret string `json:"s3Secret"`
	Region string `json:"awsRegion"`
}

type Configuration struct {
	App App `json:"app"`
	DB  Db  `json:"db"`
	S3  S3  `json:"s3"`
}

func init() {
	env := "dev"

	if len(os.Args) > 1 {
		env = os.Args[1]
	}
	log.Println("INITIALIZE", "SERVER RUN ON "+env+" ENV")
	path := env + ".config.json"
	jsonFile, _ := os.Open(path)

	byteConfig, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Panic(err.Error())
	}

	json.Unmarshal(byteConfig, &Config)
}
