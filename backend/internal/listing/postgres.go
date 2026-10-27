package listing

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type postgreSQLRepo struct {
	pool *pgxpool.Pool
}

func NewPostgreSQLRepo(pool *pgxpool.Pool) *postgreSQLRepo {
	return &postgreSQLRepo{
		pool: pool,
	}
}

func (r *postgreSQLRepo) getListings(ctx context.Context, pag pagination) ([]Listing, error) {
	q := `
	select
    l.id,
    l.title,
    l.description,
    l.price,
    l.phone_number,
    l.whatsapp_number,
    l.owner_id,
    l.status,
    l.created_at,
    l.product_condition,
    l.category_id,
    c.name AS category_name
	FROM listing l JOIN category c ON l.category_id = c.id
	where status = 'активно'
	ORDER BY l.last_raised_at DESC, l.id DESC
	LIMIT $1 OFFSET $2
	`

	rows, err := r.pool.Query(ctx, q, pag.limit, pag.offset)
	if err != nil {
		return []Listing{}, fmt.Errorf("get listings: %w", err)
	}
	defer rows.Close()

	listings := make([]Listing, 0)
	for rows.Next() {
		listing := Listing{}

		if err := rows.Scan(
			&listing.ID,
			&listing.Title,
			&listing.Description,
			&listing.Price,
			&listing.PhoneNumber,
			&listing.WhatsAppNumber,
			&listing.OwnerID,
			&listing.Status,
			&listing.CreatedAt,
			&listing.ProductCondition,
			&listing.Category.ID,
			&listing.Category.Name,
		); err != nil {
			return []Listing{}, fmt.Errorf("scan listings: %w", err)
		}
		listings = append(listings, listing)
	}

	if err := rows.Err(); err != nil {
		return []Listing{}, fmt.Errorf("scan error: %w", err)
	}

	return listings, err
}

func (r *postgreSQLRepo) getByID(ctx context.Context, id int) (Listing, error) {
	q := `
	select
		l.id ,
		title,
		description,
		price,
		phone_number,
 	whatsapp_number,
  	owner_id ,
  	status,
  	created_at,
  	product_condition,
  	category_id,
  	c.name as category_name
	from
		listing l  join category c
		on l.category_id = c.id
	where l.id = $1
	`

	var listing Listing

	row := r.pool.QueryRow(ctx, q, id)
	if err := row.Scan(
		&listing.ID,
		&listing.Title,
		&listing.Description,
		&listing.Price,
		&listing.PhoneNumber,
		&listing.WhatsAppNumber,
		&listing.OwnerID,
		&listing.Status,
		&listing.CreatedAt,
		&listing.ProductCondition,
		&listing.Category.ID,
		&listing.Category.Name,
	); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return Listing{}, ErrListingNotFound
		}

		return Listing{}, fmt.Errorf("get listing: %w", err)
	}

	return listing, nil
}
