package main

import (
	"net/http/httputil"
	"net/url"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
    // Load environment variables
    if err := godotenv.Load(); err != nil {
        panic("Error loading .env file")
    }

    r := gin.Default()

    // Route requests to Auth Service
    r.Any("/auth/*proxyPath", reverseProxy(os.Getenv("AUTH_SERVICE_URL")))

    // Route requests to Blog Service
    r.Any("/blog/*proxyPath", reverseProxy(os.Getenv("BLOG_SERVICE_URL")))
    r.Any("/comment/*proxyPath", reverseProxy(os.Getenv("BLOG_SERVICE_URL")))
	
	
    r.Any("/user/*proxyPath", reverseProxy(os.Getenv("USER_SERVICE_URL")))

    // Route requests to Notification Service
    r.Any("/notification/*proxyPath", reverseProxy(os.Getenv("NOTIFICATION_SERVICE_URL")))

    r.Run(":8080") 
}

func reverseProxy(target string) gin.HandlerFunc {
    url, _ := url.Parse(target)
    proxy := httputil.NewSingleHostReverseProxy(url)

    return func(c *gin.Context) {
        proxy.ServeHTTP(c.Writer, c.Request)
    }
}
