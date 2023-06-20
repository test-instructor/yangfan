package pytest

import (
	"github.com/test-instructor/yangfan/parsing/hrp/internal/myexec"
	"github.com/test-instructor/yangfan/parsing/hrp/internal/sdk"
)

func RunPytest(args []string) error {
	sdk.SendEvent(sdk.EventTracking{
		Category: "RunAPITests",
		Action:   "hrp pytest",
	})

	args = append([]string{"run"}, args...)
	return myexec.ExecPython3Command("httprunner", args...)
}
