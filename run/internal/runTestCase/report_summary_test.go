package runTestCase

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/test-instructor/yangfan/httprunner/hrp"
)

func TestBuildReportFromSummary_UseSummaryDuration(t *testing.T) {
	startAt := time.Now().Add(-10 * time.Second)
	s := &hrp.Summary{
		Success: true,
		Time: &hrp.TestCaseTime{
			StartAt:  startAt,
			Duration: 12.34,
		},
	}

	report := buildReportFromSummary(s)
	if assert.NotNil(t, report) && assert.NotNil(t, report.Time) {
		assert.True(t, report.Time.StartAt.Equal(startAt))
		assert.InDelta(t, 12.34, report.Time.Duration, 0.000001)
	}
}

func TestBuildReportFromSummary_FallbackDurationWhenMissing(t *testing.T) {
	startAt := time.Now().Add(-2 * time.Second)
	s := &hrp.Summary{
		Success: true,
		Time: &hrp.TestCaseTime{
			StartAt:  startAt,
			Duration: 0,
		},
	}

	report := buildReportFromSummary(s)
	if assert.NotNil(t, report) && assert.NotNil(t, report.Time) {
		assert.True(t, report.Time.StartAt.Equal(startAt))
		assert.Greater(t, report.Time.Duration, 1.0)
		assert.Less(t, report.Time.Duration, 10.0)
	}
}
