package viper

import (
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

var (
	opendb OpenDB
	option Option
)

const (
	envDevelopment = "development"
	envStaging     = "staging"
	envProduction  = "production"
)

type (
	OpenDB struct {
		// Add server
		Db *sqlx.DB
	}

	Option struct {
		configPath string
		configName string
	}
)

func viperPrepare() {
	viper.SetConfigName(option.configName)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(option.configPath)
}

func Init() error {
	configFile := getDefaultConfigFile()

	_, err := ioutil.ReadFile(configFile)
	if err != nil {
		return err
	}

	configSplit := strings.Split(configFile, "/")
	lengthSplit := len(configSplit)
	option.configPath = strings.Join(configSplit[:lengthSplit-1], "/")
	option.configName = configSplit[lengthSplit-1]

	// Viper
	viperPrepare()
	viper.ReadInConfig()
	vp := viper.Get("database")
	settingConn(vp)

	return nil
}

func settingConn(vp interface{}) {
	var err error

	vals := reflect.ValueOf(vp)
	for _, e := range vals.MapKeys() {
		connStr := vals.MapIndex(e).Interface().(string)

		switch e.Interface().(string) {

		// Add server
		case "centuryoutlet_1h":
			opendb.Db, err = sqlx.Open("mssql", connStr)
			if err != nil {
				fmt.Println("Error Connection [Repeat Order]", err)
			}
		}
	}
}

func ViperActive() {
	// Viper
	viperPrepare()
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {

		viper.ReadInConfig()
		vp := viper.Get("database")
		settingConn(vp)

	})
}

func getDefaultConfigFile() string {
	configPath := "./files/etc/example/example.development.yaml"
	namespace, _ := ioutil.ReadFile("/var/run/secrets/kubernetes.io/serviceaccount/namespace")

	env := string(namespace)

	switch env {
	case envStaging:
		configPath = "/vault/secrets/database.yaml"
	case envProduction:
		configPath = "/vault/secrets/database.yaml"
	default:
		if os.Getenv("GOPATH") == "" {
			configPath = "files/etc/request-order/request-order.development.yaml"
		}
	}

	return configPath
}

func GetConn() OpenDB {
	return opendb
}