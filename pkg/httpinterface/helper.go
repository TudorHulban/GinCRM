package httpinterface

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

// isIpv4 takes string as "192.168.1.8" and checks if IPv4
func isIpv4(theIP string) error {
	if len(theIP) == 0 {
		return errors.New("passed string for conversion is nil")
	}

	groups := strings.Split(theIP, ".")
	if len(groups) != 4 {
		return errors.New("passed string for conversion is malformed")
	}

	for i, v := range groups {
		groupNo, errParse := strconv.Atoi(v)

		if errParse != nil {
			return errors.WithMessagef(errParse, "passed string conversion fails for group: %v, parsed value: %v", i, v)
		}

		/*

		   if i == 0 && groupNo == 0 {
		   			if theIP != "0.0.0.0" {
		   				return fmt.Errorf("passed string starts with zero in group: %v, parsed value: %v", i, v)
		   			}
		   		}

		*/

		if groupNo < 0 {
			return fmt.Errorf("passed string is negative for group: %v, parsed value: %v", i, v)
		}

		if groupNo > 256 {
			return fmt.Errorf("passed string is greater than 256 for group: %v, parsed value: %v", i, v)
		}
	}

	return nil
}
