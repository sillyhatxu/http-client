package client

import (
	"github.com/google/uuid"
	"strings"
)

func GetId() string {
	id := uuid.New()
	return strings.ToUpper(strings.ReplaceAll(id.String(), "-", ""))
}
