package db

import (
	"testing"

	"github.com/pthomison/errcheck"
)

func TestDB(t *testing.T) {
	errcheck.CheckTest(nil, t)
}
