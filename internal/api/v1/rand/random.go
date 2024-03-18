package rand

import (
	"math/rand"
	"time"
)

func New() mystruct {
	source := rand.NewSource(time.Now().UnixNano())
	return mystruct{
		myrand: rand.New(source),
	}
}

func (r *mystruct) randInt(lowerBound, upperbound int64) int64 {
	return r.myrand.Int63n(upperbound-lowerBound+1) + lowerBound
}
