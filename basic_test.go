package exec_test

import (
	gloo "github.com/gloo-foo/framework"
	. "github.com/yupsh/exec"
)

func ExampleExec_basic() {
	// exec ls
	// Note: This would execute the actual ls command
	gloo.MustRun(
		Exec("ls"),
	)
	// Output:
	//
}
