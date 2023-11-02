package Config

import (
	"fmt"
	"github.com/hashicorp/vault/api"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
	"reflect"
	"time"
)

var DB *gorm.DB

// DBConfig represents db configuration
type DBConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

var httpClient = &http.Client{
	Timeout: 10 * time.Second,
}

func ReadDatafromVault() {
	vaultAddr := "http://localhost:8200"
	rootToken := "s.7a1lmH5XuRV3LsLbmcVE23fh"
	log.Println("Entered")
	client, err := api.NewClient(&api.Config{Address: vaultAddr, HttpClient: httpClient})
	if err != nil {
		panic(err)
	}
	client.SetToken(rootToken)
	//read the credentials from the client
	secrets, error := client.Logical().Read("secret/mysqlsecret")
	if error != nil {
		log.Println("Secret not found....", error)
	}

	for key, val := range secrets.Data {
		log.Println(key, val)
		log.Println(reflect.TypeOf(val))

	}

}

func BuildDBConfig() *DBConfig {
	dbConfig := DBConfig{
		Host:     "localhost",
		Port:     3306,
		User:     "root",
		Password: "vignesh",
		DBName:   "azuredb",
	}
	//dbConfig := DBConfig{}
	return &dbConfig
}
func DbURL(dbConfig *DBConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
}
