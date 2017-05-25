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

func (s *Service) PostKwet(username string, content string, tags []string, mentions []string) {
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
	//todo mentions
}
