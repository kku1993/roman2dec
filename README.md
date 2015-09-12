roman2dec: A Simple HTTP Service to Convert Roman Numerals and Decimal
======================================================================

# Build Instructions
`go build github.com/kku1993/roman2dec`

# Run the HTTP Server
`go run server/server.go`

# How to interact with the HTTP server.
1. Convert decimal to Roman Numeral
  `curl "localhost:8080/convert?d=123"`
2. Convert Roman Numeral to decimal
 	`curl "localhost:8080/convert?r=MMXV"`
3. Convert multiple things!
 	`curl "localhost:8080/convert?r=MMXV&d=123&r=VIII"`

Response body will contain the result, with each result seperated by a
newline. Each result has the form "input=result".
Standard HTTP status codes will be used to report error.
