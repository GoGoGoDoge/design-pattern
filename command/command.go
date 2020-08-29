package command

import "fmt"

type Command interface {
	Execute()
}

type Device struct{}

type OnCommand struct {
	device *Device
}

func NewOnCommand(device *Device) *OnCommand {
	return &OnCommand{
		device: device,
	}
}

func (command *OnCommand) Execute() {
	fmt.Println("On")
}

type OffCommand struct {
	device *Device
}

func NewOffCommand(device *Device) *OffCommand {
	return &OffCommand{
		device: device,
	}
}

func (command *OffCommand) Execute() {
	fmt.Println("Off")
}

// Controller is a demo how we can use commands
type Controller struct {
	command1 Command
	command2 Command
}

func NewController(command1 Command, command2 Command) Controller {
	return Controller{
		command1: command1,
		command2: command2,
	}
}

func (controller Controller) TurnOn() {
	controller.command1.Execute()
}

func (controller Controller) TurnOff() {
	controller.command2.Execute()
}
