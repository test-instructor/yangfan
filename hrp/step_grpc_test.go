package hrp

import (
	"github.com/test-instructor/grpc-plugin/demo"
	"testing"
)

var demoGrpc = tmpl("testcases/demo_hook.yaml")

func TestRunHooks(t *testing.T) {
	go demo.StartSvc()
	defer demo.StopSvc()
	buildHashicorpPyPlugin()
	defer removeHashicorpPyPlugin()
	testCase := TestCasePath(demoGrpc)
	err := NewRunner(nil).GenHTMLReport().Run(&testCase) // hrp.Run(testCase)
	if err != nil {
		t.Fatal()
	}
}
