package utils

import (
	"fmt"
	"time"
)

// RetryExBackoffOperation re-runs provided callback function with multiple attempts until either
// the Operation is successful or the timeout specified with timeoutSeconds has reached.
// The sleep between two successive retry increases by 2 times
func RetryExBackoffOperation(timeoutSeconds int, taskName string, callback func() error) (err error) {
	var (
		usedSeconds  = 0
		sleepSeconds = 1
	)
	for usedSeconds < timeoutSeconds {
		// calling callback function
		err = callback()
		if err == nil {
			// if err is nil, that means function execution was successful,
			// we are done retrying
			return nil
		}

		// sleep duration in Seconds
		sleepSeconds = sleepSeconds * 2
		if usedSeconds+sleepSeconds > timeoutSeconds {
			// with next sleep since we are exceeding provided
			// timeoutSeconds lets adjust the sleepSeconds to
			// seconds we have left between what we have used already and provided timeoutSeconds
			sleepSeconds = timeoutSeconds - usedSeconds
		}

		d := time.Duration(sleepSeconds * 1000 * 1000 * 1000)
		time.Sleep(d)
		usedSeconds = usedSeconds + sleepSeconds

	}
	return fmt.Errorf("after %d seconds, last error: %s", timeoutSeconds, err)
}

func RetryOperation(timeoutSeconds int, sleepSeconds int, taskName string, callback func() error) (err error) {
	var (
		usedSeconds = 0
	)
	for usedSeconds < timeoutSeconds {
		// calling callback function
		err = callback()
		if err == nil {
			// if err is nil, that means function execution was successful,
			// we are done retrying
			return nil
		}

		if usedSeconds+sleepSeconds > timeoutSeconds {
			// with next sleep since we are exceeding provided timeoutSeconds
			// lets adjust the sleepSeconds to
			// seconds we have left between what we have used already and provided timeoutSeconds
			sleepSeconds = timeoutSeconds - usedSeconds
		}

		d := time.Duration(sleepSeconds * 1000 * 1000 * 1000)
		time.Sleep(d)
		usedSeconds = usedSeconds + sleepSeconds
	}
	return fmt.Errorf("after %d seconds, last error: %s", timeoutSeconds, err)
}
