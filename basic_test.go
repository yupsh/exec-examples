package exec_test

import (
	yup "github.com/gloo-foo/framework"
	. "github.com/yupsh/exec"
)

func ExampleExec_basic() {
	// exec ls
	// Note: This would execute the actual ls command
	yup.MustRun(
		Exec("ls"),
	)
	// Output:
	//
}

