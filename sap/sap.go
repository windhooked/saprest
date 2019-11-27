package sap

import (
	"errors"
	"fmt"
	"os"

	"github.com/SAP/gorfc/gorfc"
)

//Rfc Instance
type Rfc struct {
	Conn   *gorfc.Connection
	Config gorfc.ConnectionParameter
}

//New Sap Rfc Instance
//func (s *Rfc) New(params map[string]string) (*Rfc, error) {
func New() (*Rfc, error) {

	var err error
	rfc := new(Rfc)

	rfc.Config =
		gorfc.ConnectionParameter{
			Dest:   os.Getenv("SAP_DEST"),
			Client: os.Getenv("SAP_CLIENT"),
			User:   os.Getenv("SAP_USER"),
			Passwd: os.Getenv("SAP_PASSWORD"),
			Lang:   os.Getenv("SAP_LANG"),
			Ashost: os.Getenv("SAP_ASHOST"),
			Sysnr:  os.Getenv("SAP_SYSNR"),
			//Saprouter: "/H/203.13.155.17/S/3299/W/xjkb3d/H/172.19.137.194/H/",
		}

	rfc.Conn, err = gorfc.ConnectionFromParams(s.Config)
	if err != nil {
		//log.Fatalln(err)
		return nil, err
	}

	defer rfc.Conn.Close()

	if err := s.Conn.Ping(); err != nil { // check if connection
		return nil, err
	}
	return rfc, nil

}

//Call RFC Call to SAP for given function
func (s *Rfc) Call(function string, params map[string]interface{}) (map[string]interface{}, error) {

	// check connection alive, open if not
	if !s.Conn.Alive() {
		s.Conn.Open()
	}

	// check connection or return if not reachable
	if err := s.Conn.Ping(); err != nil { // check if connection
		return nil, err
	}

	// now call SAP
	r, errSAP := s.Conn.Call(function, params)

	if errSAP != nil {
		return nil, errSAP
	}

	return r, nil
}
