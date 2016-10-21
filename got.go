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
