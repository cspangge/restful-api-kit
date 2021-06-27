package helpers

import "github.com/gofrs/uuid"

func GetUuid() (string, error) {
	res, err := uuid.NewV4()
	if err != nil {
		return "", err
	}
	return res.String(), nil
}
