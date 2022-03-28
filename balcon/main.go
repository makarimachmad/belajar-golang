package main

import (
	"belajar-golang/balcon/config"
	"log"
	"os"
	"strings"
	"time"
	"encoding/json"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type (
        DockerCallbackLog struct {
                ID                  int
                CorrelationID       string
                ClientID            string
                ClientName          string
                ClientMaskingID     int
                ClientMaskingNumber string
                ClientMaskingName   string
                IpAddress           string
                CallbackURL         string
                ResponseCode        string
                ResponseTime        string
                Response            string
                Header              string
                Content	            string		`json:"content"`
                Status              string
                Created_date        string
        }

	DCL	[]DockerCallbackLog

        Peubah struct {
                CreatedDate     string  `json:"created_date"`
                IsiWaktu        int     `json:"isi_waktu"`
        }

	Stat struct{
		Status	[]Statuses	`json:"statuses,omitempty"`
	}

	Statuses struct{
		Conversations	Conversation	`json:"conversation,omitempty"`
		ID		string		`json:"id,omitempty"`
		Pricings	Pricing		`json:"pricing,omitempty"`
		RecipientID	string		`json:"recipient_id,omitempty"`
		Status		string		`json:"status,omitempty"`
		Timestamp	string		`json:"timestamp,omitempty"`
	}

	Conversation struct{
		ID	string	`json:"id"`
	}

	Pricing struct{
		Billable	bool	`json:"billable"`
		PricingModel	string	`json:"pricing_model"`

	}

	Content struct {
		Contact []Contacts	`json:"contacts,omitempty"`
		Message []Messages	`json:"messages,omitempty"`
	}

	Contents []Content

	Contacts struct {
		Profiles Profile	`json:"profile,omitempty"`
		WaID     string		`json:"wa_id,omitempty"`
	}

	Profile struct {
		Name string		`json:"name,omitempty"`
	}

	Messages struct {
		From      string	`json:"from,omitempty"`
		ID        string	`json:"id,omitempty"`
		Timestamp string	`json:"timestamp,omitempty"`
		Type      string	`json:"type,omitempty"`
		Texts	  Text		`json:"text,omitempty"`
	}

	Text struct{
		Body	string	`json:"body,omitempty"`
	}

	Button struct {
		Text string		`json:"text,omitempty"`
	}

	Context struct {
		From string		`json:"from,omitempty"`
		ID   string		`json:"id,omitempty"`
	}

	BniReplyBalcon struct{
		ID		int	`json:"id,omitempty"`
		WaNumber	string	`json:"wa_number,omitempty"`
		WaName		string	`json:"wa_name,omitempty"`
		Tenor		string	`json:"tenor"`
		Reply		string	`json:"reply"`
		CreatedDate	string	`json:"created_date"`
	}
	Brbs	[]BniReplyBalcon
	
)

// TableName of bni_reply_balcons
func (BniReplyBalcon) TableName() string {
	return "bni_reply_balcons"
}

// TableName of bni_reply_balcons
func (DockerCallbackLog) TableName() string {
	return "docker_callbak_logs"
}

func main() {
	//koneksi database
        db, db2, err := connect()
        if err != nil {
                fmt.Println(err.Error())
                return
        }
	
	dcl := new(DCL)	//dockercallbacklogs data
	content := new(Content)	//contents data
	brb := new(BniReplyBalcon)

	err = db2.Table("docker_callback_logs").Where("status=?","messages").Scan(&dcl).Error	//get data dockercallbacklogs
	if err != nil {
		fmt.Println(err)
	}

	//unmarshall data content
	f, err := os.Create("data.txt")
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	for _, i := range *dcl{
		json.Unmarshal([]byte(i.Content), &content)

		if content != nil{
			for _, j := range content.Message{
				if j.Type == "text"{
					strbaru1 := strings.Contains(j.Texts.Body, "Balcon")
					strbaru2 := strings.Contains(j.Texts.Body, "balcon")

					_, err2 := f.WriteString(j.Texts.Body+"\n")

						if err2 != nil {
							log.Fatal(err2)
						}
					
					if strbaru1  || strbaru2 {
						pagar := strings.Contains(j.Texts.Body, "#")

						if pagar {
							str := strings.Split(j.Texts.Body, "#")
							
							sjk := str[0]
							tnr := str[1]
							ktr := str[2]
							
							for _, k := range content.Contact{
								fmt.Println("Wa: ", k.WaID)
								fmt.Println("Name: ", k.Profiles.Name)
								brb.WaNumber = k.WaID
								brb.WaName = k.Profiles.Name
							}
							fmt.Println("Str: ", str)
							fmt.Println("DockerID: ", i.ID)
							fmt.Println("isi text: ", sjk, " ", tnr, " ", ktr)
							brb.Tenor = tnr
							brb.Reply = ktr
							sekarang, _ := time.Parse("2006-01-02 15:04:05 Z0700 MST", time.Now().Format("2006-01-02 15:04:05 Z0700 MST"))
							brb.CreatedDate = sekarang.String()
							brb.ID = 0
							db.Create(&brb)
							fmt.Println("----------------------------------------")
					
							_, err2 := f.WriteString(j.Texts.Body+"\n")

							if err2 != nil {
								log.Fatal(err2)
							}
						}
						
					
					}
				}
			}
			
		}
	}

	brbs := new(Brbs)

	err = db.Table("bni_reply_balcons").Scan(&brbs).Error	//get data dockercallbacklogs
	if err != nil {
		fmt.Println(err)
	}
	
	xlsx := excelize.NewFile()

	sheetName := "balconBNI"
	xlsx.SetSheetName(xlsx.GetSheetName(1), sheetName)
	xlsx.SetCellValue(sheetName, "A1", "NO.")
	xlsx.SetCellValue(sheetName, "B1", "WA NUMBER")
	xlsx.SetCellValue(sheetName, "C1", "WA NAME")
	xlsx.SetCellValue(sheetName, "D1", "TENOR")
	xlsx.SetCellValue(sheetName, "E1", "REPLY")
	xlsx.SetCellValue(sheetName, "F1", "CREATED DATE")

	err = xlsx.AutoFilter(sheetName, "B1", "c1", "")

	if err != nil {
		log.Println(err)
	}

	for i, each := range *brbs {
		xlsx.SetCellValue(sheetName, fmt.Sprintf("A%d", i+2), i+1)
		xlsx.SetCellValue(sheetName, fmt.Sprintf("B%d", i+2), each.WaNumber)
		xlsx.SetCellValue(sheetName, fmt.Sprintf("C%d", i+2), each.WaName)
		xlsx.SetCellValue(sheetName, fmt.Sprintf("D%d", i+2), each.Tenor)
		xlsx.SetCellValue(sheetName, fmt.Sprintf("E%d", i+2), each.Reply)
		xlsx.SetCellValue(sheetName, fmt.Sprintf("F%d", i+2), each.CreatedDate)
	}
	
	err = xlsx.SaveAs("balcon_"+time.Now().Format("20060102150405")+".xlsx")
	if err != nil {
		fmt.Println(err)
	}
	
}

func connect() (*gorm.DB, *gorm.DB, error) {
        
	dsn := config.Config("DB_USER")+":"+config.Config("DB_PASSWORD")+"@("+config.Config("DB_HOST")+":"+config.Config("DB_PORT")+")/"+config.Config("DB_NAME")
	DB, err := gorm.Open(mysql.Open(dsn))
	if err != nil{
		fmt.Println("error koneksi db")
	}

	dsn2 := config.Config("DB_USER")+":"+config.Config("DB_PASSWORD")+"@("+config.Config("DB_HOST")+":"+config.Config("DB_PORT")+")/"+config.Config("DB2_NAME")
	DB2, err := gorm.Open(mysql.Open(dsn2))
	if err != nil{
		fmt.Println("error koneksi db")
	}

        return DB, DB2, nil
}