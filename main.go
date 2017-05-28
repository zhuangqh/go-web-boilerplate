package main

import (
    "github.com/zhuangqh/go-web-boilerplate/api/user"
    "github.com/zhuangqh/go-web-boilerplate/logs"
    "github.com/zhuangqh/go-web-boilerplate/config"
    "gopkg.in/kataras/iris.v6"
    "gopkg.in/kataras/iris.v6/adaptors/gorillamux"
)

func main () {
    app := iris.New()
    app.Adapt(gorillamux.New())

    user.New(app, "/api")

    logs.Logger.Infof("server running on %s", config.Conf.Port)
    app.Listen(config.Conf.Port)
}
