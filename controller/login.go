package controller

import (
	"net/http"
	"encoding/json"
	"chat/utils"
	"github.com/dgrijalva/jwt-go"
	"time"
	"fmt"
)

type UserRequest struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}
func Login(w http.ResponseWriter, request *http.Request) error  {

	if request.Body==nil{
		utils.JsonResp(w,-1,"不能为空",nil)

		return nil
	}
	var user UserRequest
	err:=json.NewDecoder(request.Body).Decode(&user)
	defer request.Body.Close()
	if err!=nil{
		utils.JsonResp(w,-1,err.Error(),nil)
		return nil
	}
	if len(user.UserName)<1{
		utils.JsonResp(w,-1,"用户名不能为空",nil)

		return nil
	}
	if len(user.Password)<1{
		utils.JsonResp(w,-1,"密码不能为空",nil)

		return nil
	}

	if user.UserName=="admin"&& user.Password=="admin888..."{

		token := jwt.New(jwt.SigningMethodHS256)
		claims := make(jwt.MapClaims)
		claims["exp"] = time.Now().Add(time.Hour * time.Duration(1)).Unix()
		claims["iat"] = time.Now().Unix()
		claims["username"]=user.UserName
		token.Claims = claims

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w, "Error extracting the key")

		}

		tokenString, err := token.SignedString([]byte("rJXUCzdnN8mf"))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w, "Error while signing the token")

		}
		utils.JsonResp(w,20000,"登录成功",tokenString)


	}else{
		utils.JsonResp(w,-1,"用户密码不正确",nil)

	}
return nil


}