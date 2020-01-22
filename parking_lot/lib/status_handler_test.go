package lib

import (
	"fmt"
	"testing"
)

func TestGetParkingStatusCommand(t *testing.T) {
	//assert := assert.New(t)

	cmd := GetCommandExecuter()
	cmd.Execute("create_parking_lot", "5")

	res, _ := cmd.Execute("status")

	fmt.Println(res)
	//assert.Nil(, "Command create_parking_lot should be able to execute")
}
