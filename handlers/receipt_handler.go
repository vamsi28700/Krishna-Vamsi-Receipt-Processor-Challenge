package handlers

import (
    "encoding/json"
    "net/http"
    "regexp"
    "github.com/gorilla/mux"
    "github.com/google/uuid"

    "receipt-processor/models"
    "receipt-processor/utils"
)

var receipts = make(map[string]models.Receipt)
var retailerNamePattern = regexp.MustCompile("^[\\w\\s\\-&]+$")

func ProcessReceipt(w http.ResponseWriter, r *http.Request) {
    var receipt models.Receipt
    err := json.NewDecoder(r.Body).Decode(&receipt)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Validate retailer name
    if !retailerNamePattern.MatchString(receipt.Retailer) {
        http.Error(w, "Invalid retailer name format", http.StatusBadRequest)
        return
    }

    receipt.ID = uuid.New().String()
    receipt.Points = utils.CalculatePoints(receipt)
    receipts[receipt.ID] = receipt

    response := map[string]string{"id": receipt.ID}
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}

func GetPoints(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]

    receipt, exists := receipts[id]
    if !exists {
        http.Error(w, "Receipt not found", http.StatusNotFound)
        return
    }

    response := map[string]int{"points": receipt.Points}
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}
