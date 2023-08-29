package util

import "time"

func (u Utils) GetCurTime() time.Time {
	return time.Now()
}

func (u Utils) GetCurTs() int64 {
	return u.GetCurTime().UnixNano() / 1e6
}

func (u Utils) GetCurDay() int {
	return u.GetCurTime().Day()
}
