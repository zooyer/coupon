package ruler

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
)

func TestVerify(t *testing.T) {
	const rule = `{"city": [1, 2, 3, 4], "county": [11, 22, 33]}`
	const args = `{"city": 1, "county": 11}`
	var ctx = context.Background()
	ok, reason := Verify(ctx, map[string]json.RawMessage{
		"city"
	})
	fmt.Println("is ok:", ok)
	fmt.Println("reason:", reason)
}
