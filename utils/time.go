package utils

import (
	"database/sql"
	"database/sql/driver"
	"strings"
	"time"
)

func GetLocaleJKT() (*time.Location, error) {
	loc, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		return nil, err
	}

	return loc, nil
}

type DateYYMMDD struct {
	time.Time
}

// Scan JSONDate.
func (j *DateYYMMDD) Scan(value interface{}) (err error) {
	nullTime := &sql.NullTime{}
	err = nullTime.Scan(value)
	*j = DateYYMMDD{nullTime.Time}
	return
}

// Value DateYYMMDD.
func (j DateYYMMDD) Value() (driver.Value, error) {
	y, m, d := time.Time(j.Time).Date()
	return time.Date(y, m, d, 0, 0, 0, 0, time.Time(j.Time).Location()), nil
}

// GormDataType gorm common data type
func (j DateYYMMDD) GormDataType() string {
	return "timestamp"
}

func (d *DateYYMMDD) UnmarshalJSON(b []byte) error {
	dateString := strings.Trim(string(b), `"`)
	parsedTime, err := time.Parse(FormatYYYYMMDD, dateString)
	if err != nil {
		return err
	}

	d.Time = parsedTime
	return nil
}

func (d DateYYMMDD) MarshalJSON() ([]byte, error) {
	return []byte("\"" + d.Time.Format(FormatYYYYMMDD) + "\""), nil
}

func (d DateYYMMDD) Format(s string) string {
	t := time.Time(d.Time)
	return t.Format(FormatYYYYMMDD)
}
