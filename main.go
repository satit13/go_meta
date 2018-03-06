package main

import (
	"log"
	"fmt"
	"database/sql"
	_ "github.com/lib/pq"
	//"github.com/mrtomyum/paybox_cloud/vending"
	"encoding/json"
	"strconv"
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


func must(err error) {
	if err != nil {
		fmt.Println("Error:", err)
		log.Fatal(err)
	}
}

type Meta struct {
	ClientConfig Content `json:"client_config"`
}

type Content struct {
	Logo              string `json:"logo"`
	Tax_id            string `json:"tax_id"`
	Member_url        string `json:"member_url"`
	Time_servers      []string `json:"time_servers"`
	Client            Client `json:"client"`
	Cash              Cash `json:"cash"`
	CashAcceptedValue CashAcceptedValue `json:"cash_accepted_value"`
	ChangeRefillValue ChangeRefillValue `json:"change_refill_value"`
}

type Client  struct {
	Shop_title               string `json:"shop_title"`
	Shop_website             string `json:"shop_website"`
	LogoImage_width          int `json:"logo_image_width"`
	Api_server_url           string `json:"api_server_url"`
	Cloud_server_url         string `json:"cloud_server_url"`
	Data_url                 string `json:"data_url"`
	Logo_path                string `json:"logo_path"`
	Image_path               string `json:"image_path"`
	Print_queue_receipt_mode int `json:"print_queue_receipt_mode"`
}

type Cash  struct {
	Max_change_cash int `json:"max_change_cash"`
	Max_change_coin int `json:"max_change_coin"`
	Min_change      int `json:"min_change"`
}

type CashAcceptedValue struct {
	B20   int `json:"b_20"`
	B50   int `json:"b_50"`
	B100  int `json:"b_100"`
	B500  int `json:"b_500"`
	B1000 int `json:"b_1000"`
}

type ChangeRefillValue struct {
	V1    int `json:"v1"`
	V2    int `json:"v2"`
	V5    int `json:"v5"`
	V10   int `json:"v10"`
	V20   int `json:"v20"`
	V50   int `json:"v50"`
	V100  int `json:"v100"`
	V500  int `json:"v500"`
	V1000 int `json:"v1000"`
}

type ClientTable struct {
	Id   int64
	Meta string
}


func main() {
	log.Println("Meta Testt ")
	conn := fmt.Sprintf("dbname=%s user=%s password=%s host=%s port=%s sslmode=%s",
		dbName, dbUser, dbPass, dbHost, dbPort, sslMode)

	fmt.Println(conn)
	//init db
	db, err := sql.Open("postgres", conn)
	must(err)
	err = db.Ping()
	must(err)

	err = GetClient(db,1)
	must(err)
	defer db.Close()
}

func GetClient(db *sql.DB,pid int) (error) {
	//_c := ClientTable{}
	m := Meta{}
	lccommand := "select id,meta from clients where id =  "+strconv.Itoa(pid)
	//lccommand := "select id,meta from clients where id = ? "

	fmt.Println(lccommand)
	fmt.Println("pid = ",pid)
	rs := db.QueryRow(lccommand)
	//rs := db.QueryRow(lccommand,pid)

	var _meta string
	var _id int
	rs.Scan(&_id,&_meta)
	fmt.Printf("client_id = %v\n ", _id)
	//stringMeta := _c.Meta

	_ = json.Unmarshal([]byte(_meta), &m)

	fmt.Println("meta object " , m)
	fmt.Println(m.ClientConfig.Logo)
	return nil
}