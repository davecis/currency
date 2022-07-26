package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/adrg/exrates"
	"github.com/davecis/currency-api/db"
	"github.com/davecis/currency-api/models"
)

func GetCurrencyHandler(w http.ResponseWriter, r *http.Request) {
	var currencies []models.Currency
	db.DB.Find(&currencies)
	json.NewEncoder(w).Encode(&currencies)
}

func GetCurrency(w http.ResponseWriter, r *http.Request) {

	startTime := time.Now()

	rates, err := exrates.Latest("USD", nil)
	if err != nil {

		endTime := time.Now()
		difference := startTime.Sub(endTime)

		total := float64(difference.Seconds())

		fmt.Printf("Guardando en Bitacora")
		SaveBitacora(false, total)

		w.WriteHeader(http.StatusNoContent)
		w.Write([]byte(err.Error()))
		return
	}

	endTime := time.Now()
	difference := startTime.Sub(endTime)

	total := float64(difference.Seconds())

	fmt.Printf("Guardando en Bitacora")
	SaveBitacora(true, total)

	for currency, value := range rates.Values {
		cur := models.Currency{
			Code:  currency,
			Value: value,
		}

		createdCurrency := db.DB.Create(&cur)
		err := createdCurrency.Error

		if err != nil {
			fmt.Printf("%#v\n", err.Error())
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			break
		}
	}

	w.WriteHeader(http.StatusCreated)
}

func SaveBitacora(status bool, timer float64) {

	task := models.Task{
		Time:   timer,
		Status: status,
	}

	saveBitacora := db.DB.Create(&task)
	err := saveBitacora.Error

	if err != nil {
		fmt.Printf("%#v\n", err.Error())
		return
	}

}
