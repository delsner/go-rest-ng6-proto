package controllers

import (
    "net/http"
    "encoding/json"
    "github.com/delsner/go-rest-ng6-proto/backend/services"
    "github.com/gin-gonic/gin"
    "github.com/delsner/go-rest-ng6-proto/backend/models"
)

var userService = services.GetUserServiceInstance()

func GetUserById(c *gin.Context) {
    userId := c.Param("userId")
    if userId != "" {
        res, err := json.Marshal(userService.GetById(userId))
        if err == nil {
            c.JSON(http.StatusOK, string(res))
            return
        }
    }
    c.JSON(http.StatusBadRequest, gin.H{"error": "Error"})
}

func GetAllUsers(c *gin.Context) {
    res, err := json.Marshal(userService.GetAll())
    if err == nil {
        c.JSON(http.StatusOK, string(res))
        return
    }
    c.JSON(http.StatusBadRequest, gin.H{"error": "Error"})
}

func CreateUser(c *gin.Context) {
    user := models.User{}
    // This c.ShouldBind consumes c.Request.Body and it cannot be reused. TODO: check ShouldBindJSON
    err := c.ShouldBind(&user)
    if err == nil {
        res, err := json.Marshal(userService.Create(&user))
        c.JSON(http.StatusOK, string(res))
        if err == nil {
            return
        }
    }
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
}
