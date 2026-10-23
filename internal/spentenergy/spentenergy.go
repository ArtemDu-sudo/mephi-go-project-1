package spentenergy

import (
	"errors"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

var (
	errInvalidSteps    = errors.New("количество шагов должно быть больше 0")
	errInvalidWeight   = errors.New("вес должен быть больше 0")
	errInvalidHeight   = errors.New("рост должен быть больше 0")
	errInvalidDuration = errors.New("продолжительность должна быть больше 0")
)

// validate проверяет входные параметры расчёта калорий.
func validate(steps int, weight, height float64, duration time.Duration) error {
	switch {
	case steps <= 0:
		return errInvalidSteps
	case weight <= 0:
		return errInvalidWeight
	case height <= 0:
		return errInvalidHeight
	case duration <= 0:
		return errInvalidDuration
	}
	return nil
}

// WalkingSpentCalories возвращает количество калорий, потраченных при ходьбе.
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if err := validate(steps, weight, height, duration); err != nil {
		return 0, err
	}

	speed := MeanSpeed(steps, height, duration)
	calories := (weight * speed * duration.Minutes()) / minInH

	return calories * walkingCaloriesCoefficient, nil
}

// RunningSpentCalories возвращает количество калорий, потраченных при беге.
func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if err := validate(steps, weight, height, duration); err != nil {
		return 0, err
	}

	speed := MeanSpeed(steps, height, duration)

	return (weight * speed * duration.Minutes()) / minInH, nil
}

// MeanSpeed возвращает среднюю скорость в км/ч.
func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if steps <= 0 || duration <= 0 {
		return 0
	}

	return Distance(steps, height) / duration.Hours()
}

// Distance возвращает пройденную дистанцию в километрах.
func Distance(steps int, height float64) float64 {
	stepLength := height * stepLengthCoefficient

	return float64(steps) * stepLength / mInKm
}
