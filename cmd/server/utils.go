package server

import (
	"fmt"
	"strings"
)

func encodeState(ssoclientKey, state string) string {
	return fmt.Sprintf("%s:%s", ssoclientKey, state)
}

func decodeState(encodedState string) (ssoclientKey, state string, err error) {
	splits := strings.Split(encodedState, ":")
	if len(splits) != 2 {
		return "", "", fmt.Errorf("error retrieving state")
	}

	return splits[0], splits[1], nil
}
