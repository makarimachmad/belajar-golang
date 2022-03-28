package chatactive

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/tidwall/gjson"
)

type (
        Docker struct {
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
                Contents            string      `gorm:"contents" json:"contents,omitempty"`
                Status              string
                Created_date        string
        }

        Peubah struct {
                CreatedDate     string  `gorm:"created_date" json:"created_date,omitempty"`
                IsiWaktu        int     `gorm:"isi_waktu" json:"isi_waktu,omitempty"`
		Phone		string	`gorm:"phone" json:"phone,omitempty"`
        }

        Summary struct {
                ID         int          `gorm:"id" json:"id,omitempty"`
                ClientID   string       `gorm:"client_id" json:"client_id,omitempty"`
                ClientName string       `gorm:"client_name" json:"client_name,omitempty"`
                Total      int          `gorm:"total" json:"total,omitempty"`
                Periode    string       `gorm:"periode" json:"periode,omitempty"`
        }
)

func main() {
        db, db2, err := connect()
        if err != nil {
                fmt.Println(err.Error())
                return
        }
        defer db.Close()

        rows, err := db.Query("select * from docker_callback_log where client_id = ? order by created_date asc", 64)
        if err != nil {
                fmt.Println(err.Error())
                return
        }
        defer rows.Close()

        //var zerodate = "0000-00-00 00:00:00"
        var zerodateangka = 0
        //var tampungphone, tampungname []string
	var flag = ""
        // var flagnama = ""
        var count = 0

        bantupeubah := new(Peubah)

        for rows.Next() {
                var each = Docker{}
                var err = rows.Scan(&each.ID, &each.CorrelationID, &each.ClientID, &each.ClientName, &each.ClientMaskingID, &each.ClientMaskingNumber, &each.ClientMaskingName, &each.IpAddress, &each.CallbackURL, &each.ResponseCode, &each.ResponseTime, &each.Response, &each.Header, &each.Contents, &each.Status, &each.Created_date)
                if err != nil {
                        fmt.Println(err.Error())
                        return
                }

                fmt.Println("id docker: ", each.ID)
		fmt.Println("waktu creat: ", each.Created_date)

                //ambil nilai timestamp dari multi lines json content
                timestampp := gjson.Get(each.Contents, "messages.0.timestamp")

                //cek tanggal sebelumnya kosong atau tidak
                if bantupeubah.IsiWaktu == 0 {
                        bantupeubah.IsiWaktu = zerodateangka
                }

                bantupeubah.CreatedDate = time.Unix(timestampp.Int(), 0).UTC().String()
                fmt.Println("bantupeubah: ", bantupeubah.CreatedDate)

                waid :=  gjson.Get(each.Contents, "contacts.0.wa_id")
                fmt.Println("waid: ", waid)
                nama :=  gjson.Get(each.Contents, "contacts.0.profile.name")
                fmt.Println("nama: ", nama)
  
		// for _, i := range tampungphone{
		// 	if waid.String() == i{
		// 		flag = "sudah"
		// 	}else{
		// 		flag = "belum"
		// 	}
		// }
                // for _, i := range tampungname{
                //         if nama.String() == i{
                //                 flagnama = "sudah"
		// 	}else{
		// 		flagnama = "belum"
		// 	}
                // }

                var idt int
		data, err := db2.Query("select id from session_chat_active where customer_phone = ? order by last_chat_date desc limit 1", waid.String())
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		defer data.Close()

		for data.Next(){
			data.Scan(&idt)
		}

		fmt.Println("flag: ", flag)
		if idt==0{
			_, err = db2.Exec("insert into session_chat_active values (?, ?, ?, ?, ?, ?, ?, ?)", "", &each.ClientID, &each.ClientName, waid.String(), nama.String()+" belum", bantupeubah.CreatedDate, bantupeubah.CreatedDate, time.Now().Format("2006-01-02 15:04:05"))
			if err != nil {
				fmt.Println(err.Error())
				return
			}
                        count=count+1
                        //tampungphone = append(tampungphone, waid.String())
                        // tampungname = append(tampungname, nama.String())
			fmt.Println(waid.String(), " belum ada")
		}else{	
			//ambil last chatdatenya berdasarkan numberphone
			var idt int
			var lcd string
			data, err := db2.Query("select id, last_chat_date from session_chat_active where customer_phone = ? order by last_chat_date desc limit 1", waid.String())
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			defer data.Close()

			for data.Next(){
				data.Scan(&idt, &lcd)
			}

			//fmt.Println("Data: ", data)
			fmt.Println("idt: ", idt)
			fmt.Println("last chat: ", lcd)
			//sekarang := time.Unix(int64(bantupeubah.IsiWaktu), 0)
			sekarang,_ := time.Parse("2006-01-02 15:04:05", lcd)
			expTime := int(sekarang.Add(time.Second * time.Duration(86400)).Unix())
			selisih := time.Unix(timestampp.Int(), 0).Sub(sekarang)
			scek := int(sekarang.Unix())
			bcek := timestampp.Int()
			
			fmt.Println("sekarang: ", sekarang.String())
			fmt.Println("scek: ", scek)
			fmt.Println("bcek: ", bcek)
			fmt.Println("berikutnya: ", bantupeubah.CreatedDate)
			fmt.Println("selisih: ", selisih)
			fmt.Println("expiredTime: ", expTime)
			
			if bcek > int64(expTime){
                                fmt.Println("insert")
				_, err = db2.Exec("insert into session_chat_active values (?, ?, ?, ?, ?, ?, ?, ?)", "", &each.ClientID, &each.ClientName, waid.String(), nama.String()+" expired", bantupeubah.CreatedDate, bantupeubah.CreatedDate, time.Now().Format("2006-01-02 15:04:05"))
				if err != nil {
					fmt.Println(err.Error())
					return
				}
                                count=count+1
				fmt.Println(waid.String(), " sudah expired")
                                // tampungphone = append(tampungphone, waid.String())
                                // tampungname = append(tampungname, nama.String())
			}else{
                                
                                fmt.Println("update")
				_, err = db2.Exec("update session_chat_active set last_chat_date = ? where id = ?", bantupeubah.CreatedDate, idt)
				if err != nil {
					fmt.Println(err.Error())
					return
				}
				fmt.Println(waid.String(), " update last chat")
                                // tampungphone = append(tampungphone, waid.String())
                                // tampungname = append(tampungname, nama.String())
			}
		}
                flag = ""
                //flagnama = ""
		// tampungphone = append(tampungphone, waid.String())
                // tampungname = append(tampungname, nama.String())

                bantupeubah.IsiWaktu = int(timestampp.Int())
                fmt.Println("setelah diisi bantupeubah: ", bantupeubah.IsiWaktu)
                fmt.Println("===================================================")
        }

        if err = rows.Err(); err != nil {
                fmt.Println(err.Error())
                return
        }

        active, err := db2.Query("SELECT client_id, client_name, COUNT(*) AS total_chat, MONTH(first_chat_date) AS bulan FROM session_chat_active GROUP BY MONTH(first_chat_date)")
        if err != nil {
                fmt.Println(err.Error())
                return
        }
        defer active.Close()

        summary := new(Summary)
        for active.Next() {

                var err = active.Scan(&summary.ClientID, &summary.ClientName, &summary.Total, &summary.Periode)

                if err != nil {
                        fmt.Println(err.Error())
                        return
                }
        }
        fmt.Println(summary)
        fmt.Println("total chat: ", count)

        _, err = db2.Exec("insert into session_chat_summary values (?, ?, ?, ?, ?)", "", summary.ClientID, summary.ClientName, summary.Total, summary.Periode)
        if err != nil {
                fmt.Println(err.Error())
                return
        }
}

func connect() (*sql.DB, *sql.DB, error) {
        db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/wappin_go_log_db")
        if err != nil {
                return nil, nil, err
        }

        db2, err := sql.Open("mysql", "root:@tcp(localhost:3306)/wappin_go_db")
        if err != nil {
                return nil, nil, err
        }

        return db, db2, nil
}
