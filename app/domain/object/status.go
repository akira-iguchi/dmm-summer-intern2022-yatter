package object

import (
// "fmt"

// "github.com/pkg/errors"
// "golang.org/x/crypto/bcrypt"
)

type (
	StatusID   = int64
	ForeignKey = int64

	// Account account
	Status struct {
		ID StatusID `json:"id"`

		Account *Account `json:"account"`

		AccountID ForeignKey `json:"-" db:"account_id"`

		Content *string `json:"content"`

		CreateAt DateTime `json:"create_at,omitempty" db:"create_at"`
	}
)
