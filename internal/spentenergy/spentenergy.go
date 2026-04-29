package spentenergy

import (
	"errors"
	"time"
)

const (
	mInKm                      = 1000
	minInH                     = 60
	stepLengthCoefficient      = 0.45
	walkingCaloriesCoefficient = 0.5
)

// WalkingSpentCalories рассчитывает количество калорий, сожженных при ходьбе
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, errors.New("negative or zero number of steps")
	}
	if weight <= 0 {
		return 0, errors.New("negative or zero weight")
	}
	if height <= 0 {
		return 0, errors.New("negative or zero height")
	}
	if duration <= 0 {
		return 0, errors.New("negative or zero duration")
	}
	meanSpeed := MeanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()
	return ((weight * meanSpeed * durationInMinutes) / minInH) * walkingCaloriesCoefficient, nil
}

// WalkingSpentCalories рассчитывает количество калорий, сожженных при беге
func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, errors.New("negative or zero number of steps")
	}
	if weight <= 0 {
		return 0, errors.New("negative or zero weight")
	}
	if height <= 0 {
		return 0, errors.New("negative or zero height")
	}
	if duration <= 0 {
		return 0, errors.New("negative or zero duration")
	}
	meanSpeed := MeanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()
	return (weight * meanSpeed * durationInMinutes) / minInH, nil
}

// MeanSpeed рассчитывает среднюю скорость (км/ч) на основе шагов и времени
func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 {
		return 0
	}
	distance := Distance(steps, height)
	return distance / duration.Hours()
}

// Distance рассчитывает пройденную дистанцию в километрах
func Distance(steps int, height float64) float64 {
	lenSteps := stepLengthCoefficient * height
	return float64(steps) * lenSteps / mInKm
}
