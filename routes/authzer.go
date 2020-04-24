package routes

import (
	"github.com/martini-contrib/render"
  	"github.com/go-martini/martini"
	"net/http"
	//"net/url"
	"oauth/model"
	"log"
	"fmt"
	"io/ioutil"
	"oauth/util"
	"encoding/json"
)
type AResult struct {
	Result   bool         `json:"result"`
	Code     int   `json:"code"`
	Msg      string   `json:"msg"`
	Username string   `json:"username"`
}

func Authzer(r render.Render, paramss martini.Params, res http.ResponseWriter, req *http.Request) {
	//parse request parameters
	req.ParseForm()


	//var reqParam martini.Params
	code := req.FormValue("code")
	parsedobj := req.FormValue("parsedobj")
	log.Println(code)
	fmt.Println(code)
	fmt.Println(paramss["client"])
	fmt.Println(paramss["id"])
	fmt.Println(parsedobj)

	var linkvar string
	cookies := req.Cookies()
	for _, cookie := range cookies {
		if cookie.Name == "GPIGC" {
			fmt.Println(cookie.Value)
			linkvar=cookie.Value
			fmt.Println("usersd found in cookies:")
			fmt.Println(cookie.Name + " " + cookie.Value)
		} else {
			fmt.Println("usersd not found in cookies:")
			fmt.Println(cookie.Name + " " + cookie.Value)	
		}
	}
	//acookie := http.Cookie{Name: "USERSD", Value: "username", Path: "/", Domain: "127.0.0.1", MaxAge: -1}
	//http.SetCookie(res, &acookie)
	//http://localhost:3000/oauth/token?client_id=gpitest&client_secret=gpitest
	//&grant_type=authorization_code&code=e6enYoiF&redirect_uri=my-gpi.com/gpitest
	//client_id=gpitest&client_secret=gpitest&grant_type=authorization_code&code=e6enYoiFredirect_uri=my-gpi.com
	validateUrl := util.Server.OAuth + "/oauth/token?"
	params := "client_id=" + paramss["id"] + "&client_secret=" + paramss["id"] + "&grant_type=authorization_code&code=" + code + "&redirect_uri=" + paramss["client"] + "/" + paramss["id"]
	//url.QueryEscape(util.Server.OAuth_CAS_Check+"?client_id="+client_id) + "&ticket=" + st

	fmt.Println(validateUrl+params)

	result := make(map[string]interface{})

	//return
	resp, err := http.Get(validateUrl + params)
	if err != nil {
		result["result"] = false
		result["code"] = 401
		result["msg"] = "error with credential"
		log.Println("error")
		log.Println(err)
		r.JSON(401, result)
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		result["result"] = false
		result["code"] = 401
		result["msg"] = "error reading login credential"
		log.Println("error reading response body")
		log.Println(resp.Body)
		log.Println(err)
		r.JSON(401, result)
		return
	}
	validateData := string(body)


	//r.JSON(401, result)
	//r.HTML(200, "auth", code)
	//return
	authz := new(model.TokenToReturn)
	err = json.Unmarshal([]byte(validateData), authz)
	if err != nil {
		result["result"] = false
		result["code"] = 401
		result["msg"] = "error unmarshal credential"
		log.Println("error unmarshal credential")
		log.Println(err)
		r.JSON(401, result)
		return
	}
	//r.HTML(200, authz)
	//http://localhost:3000/oauth/check?
	//access_token=be8aa907-cef9-4197-bf4b-5dcd08f8cff9&username=burumba
	//localhost:3000/oauth/check?access_token=be8aa907-cef9-4197-bf4b-5dcd08f8cff9&username=burumba
	validateUrl = util.Server.OAuth + "/oauth/check?"
	params = "access_token=" + authz.Access_token + "&username=" + linkvar
	resp, err = http.Get(validateUrl + params)
	if err != nil {
		result["result"] = false
		result["code"] = 401
		result["msg"] = "error with credential"
		log.Println("error")
		log.Println(err)
		r.JSON(401, result)
		return
	}

	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		result["result"] = false
		result["code"] = 401
		result["msg"] = "error reading login credential"
		log.Println("error reading response body")
		log.Println(resp.Body)
		log.Println(err)
		r.JSON(401, result)
		return
	}

	validateData = string(body)


	//r.JSON(401, result)
	//r.HTML(200, "auth", code)
	//return
	//authz := new(model.OAuthToken)
	//err = json.Unmarshal([]byte(validateData), authz)

	//r.JSON(200, authz)
	r.Text(200, validateData)
	return

}
