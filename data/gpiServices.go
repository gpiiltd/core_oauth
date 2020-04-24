package data

import (
	//"oauth/model"
	//"oauth/util"
	//"strings"
	"fmt"
	//"time"
)

type AllGPIServiceItems struct {
	ID          string  `json:"id"`
	Name        string  `json:"sname"`
	Description string  `json:"description"`
	DClass       string `json:"dclass"`
	IClass       string `json:"iclass"`
}
//

func (serviceItem *AllGPIServiceItems) GetAllServices(user_id string) ([]AllGPIServiceItems, error) {
	//defer Conn.Close()
	fmt.Printf("-------"+user_id+"------")
	rows, _ := Conn.Query("select id, sname, description, dclass, iclass from GPI_Oauth_clients.allservices where id NOT IN (select service_id from GPI_Oauth_clients.subscription_history where user_id='"+user_id+"')")
	
	var ServiceItems []AllGPIServiceItems
	var ServiceItem AllGPIServiceItems

	fmt.Printf("%+v\n", rows)

	for rows.Next() {
		_ = rows.Scan(&ServiceItem.ID, &ServiceItem.Name, &ServiceItem.Description, &ServiceItem.DClass, &ServiceItem.IClass)

		ServiceItems = append(ServiceItems, ServiceItem)
	}

	return ServiceItems, nil
}

func (serviceItem *AllGPIServiceItems) GetEverySingleGPIServices() ([]AllGPIServiceItems, error) {
	//defer Conn.Close()
	rows, _ := Conn.Query("select id, sname, description, dclass, iclass from GPI_Oauth_clients.allservices")

	
	var ServiceItems []AllGPIServiceItems
	var ServiceItem AllGPIServiceItems

	for rows.Next() {
		_ = rows.Scan(&ServiceItem.ID, &ServiceItem.Name, &ServiceItem.Description, &ServiceItem.DClass, &ServiceItem.IClass)

		ServiceItems = append(ServiceItems, ServiceItem)
	}

	return ServiceItems, nil
}

func GetSubscription(clientId string, user_id string) bool {
	//defer Conn.Close()
	row := Conn.QueryRow("SELECT id FROM GPI_Oauth_clients.subscription_history where user_id='"+user_id+"' and status = '1' and service_id in (select id from GPI_Oauth_clients.allservices where dclass= '"+clientId+"')");
	var client_id string

	row.Scan(&client_id)

	if client_id != "" {
		
		return true
	}
	return false

}