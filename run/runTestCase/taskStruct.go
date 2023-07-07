package runTestCase

import "github.com/test-instructor/yangfan/server/model/interfacecase"

type TaskCaseHrp struct {
	Name      string                  `json:"name"`
	Config    interfacecase.ApiConfig `json:"Config"`
	TestSteps []TestSteps             `json:"TestSteps"`
}

type TestSteps struct {
	Name     string    `json:"name"`
	TestCase TestCases `json:"TestCase"`
}

type TestCases struct {
	Name      string                  `json:"name"`
	Config    interfacecase.ApiConfig `json:"Config"`
	TestSteps []interfacecase.ApiStep `json:"TestSteps"`
}
