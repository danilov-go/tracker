package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/danilov-go/tracker/internal/personaldata"
	"github.com/danilov-go/tracker/internal/spentenergy"
)

type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

// Parse анализирует входную строку формата "шаги,длительность"
func (ds *DaySteps) Parse(datastring string) (err error) {
	dataParse := strings.Split(datastring, ",")
	if len(dataParse) != 2 {
		return errors.New("the slice length is not equal to 2")
	}
	steps, err := strconv.Atoi(dataParse[0])
	if err != nil {
		return err
	}
	if steps <= 0 {
		return errors.New("negative or zero number of steps")
	}
	ds.Steps = steps
	duration, err := time.ParseDuration(dataParse[1])
	if err != nil {
		return err
	}
	if duration <= 0 {
		return errors.New("negative or zero time")
	}
	ds.Duration = duration
	return nil
}

// ActionInfo рассчитывает показатели активности: дистанцию и сожженные калории
func (ds DaySteps) ActionInfo() (string, error) {
	distance := spentenergy.Distance(ds.Steps, ds.Height)
	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", err
	}
	strRet := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", ds.Steps, distance, calories)
	return strRet, nil
}
