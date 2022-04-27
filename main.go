package main

import (
	"github.com/gin-gonic/gin"
	"github.com/replit/database-go"
  "strconv"
  "fmt"
)

func main() {
	r := gin.Default()
  database.Set("root", "0")

  r.GET("/", func(c *gin.Context) {   
    val, err := database.Get("root")
    if err != nil {
      fmt.Println(err)
    } else {
      val_i, _ := strconv.Atoi(val)
      val_i += 1
      val = strconv.Itoa(val_i)
      database.Set("root", val)
    }
      
    c.String(200, val)
  })
  
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

  r.POST("/write", func(c *gin.Context) {
    type ReqBody struct {
      Key string    `json:"key"`
      Value string  `json:"value"`
    }

    var kv ReqBody

    if err := c.BindJSON(&kv); err != nil {
       fmt.Println("fail to parse req body")
    }
    
    database.Set(kv.Key, kv.Value)

    c.String(200, "Successful write")
  })

  r.GET("/read/:key", func(c *gin.Context) {
    key := c.Param("key")
    
    val, err := database.Get(key)
    if err != nil {
      c.String(404, "key not found")
    }

    c.String(200, val)
  })
	
  r.Run()
}