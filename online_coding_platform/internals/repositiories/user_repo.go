package repositiories

import (
	"encoding/json"

	custom_errors "github.com/online_coding_platform/internals/core/errors"

	"github.com/online_coding_platform/internals/core/domains"
)

type UserRepoMap struct {
	mp map[int64][]byte
}

type userRepoModel struct {
	Id    int64
	Name  string
	Score int64
}

func NewUserRepoMap() *UserRepoMap {
	return &UserRepoMap{
		mp: make(map[int64][]byte),
	}
}

func (repo *UserRepoMap) SaveUser(user domains.User) error {
	var (
		byteUser []byte
		err      error
	)
	model := repoUserModel(user)
	if byteUser, err = json.Marshal(&model); err != nil {
		return custom_errors.UserSaveFailed
	}
	repo.mp[model.Id] = byteUser
	return nil
}

func (repo *UserRepoMap) GetUserById(id int64) (domains.User, error) {
	var (
		userByte      []byte
		err           error
		userFound     bool
		user          domains.User
		userRepoModel userRepoModel
	)

	if userByte, userFound = repo.mp[id]; !userFound {
		return user, custom_errors.UserNotFoundError
	}
	if err = json.Unmarshal(userByte, &userRepoModel); err != nil {
		return user, custom_errors.UserNotFoundError
	}
	return domainsUserModel(userRepoModel), err

}

func (repo *UserRepoMap) GetAll() ([]domains.User, error) {
	var (
		users   = make([]domains.User, 0, len(repo.mp))
		curUser domains.User
		err     error
	)
	for curId := range repo.mp {
		if curUser, err = repo.GetUserById(curId); err != nil {
			return users, custom_errors.UserNotFoundError
		}
		users = append(users, curUser)
	}
	return users, err
}

func repoUserModel(user domains.User) userRepoModel {
	return userRepoModel{
		Id:    user.ID(),
		Name:  user.Name(),
		Score: user.Score(),
	}
}

func domainsUserModel(user userRepoModel) domains.User {
	return domains.User{}.SetID(user.Id).SetName(user.Name).SetScore(user.Score)
}
