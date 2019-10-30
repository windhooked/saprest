package main

import (
	"os"
	"log"
	"github.com/SAP/gorfc/gorfc"
)
// abapSystem save your connection info in .env
func abapSystem() gorfc.ConnectionParameter {
	return gorfc.ConnectionParameter{
		Dest:   os.Getenv("SAP_DEST"),
		Client: os.Getenv("SAP_CLIENT"),
		User:   os.Getenv("SAP_USER"),
		Passwd: os.Getenv("SAP_PASSWORD"),
		Lang:   os.Getenv("SAP_LANG"),
		Ashost: os.Getenv("SAP_ASHOST"),
		Sysnr:  os.Getenv("SAP_SYSNR"),
		//Saprouter: "/H/203.13.155.17/S/3299/W/xjkb3d/H/172.19.137.194/H/",
	}
}

//callSAP
func callSAP(function string, params map[string]interface{}) ( map[string]interface{}, error) {

	c, err := gorfc.ConnectionFromParams(abapSystem())
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	defer c.Close()

	r, errSAP := c.Call(function, params)

	if errSAP != nil {
		return nil, errSAP
	}

	return r, nil
}
