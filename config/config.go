package config

import (
	"log"
	"os"
	"strconv"
	"sync"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	Driver    string
	Name      string
	Address   string
	Port      int
	Username  string
	Password  string
	Keys3     string
	Secrets3  string
	Regions3  string
	Midserver string
}

var lock = &sync.Mutex{}
var appConfig *AppConfig

func GetConfig() *AppConfig {
	lock.Lock()
	defer lock.Unlock()

	if appConfig == nil {
		appConfig = initConfig()
	}

	return appConfig
}

func initConfig() *AppConfig {
	SECRET = os.Getenv("SECRET")
	var defaultConfig AppConfig
	if SECRET == "" {
		err := godotenv.Load("local.env")

		if err != nil {
			log.Fatal("Cannot read configuration")
			return nil
		}
	}
	SECRET = os.Getenv("SECRET")
	cnv, err := strconv.Atoi(os.Getenv("SERVERPORT"))
	if err != nil {
		log.Fatal("Cannot parse port variable")
		return nil
	}
	SERVERPORT = int16(cnv)
	defaultConfig.Name = os.Getenv("DBname")
	defaultConfig.Username = os.Getenv("DBusername")
	defaultConfig.Password = os.Getenv("DBpassword")
	defaultConfig.Address = os.Getenv("DBhost")
	cnv, err = strconv.Atoi(os.Getenv("DBport"))
	if err != nil {
		log.Fatal("Cannot parse DB Port variable")
		return nil
	}
	defaultConfig.Port = cnv
	defaultConfig.Keys3 = os.Getenv("S3_KEY")
	defaultConfig.Secrets3 = os.Getenv("S3_SECRET")
	defaultConfig.Regions3 = os.Getenv("S3_REGION")
	defaultConfig.Midserver = os.Getenv("Midtrans_server")

	return &defaultConfig
}
