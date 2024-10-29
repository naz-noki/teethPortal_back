package minioDB

import (
	"fmt"

	"github.com/google/uuid"
)

func (c *client) generateFileName(fileName string) (string, error) {
	uuid, errNewRandom := uuid.NewRandom()
	if errNewRandom != nil {
		return "", errNewRandom
	}

	return fmt.Sprintf("%s_%s", uuid.String(), fileName), nil
}
