package domains

import "sort"

type SortByScore struct {
	users   []User
	compare func(userOne, userTwo User) bool
}

func (score *SortByScore) Len() int {
	return len(score.users)
}

func (score *SortByScore) Swap(i, j int) {
	score.users[i], score.users[j] = score.users[j], score.users[i]
}

func (score *SortByScore) Less(i, j int) bool {
	return score.compare(score.users[i], score.users[j])
}

func SortScoreAscending(user []User) []User {
	sorting := &SortByScore{
		users: user,
		compare: func(userOne, userTwo User) bool {
			return (userOne.Score() < userTwo.Score())
		},
	}
	sort.Sort(sorting)
	return sorting.users
}

func SortScoreDescending(user []User) []User {
	sorting := &SortByScore{
		users: user,
		compare: func(userOne, userTwo User) bool {
			return (userOne.Score() > userTwo.Score())
		},
	}
	sort.Sort(sorting)
	return sorting.users
}
