package runTestCase

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/test-instructor/yangfan/httprunner/hrp"
)

// isAPIStep checks if the step type represents an API call
func isAPIStep(stepType string) bool {
	return stepType == "request" || stepType == "api" || strings.Contains(stepType, "request")
}

// isStepCollection checks if the step type represents a step collection (nested testcase)
func isStepCollection(stepType string) bool {
	return stepType == "testcase"
}

// runCasesWithProgress executes hrp testcases and returns the aggregated
// Summary JSON. It also updates Redis progress counters after each
// step is completed.
func runCasesWithProgress(runner *hrp.HRPRunner, reportID uint, testcases ...hrp.ITestCase) (data []byte, err error) {
	// Create an aggregated summary; this mirrors hrp.RunJsons behaviour
	// but runs in this package so we can hook progress updates into Redis.
	s := hrp.NewSummary()

	// Set step complete callback to update progress in real-time
	if reportID != 0 {
		runner.SetOnStepComplete(func(stepType string, stepResult *hrp.StepResult) {
			// Update API progress when an API step completes
			if isAPIStep(stepType) {
				incrReportProgress(reportID, "api_executed", 1)
			}
			// Update step progress when a step collection completes
			if isStepCollection(stepType) {
				incrReportProgress(reportID, "step_executed", 1)
			}
		})
	}

	// Load all testcases
	testCases, err := hrp.LoadTestCases(testcases...)
	if err != nil {
		log.Error().Err(err).Msg("runCasesWithProgress: failed to load testcases")
		return nil, err
	}

	var runErr error

	for _, tc := range testCases {
		// Each testcase has its own case runner
		caseRunner, err := hrp.NewCaseRunner(*tc, runner)
		if err != nil {
			log.Error().Err(err).Msg("runCasesWithProgress: init case runner failed")
			return nil, err
		}

		it := caseRunner.GetParametersIterator()
		for it.HasNext() {
			sessionRunner := caseRunner.NewSession()
			caseSummary, err1 := sessionRunner.Start(it.Next())
			if err1 != nil {
				log.Error().Err(err1).Msg("runCasesWithProgress: run testcase failed")
				runErr = err1
			}
			if caseSummary != nil {
				// Align case name with config
				caseSummary.Name = tc.Config.Get().Name
				s.AddCaseSummary(caseSummary)

				// Update case progress when a testcase iteration completes
				if reportID != 0 {
					incrReportProgress(reportID, "case_executed", 1)
				}
			}
			// if failfast is enabled, we don't break the loop here.
			// Instead, we let HRPRunner.Skip handle the skipping of subsequent steps/cases.
		}
		// if failfast is enabled, we don't break the loop here.

		// Finalize summary timing
		if s.Time != nil {
			s.Time.Duration = time.Since(s.Time.StartAt).Seconds()
		}

		// Marshal summary to JSON for downstream reuse (unmarshal into AutoReport)
		sj, marshalErr := json.Marshal(s)
		if marshalErr != nil {
			// Prefer returning marshal error explicitly
			log.Error().Err(marshalErr).Msg("runCasesWithProgress: marshal summary failed")
			if runErr == nil {
				return nil, errors.Wrap(marshalErr, "marshal summary failed")
			}
			// If execution already had errors, keep runErr as primary but still log
		}

		// Only mark TTL refresh once execution is done
		if reportID != 0 {
			finalizeReportProgress(reportID)
		}

		return sj, runErr
	}
	return
}
