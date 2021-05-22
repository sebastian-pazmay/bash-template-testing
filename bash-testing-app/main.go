package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"bash-testing-app/testSuites"
)

type TestCase struct {
	Name  string  `yaml:"Name"`
	Command   string  `yaml:"Command"`
	Argument   string  `yaml:"Argument"`
	ExpectedResult   string  `yaml:"ExpectedResult"`
}

type TestSuite struct {
	TestCases []TestCase  `yaml:"testCases"`
}

func main() {
	fileTestSuite, err1 := ioutil.ReadFile("testSuites/linuxTestSuite.yaml")
	if err1 != nil {
		panic(err1)
	}
	var TestSuite TestSuite
	err := yaml.Unmarshal(fileTestSuite, &TestSuite)
	if err != nil {
		panic(err)
	}
	for indexTestCase := range TestSuite.TestCases {
		testSuites.ExecuteTestCase(indexTestCase,TestSuite.TestCases[indexTestCase].Name,TestSuite.TestCases[indexTestCase].Command,TestSuite.TestCases[indexTestCase].Argument,TestSuite.TestCases[indexTestCase].ExpectedResult)
	}
}
