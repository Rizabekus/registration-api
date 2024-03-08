package errortypes

import "errors"

var ErrNoUserID = errors.New("No user ID found for the provided cookie")
