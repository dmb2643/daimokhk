package listing

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"strconv"
)

type Handler struct {
	service service
}

func NewHandler(service service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) GetListings(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()

	rawPage := q.Get("page")
	if rawPage == "" {
		rawPage = "1"
	}

	page, err := strconv.Atoi(rawPage)
	if err != nil || page < 1 {
		WriteError(w, "page must be an integer than greater or equal to 1", http.StatusBadRequest)
		return
	}

	rawPerPage := q.Get("per_page")
	if rawPerPage == "" {
		rawPerPage = "20"
	}

	perPage, err := strconv.Atoi(rawPerPage)
	if err != nil || perPage < 10 || perPage > 100 {
		WriteError(w, "per_page must be an integer greated than or equal 10 and less than or equal 100", http.StatusBadRequest)
		return
	}

	pagination := pagination{
		limit:  perPage,
		offset: (page - 1) * perPage,
	}

	listings, err := h.service.getListings(r.Context(), pagination)
	if err != nil {
		WriteError(w, "internal server error", http.StatusInternalServerError)
		slog.Error("handler get listings: " + err.Error())
		return
	}

	resp := make([]ListingDTO, 0)
	for _, v := range listings {
		listing := ListingDTO{}

		listing.ID = v.ID
		listing.Category.ID = v.Category.ID
		listing.Category.Name = v.Category.Name
		listing.CreatedAt = v.CreatedAt
		listing.Description = v.Description
		listing.OwnerID = v.OwnerID
		listing.PhoneNumber = v.PhoneNumber
		listing.Price = v.Price
		listing.ProductCondition = v.ProductCondition
		listing.Status = v.Status
		listing.Title = v.Title
		listing.WhatsAppNumber = v.WhatsAppNumber

		resp = append(resp, listing)
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		slog.Error("encode listings: " + err.Error())
	}
}

func (h *Handler) GetListingByID(w http.ResponseWriter, r *http.Request) {
	url := r.PathValue("id")
	id, err := strconv.Atoi(url)

	if err != nil {
		WriteError(w, "id must be an integer greater than or equal to 1", http.StatusBadRequest)
		return
	}

	if id < 1 {
		WriteError(w, "id must be greater or equal to 1", http.StatusBadRequest)
		return
	}

	if id > 2147483647 {
		WriteError(w, "ID must be less than or equal to 2147483647", http.StatusBadRequest)
		return
	}

	listing, err := h.service.getListingByID(r.Context(), id)
	if err != nil {
		if errors.Is(err, ErrListingNotFound) {
			WriteError(w, ErrListingNotFound.Error(), http.StatusNotFound)
			return
		}

		if errors.Is(err, ErrListigRemovedByModerator) {
			WriteError(w, ErrListigRemovedByModerator.Error(), http.StatusForbidden)
			return
		}

		WriteError(w, "internal server error", http.StatusInternalServerError)
		slog.Error("hadnler: get listing:" + err.Error())
		return
	}

	resp := ListingDTO{
		ID:               listing.ID,
		Title:            listing.Title,
		Description:      listing.Description,
		Price:            listing.Price,
		PhoneNumber:      listing.PhoneNumber,
		WhatsAppNumber:   listing.WhatsAppNumber,
		OwnerID:          listing.OwnerID,
		Status:           listing.Status,
		CreatedAt:        listing.CreatedAt,
		ProductCondition: listing.ProductCondition,
		Category:         CategoryDto(listing.Category),
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		slog.Error("encode listing: " + err.Error())
	}
}
