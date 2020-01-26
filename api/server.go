package api

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v7"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"

	"github.com/mahendrahegde/url-shortner-golang/api/middlewares"
	"github.com/mahendrahegde/url-shortner-golang/api/models"

	//swagger
	_ "github.com/mahendrahegde/url-shortner-golang/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Env struct {
	PORT          string
	DB_HOST       string
	DB_PORT       string
	DB_USER       string
	DB_PASS       string
	DB_NAME       string
	REDIS_HOST    string
	REDIS_PORT    string
	START_SEQ     string
	API_VERSION   string
	SERVER_DOMAIN string
}

type Server struct {
	Db          *gorm.DB
	RouterGroup *gin.RouterGroup
	Cache       *redis.Client
	ENV         Env
}

const (
	PORT        = "PORT"
	dbHost      = "DB_HOST"
	dbPort      = "DB_PORT"
	dbUser      = "DB_USER"
	dbPass      = "DB_PASS"
	dbName      = "DB_NAME"
	redisHost   = "REDIS_HOST"
	redisPort   = "REDIS_PORT"
	API_VERSION = "API_VERSION"
)

// @title  API Docs
// @version V1
// @description This is a sample url shortner service.
// @BasePath /v1
func (server *Server) Start() {
	//env setup
	if err := godotenv.Load(); err != nil {
		log.Print("Error loading .env file")
	}
	log.Println("envs loaded")
	server.ENV = Env{PORT: os.Getenv("PORT"), DB_HOST: os.Getenv("DB_HOST"), DB_PORT: os.Getenv("DB_PORT"), DB_USER: os.Getenv("DB_USER"), DB_PASS: os.Getenv("DB_PASS"), DB_NAME: os.Getenv("DB_NAME"), REDIS_HOST: os.Getenv("REDIS_HOST"), REDIS_PORT: os.Getenv("REDIS_PORT"), START_SEQ: os.Getenv("START_SEQ"), API_VERSION: os.Getenv("API_VERSION"), SERVER_DOMAIN: os.Getenv("SERVER_DOMAIN")}

	//db setup
	connectionString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", server.ENV.DB_HOST, server.ENV.DB_PORT, server.ENV.DB_USER, server.ENV.DB_NAME, server.ENV.DB_PASS)
	log.Println("connecting to.." + connectionString)
	db, err := gorm.Open("postgres", connectionString)
	if err != nil {
		log.Fatal("falied to connect to db,", err.Error())
	}
	log.Println("connected to DB...")
	server.Db = db
	defer server.Db.Close()
	models.Migrate(server.Db)

	//redis
	redisString := fmt.Sprintf("%s:%s", server.ENV.REDIS_HOST, server.ENV.REDIS_PORT)
	client := redis.NewClient(&redis.Options{
		Addr:     redisString,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	if _, err := client.Ping().Result(); err != nil {
		log.Fatal("unable to connect to redis,", redisString)
	}
	server.Cache = client

	//swagger
	url := ginSwagger.URL(fmt.Sprintf("http://localhost:%s/swagger/doc.json", server.ENV.PORT)) // The url pointing to API definition
	router := gin.Default()
	router.Use(middlewares.DummyMiddleware())
	server.RouterGroup = router.Group(server.ENV.API_VERSION)
	server.InitRotes()
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	if err := router.Run(":" + server.ENV.PORT); err != nil {
		log.Fatal("unable to start server at port", server.ENV.PORT)
	}
}
