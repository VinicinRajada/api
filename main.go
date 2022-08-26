package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/quaglio/oficina/controller"
	"github.com/quaglio/oficina/database"
)

func main() {
	if err := database.Connect(); err != nil {
		panic(err)
	}

	db := database.GrabDB()
	st := db.MustBegin()
	sql := "CREATE TABLE IF NOT EXISTS usuarios (nome VARCHAR(128), RA VARCHAR(12)"
	if err := st.MustExec(sql); err != nil {
		panic(err)
	}

	r := gin.Default()

	r.POST("/usuario", func(c *gin.Context) {
		if err := controller.PostAluno(c); err != nil {
			panic(err)
		}
	})

	r.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))

}
