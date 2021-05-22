package testSuites

import (
	"fmt"
	"os/exec"
	"os/user"
	"strings"
)

func ExecuteTestCase(Index int, Name string, Command string, Argument string, ExpectedResult string) {
	var resultBool string
	currentUser, err := user.Current()
	fmt.Println(currentUser)
	cmdOut, err := exec.Command(Command,Argument).Output()
	if err != nil {
		fmt.Printf("%s", err)
	}
	cmdOutStr := string(cmdOut[:])
		if cmdOutStrContainsOSVer := strings.Contains(cmdOutStr, ExpectedResult) ; cmdOutStrContainsOSVer {
		resultBool = "PASS"
	} else {
		resultBool = "FAIL"
	}
	fmt.Printf("\nTest Case #%+v\n",Index+1)
	fmt.Printf("%+v: %+v\n",Name,resultBool)
}
