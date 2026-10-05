// Package uuidutil
package uuidutil

import (
	"fmt"
	"uuid"
)

func ParseUUID(value any) (uuid.UUID, error) {
	switch v := value.(type) {
	case uuid.UUID:
		return v, nil

	case string:
		return uuid.Parse(string(v))

	case nil:
		return uuid.Nil(), fmt.Errorf("UUID value is nil")

	default:
		return uuid.Nil(), fmt.Errorf("unsupported UUID type: %T", value)
	}
}
