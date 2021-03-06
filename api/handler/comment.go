package handler

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"github.com/132yse/acgzone-server/api/def"
	"encoding/json"
	"github.com/132yse/acgzone-server/api/db"
	"log"
	"strconv"
)

func AddComment(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	req, _ := ioutil.ReadAll(r.Body)
	cbody := &def.Comment{}


	if err := json.Unmarshal(req, cbody); err != nil {
		sendErrorResponse(w, def.ErrorRequestBodyParseFailed)
		return
	}

	if resp, err := db.AddComment(cbody.Content, cbody.Pid, cbody.Uid); err != nil {
		log.Printf("%s", err)
		sendErrorResponse(w, def.ErrorDB)
		return
	} else {
		res := def.Comment{Id: resp.Id, Content: resp.Content, Time: resp.Time, Pid: resp.Pid, Uid: resp.Uid}
		sendCommentResponse(w, res, 201)
	}

}

func GetComments(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	pid, _ := strconv.Atoi(r.URL.Query().Get("pid"))
	uid, _ := strconv.Atoi(r.URL.Query().Get("uid"))
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	pageSize, _ := strconv.Atoi(r.URL.Query().Get("pageSize"))
	resp, err := db.GetComments(pid,uid, page, pageSize)
	if err != nil {
		sendErrorResponse(w, def.ErrorDB)
		log.Printf("%s", err)
		return
	} else {
		res := &def.Comments{Posts: resp}
		sendCommentsResponse(w, res, 201)
	}
}

func DeleteComment(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	pid, _ := strconv.Atoi(p.ByName("id"))
	err := db.DeleteComment(pid)
	if err != nil {
		sendErrorResponse(w, def.ErrorDB)
		return
	} else {
		sendErrorResponse(w, def.Success)
	}
}
