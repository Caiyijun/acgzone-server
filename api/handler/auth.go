package handler

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"github.com/132yse/acgzone-server/api/def"
)

func Auth(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	cookie, err := r.Cookie("uqq")
	if err != nil || cookie == nil {
		sendErrorResponse(w, def.ErrorNotAuthUser)
		return
	} else {
		sendErrorResponse(w, def.Success)
	}

}
