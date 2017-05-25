package service

import (
	"../types"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
)

type Service struct {
	Kwets []types.Kwet
	Users map[string]types.User
	db gorm.DB
}

func (s *Service) SetupDatabase() { //todo arguments dialect database
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		log.Panic("Failed to open database!")
	}

	db.AutoMigrate(&types.User{}, &types.Kwet{}, &types.Tag{}, &types.UserFollowing{})
	s.db = *db
}


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