package ruler

import (
	"context"
	"encoding/json"
	"sync"
)

type RulerArgs struct {
	Key  string          `json:"key"`
	Args json.RawMessage `json:"args"`
}

type Ruler interface {
	Verify(ctx context.Context, args RulerArgs) (valid bool, reason string)
}

type RulerFunc func(data []byte) (ruler Ruler, err error)

type RuleEngine interface {
}

var rulerBatchIndex = [][]string{
	{"city", "county"},
	{"abc", "tag", "ice"},
}

var rulerIndex = make(map[string]int)

var rulers = make(map[string]RulerFunc)

func init() {
	for i, keys := range rulerBatchIndex {
		for _, key := range keys {
			rulerIndex[key] = i
		}
	}
}

type batchArgs struct {
	Key  string          `json:"key"`
	Args json.RawMessage `json:"rule"`
}

func doVerify() {

}

type result struct {
	Ok     bool   `json:"ok"`
	Error  error  `json:"error"`
	Reason string `json:"reason"`
}

func Verify(ctx context.Context, rule, args map[string]json.RawMessage) (ok bool, reason []string) {
	var batch = make([][]batchArgs, len(rulerBatchIndex))
	var other []batchArgs

	for key, args := range args {
		var args = batchArgs{
			Key:  key,
			Args: args,
		}

		if index, exists := rulerIndex[key]; exists {
			batch[index] = append(batch[index], args)
			continue
		}

		other = append(other, args)
	}

	if len(other) > 0 {
		batch = append(batch, other)
	}

	var results = make([]result, len(batch))
	for _, batch := range batch {
		if len(batch) == 0 {
			continue
		}

		var wg sync.WaitGroup
		var res = make([]result, len(batch))
		for i, args := range batch {
			wg.Add(1)
			go func(i int, args batchArgs) {
				defer wg.Done()
				var fn = rulers[args.Key]
				if fn == nil {
					fn = NewDefaultRuler
				}
				ruler, err := fn(args.Args)
				if err != nil {
					res[i].Error = err
					return
				}
				res[i].Ok, res[i].Reason = ruler.Verify(context.Background(), RulerArgs{
					Key: args.Key,
					//Args: args.Rule,
				})
			}(i, args)
		}

		results = append(results, res...)
	}

	ok = true

	for _, result := range results {
		if !result.Ok {
			ok = false
		}
		if result.Error != nil {
			reason = append(reason, result.Error.Error())
			continue
		}
		reason = append(reason, result.Reason)
	}

	return
}
