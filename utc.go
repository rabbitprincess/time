package time

import "time"

func UTC_now() time.Time {
	return time.Now().UTC()
}

func UTC_custom(_sec int64, _nsec int64) time.Time {
	return time.Unix(_sec, _nsec).UTC()
}

func UTC_truncateYear(_sec int64) time.Time {
	pt_time := UTC_custom(_sec, 0)
	pt_time__truncate := time.Date(pt_time.Year(), 1, 1, 0, 0, 0, 0, time.UTC)
	return pt_time__truncate
}

func UTC_truncateMonth(_sec int64) time.Time {
	pt_time := UTC_custom(_sec, 0)
	pt_time__truncate := time.Date(pt_time.Year(), pt_time.Month(), 1, 0, 0, 0, 0, time.UTC)
	return pt_time__truncate
}

func UTC_truncateWeek(_sec int64) time.Time {
	return UTC_truncate(_sec, 1*60*60*24*7)
}

func UTC_truncateDay(_sec int64) time.Time {
	return UTC_truncate(_sec, 1*60*60*24)
}

func UTC_truncateHour(_sec int64) time.Time {
	return UTC_truncate(_sec, 1*60*60)
}

func UTC_truncateMin(_sec int64) time.Time {
	return UTC_truncate(_sec, 1*60)
}

func UTC_truncateSec(_sec int64) time.Time {
	return UTC_truncate(_sec, 1)
}

func UTC_truncate(_sec int64, _truncate_sec int64) time.Time {
	timeCustom := time.Unix(_sec, 0).UTC()
	timeCustom = timeCustom.Truncate(time.Second * time.Duration(_truncate_sec))
	return timeCustom
}
