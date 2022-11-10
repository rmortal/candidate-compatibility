package model

type ResultDto struct {
	ScoredApplicants []Score `json:"scoredApplicants"`
}
