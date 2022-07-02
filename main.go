package main

import (
	"blog-service/internal/routers"
	"net/http"
	"time"
)

// func main() {
// 	r := gin.Default()
// 	r.GET("/ping", func(c *gin.Context) {
// 			c.JSON(200, gin.H{"message": "pong"})
// 	})
// 	r.Run()
// }
// curl http://localhost:8080/ping


func main() {
	router := routers.NewRouter()
	s := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}