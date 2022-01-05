package wallet

import (
	"testing"
)


func TestService_FindAccoundById_Method_NotFound(t *testing.T) {
	svc := Service{}
	svc.RegisterAccount("1234567890")

	account, err := svc.FindAccountByID(4567)
	if err == nil {
		t.Errorf("got > %v want > nil", account)
	}
}

// func TestService_FindAccountByID_notFound(t *testing.T) {

// 	accounts := []*types.Account{
// 		{ID: 1},
// 		{ID: 2},
// 		{ID: 3},
// 	}
// 	result := FindAccountByID(accounts, 5)

// 	if len(result) != 0 {
// 		t.Error("account not found")
// 	}
// }

//Удачи тебе с этой хренью 
//#AKMALKULIEV помоги