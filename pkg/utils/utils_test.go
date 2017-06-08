package utils

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_RetryExBackoffOperation_NoRetry(t *testing.T) {
	startTime := time.Now()
	RetryExBackoffOperation(10, "TestRetry", func() error {
		return nil
	})
	execTime := time.Since(startTime)
	assert.InDelta(t, 0, execTime.Seconds(), 0.1, "Retry timeout was incorrect")
}

func Test_RetryExBackoffOperation_RetryTwice(t *testing.T) {
	var count int
	startTime := time.Now()
	RetryExBackoffOperation(1, "TestRetry", func() error {
		count = count + 1
		if count <= 2 {
			return fmt.Errorf("Failed")
		}
		return nil
	})
	execTime := time.Since(startTime)
	assert.InDelta(t, 1, execTime.Seconds(), 0.1, "Retry timeout was incorrect")
}
func Test_RetryExBackoffOperation_Timeout(t *testing.T) {
	startTime := time.Now()
	RetryExBackoffOperation(1, "TestRetry", func() error {
		return fmt.Errorf("Failed")
	})
	execTime := time.Since(startTime)
	assert.InDelta(t, 1, execTime.Seconds(), 0.1, "Retry timeout was incorrect")
}

func Test_RetryOperation_NoRetry(t *testing.T) {
	startTime := time.Now()
	RetryOperation(3, 1, "TestRetry", func() error {
		return nil
	})
	execTime := time.Since(startTime)
	assert.InDelta(t, 0, execTime.Seconds(), 0.1, "Retry timeout was incorrect")
}

func Test_RetryOperation_RetryTwice(t *testing.T) {
	var count int
	startTime := time.Now()
	RetryOperation(3, 1, "TestRetry", func() error {
		count = count + 1
		if count <= 2 {
			return fmt.Errorf("Failed")
		}
		return nil
	})
	execTime := time.Since(startTime)
	assert.InDelta(t, 2, execTime.Seconds(), 0.1, "Retry timeout was incorrect")
}
func Test_RetryOperation_Timeout(t *testing.T) {
	startTime := time.Now()
	RetryOperation(3, 1, "TestRetry", func() error {
		return fmt.Errorf("Failed")
	})
	execTime := time.Since(startTime)
	assert.InDelta(t, 3, execTime.Seconds(), 0.1, "Retry timeout was incorrect")
}
