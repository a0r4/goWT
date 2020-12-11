# goWT
Some situation of usage JWT token could cause to be vulnerabilities. Most popular scenarios:
- "none" Algorithm usage
- Weak key usage for Token signing
- Support multiple signiture methods (RSA-HMAC) (**Feature**)

More detail about this topic visit [auth0 blog](https://auth0.com/blog/critical-vulnerabilities-in-json-web-token-libraries/)

# usage
`./goWt -jwt $JWT_TOKEN -attackType noneAlg : Change JWT Token's algorithm as "none"` 
<br>
`./goWT -jwt $JWT_TOKEN -attackType dictionary -wordlist $FILE_PATH : Try to find JWT Token's secret with provided wordlist`
<br>
`./goWT -jwt $JWT_TOKEN -attackType showJwt : Decode and show JWT Token's part`
