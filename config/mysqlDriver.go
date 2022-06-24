package config

import (
	"fmt"
	"os"

	"project/e-comerce/migration"

	"github.com/joho/godotenv"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB()*gorm.DB{

	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	config := map[string]string{
		"DB_Username": os.Getenv("DB_Username"),
		"DB_Password": os.Getenv("DB_Password"),
		"DB_Port":     os.Getenv("DB_Port"),
		"DB_Host":     os.Getenv("DB_Host"),
		"DB_Name":     os.Getenv("DB_Name"),
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=UTC", 
		config["DB_Username"],
		config["DB_Password"],
		config["DB_Host"],
		config["DB_Port"],
		config["DB_Name"])

		var e error
		db, e := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
		if e != nil{
			panic(e)
		}
		
		migration.Migration(db)
		
		return db
}


