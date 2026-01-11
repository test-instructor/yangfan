package runTestCase

type TestCase interface {
	LoadCase() (err error)
	RunCase() (err error)
	Report() (reports interface{}, err error)
}

func RunTestCase(tc TestCase) (reports interface{}, err error) {
	err = tc.LoadCase()
	if err != nil {
		return
	}
	err = tc.RunCase()
	if err != nil {
		return
	}
	report, err := tc.Report()
	return report, nil
}
