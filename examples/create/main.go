package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"syscall"

	openproj "github.com/manuelbcd/go-openproject"
	"golang.org/x/term"
)

func main() {
	r := bufio.NewReader(os.Stdin)

	fmt.Print("OpenProject URL: ")
	openProjURL, _ := r.ReadString('\n')

	fmt.Print("OpenProject Username (Suggested 'apikey'): ")
	username, _ := r.ReadString('\n')
	if username == "" {
		username = "apikey"
	}

	fmt.Print("OpenProject Password: ")
	bytePassword, _ := term.ReadPassword(int(syscall.Stdin))
	password := string(bytePassword)

	tp := openproj.BasicAuthTransport{
		Username: strings.TrimSpace(username),
		Password: strings.TrimSpace(password),
	}

	client, err := openproj.NewClient(tp.Client(), strings.TrimSpace(openProjURL))
	if err != nil {
		fmt.Printf("\nerror: %v\n", err)
		return
	}

	i := openproj.WorkPackage{
		Subject: "This is my test work package",
		Description: &openproj.WPDescription{
			Format: "textile",
			Raw:    "This is just a demo workpackage description",
		},
	}

	wpResponse, resp, err := client.WorkPackage.Create(&i, "demo-project")
	if err != nil {
		body, err := io.ReadAll(resp.Body)
		fmt.Print(string(body))
		panic(err)
	}

	// Output with particular fields from response
	fmt.Printf("\n\nSubject: %s \nDescription: %s\n\n", wpResponse.Subject, wpResponse.Description.Raw)

	// Raw output of the whole object (only for debug)
	// fmt.Printf(prettyPrint(wpResponse))
}

//func prettyPrint(i interface{}) string {
//	s, _ := json.MarshalIndent(i, "", "\t")
//	return string(s)
//}
