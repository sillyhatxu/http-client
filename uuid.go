package client

import (
	"github.com/google/uuid"
	"strings"
)

func GetId() string {
	id := uuid.New()
	return strings.ToUpper(strings.ReplaceAll(id.String(), "-", ""))
<<<<<<< HEAD
}
=======
}
>>>>>>> 35e5e3e849e8d8001329135bbe60ca024d20839a
