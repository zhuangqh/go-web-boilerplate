package db

import (
    "gopkg.in/mgo.v2"
    "time"
    "log"
    "github.com/zhuangqh/go-web-boilerplate/config"
)

var MgoSession *mgo.Session
var Database string

func init() {
    Database = config.Conf.Database

    mgoDBDialInfo := &mgo.DialInfo{
        Addrs:    []string{config.Conf.DBConf.Host},
        Username:  config.Conf.DBConf.Username,
        Password:  config.Conf.DBConf.Password,
        Timeout:  30 * time.Second,
        Database: config.Conf.Database,
    }

    var err error
    MgoSession, err = mgo.DialWithInfo(mgoDBDialInfo)
    if err != nil {
        log.Fatalf("cannot dial mongodb: %s\n", err)
    }

    MgoSession.SetMode(mgo.Monotonic, true)
}
