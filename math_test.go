package main

import (
	"testing"
)

type UserTests struct { Test *testing.T}
func TestRunner(t *testing.T) {

	t.Run("A=create", func(t *testing.T) {
		test:= UserTests{Test: t}
		test.TestCreateRegularUser()
		//test.TestCreateConfirmedUser()
		//test.TestCreateMasterUser()
		//test.TestCreateUserTwice()
	})
	t.Run("A=login", func(t *testing.T) {
		//test:= UserTests{Test: t}
		//test.TestLoginRegularUser()
		//test.TestLoginConfirmedUser()
		//test.TestLoginMasterUser()
	})

}

func (t *UserTests) TestCreateRegularUser() {
	//Function.GetRow(1, 2)
	//registerRegularUser := util.TableTest{
	//	//Method:      "POST",
	//	//Path:        "/iot/users",
	//	//Status:      http.StatusOK,
	//	//Name:        "registerRegularUser",
	//	//Description: "register Regular User has to return 200",
	//	//Body: SerializeUser(RegularUser),
	//}
	//response := util.SpinSingleTableTests(t.Test, registerRegularUser)
	//util.LogIfVerbose(color.BgCyan, "IOT/USERS/TEST", response)
}