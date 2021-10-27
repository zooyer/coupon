package ruler

import (
	"context"
	"encoding/json"
	"fmt"
)

type defaultRuler struct {
	value interface{}
}

func NewDefaultRuler(data []byte) (_ Ruler, err error) {
	var ruler defaultRuler
	if err = json.Unmarshal(data, &ruler.value); err != nil {
		return
	}
	return ruler, nil
}

func (d defaultRuler) inSlice(slice []interface{}, v interface{}) bool {
	for _, val := range slice {
		if val == v {
			return true
		}
	}
	return false
}

func (d defaultRuler) Verify(ctx context.Context, args RulerArgs) (valid bool, reason string) {
	var v interface{}
	if err := json.Unmarshal(args.Args, &v); err != nil {
		return false, err.Error()
	}

	switch value := d.value.(type) {
	case []interface{}:
		if len(value) > 0 {
			switch value[0].(type) {
			case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64, string, bool:
				if d.inSlice(value, d.value) {
					return true, "ok"
				}
				return false, fmt.Sprintf("%v not in [%v]", d.value, value)
			default:
			}
		}
	}

	return false, "unknown"
}
