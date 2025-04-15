package main

import (
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	inventoryURL, _ := url.Parse("http://inventory-service:8081")
	inventoryProxy := httputil.NewSingleHostReverseProxy(inventoryURL)
	r.Any("/products/*proxyPath", func(c *gin.Context) {
		inventoryProxy.ServeHTTP(c.Writer, c.Request)
	})

	orderURL, _ := url.Parse("http://order-service:8082")
	orderProxy := httputil.NewSingleHostReverseProxy(orderURL)
	r.Any("/orders/*proxyPath", func(c *gin.Context) {
		orderProxy.ServeHTTP(c.Writer, c.Request)
	})

	r.Run(":8080")
}
