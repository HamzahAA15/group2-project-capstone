package validation

import (
	"errors"
)

func CheckEmpty(input ...interface{}) error {
	for _, value := range input {
		switch value {
		case "":
			return errors.New("make sure your input not empty")
		case 0:
			return errors.New("make sure your input not zero")
		case nil:
			return errors.New("make sure your input not nil")
		}
	}

	return nil
}
