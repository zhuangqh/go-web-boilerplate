package util

import (
    "net/http"
    "gopkg.in/kataras/iris.v6"
    "github.com/zhuangqh/go-web-boilerplate/db"
    "github.com/zhuangqh/go-web-boilerplate/logs"
)

func OPErrorHandler(ctx *iris.Context, error *db.OPError) {
    if error == nil {
        ctx.JSON(http.StatusOK, nil)
        return
    }

    status := error.Code / 100

    if status == 5 {
        ctx.EmitError(error.Code)
        logs.Logger.Error(error.Err)
    } else if error.Err != nil {
        ctx.Text(error.Code, error.Err.Error())
        logs.Logger.Warn(error.Err)
    } else {
        ctx.EmitError(error.Code)
    }
}
