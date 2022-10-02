package main

import (
	"io/ioutil"
	"log"

	"github.com/go-resty/resty/v2"
)

const (
	URI_BPM = "API/"
)

func init() {
	log.SetOutput(ioutil.Discard)
}

// TODO: Need chanage to Variables of Environment
var (
	user_ppassword = "bpm"
	server_ip_port = "localhost:8082"
)

// sources := fmt.Sprintf(server_addr,
//  // os.Getenv("BPM_SERVER_ADDR"),
//  os.Getenv("b.server"),
// )

//
//  FormInput
//  @Description: CRUD必要的 JSON 外層結構
//
type FormInput struct {
	ModelInput *interface{} `json:"modelInput,omitempty"`
}

type BPMClient struct {
	serverUri  string
	apiUri     string
	username   string
	password   string
	request    *resty.Request
	token      string
	jSessionId string // JSESSIONID
}
