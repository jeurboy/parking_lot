package lib

// "github.com/davecgh/go-spew/spew"
import (
	"errors"
	"fmt"

	"parking_lot/models"

	"github.com/thoas/go-funk"
)

var PS models.ParkingStore

type CommandHandlerInterface interface {
	ExecuteHandler(*models.ParkingStore, ...string) (string, error)
	ValidateHandler(...string) error
}

type CommandInterface interface {
	Execute(string, ...string) (string, error)
	Validate(string, ...string) error
	SetHandler(CommandHandlerInterface) (res bool)
	GetName() string
}

type Command struct {
	Name    string
	Handler *CommandHandlerInterface

	NextCommand []Command
}

func (Command) getNextStep(text string) (cmd Command, err error) {
	return
}

func init() {
	//fmt.Println("Enter command")
}

func GetCommandExecuter() (cmd CommandInterface) {
	Initial := &Command{
		Name: "init",
	}

	CreateParkingLot := &Command{
		Name: "create_parking_lot",
	}

	CreateParkingLot.SetHandler(CreateHandler{})

	Leave := &Command{
		Name: "leave",
	}
	Leave.SetHandler(LeaveHandler{})

	Park := &Command{
		Name: "park",
	}

	Park.SetHandler(ParkHandler{})

	Status := &Command{
		Name: "status",
	}

	Status.SetHandler(StatusHandler{})

	SearchByRegisNoAndColor := &Command{
		Name: "registration_numbers_for_cars_with_colour",
	}
	SearchByRegisNoAndColor.SetHandler(NumberColorHandler{})

	SearchByColor := &Command{
		Name: "slot_numbers_for_cars_with_colour",
	}
	SearchByColor.SetHandler(SlotColorHandler{})

	SearchByRegisNo := &Command{
		Name: "slot_number_for_registration_number",
	}
	SearchByRegisNo.SetHandler(SlotNumberHandler{})

	Initial.NextCommand = []Command{
		*CreateParkingLot,
		*Park,
		*Leave,
		*Status,
		*SearchByColor,
		*SearchByRegisNo,
		*SearchByRegisNoAndColor,
	}

	// Initial.SetHandler(CreaterHandler{})

	cmd = Initial

	return
}

func (c *Command) Validate(command string, args ...string) (err error) {
	var match bool
	funk.ForEach(c.NextCommand, func(c Command) {
		if err != nil {
			return
		}

		if c.GetName() != command {
			return
		}

		match = true
		fmt.Printf("Validation on command %s\n", c.GetName())

		_, err = c.GetHandler()
		if err != nil {
			fmt.Printf("No handler on command %s\n", c.GetName())
			return
		}

		Handler, _ := c.GetHandler()
		if err = Handler.ValidateHandler(args...); err != nil {
			fmt.Println("ValidateHandler ", c.GetName(), err.Error())
			return
		}

	})

	if match == false {
		err = errors.New("No command match")
		fmt.Println(err.Error())
	}

	return
}

func (c *Command) Execute(command string, args ...string) (res string, err error) {
	funk.ForEach(c.NextCommand, func(c Command) {
		if err != nil {
			return
		}

		if c.GetName() != command {
			return
		}

		Handler, err := c.GetHandler()

		//fmt.Printf("Execution on command %s\n", c.GetName())

		if err != nil {
			fmt.Printf("No handler on command %s\n", c.GetName())
			return
		}

		if res, err = Handler.ExecuteHandler(&PS, args...); err != nil {
			//fmt.Println("Error :", err.Error())
			res = err.Error()
			return
		}

	})

	return
}
func (c *Command) GetName() string {
	return c.Name
}
func (c *Command) GetHandler() (CommandHandlerInterface, error) {
	if c.Handler == nil {
		return nil, errors.New("No handler")
	}
	return *c.Handler, nil
}
func (c *Command) SetHandler(handler CommandHandlerInterface) (res bool) {
	res = true
	c.Handler = &handler

	return
}
