package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// func viewHandler(w http.ResponseWriter, r *http.Request) {
// 	placeholder := []byte("signature list goes here")
// 	_, err := w.Write(placeholder)
// 	check(err)
// }

// func viewHandler(w http.ResponseWriter, r *http.Request) {
// 	text := "Here's my template!\n"
// 	tmpl, err := template.New("text").Parse(text)
// 	check(err)
// 	err = tmpl.Execute(w, nil)
// 	check(err)
// }

func getStrings(fileName string) []string {
	var lines []string
	file, err := os.Open(fileName)
	if os.IsNotExist(err) {
		return nil
	}
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	check(scanner.Err())
	return lines
}

type Guestbook struct {
	SignatureCount int
	Signatures     []string
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	signatures := getStrings("signatures.txt")
	guestbook := Guestbook{
		SignatureCount: len(signatures),
		Signatures:     signatures,
	}
	html, err := template.ParseFiles("view.html")
	check(err)
	err = html.Execute(w, guestbook)
	check(err)
}

func newHandler(w http.ResponseWriter, r *http.Request) {
	html, err := template.ParseFiles("new.html")
	check(err)
	err = html.Execute(w, nil)
	check(err)
}

func createHandler(w http.ResponseWriter, r *http.Request) {
	signature := r.FormValue("signature")
	options := os.O_WRONLY | os.O_APPEND | os.O_CREATE
	file, err := os.OpenFile("signatures.txt", options, os.FileMode(0600))
	check(err)
	_, err = fmt.Fprintln(file, signature)
	check(err)
	err = file.Close()
	check(err)
	http.Redirect(w, r, "/guestbook", http.StatusFound)
}

func main() {
	http.HandleFunc("/guestbook", viewHandler)
	http.HandleFunc("/guestbook/new", newHandler)
	http.HandleFunc("/guestbook/create", createHandler)
	err := http.ListenAndServe(":3000", nil)
	log.Fatal(err)
}

// --------- TEXT/HTML-----------

// func executeTemplate(text string, data interface{}) {
// 	tmpl, err := template.New("test").Parse(text)
// 	check(err)
// 	err = tmpl.Execute(os.Stdout, data)
// 	check(err)
// }

// type Part struct {
// 	Name  string
// 	Count int
// }

// func main() {

// 	text := "Here's my template!\nAction: {{.}}\nTemplate end\n"
// 	tmpl, err := template.New("test").Parse(text)
// 	check(err)
// 	err = tmpl.Execute(os.Stdout, "ABC")
// 	check(err)
// 	err = tmpl.Execute(os.Stdout, 42)
// 	check(err)
// 	err = tmpl.Execute(os.Stdout, true)
// 	check(err)

// 	executeTemplate("Dot is: {{.}}!\n", "Minh Toan")
// 	executeTemplate("Dot is: {{.}}!\n", 123.5)

// 	executeTemplate("start {{if .}}Dot is true!{{end}} finish\n", true)
// 	executeTemplate("start {{if .}}Dot is true!{{end}} finish\n", false)

// 	templateText := "Before loop: {{.}}\n{{range .}}In loop: {{.}}\n{{end}}After loop: {{.}}\n"
// 	executeTemplate(templateText, []string{"do", "re", "mi"})

// 	templateText = "Name: {{.Name}}\nCount: {{.Count}}\n"
// 	executeTemplate(templateText, Part{Name: "Fuses", Count: 5})
// 	executeTemplate(templateText, Part{Name: "Cables", Count: 10})
// }
