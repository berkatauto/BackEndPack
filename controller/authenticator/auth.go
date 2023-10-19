package controller

import (
	"fmt"

	"github.com/whatsauth/watoken"
)

var pvtKey = "732yrfgew768a8t7hfasiudf"
var pbcKey = "g532g7rfgdshusdvdbfhseuk"

func Auth() {
	userid := "admin"
	tokenstring, _ := watoken.Encode(userid, pvtKey)
	fmt.Println(tokenstring)
	//decode token to get userid
	useridstring := watoken.DecodeGetId(pbcKey, tokenstring)
	if useridstring == "" {
		fmt.Println("expire token")
	}
}
