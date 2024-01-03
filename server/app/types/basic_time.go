package types

import (
	"encoding/json"
	"time"
)

type BasicTime string

func (basicTime *BasicTime) UnmarshalJSON(date []byte) error {

	var rawDate string
	if err := json.Unmarshal(date, &rawDate); err != nil {
		return err
	}

	defaultFormat, err := time.Parse("2006-01-02", rawDate)
	if err != nil {
		return err
	}

	basicFormat := defaultFormat.Format("2006-01-02")

	*basicTime = BasicTime(basicFormat)

	return nil
}

func (basicTime BasicTime) MarshalJSON() ([]byte, error) {
	return []byte(`"` + basicTime + `"`), nil
}
