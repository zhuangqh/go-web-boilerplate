package db

import (
    "github.com/zhuangqh/go-web-boilerplate/config"
    _ "github.com/jinzhu/gorm/dialects/mysql"
    "github.com/jinzhu/gorm"
    "log"
    "fmt"
    "os"
)

var DB *gorm.DB

func init() {
    connPath := fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
        config.Conf.Username, config.Conf.Password, config.Conf.Database)

    var err error
    DB, err = gorm.Open("mysql", connPath)
    if err != nil {
        log.Fatalf("cannot dial database: %s\n", err)
        os.Exit(1)
    }

    // DB.LogMode(true)
}
