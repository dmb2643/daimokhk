package listing

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func TestGetByID(t *testing.T) {
	ctx := context.Background()
	connectContext, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	id := 2

	psString := os.Getenv("POSTGRESQL_STRING")
	if psString == "" {
		t.Skip("PostgreSQL string is not set")
	}

	pool, err := pgxpool.New(connectContext, psString)
	if err != nil {
		t.Fatalf("create pool: %v", err)
	}

	defer pool.Close()

	postgreDB := NewPostgreSQLRepo(pool)

	listing, err := postgreDB.getByID(connectContext, id)
	if err != nil {
		t.Fatalf("get listing: %v", err)
	}

	if listing.ID != 2 {
		t.Errorf("ID: get %d, want %d", listing.ID, id)
	}

	if listing.Title != "видеокарта rtx 3080" {
		t.Errorf("Title: get %s, want %s", listing.Title, "видеокарта rtx 3080")
	}

	if listing.Description != "Использовалась около 2 лет" {
		t.Errorf("Description: get %s, want %s", listing.Description, "Использовалась около 2 лет")
	}

	if listing.Price != 45000 {
		t.Errorf("Price: get %d, want %d", listing.Price, 45000)
	}

	if listing.PhoneNumber != "898123114" {
		t.Errorf("PhoneNubmer: get %s, want %s", listing.PhoneNumber, "898123114")
	}

	if listing.WhatsAppNumber != nil {
		t.Errorf("WhatsAppNumber: get %v, want %v", listing.WhatsAppNumber, nil)
	}

	if listing.OwnerID != 1 {
		t.Errorf("OwnerID: get %d, want %d", listing.OwnerID, 1)
	}

	if listing.Status != "активно" {
		t.Errorf("Status: get %s, want %s", listing.Status, "активно")
	}

	if listing.Category.ID != 1 {
		t.Errorf("CategoryID: get %d, want %d", listing.Category.ID, 1)
	}

	if listing.Category.Name != "video card" {
		t.Errorf("CategoryName: get %s, want %s", listing.Category.Name, "video card")
	}

	if listing.ProductCondition != "б/у" {
		t.Errorf("ProductConditon: get %s, want %s", listing.ProductCondition, "б/у")
	}

	expectedTimeStr := "2026-10-14 14:20:11.824538+00"
	layout := "2006-01-02 15:04:05.999999-07"

	expectedTime, err := time.Parse(layout, expectedTimeStr)
	if err != nil {
		t.Fatalf("parse time: %v", err)
	}

	if !listing.CreatedAt.Equal(expectedTime) {
		t.Errorf("CreatedAt: get %v, want %s", listing.CreatedAt, expectedTime)
	}

}
