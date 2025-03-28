package helpers

import (
	"fmt"
	"time"
)

func GenerateInvoiceNumber() string {
	timestamp := time.Now().UnixNano() / int64(time.Millisecond) // Timestamp dalam millisecond
	return fmt.Sprintf("INV-%d", timestamp)
}
