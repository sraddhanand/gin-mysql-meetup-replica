package models

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

// Setup initializes the database instance
func Setup() {
	var err error
	// fmt.Println(setting.DatabaseSetting.Host)
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DATABASE_USER"),
		os.Getenv("DATABASE_PASS"),
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_NAME"))
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}

	err1 := db.AutoMigrate(&User{}, &Group{}, &GroupEvent{}, &Story{})
	if err1 != nil {
		log.Fatalf("migrations failed: %v", err)
	}
	log.Println("Automation Completed")

	// gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
	// 	return setting.DatabaseSetting.TablePrefix + defaultTableName
	// }

	// db.SingularTable(true)
	// // db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	// // db.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
	// // db.Callback().Delete().Replace("gorm:delete", deleteCallback)
	// db.DB().SetMaxIdleConns(10)
	// db.DB().SetMaxOpenConns(100)
}

func GetDB() *gorm.DB {
	return db
}

// CloseDB closes database connection (unnecessary)
// func CloseDB() {
// 	defer db.Close()
// }
