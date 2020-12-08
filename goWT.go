package main

import (
	//"os"
	"fmt"
	//"flag"
	"github.com/fatih/color"
	"github.com/dgrijalva/jwt-go"
)

var parser jwt.Parser

var r = color.New(color.FgRed)
var y = color.New(color.FgYellow)
var c = color.New(color.FgCyan) 
var m = color.New(color.FgMagenta)

func main() {
	jwtSample := "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWUsImp0aSI6IjBiMDMyYTE0LTg2YTAtNGY0Ny05MDIzLTI3MWQ4ZjQ3NjdkNiIsImlhdCI6MTYwNzQwMjc3OCwiZXhwIjoxNjA3NDA2Mzc4fQ.NT0nnKnUVG9YyfPemYXJXKDX7hM_ylfct8PHU9_DbAg"

	token, headers, claims, valid := parseToken(jwtSample)

	showToken(headers, claims, valid)
	fmt.Println(token.Raw)

	signToken(claims)

}

func parseToken(jwtToken string) (*jwt.Token, map[string] interface{}, jwt.MapClaims, bool) {
	token, _, err := parser.ParseUnverified(jwtToken, jwt.MapClaims{})

	if err != nil {
		r.Println("Token was not parsed:", color.New(color.FgRed))
		return nil, nil, nil, false
	}

	headers := token.Header
	claims := token.Claims
	valid := token.Valid

	return token, headers, claims.(jwt.MapClaims), valid
}

func signToken(claims jwt.MapClaims) {
	method := jwt.NewWithClaims(jwt.SigningMethodNone, claims)
	newToken, err := method.SignedString(jwt.UnsafeAllowNoneSignatureType)

	if (err != nil){
		r.Println("Token was not signed:" + err.Error())
		return
	}

	fmt.Println(newToken)

	_, headers, claims, valid := parseToken(newToken)

	showToken(headers, claims, valid)

}

func showToken(headers map[string]interface{}, claims jwt.MapClaims, valid bool) {
	c.Println("***** Headers *****")
	for k,v := range headers {
		m.Println(k + ":", v)
	}

	c.Println("***** Claims *****")
	for k,v := range claims {
		m.Println(k + ":", v)
	}

	c.Println("***** Validation *****")
	m.Println(valid)
}