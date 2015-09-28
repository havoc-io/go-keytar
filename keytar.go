package keytar

import (
	// System imports
	"errors"
)

// Error definitions
var (
	ErrUnsupported = errors.New("operation unsupported on this platform")
	ErrUnknown  = errors.New("unknown keychain failure")
	ErrNotFound = errors.New("keychain entry not found")
	ErrInvalidValue = errors.New("an invalid value was provided")
)

// All strings passed to this interface must be encoded in UTF-8.  GetPassword
// MAY return a value which is not UTF-8 encoded if the original keychain entry
// was created by another service which stored the password in a non-UTF-8
// encoding.
type Keychain interface {
	AddPassword(service, account, password string) error
	GetPassword(service, account string) (string, error)
	DeletePassword(service, account string) error
}

// Create a common replace function.  It'd be nice if this common implementation
// could be baked into the Keychain interface, but such is Go...
func ReplacePassword(k Keychain, service, account, newPassword string) error {
	// Delete the password.  We ignore errors, because the password may not
	// exist.  Unfortunately, not every platform returns enough information via
	// its delete call to determine the reason for failure, so we can't check
	// that errors were ErrNotFound, but if there's a more serious problem,
	// AddPassword should pick it up.
	k.DeletePassword(service, account)

	// Add the new password
	return k.AddPassword(service, account, newPassword)
}
