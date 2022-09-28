package main

import (
	"fmt"
	"math/rand"
	"net/http"
)

var fortunes = []string{
	"Your ability for accomplishment will follow with success.",
	"A stranger, is a friend you have not spoken to yet.",
	"A dream you have will come true.",
	"A pleasant surprise is waiting for you.",
	"The greatest risk is not taking one.",
	"For hate is never conquered by hate. Hate is conquered by love.",
	"Your life will be happy and peaceful.",
}

func main() {
	http.HandleFunc("/fortune", FortuneHandler)
	http.HandleFunc("/", HelloHandler)
	http.ListenAndServe(":8080", nil)
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}
func FortuneHandler(w http.ResponseWriter, r *http.Request) {
	randomIndex := rand.Intn(len(fortunes))
	fortune := fortunes[randomIndex]
	htmlBody := `
		<html>
			<head>
				<title>Fortune</title>
				<style>
					body {
						background-color: #111333;
						color: #fff;
					}
					p {
						background-color: #ffffff;
						color: #000000;
					}
				</style>
			</head>
			<body>
				<h1>Your fortune</h1>
				<b>Guaranteed to be true</b>
				<p>%s</p>
			</body>
	`
	fmt.Fprintf(w, htmlBody, fortune)
}
