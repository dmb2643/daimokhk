package listing

import "fmt"

var (
	ErrListingNotFound          = fmt.Errorf("listing not found")
	ErrListigRemovedByModerator = fmt.Errorf("listing was removed by moderator")
)
