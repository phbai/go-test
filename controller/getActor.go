package controller

import (
    "github.com/gin-gonic/gin"
    "github.com/phbai/api/logic"
    "github.com/phbai/api/util"
    "github.com/bitly/go-simplejson"
)

func ActorController(c *gin.Context) {
    result := make(chan *simplejson.Json)
    var params util.Actor

    params.PageIndex = c.Query("index");
    params.PageSize = c.Query("size");

    go logic.GetActor(params, result)
    // res := <- result

    c.JSON(200, gin.H{
        "result": <- result,
    })
}