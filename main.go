package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Request struct {
	ExecutableID string  `json:"executableId"`
	Latitude     float64 `json:"latitude"`
	Longitude    float64 `json:"longitude"`
}

func main() {
	http.HandleFunc("/request", requestHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func requestHandler(w http.ResponseWriter, r *http.Request) {
	// Ensure the request method is POST
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var req Request
	// Decode the request body
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Validate the request payload
	if req.ExecutableID == "" || req.Latitude == 0 || req.Longitude == 0 {
		http.Error(w, "Missing executableId, latitude, or longitude", http.StatusBadRequest)
		return
	}

	// Step 1: Determine Location
	region := determineLocation(req.Latitude, req.Longitude)

	// Step 2: Spin up VM and run executable
	instanceName, err := spinUpVMAndRunExecutable(region, req.ExecutableID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Step 4: Retrieve VM IP Address
	ipAddress, err := retrieveVMIPAddress(instanceName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return the final response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ipAddress)
}
