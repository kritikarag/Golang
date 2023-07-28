package main

import(
	"net/http"
	"github.com/gin-gonic/gin"
)

func getapi(c *gin.context)(string){
	var mp map[int]string

	mp[1] = "ABC"
	mp[2] = "DEF"
	mp[3] = "GHI"

	var id = c.Param("id")

	value, err:= mp[id]

	if err!= NULL{
		return err
	}

	c.JSON(http.ResponseOK,gin.H{"val:":"value"})

}