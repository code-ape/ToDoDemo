package main

import (
	log "github.com/Sirupsen/logrus"
)

type ToDo struct {
	ID   string
	Text string
	User string
}

type AddToDoReq struct {
	User  string              `json:"user" binding:"required"`
	Token string              `json:"token" binding:"required"`
	ToDos []map[string]string `json:"to_dos" binding:"required"`
}

type GetToDosReq struct {
	User  string `json:"user" binding:"required"`
	Token string `json:"token" binding:"required"`
}

type DeleteToDoReq struct {
	User  string   `json:"user" binding:"required"`
	Token string   `json:"token" binding:"required"`
	IDs   []string `json:"ids" binding:"required"`
}

var to_do_list map[string]*ToDo

func init() {
	to_do_list = make(map[string]*ToDo)
}

func AddToDos(tds []map[string]string) bool {
	success := true
	for _, tdo := range tds {
		user, ue := tdo["user"]
		text, te := tdo["text"]
		if ue && te {
			_ = AddToDo(user, text)
		} else {
			log.Error("A todo isn't formatted correctly.")
			success = false
		}
	}
	return success
}

func AddToDo(user string, text string) bool {
	token := RandomHexToken()
	td := &ToDo{ID: token, Text: text, User: user}
	to_do_list[token] = td
	log.WithField("User", td.User).WithField("ToDo", td.Text).Info("To-do added.")
	return true
}

func GetToDos() []map[string]string {
	data := make([]map[string]string, len(to_do_list))

	i := 0
	for _, td := range to_do_list {
		tdo := map[string]string{"user": td.User, "text": td.Text, "id": td.ID}
		data[i] = tdo
		i++
	}
	log.Infof("Serving ToDos: %+v", data)
	return data
}

func DeleteToDos(ids []string) bool {
	for _, id := range ids {
		_ = DeleteToDo(id)
	}
	return true
}

func DeleteToDo(id string) bool {
	td, exists := to_do_list[id]
	if !exists {
		log.WithField("ID", id).Error("ToDo with ID doesn't exist.")
		return false
	}
	delete(to_do_list, id)
	log.WithField("ToDo", td.Text).Info("Deleted ToDo.")
	return true
}
