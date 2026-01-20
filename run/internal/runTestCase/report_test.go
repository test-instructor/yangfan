package runTestCase

import (
	"testing"

	"github.com/test-instructor/yangfan/httprunner/hrp"
)

func TestBuildReportFromSummary_UsesCaseSkipped(t *testing.T) {
	s := hrp.NewSummary()
	cs := hrp.NewCaseSummary()
	cs.Name = "case1"
	cs.Success = true
	cs.Skipped = true
	cs.Stat.Total = 2
	cs.Stat.Skipped = 0
	s.AddCaseSummary(cs)

	report := buildReportFromSummary(s)
	if report == nil || report.Stat == nil || report.Stat.Testcases == nil {
		t.Fatalf("expected report stat")
	}
	if report.Stat.Testcases.Skip != 1 {
		t.Fatalf("expected report.stat.testcases.skip=1, got %d", report.Stat.Testcases.Skip)
	}
	if len(report.Details) != 1 {
		t.Fatalf("expected 1 detail, got %d", len(report.Details))
	}
	if !report.Details[0].Skip {
		t.Fatalf("expected report.details[0].skip=true")
	}
}
