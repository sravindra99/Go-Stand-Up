package utils

import (
	"fmt"
	"strings"
)

type Session struct {
	UserMap map[string]bool
	done    bool
}

func (s *Session) BuildUserMap(input string) error {
	splitInput := strings.Fields(input)
	s.UserMap = make(map[string]bool, len(splitInput))

	for _, val := range splitInput {
		_, ok := s.UserMap[val]
		if ok {
			return fmt.Errorf("repeating names present")
		}
		s.UserMap[val] = true
	}
	return nil
}

func (s *Session) PickRandomUser() (string, error) {
	// Random picking logic not manually written. Iterating over maps is pseudorandom in itself
	var res string

	if len(s.UserMap) < 1 {
		return "", fmt.Errorf("no one left to pick")
	}
	if len(s.UserMap) == 1 {
		for k, _ := range s.UserMap {
			res = k
		}
		delete(s.UserMap, res)
		return res, nil
	}

	for k, _ := range s.UserMap {

		res = k
		break

	}

	if res != "" {
		delete(s.UserMap, res)
		return res, nil
	}

	return "", fmt.Errorf("oops something went wrong choosing a random person")
}

func (s *Session) IsDone() bool {
	return len(s.UserMap) > 0
}
