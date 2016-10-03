// +build !windows
// +build !cgo

package luser

import (
	"os"
	"os/user"
	"strconv"
)

// The default stub uses $USER and $HOME to fill in the user.  A more reliable
// method will be tried before falling back to the stub.
func currentUser() (*User, error) {
	uid := os.Getuid()
	if uid >= 0 {
		if u, err := lookupId(strconv.Itoa(uid)); err == nil {
			return u, nil
		}
	}

	if u, err := user.Current(); err == nil {
		return luser(u), nil
	}

	return nil, ErrCurrentUser
}
