package Task

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type ca struct {
	CompanyName string `json:"companyname"`
	CPort       string `json:"cport"`
	CreditCode  string `json:"creditcode"`
	DPort       string `json:"dport"`
	Ip          string `json:"ip"`
	TaxCode     string `json:"taxcode"`
	TaxpayerID  string `json:"taxpayerid"`
	UsbPort     string `json:"usbport"`
	VidPid      string `json:"vidpid"`
}

type Task struct {
	CA           ca     `json:"ca"`
	Data         string `json:"data"`
	Env          string `json:"env"`
	Message      string `json:"message"`
	SerialNumber string `json:"serialnumber"`
	Status       string `json:"status"`
	Uuid		 string `json:"uuid"`
	IP		 	 string `json:"IP"`
}

func (task *Task) String() string {
	b, err := json.Marshal(*task)
	if err != nil {
		return fmt.Sprintf("%+v", *task)
	}
	var out bytes.Buffer
	err = json.Indent(&out, b, "", "    ")
	if err != nil {
		return fmt.Sprintf("%+v", *task)
	}
	return out.String()
}
