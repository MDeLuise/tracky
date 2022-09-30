package services

import "tracky/models"

func CalcMean(t *models.Target) float64 {
	var sum float64 = 0
	for _, obs := range t.Observations {
		sum += obs.Value
	}
	return sum / float64(len(t.Observations))
}

func CalcMeanAt(t *models.Target, numberOfLastValues int) float64 {
	var sum float64 = 0
	elements := t.Observations
	if numberOfLastValues < len(elements) {
		elements = elements[len(elements)-numberOfLastValues:]
	}
	for _, obs := range elements {
		sum += obs.Value
	}
	return sum / float64(len(elements))
}

func CalcLastIncr(t *models.Target) float64 {
	if len(t.Observations) == 1 {
		return t.Observations[0].Value
	}
	if len(t.Observations) == 0 {
		return 0
	}
	return t.Observations[len(t.Observations)-1].Value -
		t.Observations[len(t.Observations)-2].Value
}
