package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func ShowMeYouToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.RequestURI)
		token := ""
		if r.Header["Access-Token"] != nil {
			token = r.Header["Access-Token"][0]
		}
		if token != "" && ParseToken(token) {
			next.ServeHTTP(w, r)
		} else {
			res := InvalidTokenResponse{}
			res.Success = false
			res.Error = "Your token is invalid"
			jData, _ := json.Marshal(res)
			w.Header().Set("Content-Type", "application/json")
			w.Write(jData)
		}
	})
}

type InvalidTokenResponse struct {
	Success bool
	Error   string
}
