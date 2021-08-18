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
