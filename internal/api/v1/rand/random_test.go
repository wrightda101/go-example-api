package rand

import (
	"fmt"
	"testing"

	"gotest.tools/assert"
)

// Notice that we do more in depth testing of the randInt function here
func Test_mystruct_randInt(t *testing.T) {
	r := New()
	var result, lower, upper int64
	lower = 10
	upper = 20
	t.Run(fmt.Sprintf("test between limits -- lower: %d, upper: %d", lower, upper), func(t *testing.T) {
		for i := 0; i < 100; i++ {
			result = r.randInt(lower, upper)
			assert.Equal(t, true, result >= lower)
			assert.Equal(t, true, result <= upper)
		}
	})
	lower = 99990
	upper = 99999
	t.Run(fmt.Sprintf("test between limits -- lower: %d, upper: %d", lower, upper), func(t *testing.T) {
		for i := 0; i < 100; i++ {
			result = r.randInt(lower, upper)
			assert.Equal(t, true, result >= lower)
			assert.Equal(t, true, result <= upper)
		}
	})
}
