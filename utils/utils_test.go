package utils

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRunBashCmd(t *testing.T) {
	Convey("RunBashCmd", t, func() {
		Convey("should successfully run invoked with a valid bash command string", func() {
			validBashCmd := "echo this should pass"
			expected := "this should pass"
			actual, err := RunBashCmd(validBashCmd)
			So(err, ShouldBeNil)
			So(actual, ShouldResemble, expected)
		})
		Convey("should fail when invoked with an invalid bash command string", func() {
			invalidBashCmd := "whatwasithinking"
			expectedOut := ""
			expectedErr := fmt.Errorf("ERROR: bash: whatwasithinking: command not found\n: exit status 127")
			actualOut, actualErr := RunBashCmd(invalidBashCmd)
			So(actualErr, ShouldResemble, expectedErr)
			So(actualOut, ShouldResemble, expectedOut)
		})
	})
}
