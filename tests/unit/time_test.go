package unit

import (
	"testing"
	"time"
)

const duration_time = time.Hour * 2

func TestTime(t *testing.T) {
	t.Log(duration_time.Seconds())
}
