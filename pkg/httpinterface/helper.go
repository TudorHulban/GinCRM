package httpinterface

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/TudorHulban/log"
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

		if i == 0 && groupNo == 0 {
			return fmt.Errorf("passed string starts with zero in group: %v, parsed value: %v", i, v)
		}

		if groupNo < 0 {
			return fmt.Errorf("passed string is negative for group: %v, parsed value: %v", i, v)
		}

		if groupNo > 256 {
			return fmt.Errorf("passed string is greater than 256 for group: %v, parsed value: %v", i, v)
		}
	}

	return nil
}

// CreateConfig Package Helper for creating Gin configuration.
// Holds also validators.
// socket like "0.0.0.0:8080"
func CreateConfig(socket string, ldVersion string, logLevel int) (*Config, error) {
	// input validation
	ipv4 := socket[:strings.Index(socket, ":")+1]

	// check is IPV4
	errParseIP := isIpv4(ipv4)
	if errParseIP != nil {
		return nil, errors.WithMessage(errParseIP, "provided Gin listening port could not be parsed")
	}

	port, errParsePort := strconv.Atoi(socket[strings.Index(socket, ":")+1:])
	if errParsePort != nil {
		return nil, errors.WithMessage(errParsePort, "provided Gin listening port could not be parsed")
	}

	return &Config{
		IPV4Address:   ipv4,
		Port:          uint16(port),
		BinaryVersion: ldVersion,
		GLogger:       log.New(logLevel, os.Stderr, true),
	}, nil
}
