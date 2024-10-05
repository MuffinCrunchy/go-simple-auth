package services

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"simple-auth/models"
	"time"
)

var dataFile = "repository/data.json"

func Login(customerID string) error {
	data, err := ioutil.ReadFile(dataFile)
	if err != nil {
		return err
	}

	var database map[string][]models.Customer
	err = json.Unmarshal(data, &database)
	if err != nil {
		return err
	}

	customers := database["customers"]
	for i, customer := range customers {
		if customer.ID == customerID {
			if customer.LoggedIn {
				return errors.New("customer already logged in")
			}

			customers[i].LoggedIn = true
			logAction(customerID, "login")
			return saveData(database)
		}
	}

	return errors.New("customer not found")
}

func saveData(data map[string][]models.Customer) error {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(dataFile, jsonData, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}
