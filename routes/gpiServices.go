package routes

	import (
	"github.com/martini-contrib/render"
	"fmt"
	"log"
	"time"
	"strings"
	"oauth/data"
	"net/http"
	"encoding/json"
	"oauth/model"
	"oauth/mailer"
	//"oauth/util"
)

//func Authzer(r render.Render, paramss martini.Params, res http.ResponseWriter, req *http.Request) {
func GPIServices(r render.Render, res http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	uid := req.FormValue("uid")
	fmt.Printf("----gpi----"+uid+"-----services-----")
	item := data.AllGPIServiceItems{}
	items, err := item.GetAllServices(uid)
	//res.Header().Set("Access-Control-Allow-Origin", "*")
	//res.Header().Add("Content-Type", "application/json")
	fmt.Printf(uid)
	if err != nil {
		// res.WriteHeader(http.StatusInternalServerError)
		// fmt.Println(err.Error())
		// payload := responsePayload{}
		// payload.Message = string(err.Error())
		// payload.Code = http.StatusInternalServerError
		// payload.Error = err
		// json.NewEncoder(w).Encode(payload)
		//return
	}
	
	r.JSON(200, items)
	return
}

func GPISubscribe(r render.Render, res http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	uid := req.FormValue("uid")
	username := req.FormValue("username")

	result := make(map[string]interface{})
	result["username"] = username
	result["uid"] = uid

	r.HTML(200, "subscribe", result)

	return
}

func SendMail(r render.Render, res http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	//from := req.FormValue("from")
	to := req.FormValue("to")
	from := req.FormValue("from")
	subject := req.FormValue("subject")
	names := req.FormValue("names")


	eMail := mailer.NewRequest(from, to, subject, []string{})
	data := mailer.Data{}
	//data.Item = ""
	data.Names = names
	data.CreatedAt = time.Now()
	eMail.Send("templates/notification.html", data)

	result := make(map[string]interface{})
	result["subject"] = subject
	result["names"] = names

	r.JSON(200, result)

	return
}

func GPISubscriptionServices(r render.Render, res http.ResponseWriter, req *http.Request) {
	req.ParseForm()
    log.Printf("%+v\n", req.Form)
    fmt.Printf("%+v\n", req.Form)
    productsSelected := req.Form["subcheck"]
    username := req.FormValue("username")
    uid := req.FormValue("uid")
    //log.Println(contains(productsSelected, "1"))
    fmt.Println(contains(productsSelected, "", username, uid))

    _, sub_clientId, v2 := data.GetCodeInRed(username)

    var st = new(model.ServiceTicket)
    if err := json.Unmarshal([]byte(sub_clientId), st); err != nil {
       // return err
    }

    fmt.Printf("%+v\n", st)

    fmt.Println("start - "+sub_clientId+" - end")
    fmt.Println("start - "+v2+" - end")

    redirectToServiced(r, st, username)
}

func contains(slice []string, item string, username string, uid string) bool {
    set := make(map[string]struct{}, len(slice))
    start_date := time.Now().Unix()
    end_date := time.Now().Unix()
    status := "1"
    for _, s := range slice {
        set[s] = struct{}{}
        //fmt.Printf("%+v\n", set[s])
        //fmt.Printf(s)
        
        service_id := s

        data.Conn.Exec("insert into GPI_Oauth_clients.subscription_history ( user_id, start_date, end_date, service_id, status)" +
			" values ( '" + uid + "', '" + string(start_date) + "', '" + string(end_date)+ "', '" +
			service_id + "', '" + status + "')")

    }
    _, ok := set[item]
    return ok
}

func redirectToServiced(r render.Render, st *model.ServiceTicket, name string) {
	needAnd := strings.Contains(st.Service, "?")
	sep := "?"
	if needAnd {
		sep = "&"
	}
	redirectUrl := st.Service + sep + "ticket=" + st.St + "&parsedobj="+name
	log.Println("redirect to serviceb:" + redirectUrl)
	fmt.Println("redirect to serviceb:" + redirectUrl)
	r.Redirect(redirectUrl)
	//r.Redirect("callRoute/"+redirectUrl)
}