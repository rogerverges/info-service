package main

import (   
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type cars struct {
	Brand, Model, HorsePower string
	ID                       string
}

func main() {
	
	http.HandleFunc("/services/v1/cars/", getCars)
	direccion := ":8080"
	fmt.Println("Servidor listo escuchando en " + direccion)
	log.Fatal(http.ListenAndServe(direccion, nil))

}

func getCars(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		fmt.Printf("roger")
		ID := r.URL.Path[18:]

		obtenerCoches(ID, w)

	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

func obtenerBaseDeDatos() (db *sql.DB, e error) {
	usuario := "root"
	pass := "roger"
	host := "tcp(database)"
	nombreBaseDeDatos := "entrust"
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@%s/%s", usuario, pass, host, nombreBaseDeDatos))
	if err != nil {
		return nil, err
	}
	return db, nil
}

func obtenerCoches(ID string, w http.ResponseWriter) ([]cars, error) {

	db, err := obtenerBaseDeDatos()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	filas, err := db.Query("SELECT * FROM coche WHERE id = " + ID)

	if err != nil {
		return nil, err
	}
	defer filas.Close()

	var receivedCar cars

	for filas.Next() {
		err = filas.Scan(&receivedCar.ID, &receivedCar.Brand, &receivedCar.Model, &receivedCar.HorsePower)
		if err != nil {
			return nil, err
		}
		json.NewEncoder(w).Encode(receivedCar)
		json.NewEncoder(w).Encode(receivedCar)
		
	}
	return nil, nil
}
