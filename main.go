package main

import (
	"chat/controller"
	"net/http"

	"gopkg.in/yaml.v2"

	"chat/model"
	"io/ioutil"
	"log"
	"chat/utils"
)

func init() {
	config := new(model.Config)
	yamlFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Print(err)
	}
	err = yaml.Unmarshal(yamlFile, config)

	if err != nil {
		log.Print(err)
	}
	log.Print(config.User.Account)

}
func main() {


	http.HandleFunc("/login", utils.Wrapper(controller.Login))
	http.HandleFunc("/list", utils.Wrapper(controller.List))
	http.ListenAndServe(":5600",nil)
}
