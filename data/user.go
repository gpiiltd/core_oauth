package data

import (
	"oauth/model"
	"oauth/util"
	//"strings"
	"fmt"
	"log"
)

func GetEnterprise(user_id string) *model.GPICompanyDetails {
	fmt.Println(user_id)
	row := Conn.QueryRow("select id, account_owner, account_id, company, companyname FROM companies where account_id='" + user_id + "'")
	var id int
	var account_owner string
	var account_id string
	var company string
	var companyname string
	
	row.Scan(&id, &account_owner, &account_id, &company, &companyname)

	if company != "" {
		var u = new(model.GPICompanyDetails)
		u.ID = id
		u.AccountOwner = account_owner
		u.AccountID = account_id
		u.Company = company
		u.CompanyName = companyname
		
		return u
	}
	
	return nil
}
func GetLoginDetails(clientId string) *model.GPIUser {
	//defer Conn.Close()
	row := Conn.QueryRow("select id, myemail, username, fname, lname, dob, gender, address, city, picture, company FROM users where username='" + clientId + "'")
	var id int
	var myemail string
	var username string
	var fname string
	var lname string
	var dob string
	var gender string
	var picture string
	var address string
	var city string
	var company string

	row.Scan(&id, &myemail, &username, &fname, &lname, &dob, &gender, &address, &city, &picture, &company)

	if username != "" {
		var u = new(model.GPIUser)
		u.ID = id
		u.Lastname = lname
		u.Firstname = fname
		u.Dob = dob
		u.Sex = gender
		u.Email = myemail
		u.Address = address
		u.City = city
		u.Username = username
		u.Picture = util.Server.SSO+"/profilepics/"+picture
		u.Company = company
		fmt.Println("Looged in: "+username)
		log.Println("Looged in: "+username)
		return u
	}
	
	return nil

}
