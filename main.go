package main

import (
	"log"
	"fmt"
	"database/sql"
	_ "github.com/lib/pq"
	//"github.com/mrtomyum/paybox_cloud/vending"
	"encoding/json"
)

const (
	sslMode = "disable"
	dbPort = "5432"
	dbUser = "paybox"
	//dbHost = "paybox.work"
	//dbPass = ""
	//dbName = "paybox_vending"
	dbHost = "9tom.me"
	dbPass = "paybox"
	dbName = "paybox"
)

func main() {
	log.Println("Paybox Clients Service")
	conn := fmt.Sprintf("dbname=%s user=%s password=%s host=%s port=%s sslmode=%s",
		dbName, dbUser, dbPass, dbHost, dbPort, sslMode)

	fmt.Println(conn)
	//init db
	db, err := sql.Open("postgres", conn)
	must(err)
	err = db.Ping()
	must(err)
	GetClient(db,0)
	defer db.Close()
}

func must(err error) {
	if err != nil {
		fmt.Println("Error:", err)
		log.Fatal(err)
	}
}

type Meta struct {
	ClientConfig  Content `json:"client_config"`
}

type Content struct {
	Logo         string `json:"logo"`
	Tax_id       string `json:"tax_id"`
	Member_url   string `json:"member_url"`
	Time_servers []string `json:"time_servers"`
	Client       Client `json:"client"`
}
type Client  struct {
	Shop_title               string `json:"shop_title"`
	Shop_website             string `json:"shop_website"`
	LogoImage_width          string `json:"logo_image_width"`
	Api_server_url           string `json:"api_server_url"`
	Cloud_server_url         string `json:"cloud_server_url"`
	Data_url                 string `json:"data_url"`
	Logo_path                string `json:"logo_path"`
	Image_path               string `json:"image_path"`
	Print_queue_receipt_mode int `json:"print_queue_receipt_mode"`
}

type ClientTable struct {
	Id int64
	Meta sql.NullString
}
func GetClient(db *sql.DB, id int64) (error) {
	_c:=ClientTable{}
	m := Meta{}

	rs := db.QueryRow("select id,meta from clients where id = ?",id)
	rs.Scan(&_c)
	stringMeta := _c.Meta.String



	_ = json.Unmarshal([]byte(stringMeta), &m)
	fmt.Println(m)
	return nil
}