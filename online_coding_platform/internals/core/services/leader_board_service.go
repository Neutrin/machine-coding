package services

import (
	"github.com/online_coding_platform/internals/core/domains"
	"github.com/online_coding_platform/internals/core/ports"
)

type LeaderBoardService struct {
	userRepo ports.UserRepo
}

func NewLeaderBoardService(userRep ports.UserRepo) *LeaderBoardService {
	return &LeaderBoardService{
		userRepo: userRep,
	}
}

func (service *LeaderBoardService) GenerateLeaderboard(order int64) []domains.LeaderBoard {
	leaderboard := make([]domains.LeaderBoard, 0)
	users, _ := service.userRepo.GetAll()
	if order == 1 {
		for _, curUser := range domains.SortScoreAscending(users) {
			leaderboard = append(leaderboard, domains.NewLeaderBoard(curUser))
		}
	} else {
		for _, curUser := range domains.SortScoreDescending(users) {
			leaderboard = append(leaderboard, domains.NewLeaderBoard(curUser))
		}
	}
	return leaderboard
}
