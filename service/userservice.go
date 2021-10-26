package service

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strconv"
	"time"

	"wxcloudrun-golang/db/dao"
	"wxcloudrun-golang/db/model"
)

func AddUser(w http.ResponseWriter, r *http.Request) {
	log.Print("received a AddUser request.")
	if r.Method != http.MethodPost {
		fmt.Fprintf(w, "AddUser only support POST method")
		return
	}

	decoder := json.NewDecoder(r.Body)
	body := make(map[string]interface{}, 0)
	if err := decoder.Decode(&body); err != nil {
		fmt.Fprintf(w, "%+v", err)
		return
	}
	defer r.Body.Close()

	user := new(model.UserModel)
	var ok bool
	var data interface{}
	if data, ok = body["name"]; ok {
		if user.Name, ok = data.(string); !ok {
			fmt.Fprintf(w, "name need string type but %+v", reflect.TypeOf(data))
			return
		}
	}
	if data, ok = body["email"]; ok {
		if user.Email, ok = data.(string); !ok {
			fmt.Fprintf(w, "email need string type but %+v", reflect.TypeOf(data))
			return
		}
	}
	if data, ok = body["phone"]; ok {
		if user.Phone, ok = data.(string); !ok {
			fmt.Fprintf(w, "phone need string type but %+v", reflect.TypeOf(data))
			return
		}
	}
	if data, ok = body["description"]; ok {
		if user.Description, ok = data.(string); !ok {
			fmt.Fprintf(w, "description need string type but %+v", reflect.TypeOf(data))
			return
		}
	}
	if data, ok = body["age"]; ok {
		var age float64
		if age, ok = data.(float64); !ok {
			fmt.Fprintf(w, "age need int type but %+v", reflect.TypeOf(data))
			return
		}
		user.Age = int32(age)
	}

	now := time.Now()
	user.CreateTime = now
	user.UpdateTime = now

	id, err := dao.Imp.AddUser(user)
	if err != nil {
		fmt.Fprintf(w, "%+v", err)
		return
	}

	fmt.Fprintf(w, "add user success id[%d]", id)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	log.Print("received a DeleteUser request.")
	if r.Method != http.MethodDelete {
		fmt.Fprintf(w, "DeleteUser only support DELETE method")
		return
	}

	query := r.URL.Query()
	ids, ok := query["id"]
	if !ok {
		fmt.Fprintf(w, "DeleteUser need param id")
		return
	}

	if len(ids) != 1 {
		fmt.Fprintf(w, "DeleteUser only support query one user once")
		return
	}

	id, err := strconv.Atoi(ids[0])
	if err != nil {
		fmt.Fprintf(w, "%+v", err)
		return
	}

	err = dao.Imp.DeleteUserById(int32(id))
	if err != nil {
		fmt.Fprintf(w, "%+v", err)
		return
	}

	fmt.Fprintf(w, "delete success")
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	log.Print("received a UpdateUser request.")
	if r.Method != http.MethodPut {
		fmt.Fprintf(w, "UpdateUser only support PUT method")
		return
	}

	decoder := json.NewDecoder(r.Body)
	body := make(map[string]interface{}, 0)
	if err := decoder.Decode(&body); err != nil {
		fmt.Fprintf(w, "%+v", err)
		return
	}
	defer r.Body.Close()

	user := new(model.UserModel)
	user.Id = -1
	var ok bool
	var data interface{}
	if data, ok = body["name"]; ok {
		if user.Name, ok = data.(string); !ok {
			fmt.Fprintf(w, "name need string type but %+v", reflect.TypeOf(data))
			return
		}
	}
	if data, ok = body["email"]; ok {
		if user.Email, ok = data.(string); !ok {
			fmt.Fprintf(w, "email need string type but %+v", reflect.TypeOf(data))
			return
		}
	}
	if data, ok = body["phone"]; ok {
		if user.Phone, ok = data.(string); !ok {
			fmt.Fprintf(w, "phone need string type but %+v", reflect.TypeOf(data))
			return
		}
	}
	if data, ok = body["description"]; ok {
		if user.Description, ok = data.(string); !ok {
			fmt.Fprintf(w, "description need string type but %+v", reflect.TypeOf(data))
			return
		}
	}
	if data, ok = body["age"]; ok {
		var age float64
		if age, ok = data.(float64); !ok {
			fmt.Fprintf(w, "age need int type but %+v", reflect.TypeOf(data))
			return
		}
		user.Age = int32(age)
	}
	if data, ok = body["id"]; ok {
		var id float64
		if id, ok = data.(float64); !ok {
			fmt.Fprintf(w, "id need int type but %+v", reflect.TypeOf(data))
			return
		}
		user.Id = int32(id)
	}

	if user.Id < 0 {
		fmt.Fprintf(w, "id[%d] not exist", user.Id)
		return
	}

	now := time.Now()
	user.UpdateTime = now

	err := dao.Imp.UpdateUserById(user.Id, user)
	if err != nil {
		fmt.Fprintf(w, "%+v", err)
		return
	}

	fmt.Fprintf(w, "update user success id[%d]", user.Id)

}

func QueryUser(w http.ResponseWriter, r *http.Request) {
	log.Print("received a QueryUser request.")
	if r.Method != http.MethodGet {
		fmt.Fprintf(w, "QueryUser only support GET method")
		return
	}

	query := r.URL.Query()
	ids, ok := query["id"]
	if !ok {
		fmt.Fprintf(w, "QueryUser need param id")
		return
	}

	if len(ids) != 1 {
		fmt.Fprintf(w, "QueryUser only support query one user once")
		return
	}

	id, err := strconv.Atoi(ids[0])
	if err != nil {
		fmt.Fprintf(w, "%+v", err)
		return
	}

	user, err := dao.Imp.QueryUserById(int32(id))
	if err != nil {
		fmt.Fprintf(w, "%+v", err)
		return
	}

	fmt.Fprintf(w, "%+v", *user)
}