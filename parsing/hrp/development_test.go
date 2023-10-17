package hrp

import (
	"testing"
)

var yangfanHooks = tmpl("testcases/yangfan_hooks.yml")

func TestRunHook(t *testing.T) {
	buildHashicorpPyPlugin()
	defer removeHashicorpPyPlugin()
	testCase := TestCasePath(yangfanHooks)
	err := NewRunner(nil).GenHTMLReport().Run(&testCase) // hrp.Run(testCase)
	if err != nil {
		t.Fatal()
	}
}
