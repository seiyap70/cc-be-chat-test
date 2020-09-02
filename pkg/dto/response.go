package dto

type Response struct{
	Code int `json:"code"`
	Message string `json:"message"`
	Result interface{} `json:"result"`
}

type ChatServerInfo struct {
	AccessUrl string `json:"accessUrl"`
}