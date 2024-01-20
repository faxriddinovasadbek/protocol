package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/k0kubun/pp"
)

type User struct {
	Name string `json:"name"`
	ID   uint32 `json:"id"`
}

func GetMetod() {

	resp, err := http.Get("https://httpbin.org/get")
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return
	}

	pp.Printf(string(data))
}

func DeleteMetod() {

	client := &http.Client{}

	req, err := http.NewRequest(http.MethodDelete, "https://httpbin.org/delete", nil)
	if err != nil {
		panic(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	pp.Printf(string(body))
}

func PatchMetod() {
	url := "https://httpbin.org/patch"
	payload := map[string]string{"name": "New Name"}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest(http.MethodPatch, url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	pp.Printf(string(data))
}

func PostMetod() {
	var u = User{
		Name: "Alex",
		ID:   1,
	}

	bytesRepresentation, err := json.Marshal(u)
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := http.Post("https://httpbin.org/post", "application/json", bytes.NewBuffer(bytesRepresentation))
	if err != nil {
		log.Fatalln(err)
	}

	bytesResp, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return
	}

	pp.Printf(string(bytesResp))

}

func PutMetod() {

	url := "https://httpbin.org/put"
	data := []byte(`{"name": "John", "age": 30}`)

	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(data))
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	bytesResp, err := io.ReadAll(resp.Body)

	pp.Printf(string(bytesResp))

}

func main() {

	pp.Println("get")
	GetMetod()

	pp.Println("delete")
	DeleteMetod()

	pp.Println("put")
	PutMetod()

	pp.Println("post")
	PostMetod()

	pp.Println("patch")
	PatchMetod()

}
