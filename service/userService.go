package service

import (
	"../types"
)

func (s *Service) GetUsers(count int, offset int) []types.User  {
	var results []types.User
	s.db.Limit(&count).Offset(&offset).Find(&results)
	return results
}

func (s *Service) AddUser(user types.User) bool  {
	if &user == nil {
		return false
	}
	if user.Id != 0{
		return false //ID may not be set, auto generated
	}

	//username must be unique
	var count int
	s.db.Model(&types.User{}).Where("username = ?", user.Username).Count(&count)
	if count > 0 {
		return false
	}

	//store the user
	s.db.Create(&user)
	return true
}

func (s *Service) RemoveUser(user types.User) bool {
	var count int
	s.db.Model(&types.User{}).Where("username = ?", user.Username).Count(&count)
	if count > 0 {
		return false
	}
	s.db.Delete(&user)
	return true
}

func (s *Service) GetUserByUsername(username string) *types.User {
	var user types.User
	s.db.Where("username = ?", username).First(&user)
	return &user
}

func (s *Service) GetUserById(id int) *types.User {
	var user types.User
	s.db.Where("id = ?", id).First(&user)
	return &user
}

func (s *Service) GetFollowingByUsername(username string) *[]string {
	var result []string
	var followers []int
	s.db.Model(&types.UserFollowing{}).Where("follower_id = (SELECT id FROM users WHERE username = ?)", username).
		Pluck("following_id", &followers)
	s.db.Model(&types.User{}).Where(followers).Pluck("username", &result)
	return &result
}

func (s *Service) GetFollowersByUsername(username string) *[]string {
	var result []string
	var followers []int
	s.db.Model(&types.UserFollowing{}).Where("following_id = (SELECT id FROM users WHERE username = ?)", username).
		Pluck("follower_id", &followers)
	s.db.Model(&types.User{}).Where(followers).Pluck("username", &result)
	return &result
}
