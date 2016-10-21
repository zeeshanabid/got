package got

import "time"

type Got struct {
	time.Time
}

func New(t time.Time) *Got {
	return &Got{t}
}

func Now() *Got {
	return &Got{time.Now()}
}

func UTC() *Got {
	return New(time.Now().UTC())
}

func Unix(sec int64, nsec int64) *Got {
	return New(time.Unix(sec, nsec))
}

func Epoch() *Got {
	return Unix(0, 0)
}

func (t *Got) UTC() *Got {
	return New(t.Time.UTC())
}

func (t *Got) BeginningOfMinute() *Got {
	return New(t.Truncate(time.Minute))
}

func (t *Got) EndOfMinute() *Got {
	return New(t.Truncate(time.Minute).Add(time.Minute - time.Nanosecond))
}

func (t *Got) NextMinute() *Got {
	return New(t.Add(time.Minute)).BeginningOfMinute()
}

func (t *Got) PreviousMinute() *Got {
	return New(t.Add(-time.Minute)).BeginningOfMinute()
}

func (t *Got) HalfOfMinute() *Got {
	return New(t.Truncate(time.Minute).Add(30 * time.Second))
}

func (t *Got) BeginningOfHour() *Got {
	return New(t.Truncate(time.Hour))
}

func (t *Got) EndOfHour() *Got {
	return New(t.Truncate(time.Hour).Add(time.Hour - time.Nanosecond))
}

func (t *Got) NextHour() *Got {
	return New(t.Add(time.Hour)).BeginningOfHour()
}

func (t *Got) PreviousHour() *Got {
	return New(t.Add(-time.Hour)).BeginningOfHour()
}

func (t *Got) HalfOfHour() *Got {
	return New(t.Truncate(time.Hour).Add(30 * time.Minute)).BeginningOfMinute()
}
