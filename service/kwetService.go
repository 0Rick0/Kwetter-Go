package service

import (
	"../types"
	"fmt"
)

func (s *Service) GetKwetsOfUser(username string, count int, offset int) *[]types.Kwet {
	var kwets []types.Kwet
	s.db.Limit(count).Offset(offset).
		Preload("Sender").Preload("Tags").Preload("Mentions").Preload("LikedBy").
		Where("sender_id = (SELECT id FROM users WHERE username = ?)", username).Find(&kwets)
	return &kwets
}

func (s *Service) GetKwetsOfFollowed(username string, count int, offset int) *[]types.Kwet {
	var kwets []types.Kwet
	var user types.User
	//First get the userid of the user, to minify the query complexity
	s.db.Model(&types.User{}).Where("username = ?", username).Find(&user)

	//Now get the kwets, first set limit/offset, then preload te extras, then select and finally order and find
	s.db.Limit(count).Offset(offset).
		Preload("Sender").Preload("Tags").Preload("Mentions").Preload("LikedBy").
		Where("sender_id in (SELECT following_id FROM user_followings WHERE follower_id = ?)", user.Id).
		Order("created_at desc").Find(&kwets)
	return &kwets
}

func (s *Service) PostKwet(username string, content string, tags []string, mentions []string) *types.Kwet{
	fmt.Printf("Posting '%s' for %s with %d tags and %d mentions\n", content, username, len(tags), len(mentions))
	var user types.User
	s.db.Where("username = ?", username).First(&user)

	var kwet = new(types.Kwet)
	kwet.Text = content
	kwet.Sender = user

	s.db.Create(&kwet)
	for _, tag := range tags {
		var tago = new(types.Tag)
		tago.Text = tag
		tago.KwetID = kwet.Id
		s.db.Create(&tago)
	}

	return kwet
	//todo mentions
}

func (s *Service) GetTrends(count int) *[]string  {
	var trends []string
	s.db.Model(&types.Tag{}).Limit(count).
		Group("text").Order("count(kwet_id) desc").Pluck("text", &trends)
	return &trends
}