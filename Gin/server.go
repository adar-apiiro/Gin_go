package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

type Item struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}

var items []Item

func main() {
    r := gin.Default()

    r.GET("/items", GetItems)
    r.POST("/items", AddItem)

    r.Run(":8080")
}

func GetItems(c *gin.Context) {
    c.JSON(http.StatusOK, items)
}

func AddItem(c *gin.Context) {
    var newItem Item
    if err := c.ShouldBindJSON(&newItem); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    newItem.ID = len(items) + 1
    items = append(items, newItem)
    c.JSON(http.StatusCreated, newItem)
}
