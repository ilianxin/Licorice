package main

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	db := initDB()
	connectToETH(db)
	createRestAPI(db)

}

func connectToETH(db *gorm.DB) {
	ethClient, err := ethclient.Dial(os.Getenv("RPC_URL"))
	if err != nil {
		log.Fatal("Error connecting to Ethereum")
	}

	contractAddress := common.HexToAddress(os.Getenv("CONTRACT_ADDRESS"))

	return ethClient
}

func createRestAPI(db *gorm.DB) {
	router := gin.Default()

	router.GET("/user/:id", func(c *gin.Context) {
		idStr := c.Param("id")

		id, err := strconv.Atoi(idStr)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}

		user := getUser(db, id)
		if user.ID == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		c.JSON(http.StatusOK, user)

	})

	router.POST("/user", func(c *gin.Context) {
		var user User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		user = createUser(db, user)
		c.JSON(http.StatusOK, gin.H{"message": &user})
	})

	router.Run(":8080")
}

func initDB() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := gorm.Open(mysql.Open(os.Getenv("DB_URL")), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connecting to database")
	}

	db.AutoMigrate(&User{})

	return db

}

func createUser(db *gorm.DB, user User) User {
	db.Create(&user)
	return user
}

func getUser(db *gorm.DB, id int) User {
	var user User
	db.First(&user, id)
	return user
}

type User struct {
	ID    int    `json:"id" gorm:"primaryKey"`
	Name  string `json:"name" gorm:"not null"`
	Email string `json:"email" gorm:"not null"`
	Age   int    `json:"age" gorm:"not null"`
}

const contractABI = `
[
	{
		"inputs": [],
		"name": "getUser",
		"outputs": [
			{
				"internalType": "string",
				"name": "",
				"type": "string"
			}
		],
		"stateMutability": "view",
		"type": "function"
	}
]
`
