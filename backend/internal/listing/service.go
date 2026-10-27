package listing

import (
	"context"
	"errors"
	"fmt"
)

type service struct {
	repo postgreSQLRepo
}

func NewService(repo postgreSQLRepo) *service {
	return &service{
		repo: repo,
	}
}

func (s *service) getListings(ctx context.Context, pag pagination) ([]Listing, error) {
	Listings, err := s.repo.getListings(ctx, pag)
	if err != nil {
		return []Listing{}, err
	}

	return Listings, err
}

func (s *service) getListingByID(ctx context.Context, ID int) (Listing, error) {
	listing, err := s.repo.getByID(ctx, ID)
	if err != nil {
		if errors.Is(err, ErrListingNotFound) {
			return Listing{}, ErrListingNotFound
		}

		return Listing{}, fmt.Errorf("service get listing: %w", err)
	}

	if listing.Status == statusPostponed {
		return Listing{}, ErrListingNotFound
	}

	if listing.Status == statusRemoved {
		return Listing{}, ErrListigRemovedByModerator
	}

	return listing, nil
}
