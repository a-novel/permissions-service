package models

import "github.com/samber/lo"

type Scope string

type Scopes []Scope

const (
	CanVotePost              Scope = "forum:post:vote"
	CanPostImproveRequest    Scope = "forum:improve-request:post"
	CanPostImproveSuggestion Scope = "forum:improve-suggestion:post"

	CanUseOpenAIPlayground Scope = "openai:playground"
)

func (s *Scopes) Has(scope Scope) bool {
	_, ok := lo.Find(*s, func(item Scope) bool {
		return item == scope
	})

	return ok
}
