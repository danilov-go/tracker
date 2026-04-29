package trainings

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/danilov-go/tracker/internal/personaldata"
	"github.com/danilov-go/tracker/internal/spentenergy"
)

type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

// Parse анализирует входную строку в формате "steps,type,duration"
func (t *Training) Parse(datastring string) (err error) {
	dataParse := strings.Split(datastring, ",")
	if len(dataParse) != 3 {
		return errors.New("the slice length is not equal to 3")
	}
	steps, err := strconv.Atoi(dataParse[0])
	if err != nil {
		return err
	}
	if steps <= 0 {
		return errors.New("negative or zero number of steps")
	}
	t.Steps = steps
	t.TrainingType = dataParse[1]
	duration, err := time.ParseDuration(dataParse[2])
	if err != nil {
		return err
	}
	if duration <= 0 {
		return errors.New("negative or zero time")
	}
	t.Duration = duration
	return nil
}

// ActionInfo рассчитывает основные показатели тренировки
func (t Training) ActionInfo() (string, error) {
	distance := spentenergy.Distance(t.Steps, t.Height)
	meanSpeed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)
	var calories float64
	var err error
	switch t.TrainingType {
	case "Бег":
		calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	case "Ходьба":
		calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	default:
		return "", errors.New("неизвестный тип тренировки.")
	}
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", t.TrainingType, t.Duration.Hours(), distance, meanSpeed, calories), nil
}
