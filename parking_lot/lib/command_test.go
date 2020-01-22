package lib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitCommandExecuter(t *testing.T) {
	assert := assert.New(t)

	cmd := GetCommandExecuter()
	assert.Equal(cmd.GetName(), "init", "Account owner is new knocking")
}

func TestCreateParkingValidationCommand(t *testing.T) {
	assert := assert.New(t)

	cmd := GetCommandExecuter()
	assert.NotNil(cmd.Validate("park_dog"), "Command park_dog shouldn't be able to execute")

	assert.NotNil(cmd.Validate("create_parking_lot"), "Command create_parking_lot needs one parameter at least")

	assert.Nil(cmd.Validate("create_parking_lot", "0"), "Command create_parking_lot should be able to execute")
}

func TestCreateParkingExecutionCommand(t *testing.T) {
	assert := assert.New(t)

	cmd := GetCommandExecuter()

	_, err := cmd.Execute("create_parking_lot", "5")
	assert.Nil(err, "Command create_parking_lot should be able to execute")

}
