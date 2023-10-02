package ports

import "github.com/online_coding_platform/internals/core/domains"

type ContestRepositiories interface {
	Save(contest domains.Contest)
	All() []domains.Contest
	Level(level int64) ([]domains.Contest, error)
	ById(id int64) (domains.Contest, error)
}
