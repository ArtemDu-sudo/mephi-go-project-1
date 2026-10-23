package trainings

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

// Training хранит данные о тренировке.
type Training struct {
	Steps        int           // количество шагов за тренировку
	TrainingType string        // тип тренировки (Бег или Ходьба)
	Duration     time.Duration // длительность тренировки
	personaldata.Personal
}

// Parse разбирает строку формата "3456,Ходьба,3h00m" и записывает данные в t.
func (t *Training) Parse(datastring string) (err error) {
	parts := strings.Split(datastring, ",")
	if len(parts) != 3 {
		return errors.New("неверный формат данных: ожидается \"шаги,тип,длительность\"")
	}

	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return fmt.Errorf("ошибка разбора количества шагов: %w", err)
	}
	if steps <= 0 {
		return errors.New("количество шагов должно быть больше 0")
	}

	duration, err := time.ParseDuration(parts[2])
	if err != nil {
		return fmt.Errorf("ошибка разбора длительности: %w", err)
	}
	if duration <= 0 {
		return errors.New("длительность тренировки должна быть больше 0")
	}

	t.Steps = steps
	t.TrainingType = parts[1]
	t.Duration = duration

	return nil
}

// ActionInfo формирует строку с информацией о тренировке.
func (t Training) ActionInfo() (string, error) {
	distance := spentenergy.Distance(t.Steps, t.Height)
	speed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)

	var (
		calories float64
		err      error
	)

	switch t.TrainingType {
	case "Бег":
		calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	case "Ходьба":
		calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	default:
		return "", errors.New("неизвестный тип тренировки")
	}
	if err != nil {
		return "", err
	}

	return fmt.Sprintf(
		"Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
		t.TrainingType, t.Duration.Hours(), distance, speed, calories,
	), nil
}
