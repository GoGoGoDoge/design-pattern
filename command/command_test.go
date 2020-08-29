package command_test

import (
	"command"
)

func ExampleCommand() {
	tv := &command.Device{}
	command1 := command.NewOnCommand(tv)
	command2 := command.NewOffCommand(tv)

	toyController := command.NewController(command1, command2)

	toyController.TurnOn()
	toyController.TurnOff()
	// output:
	// On
	// Off
}
