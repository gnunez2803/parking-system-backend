package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// ... data structures (ParkingSpot, Reservation, User) ...

func main() {
	router := gin.Default()

	x := 123
	fmt.Println(x)
	edgar := "my name is Edgar"
	fmt.Println(edgar)
	y := time.Now()
	fmt.Println(y)
	router.POST("/reservations", createReservation)
	router.GET("/available", getAvailableLocations)
	router.GET("/pricing/:spotID", getPricing)
	router.GET("/user/:userID/reservations", getUserReservations)

	router.Run(":8080")
}

func createReservation(c *gin.Context) {
	fmt.Println("created reservations")
	c.JSON(http.StatusCreated, gin.H{"message": "Reservation created"})
}

func getAvailableLocations(c *gin.Context) {
	fmt.Println("locations")
	c.JSON(http.StatusOK, gin.H{"message": "availibility"})
}

func getPricing(c *gin.Context) {
	fmt.Println("pricing")
	c.JSON(http.StatusOK, gin.H{"message": "Pricing"})
}

func getUserReservations(c *gin.Context) {
	fmt.Println("here are my reservations")
	c.JSON(http.StatusOK, gin.H{"message": "get my reservations"})
}
