package Config

import (
	"fmt"
	"github.com/hashicorp/vault/api"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
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

func ReadDataFromVault() []string {
	vaultAddr := "http://127.0.0.1:8200"
	rootToken := "s.1hcA47eSwj5sHHgrbuJQv80F"
	log.Println("Entered")
	client, err := api.NewClient(&api.Config{Address: vaultAddr, HttpClient: httpClient})
	if err != nil {
		panic(err)
	}
	client.SetToken(rootToken)

	//read the credentials from the client
	secrets, error := client.Logical().Read("/secret/data/mysqlsecret")
	if error != nil {
		log.Println("Secret not found....", error)
	}
	newData := make([]string, 2)
	i := 0
	for _, value := range secrets.Data {

		for k, v := range value.(map[string]interface{}) {
			if k == "username" || k == "password" {
				fmt.Println(v)
				newData[i] = v.(string)
				i++
			}

		}
	}
	return newData

}

func BuildDBConfig() *DBConfig {

	data := ReadDataFromVault()

	dbConfig := DBConfig{
		Host:     "localhost",
		Port:     3306,
		User:     data[1],
		Password: data[0],
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
