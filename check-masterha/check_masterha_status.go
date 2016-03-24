package main

import (
	"strings"

	"github.com/mackerelio/checkers"
)

type statusChecker struct {
	subcommand
}

func (c statusChecker) Execute(args []string) error {
	c.Executer = &c

	checker, err := c.executeAll()
	if err != nil {
		return err
	}

	checker.Name = "MasterHA"
	checker.Exit()
	return nil
}

func (c statusChecker) makeCommandName() string {
	return "masterha_check_status"
}

func (c statusChecker) makeCommandArgs() []string {
	return make([]string, 0, 2)
}

func (c statusChecker) parse(out string) (checkers.Status, string) {
	lines := strings.Split(out, "\n")
	errors := make([]string, 0, 0)

	for _, line := range lines {
		if line != "" && !strings.Contains(line, "running(0:PING_OK)") {
			errors = append(errors, line)
		}
	}
	if len(errors) == 0 {
		return checkers.OK, "running(0:PING_OK)"
	}

	msg := strings.Join(errors, "\n")
	return checkers.CRITICAL, msg
}
