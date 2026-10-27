package listing

import "time"

type listingStatus string

const (
	statusActive    listingStatus = "активно"
	statusClosed    listingStatus = "закрыто"
	statusPostponed listingStatus = "отложено"
	statusRemoved   listingStatus = "снято модератором"
)

type listingCondition string

const (
	conditionNew  listingCondition = "новый"
	conditionUsed listingCondition = "б/у"
)

type Category struct {
	ID   int
	Name string
}

type Listing struct {
	ID               int
	Title            string
	Description      string
	Price            int
	PhoneNumber      string
	WhatsAppNumber   *string
	OwnerID          int
	Status           listingStatus
	CreatedAt        time.Time
	ProductCondition listingCondition
	Category         Category
}

type pagination struct {
	limit  int
	offset int
}
