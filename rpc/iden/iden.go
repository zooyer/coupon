package iden

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func GenID() (id int64, err error) {
	return rand.Int63()%10000000 + 1000000, nil
}
