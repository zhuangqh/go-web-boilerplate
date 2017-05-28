package logs

import (
    "github.com/cihub/seelog"
    "fmt"
    "github.com/zhuangqh/go-web-boilerplate/config"
)

var Logger seelog.LoggerInterface

func init() {
    Logger = seelog.Disabled

    logger, err := seelog.LoggerFromConfigAsFile(config.Conf.LogConfig)

    if err != nil {
        fmt.Println(err.Error())
        return
    }

    Logger = logger
}

