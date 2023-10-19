package controller

import (
	"net/http"

	"github.com/whatsauth/watoken"
)

var Privatekey = "732yrfgew768a8t7hfasiudf"

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Metode tidak diizinkan", http.StatusMethodNotAllowed)
		return
	}

	username := r.FormValue("user_name")
	password := r.FormValue("user_pass")

	if checkCredentials(username, password) {
		tokenString, err := watoken.Encode(username, Privatekey)
		if err != nil {
			http.Error(w, "Token Generating Fail", http.StatusInternalServerError)
			return
		}
		w.Write([]byte(tokenString))
	} else {
		http.Error(w, "Login gagal", http.StatusUnauthorized)
	}
}

func checkCredentials(username, password string) bool {
	return username == "user" && password == "pass"
}
