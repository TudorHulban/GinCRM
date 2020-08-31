package httpinterface

import (
	"os"
	"strconv"
	"strings"

	"github.com/TudorHulban/log"
	"github.com/pkg/errors"
)

// CreateConfig Package Helper for creating Gin configuration.
// Holds also validators.
// socket like "0.0.0.0:8080"
func CreateConfig(socket string, ldVersion string) (*Config, error) {
	// input validation
	port, errParse := strconv.Atoi(socket[strings.Index(socket, ":")+1:])
	if errParse != nil {
		return nil, errors.WithMessage(errParse, "provided Gin listening port could not be parsed")
	}

	return &Config{
		Port:          uint16(port),
		BinaryVersion: ldVersion,
		GLogger:       log.New(3, os.Stderr, true),
	}, nil
}
