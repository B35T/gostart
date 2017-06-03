package test

import (
	"testing"
	"os/exec"
	"GO_START/src/controllers/commands"

)

func TestCommamdGetPWD(t *testing.T)  {
	cmd := exec.Command("pwd")

	getPWD,_ := commands.GetPWD()
	input, err := cmd.CombinedOutput()
	if err != nil {
		t.Error("pwd not true")
	}

	if string(input) != getPWD {
		t.Error("pwd not true")
	}

}
