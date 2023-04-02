package main

import (
	"fmt"

	"github.com/jutionck/golang-db-sinar-harapan-makmur-orm/config"
	"github.com/jutionck/golang-db-sinar-harapan-makmur-orm/model"
)

func main() {
	c, err := config.NewConfig()
	if err != nil {
		fmt.Println(err)
	}

	dbConn, _ := config.NewDbConnection(c)
	db := dbConn.Conn()

	// dbConn.Migrate(
	// 	&model.UserCredential{},
	// 	&model.Customer{},
	// 	&model.Brand{},
	// 	&model.Vehicle{},
	// 	&model.Employee{},
	// 	&model.Transaction{},
	// )

	// Update kendaraan
	// Vehicle -> 726924a7-7da1-4791-8a1b-8d34d1e58686
	// Customer -> d34661d9-fbae-4c9f-ab07-e1b277224367

	var customer model.Customer
	if err := db.Debug().Preload("Vehicles").Where("id=?", "d34661d9-fbae-4c9f-ab07-e1b277224367").First(&customer).Error; err != nil {
		fmt.Println(err)
	}

	var newVehicle model.Vehicle
	if err := db.Debug().Where("id=?", "4d10acb7-33c9-4787-b186-374ddde39ae2").First(&newVehicle).Error; err != nil {
		fmt.Println(err)
	}

	var oldVehicleID = "726924a7-7da1-4791-8a1b-8d34d1e58686"
	var newVehicles []model.Vehicle
	for _, cv := range customer.Vehicles {
		if cv.ID != oldVehicleID {
			newVehicles = append(newVehicles, cv)
		} else {
			newVehicles = append(newVehicles, newVehicle)
		}
	}

	if err := db.Debug().Model(&customer).Association("Vehicles").Replace(newVehicles); err != nil {
		fmt.Println(err)
	}
}
