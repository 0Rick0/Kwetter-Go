package service

import (
	"../types"
)

func (s *Service) GetKwetsOfUser(username string, count int, offset int) *[]types.Kwet {
	var kwets []types.Kwet
	s.db.Limit(count).Offset(offset).
		Where("sender_id = (SELECT id FROM users WHERE username = ?)", username).Find(&kwets)
	return &kwets
}
