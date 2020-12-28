![alt text](https://github.com/a0r4/goJAT/blob/master/goJAT.png)

# goJAT - Go JWT Attack Tool
Some situation of usage JWT token could cause to be vulnerabilities. Most popular scenarios:
- "none" Algorithm usage
- Weak key usage for Token signing

More detail about this topic visit [auth0 blog](https://auth0.com/blog/critical-vulnerabilities-in-json-web-token-libraries/)

# build
```bash
git clone https://github.com/a0r4/goJAT.git
cd goJAT
go buid goJAT.go
```

# usage
`./goJAT -jwt $JWT_TOKEN -attackType noneAlg -claims '{"$claim_key":"$claim_value"}' : Change JWT Token's algorithm as "none" and update or insert claims` 
<br>
`./goJAT -jwt $JWT_TOKEN -attackType dictionary -wordlist $FILE_PATH : Try to find JWT Token's secret with provided wordlist`
<br>
`./goJAT -jwt $JWT_TOKEN -attackType showJwt : Decode and show JWT Token's part`
<br>
`./goJAT -jwt $JWT_TOKEN -attackType bruteForce -alphabet $CHARACTERS -passMaxLength $LENGTH : Try to find JWT Token's secret with brute force attack`
