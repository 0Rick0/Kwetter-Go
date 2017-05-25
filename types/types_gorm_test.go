package types

import (
	"testing"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func TestRelations(t *testing.T) {
	db, err := gorm.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatal("Could not open memory database!")
	}

	db.CreateTable(&User{}, &Kwet{}, &UserFollowing{}, &Tag{})

	var user1 = User{Username:"test_user1"}
	db.Create(&user1)
	var user2 = User{Username:"test_user2"}
	db.Create(&user2)
	var user3 = User{Username:"test_user3"}
	db.Create(&user3)
	var user4 = User{Username:"test_user4"}
	db.Create(&user4)
	db.Where("username = ?", "test_user1").Find(&user1)
	db.Where("username = ?", "test_user2").Find(&user2)

	var kwet = Kwet{Sender: user2,  Text:"Shalalala"}
	kwet.Mentions = append(kwet.Mentions, user3)
	db.Create(&kwet)

	db.Where("username = ?", "test_user2").Find(&user2)

	var findKwet = kwet

	db.Where("sender_id = ?", user2.Id).Find(&findKwet)
	if findKwet.Text != kwet.Text {
		t.Fatal("Text not equal")
	}

	if len(findKwet.Mentions) != 1 {
		t.Fatalf("%d not proper amount of mentions!", len(findKwet.Mentions))
	}
}
