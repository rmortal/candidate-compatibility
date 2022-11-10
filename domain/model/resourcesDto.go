package model

type ResourcesDto struct {
	Team       []Resource `json:"Team"`
	Applicants []Resource `json:"Applicants"`
}
