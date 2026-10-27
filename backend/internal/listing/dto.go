package listing

import (
	"encoding/json"
	"net/http"
	"time"
)

type CategoryDto struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type ListingDTO struct {
	ID               int              `json:"id"`
	Title            string           `json:"title"`
	Description      string           `json:"description"`
	Price            int              `json:"price"`
	PhoneNumber      string           `json:"phone_number"`
	WhatsAppNumber   *string          `json:"whatsapp_number"`
	OwnerID          int              `json:"owner_id"`
	Status           listingStatus    `json:"status"`
	CreatedAt        time.Time        `json:"created_at"`
	ProductCondition listingCondition `json:"product_condition"`
	Category         CategoryDto      `json:"category"`
}

type ErrorResponce struct {
	Msg string `json:"message"`
}

func WriteError(w http.ResponseWriter, msg string, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(ErrorResponce{Msg: msg})
}
