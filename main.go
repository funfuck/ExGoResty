package main

import (
	"github.com/go-resty/resty"
	"fmt"
	"gopkg.in/mgo.v2/bson"
)

type Person struct {
	ID        	bson.ObjectId `bson:"_id,omitempty"`
	FirstName      	string
	LastName	string
	Email 		string
	Password 	string
}

type Login struct {
	Email string
	Password string
}

var token = "fake"

func postRegister()  {
	tmp := "333"
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetBody(Person{FirstName:tmp, LastName:tmp, Email:tmp, Password:tmp}).
		Post("http://localhost:8080/v1/member/register")

	showResponse(resp, err)
}

func postLogin()  {
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetBody(Login{Email:"111", Password:"111"}).
		Post("http://localhost:8080/v1/member/login")

	token = resp.RawResponse.Header.Get("Token")

	showResponse(resp, err)
}

func getMember(){
	resp, err := resty.R().Get("http://localhost:8080/v1/member/" + token)
	showResponse(resp, err)
}

func putMember()  {
	resp, err := resty.R().SetBody(Person{
		FirstName: "111",
		LastName: "111",
	}).Put("http://localhost:8080/v1/member/" + token)

	showResponse(resp, err)
}

func showResponse(resp *resty.Response, err error)  {

	// explore response object
	fmt.Printf("\nError: %v", err)
	fmt.Printf("\nResponse Status Code: %v", resp.StatusCode())
	fmt.Printf("\nResponse Status: %v", resp.Status())
	fmt.Printf("\nResponse Time: %v", resp.Time())
	fmt.Printf("\nResponse Recevied At: %v", resp.ReceivedAt())
	fmt.Printf("\nResponse Body: %v", resp)
}

func main() {
	postLogin()
	getMember()
	//putMember()
	//postRegister()
}
