package gateway

import (
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

func Start() {
	r := gin.Default()

	r.Any("/inventory/*path", reverseProxy("http://localhost:8001"))
	r.Any("/orders/*path", reverseProxy("http://localhost:8002"))

	r.Run(":8000") // API Gateway listens on 8000
}

func reverseProxy(target string) gin.HandlerFunc {
	return func(c *gin.Context) {
		remote, err := url.Parse(target)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid target"})
			return
		}
		proxy := httputil.NewSingleHostReverseProxy(remote)
		c.Request.URL.Path = c.Param("path")
		proxy.ServeHTTP(c.Writer, c.Request)
	}
}
