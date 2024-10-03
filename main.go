package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

// Creating Struct for create an api interface for booking data
type Booking struct {
	ID       int    `json:"id"`
	Doctor   string `json:"doctor"`
	Patient  string `json:"patient"`
	Date     string `json:"date"`
	Time     string `json:"time"`
}

// Declaring slice to hold the bokking data
var bookings []Booking

// Creating new booking or POST
func createBooking(w http.ResponseWriter, r *http.Request) {
	var newBooking Booking
	_ = json.NewDecoder(r.Body).Decode(&newBooking)
	newBooking.ID = len(bookings) + 1
	bookings = append(bookings, newBooking)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newBooking)
}

// Read the existing booking details by its ID or GET by ID
func getBooking(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	for _, booking := range bookings {
		if booking.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(booking)
			return
		}
	}
	http.Error(w, "Booking could not found", http.StatusNotFound)
}

// Read all the bookings or GET
func getAllBookings(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bookings)
}

// Update the exisitng booking or PUT
func updateBooking(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	for i, booking := range bookings {
		if booking.ID == id {
			_ = json.NewDecoder(r.Body).Decode(&bookings[i])
			bookings[i].ID = id // Ensure ID remains unchanged
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(bookings[i])
			return
		}
	}
	http.Error(w, "Booking could not found", http.StatusNotFound)
}

// Delete the existing booking or DELETE
func deleteBooking(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	for i, booking := range bookings {
		if booking.ID == id {
			bookings = append(bookings[:i], bookings[i+1:]...)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(bookings)
			return
		}
	}
	http.Error(w, "Booking could not found", http.StatusNotFound)
}

func main() {
	
	r := mux.NewRouter()


	r.HandleFunc("/bookings", createBooking).Methods("POST")
	r.HandleFunc("/bookings/{id}", getBooking).Methods("GET")
	r.HandleFunc("/bookings", getAllBookings).Methods("GET")
	r.HandleFunc("/bookings/{id}", updateBooking).Methods("PUT")
	r.HandleFunc("/bookings/{id}", deleteBooking).Methods("DELETE")


	fmt.Println("Server is running on port 8000...")
	log.Fatal(http.ListenAndServe(":8000", r))
}