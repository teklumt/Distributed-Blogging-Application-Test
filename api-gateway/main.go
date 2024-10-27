package main

import (
	"api-gateway/middleware"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v2"
)

type Config struct {
    Services struct {
        Auth         string `yaml:"auth"`
        User         string `yaml:"user"`
        Blog         string `yaml:"blog"`
        Notification string `yaml:"notification"`
    } `yaml:"services"`
    JWTSecret string `yaml:"jwt_secret"`
}

var config Config

func initConfig() {
    data, err := ioutil.ReadFile("config.yaml")
    if err != nil {
        log.Fatalf("Error reading config file: %v", err)
    }
    if err := yaml.Unmarshal(data, &config); err != nil {
        log.Fatalf("Error parsing config file: %v", err)
    }
    os.Setenv("JWT_SECRET", config.JWTSecret)
}

func main() {
    initConfig()
    r := gin.Default()

    // Apply CORS 
    r.Use(middleware.CORSMiddleware())
    r.POST("/auth/*path", AuthHandler)

    // Apply JWT middleware
    r.Use(middleware.JWTAuthMiddleware())

    // Define routes and handlers
    r.Any("/users/*path", UserHandler)
    r.Any("/blogs/*path", BlogHandler)
    r.Any("/notifications/*path", NotificationHandler)

    // Start the server
    r.Run(":8080")
}


func AuthHandler(c *gin.Context) {
    proxyRequest(c, config.Services.Auth + c.Request.URL.Path)
}

func UserHandler(c *gin.Context) {
    proxyRequest(c, config.Services.User + c.Request.URL.Path)
}

func BlogHandler(c *gin.Context) {
    proxyRequest(c, config.Services.Blog + c.Request.URL.Path)
}

func NotificationHandler(c *gin.Context) {
    proxyRequest(c, config.Services.Notification + c.Request.URL.Path)
}

func proxyRequest(c *gin.Context, targetURL string) {
    resp, err := http.Post(targetURL, c.Request.Header.Get("Content-Type"), c.Request.Body)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to reach service"})
        return
    }
    c.DataFromReader(resp.StatusCode, resp.ContentLength, resp.Header.Get("Content-Type"), resp.Body, nil)
}
