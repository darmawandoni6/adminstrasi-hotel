package conversi

import "time"

func Date(value string) (time.Time, error) {
	location, _ := time.LoadLocation("Asia/Jakarta")

	format := "2006-01-02 15:04:05"

	date, err := time.ParseInLocation(format, value, location)

	if err != nil {
		return time.Now(), err
	}

	return date, nil
}

func Subtract(value1, value2 string) (int, error) {
	location, _ := time.LoadLocation("Asia/Jakarta")

	format := "2006-01-02 15:04:05"

	date1, err := time.ParseInLocation(format, value1, location)
	if err != nil {
		return 0, err
	}
	date2, err := time.ParseInLocation(format, value2, location)

	if err != nil {
		return 0, err
	}

	day := date1.Sub(date2).Hours() / 24

	return int(day), nil
}
