// The main package provides a HTTP server that can interact with a client to
// convert decimal and Roman Numerals.
//
// Examples:
// 	1. Convert decimal to Roman Numeral
//		`curl "localhost:8080/convert?d=123"`
//	2. Convert Roman Numeral to decimal
//		`curl "localhost:8080/convert?r=MMXV"`
//	3. Convert multiple things!
//		`curl "localhost:8080/convert?r=MMXV&d=123&r=VIII"`
//
// Response body will contain the result, with each result seperated by a
// newline. Each result has the form "input=result".
// Standard HTTP status codes will be used to report error.
package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	conv "github.com/kku1993/roman2dec"
)

func main() {
	http.HandleFunc("/convert", func(w http.ResponseWriter, r *http.Request) {
		queries := r.URL.Query()
		response := ""
		responseCode := http.StatusBadRequest

		// Defer writing the response.
		defer func() {
			w.WriteHeader(responseCode)
			fmt.Fprintf(w, "%s", response)
		}()

		for q, v := range queries {
			for _, val := range v {
				if q == "d" {
					// Convert decimal to Roman Numeral.
					i, err := strconv.Atoi(val)
					if err != nil {
						response = fmt.Sprintf("%v\n", err)
						return
					}

					d, err := conv.Dec2Roman(i)
					if err != nil {
						response = fmt.Sprintf("%v\n", err)
						return
					}

					response += fmt.Sprintf("%v=%v\n", val, d)
				} else if q == "r" {
					// Convert Roman Numeral to decimal.

					// Convert to Roman Number first.
					r, err := conv.RomanString2Number(val)
					if err != nil {
						response = fmt.Sprintf("%v\n", err)
						return
					}

					response += fmt.Sprintf("%v=%d\n", val, conv.Roman2Dec(r))
				} else {
					response = fmt.Sprintf("%s=%s is not a valid argument\n", q, val)
					return
				}
			}
		}

		responseCode = http.StatusOK
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
