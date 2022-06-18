package main

import (
	"demoGo/model/inck"
	"fmt"
	"github.com/cengsin/oracle"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

/**
oracle 访问
*/
func db_oracle() {
	dsn := "djwkerp/djwk888wK@192.168.191.202:1521/data"
	log.Panicln("init database connect......")
	db, err := gorm.Open(oracle.Open(dsn), &gorm.Config{
		Logger: logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
			SlowThreshold: 1 * time.Millisecond,
			LogLevel:      logger.Warn,
			Colorful:      true,
		}),
	})
	if err != nil {
		log.Fatalln(err)
	}
	if e := db.AutoMigrate(inck.GroupDoc{}); e != nil {
		log.Fatalln(e.Error())
	}
	groupDoc := new(inck.GroupDoc)
	db.First(&groupDoc)
	fmt.Print(groupDoc)
}

func main() {
	db_oracle()
}
