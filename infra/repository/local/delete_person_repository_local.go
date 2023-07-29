package personRepositoryLocal

import (
	"errors"
	"strings"
)

func DeleteByName(name string) error {
	for index, person := range People {
		if strings.ToLower(person.GetName()) == strings.ToLower(name) {
			People = append(People[:index], People[index+1:]...)
			return nil
		}
	}
	return errors.New("DeleteByName() - People not found!")
}
