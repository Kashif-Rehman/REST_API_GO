package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// User represents a user with first name, last name, and email.
type User struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

// Profile represents a user profile with department, designation, and associated employee.
type Profile struct {
	Department  string `json:"department"`
	Designation string `json:"designation"`
	Employee    User   `json:"employee"`
}

// Slice to store profiles.
var profiles = []Profile{}

func main() {
	// Create a new Gorilla mux router.
	r := mux.NewRouter()

	// Define routes for curd profile.
	r.HandleFunc("/profiles", addItem).Methods("POST")
	r.HandleFunc("/profiles", getItems).Methods("GET")
	r.HandleFunc("/profiles/{id}", getItem).Methods("GET")
	r.HandleFunc("/profiles/{id}", updateItem).Methods("PUT")
	r.HandleFunc("/profiles/{id}", deleteItem).Methods("DELETE")

	// Start the HTTP server on port 5000 with the Gorilla mux router.
	http.ListenAndServe(":5000", r)
}

// addItem is an HTTP handler function to add a new profile.

func getItem(w http.ResponseWriter, r *http.Request) {
	idParm := mux.Vars(r)["id"]

	id, err := strconv.Atoi(idParm)

	if err != nil {

		w.WriteHeader(400)
		w.Write([]byte("ID could not be converted to Integer"))
		return

	}

	if id >= len(profiles) {
		w.WriteHeader(404)
		w.Write([]byte("ID not found"))

		return
	}

	profile := profiles[id]

	w.Header().Set("Content-Type", "application/json")

	// Use the JSON encoder to encode the 'profiles' slice and send it as the response.
	json.NewEncoder(w).Encode(profile)

}

func updateItem(w http.ResponseWriter, r *http.Request) {
	idParm := mux.Vars(r)["id"]

	id, err := strconv.Atoi(idParm)

	if err != nil {

		w.WriteHeader(400)
		w.Write([]byte("ID could not be converted to Integer"))
		return

	}

	if id >= len(profiles) {
		w.WriteHeader(404)
		w.Write([]byte("ID not found"))

		return
	}

	var newProfile Profile

	// Decode the JSON data from the request body into the newProfile variable.
	err = json.NewDecoder(r.Body).Decode(&newProfile)
	if err != nil {
		// If there is an error in decoding, send a Bad Request response.
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Set the response header to indicate JSON content.
	w.Header().Set("Content-Type", "application/json")

	// Append the new profile to the existing profiles slice.

	profiles[id] = newProfile
	// Encode the updated profiles slice and send it as the response.
	json.NewEncoder(w).Encode(profiles)

}

func addItem(w http.ResponseWriter, r *http.Request) {
	var newProfile Profile

	// Decode the JSON data from the request body into the newProfile variable.
	err := json.NewDecoder(r.Body).Decode(&newProfile)
	if err != nil {
		// If there is an error in decoding, send a Bad Request response.
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Set the response header to indicate JSON content.
	w.Header().Set("Content-Type", "application/json")

	// Append the new profile to the existing profiles slice.
	profiles = append(profiles, newProfile)

	// Encode the updated profiles slice and send it as the response.
	json.NewEncoder(w).Encode(profiles)
}

func getItems(w http.ResponseWriter, r *http.Request) {
	// Set the response header to indicate JSON content type.
	w.Header().Set("Content-Type", "application/json")

	// Use the JSON encoder to encode the 'profiles' slice and send it as the response.
	json.NewEncoder(w).Encode(profiles)
}

func deleteItem(w http.ResponseWriter, r *http.Request) {
	idParm := mux.Vars(r)["id"]

	id, err := strconv.Atoi(idParm)

	if err != nil {

		w.WriteHeader(400)
		w.Write([]byte("ID could not be converted to Integer"))
		return

	}

	if id >= len(profiles) {
		w.WriteHeader(404)
		w.Write([]byte("ID not found"))

		return
	}

	profiles = deleteElement(profiles, id)

	w.Write([]byte("Item deleted"))

}

func deleteElement(slice []Profile, index int) []Profile {
	// Check if the index is out of bounds
	if index < 0 || index >= len(slice) {
		return slice // No changes needed, index is out of bounds
	}

	// Create a new slice with the element at the specified index removed
	result := append(slice[:index], slice[index+1:]...)

	return result
}
