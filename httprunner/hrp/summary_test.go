package hrp

import "testing"

func TestTestCaseSummary_UpdateSkippedFromInterfaceRecords(t *testing.T) {
	t.Run("success_and_skip_is_success", func(t *testing.T) {
		cs := NewCaseSummary()
		cs.Success = true
		cs.Records = []*StepResult{
			{StepType: StepTypeRequest, Success: true, Skipped: false},
			{StepType: StepTypeAPI, Success: false, Skipped: true},
		}
		cs.UpdateSkippedFromInterfaceRecords()
		if cs.Skipped {
			t.Fatalf("expected skipped=false")
		}
	})

	t.Run("all_skip_is_skip", func(t *testing.T) {
		cs := NewCaseSummary()
		cs.Success = true
		cs.Records = []*StepResult{
			{StepType: StepTypeRequest, Success: false, Skipped: true},
			{StepType: StepTypeAPI, Success: false, Skipped: true},
		}
		cs.UpdateSkippedFromInterfaceRecords()
		if !cs.Skipped {
			t.Fatalf("expected skipped=true")
		}
	})

	t.Run("any_failure_is_fail", func(t *testing.T) {
		cs := NewCaseSummary()
		cs.Success = false
		cs.Records = []*StepResult{
			{StepType: StepTypeRequest, Success: false, Skipped: true},
			{StepType: StepTypeAPI, Success: false, Skipped: true},
		}
		cs.UpdateSkippedFromInterfaceRecords()
		if cs.Skipped {
			t.Fatalf("expected skipped=false")
		}
	})

	t.Run("no_interface_records_keep_existing", func(t *testing.T) {
		cs := NewCaseSummary()
		cs.Success = true
		cs.Skipped = true
		cs.Records = []*StepResult{
			{StepType: StepTypeThinkTime, Success: true, Skipped: false},
		}
		cs.UpdateSkippedFromInterfaceRecords()
		if !cs.Skipped {
			t.Fatalf("expected skipped=true")
		}
	})
}

func TestSummary_AddCaseSummary_UsesCaseSkipped(t *testing.T) {
	s := NewSummary()
	cs := NewCaseSummary()
	cs.Success = true
	cs.Skipped = true
	cs.Stat.Total = 2
	cs.Stat.Skipped = 0
	s.AddCaseSummary(cs)
	if s.Stat.TestCases.Skipped != 1 {
		t.Fatalf("expected testcases.skipped=1, got %d", s.Stat.TestCases.Skipped)
	}
	if s.Stat.TestCases.Success != 0 {
		t.Fatalf("expected testcases.success=0, got %d", s.Stat.TestCases.Success)
	}
}
