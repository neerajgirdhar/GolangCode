package main

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
)

type trainsSchedule struct {
	ID                 string  `json:"id"`
	TrainName          string  `json:"trainname"`
	StartingStation    string  `json:"startingstation"`
	TerminatingStation string  `json:"terminatingstation"`
	Days               string  `json:"days"`
	Fare               int     `json:"fare"`
	StartTime          string  `json:"starttime"`
	EndTime            string  `json:"endtime"`
	Tickets            Tickets `json:"tickets"`
}

type bookings struct {
	BookingID      int    `json:"bookingid"`
	CustomerUserID string `json:"customeruserid"`
	CustomerName   string `json:"customername"`
	TrainID        string `json:"trainid"`
	TotalTickets   int    `json:"totaltickets"`
	Date           string `json:"date"`
	TicketType     string `json:"tickettype"`
}

var allBookings []bookings

type Tickets struct {
	AC2Tier               string `json:"ac2tier"`
	AC3Tier               string `json:"ac3tier"`
	General               string `json:"general"`
	AC2TierSeatsAvailable int    `json:"ac2TierSeatsAvailable"`
	AC3TierSeatsAvailable int    `json:"ac3TierSeatsAvailable"`
	GeneralSeatsAvailable int    `json:"generalSeatsAvailable"`
}

var tickets = Tickets{"Y", "Y", "T", 600, 200, 0}

var alltrains = []trainsSchedule{
	{"VB23010", "VandeBharat", "Chandigarh", "Jodhpur", "S-M-T-W-T-F-S", 1000, "15:15", "23:30", tickets},
	{"VB23011", "VandeBharat", "Jodhpur", "Chandigarh", "S-M-T-W-T-F-S", 1000, "6:15", "14:30", tickets},

	{"ST45689", "Shatabdi Express", "Chandigarh", "Delhi", "S-M-T-W-T-F-S", 800, "18:15", "20:30", tickets},
}

func getAllTrains(c *gin.Context) {
	fmt.Println("getAllTrains", alltrains)
	c.IndentedJSON(http.StatusOK, alltrains)

	return
}

func addNewTrain(c *gin.Context) {
	//created a new struct of the same type(trainsSchedule)
	var train trainsSchedule
	if err := c.BindJSON(&train); err != nil {
		return
	}

	//mrerged or appended the new train strcy with the existing struct
	alltrains = append(alltrains, train)
	c.IndentedJSON(http.StatusOK, alltrains)
	return
}

func getTrainByID(c *gin.Context) {
	fmt.Println("In getTrainByID")
	id := c.Param("id")

	train, err := getTrain(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, train)
	return
}

func getTrain(id string) (*trainsSchedule, error) {
	for i, train := range alltrains {
		if train.ID == id {
			return &alltrains[i], nil
		}

	}
	return nil, errors.New("Train With This ID Not Present | Try With Valid ID")
}

func getTrainByStation(c *gin.Context) {
	startingstation, _ := c.GetQuery("startingstation")
	terminatingstation, _ := c.GetQuery("terminatingstation")

	if startingstation == "" || terminatingstation == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Missing Startingstation or Terminatingstation"})
		return
	}

	train, err := getTrainBetweenStation(startingstation, terminatingstation)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, train)
	return
}

func getTrainBetweenStation(startingstation, terminatingstation string) (*trainsSchedule, error) {

	for i, train := range alltrains {
		if train.StartingStation == startingstation && train.TerminatingStation == terminatingstation {
			return &alltrains[i], nil
		}

	}
	return nil, errors.New("Train With This startingstation and terminatingstation not present | Try With Valid startingstation and  terminatingstation")

}

func bookTicketsAndUpdateCount(c *gin.Context) {
	var newBooking bookings
	if err := c.BindJSON(&newBooking); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Not Valid Booking JSON"})
		return
	}
	newAllTrains := &alltrains
	err := updateBookingInfoIntoAvailability(newAllTrains, newBooking.Date, newBooking.TrainID, newBooking.TicketType, newBooking.TotalTickets)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newBooking.BookingID = rand.Int()
	allBookings = append(allBookings, newBooking)
	c.IndentedJSON(http.StatusOK, allBookings)
	return
}

func updateBookingInfoIntoAvailability(alltrains1 *[]trainsSchedule, bookingdate, trainid, tickettype string, ticketsbooked int) error {
	fmt.Println("bookingdate ", bookingdate)
	fmt.Println("trainid ", trainid)
	fmt.Println("tickettype ", tickettype)
	fmt.Println("ticketsbooked ", ticketsbooked)
	checkValidTrain := false

	var updatedtrainsInfo []trainsSchedule

	for _, train := range *alltrains1 {

		if train.ID == trainid {
			checkValidTrain = true

			if tickettype == "ac2tier" {
				fmt.Println("ac2tier ")
				if train.Tickets.AC2TierSeatsAvailable >= ticketsbooked {
					fmt.Println("tickets deducted from ac2tier quota")
					train.Tickets.AC2TierSeatsAvailable = train.Tickets.AC2TierSeatsAvailable - ticketsbooked
					fmt.Println("*********** remaining tickets after deduction from ac2tier quota", train.Tickets.AC2TierSeatsAvailable)
				} else {
					return errors.New("Tickets are not available in this category")
				}
			}
			if tickettype == "ac3tier" {
				if train.Tickets.AC3TierSeatsAvailable >= ticketsbooked {
					fmt.Println("tickets deducted from ac3tier quota")
					train.Tickets.AC3TierSeatsAvailable = train.Tickets.AC3TierSeatsAvailable - ticketsbooked
					fmt.Println("*********** remaining tickets after deduction from ac3tier quota", train.Tickets.AC3TierSeatsAvailable)
				} else {
					return errors.New("Tickets are not available in this category")
				}
			}
			if tickettype == "general" {

				if train.Tickets.GeneralSeatsAvailable >= ticketsbooked {
					fmt.Println("tickets deducted from general quota")
					train.Tickets.GeneralSeatsAvailable = train.Tickets.GeneralSeatsAvailable - ticketsbooked
					fmt.Println("*********** remaining tickets after deduction from general quota", train.Tickets.GeneralSeatsAvailable)

				} else {
					fmt.Println("tickets are not available in this category")
					return errors.New("Tickets are not available in this category")
				}
			}
			if tickettype != "general" && tickettype != "ac3tier" && tickettype != "ac2tier" {
				return errors.New("This  tickettype Category is not available in the Train")
			}
			updatedtrainsInfo = append(updatedtrainsInfo, train)
			//break
		} else {
			updatedtrainsInfo = append(updatedtrainsInfo, train)
		}

	}
	if checkValidTrain == false {
		return errors.New("Invalid Train ID")
	}
	*alltrains1 = updatedtrainsInfo

	return nil
}

func main() {
	router := gin.Default()
	router.GET("/trains", getAllTrains)
	router.POST("/addNewTrain", addNewTrain)
	router.GET("/getTrainById/:id", getTrainByID)
	router.GET("/getTrainBetweenStations", getTrainByStation)
	router.POST("/bookTickets", bookTicketsAndUpdateCount)
	router.Run("localhost:8080")
}
