package service

import (
	"candidate-compatibility/domain/model"
)

type AnalysisService struct{}

func (as AnalysisService) AnalyzeCompatibility(dto model.ResourcesDto) model.ResultDto {

	avg := getTeamAverages(dto)
	var scores []model.Score

	for _, applicant := range dto.Applicants {
		weightedIntelligence := getWeightedPercentage(applicant.Attributes.Intelligence, avg.Intelligence)
		weightedStrength := getWeightedPercentage(applicant.Attributes.Strength, avg.Strength)
		weightedEndurance := getWeightedPercentage(applicant.Attributes.Endurance, avg.Endurance)
		weighjtedSpicyFoodTolerance := getWeightedPercentage(applicant.Attributes.SpicyFoodTolerance, avg.SpicyFoodTolerance)

		scores = append(scores, model.Score{
			Name:  applicant.Name,
			Score: weightedIntelligence + weightedStrength + weightedEndurance + weighjtedSpicyFoodTolerance,
		})
	}

	return model.ResultDto{
		ScoredApplicants: scores,
	}
}

func getTeamAverages(dto model.ResourcesDto) model.AttributeSet {
	/* Get team member averages - optimize as time allows */
	teamIntelligence := 0.00
	teamStrength := 0.00
	teamEndurance := 0.00
	teamSpicyFoodTolerance := 0.00

	for _, resource := range dto.Team {
		teamIntelligence += resource.Attributes.Intelligence
		teamStrength += resource.Attributes.Strength
		teamEndurance += resource.Attributes.Endurance
		teamSpicyFoodTolerance += resource.Attributes.SpicyFoodTolerance
	}
	teamSize := float64(len(dto.Team))

	return model.AttributeSet{
		Intelligence:       teamIntelligence / teamSize,
		Strength:           teamStrength / teamSize,
		Endurance:          teamEndurance / teamSize,
		SpicyFoodTolerance: teamSpicyFoodTolerance / teamSize,
	}
}

func getWeightedPercentage(value float64, average float64) float64 {
	percentage := value / average
	if percentage > 1 {
		return 0.25
	}
	return percentage / 4
}
