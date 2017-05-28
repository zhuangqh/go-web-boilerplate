package user

import (
    "github.com/zhuangqh/go-web-boilerplate/util"
    "net/http"
    "gopkg.in/kataras/iris.v6"
    "github.com/zhuangqh/go-web-boilerplate/db"
    "github.com/zhuangqh/go-web-boilerplate/logs"
)

type Controller struct {
    context    *iris.Framework
}

func New(context *iris.Framework, prefix string) Controller {
    c := Controller{context}
    c.Routes(prefix)
    return c
}

func (this *Controller) Routes(prefix string) {
    router := this.context.Party(prefix + "/user")

    router.Get("", this.GetMany)
    router.Post("", this.CreateOne)
}

func (this *Controller) GetMany(ctx *iris.Context) {
    data, err := GetMany()

    if err != nil {
        util.OPErrorHandler(ctx, err)
        return
    }

    ctx.JSON(http.StatusOK, data)
}

func (this *Controller) CreateOne(ctx *iris.Context) {
    var param Param
    if err := ctx.ReadJSON(&param); err != nil {
        util.OPErrorHandler(ctx, &db.OPError{400, err})
        return
    }

    data, err := CreateOne(&param)
    if err != nil {
        util.OPErrorHandler(ctx, err)
        return
    }

    logs.Logger.Infof("Create a new user %+v", data)
}
