package util

import (
	"fmt"
	"time"
)

func NewID() string {
	return fmt.Sprintf("%x", time.Now().UnixNano())
}
