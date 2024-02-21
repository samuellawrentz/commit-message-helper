package data

import "commit-helper/models"

type AppData struct {
    UserDetails models.UserDetails 
    TicketData models.Storage
}

var data *AppData

func GetData() *AppData {
    if data == nil {
        data = &AppData{}
    }
    return data
}

func (d *AppData) SetUserDetails(userDetails models.UserDetails) {
    d.UserDetails = userDetails
}

func (d *AppData) SetTicketData(ticketData models.Storage) {
    d.TicketData = ticketData
}
