package quote

import (
	"github.com/piquette/finance-go/equity"
	"testing"
)

func TestName(t *testing.T) {
	q, err := equity.Get("AAPL")
	if err != nil {
		// Uh-oh.
		panic(err)
	}
	// Print out the quote
	t.Log(q)
}
