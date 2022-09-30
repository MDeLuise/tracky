package services

import (
	"testing"
	"tracky/models"
)

const EPSILON float64 = 0.00000001

func floatEquals(a, b float64) bool {
	return (a-b) < EPSILON && (b-a) < EPSILON
}

func createObservations(values []float64) models.Observations {
	var toReturn []models.Observation = make([]models.Observation, 0)
	for _, val := range values {
		toReturn = append(toReturn, models.Observation{
			Value: val,
		})
	}
	return toReturn
}

func Test_MeanShouldBeCorrect1(t *testing.T) {
	elements := []float64{1, 2.5, 3.7, 0.007, 0.788, 0.789}
	target := &models.Target{
		Name:         "foo",
		Observations: createObservations(elements),
	}
	mean := CalcMean(target)
	if !floatEquals(1.464, mean) {
		t.Errorf("mean should be 1.464, but is %v", mean)
	}
}

func Test_MeanAtShouldBeCorrect1(t *testing.T) {
	elements := []float64{1, 2.5, 3.7, 0.007, 0.788, 0.789}
	target := &models.Target{
		Name:         "foo",
		Observations: createObservations(elements),
	}
	meanAt2 := CalcMeanAt(target, 2)
	if !floatEquals(0.7885, meanAt2) {
		t.Errorf("mean at 2 should be 0.7885, but is %v", meanAt2)
	}
}

func Test_MeanAtShouldBeCorrect2(t *testing.T) {
	elements := []float64{1, 2.5, 3.7, 0.007, 0.788, 0.789}
	target := &models.Target{
		Name:         "foo",
		Observations: createObservations(elements),
	}
	meanAt2 := CalcMeanAt(target, 10)
	if !floatEquals(1.464, meanAt2) {
		t.Errorf("mean at 10 should be 1.464, but is %v", meanAt2)
	}
}

func Test_LastIncrShouldBeCorrect1(t *testing.T) {
	elements := []float64{1, 2.5, 3.7, 0.007, 0.788, 0.789}
	target := &models.Target{
		Name:         "foo",
		Observations: createObservations(elements),
	}
	lastIncr := CalcLastIncr(target)
	if !floatEquals(0.001, lastIncr) {
		t.Errorf("last incr should be 0.001, but is %v", lastIncr)
	}
}

func Test_LastIncrShouldBeCorrect2(t *testing.T) {
	elements := []float64{1, 2.5, 3.7, 0.007, 2, 0.789}
	target := &models.Target{
		Name:         "foo",
		Observations: createObservations(elements),
	}
	lastIncr := CalcLastIncr(target)
	if !floatEquals(-1.211, lastIncr) {
		t.Errorf("last incr should be -1.211, but is %v", lastIncr)
	}
}

func Test_LastIncrShouldBeCorrect3(t *testing.T) {
	elements := []float64{1}
	target := &models.Target{
		Name:         "foo",
		Observations: createObservations(elements),
	}
	lastIncr := CalcLastIncr(target)
	if !floatEquals(1, lastIncr) {
		t.Errorf("last incr should be 1, but is %v", lastIncr)
	}
}
