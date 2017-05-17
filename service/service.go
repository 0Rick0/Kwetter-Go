package service

import (
	"../types"
)

type Service struct {
	Kwets []types.Kwet
	Users map[string]types.User
}


func (s *Service) GetUsers(count int, offset int) []types.User  {
	if len(s.Users) <= count + offset {
		return nil
	}
	var values = make([]types.User, len(s.Users))
	for _, u := range s.Users {
		values = append(values, u)
	}
	return values[offset:offset+count]
}

func (s *Service) AddUser(user types.User) bool  {
	if &user == nil {
		return false
	}
	if _, ok := s.Users[user.Username]; ok{
		return false
	}
	s.Users[user.Username] = user
	return true
}

func (s *Service) RemoveUser(user types.User) bool {
	if _, ok := s.Users[user.Username]; !ok{
		return false
	}
	delete(s.Users, user.Username)
	return true
}

func (s *Service) GetUserByUsername(username string) *types.User {
	for _, user := range s.Users{
		if user.Username == username{
			return &user
		}
	}
	return nil
}

func (s *Service) GetUserById(id int) *types.User {
	for _, user := range s.Users{
		if user.Id == id{
			return &user
		}
	}
	return nil
}