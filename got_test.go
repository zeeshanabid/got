package got

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var format = "2006-01-02 15:04:05.999999999"

func TestNew(t *testing.T) {
	gt := New(time.Date(2016, 10, 21, 16, 20, 8, 123456789, time.UTC))
	assert.Equal(t, "2016-10-21 16:20:08.123456789", gt.Format(format))
}

func TestGotBeginningOfMinute(t *testing.T) {
	gt := New(time.Date(2016, 10, 21, 16, 20, 8, 123456789, time.UTC)).BeginningOfMinute()
	assert.Equal(t, "2016-10-21 16:20:00", gt.Format(format))
}

func TestGotEndOfMinute(t *testing.T) {
	gt := New(time.Date(2016, 10, 21, 16, 20, 8, 123456789, time.UTC)).EndOfMinute()
	assert.Equal(t, "2016-10-21 16:20:59.999999999", gt.Format(format))
}

func TestGotNextMinute(t *testing.T) {
	gt := New(time.Date(2016, 10, 21, 16, 20, 8, 123456789, time.UTC)).NextMinute()
	assert.Equal(t, "2016-10-21 16:21:00", gt.Format(format))

	gt = New(time.Date(2016, 10, 21, 16, 00, 00, 00, time.UTC)).NextMinute()
	assert.Equal(t, "2016-10-21 16:01:00", gt.Format(format))
}

func TestGotPreviousMinute(t *testing.T) {
	gt := New(time.Date(2016, 10, 21, 16, 20, 8, 123456789, time.UTC)).PreviousMinute()
	assert.Equal(t, "2016-10-21 16:19:00", gt.Format(format))

	gt = New(time.Date(2016, 10, 21, 16, 00, 00, 00, time.UTC)).PreviousMinute()
	assert.Equal(t, "2016-10-21 15:59:00", gt.Format(format))
}

func TestGotHalfOfMinute(t *testing.T) {
	gt := New(time.Date(2016, 10, 21, 16, 20, 8, 123456789, time.UTC)).HalfOfMinute()
	assert.Equal(t, "2016-10-21 16:20:30", gt.Format(format))
}
