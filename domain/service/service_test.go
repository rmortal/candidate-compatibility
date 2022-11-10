package service

import (
	"candidate-compatibility/domain/model"
	"fmt"
	"testing"
)

func TestAnalyzeCompatibility(t *testing.T) {
	teamMember1 := model.Resource{
		Name: "Eddie",
		Attributes: model.AttributeSet{
			Intelligence:       1,
			Strength:           5,
			Endurance:          3,
			SpicyFoodTolerance: 1,
		},
	}

	teamMember2 := model.Resource{
		Name: "Will",
		Attributes: model.AttributeSet{
			Intelligence:       9,
			Strength:           4,
			Endurance:          1,
			SpicyFoodTolerance: 6,
		},
	}

	teamMember3 := model.Resource{
		Name: "Mike",
		Attributes: model.AttributeSet{
			Intelligence:       3,
			Strength:           2,
			Endurance:          9,
			SpicyFoodTolerance: 5,
		},
	}

	applicant1 := model.Resource{
		Name: "John",
		Attributes: model.AttributeSet{
			Intelligence:       4,
			Strength:           5,
			Endurance:          2,
			SpicyFoodTolerance: 1,
		},
	}

	applicant2 := model.Resource{
		Name: "Jane",
		Attributes: model.AttributeSet{
			Intelligence:       7,
			Strength:           4,
			Endurance:          3,
			SpicyFoodTolerance: 2,
		},
	}

	applicant3 := model.Resource{
		Name: "Joe",
		Attributes: model.AttributeSet{
			Intelligence:       1,
			Strength:           1,
			Endurance:          1,
			SpicyFoodTolerance: 10,
		},
	}

	payload := model.ResourcesDto{
		Team:       []model.Resource{teamMember1, teamMember2, teamMember3},
		Applicants: []model.Resource{applicant1, applicant2, applicant3},
	}

	for _, resource := range payload.Team {
		fmt.Println(resource.Name)
	}

	// TODO: Update tests for floating point numbers
}

// TODO: Add coverage for "helper methods" in analysisservice.go
