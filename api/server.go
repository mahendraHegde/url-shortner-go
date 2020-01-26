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

	"github.com/mahendrahegde/url-shortner-golang/api/models"

	//swagger
	_ "github.com/mahendrahegde/url-shortner-golang/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Server struct {
	Db          *gorm.DB
	RouterGroup *gin.RouterGroup
	Cache       *redis.Client
}

const (
	PORT      = "PORT"
	dbHost    = "DB_HOST"
	dbPort    = "DB_PORT"
	dbUser    = "DB_USER"
	dbPass    = "DB_PASS"
	dbName    = "DB_NAME"
	redisHost = "REDIS_HOST"
	redisPort = "REDIS_PORT"
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

	//db setup
	connectionString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", GetEnv(dbHost), GetEnv(dbPort), GetEnv(dbUser), GetEnv(dbName), GetEnv(dbPass))
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
	redisString := fmt.Sprintf("%s:%s", GetEnv(redisHost), GetEnv(redisPort))
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
	url := ginSwagger.URL(fmt.Sprintf("http://localhost:%s/swagger/doc.json", GetEnv(PORT))) // The url pointing to API definition
	router := gin.Default()
	server.RouterGroup = router.Group("/v1")
	server.InitRotes()
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	if err := router.Run(":" + GetEnv(PORT)); err != nil {
		log.Fatal("unable to start server at port", GetEnv(PORT))
	}
}

func GetEnv(name string) string {
	return os.Getenv(name)
}
