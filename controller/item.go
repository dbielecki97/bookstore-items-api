package controller

import (
	"encoding/json"
	"github.com/dbielecki97/bookstore-items-api/domain/item"
	"github.com/dbielecki97/bookstore-items-api/service"
	"github.com/dbielecki97/bookstore-items-api/utils/resp"
	"github.com/dbielecki97/bookstore-oauth-go/oauth"
	"github.com/dbielecki97/bookstore-utils-go/errs"
	"net/http"
)

var (
	ItemController itemController = &defaultItemController{}
)

type itemController interface {
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
}

type defaultItemController struct {
}

func (c *defaultItemController) Create(w http.ResponseWriter, r *http.Request) {
	if restErr := oauth.AuthenticateRequest(r); restErr != nil {
		resp.Error(w, restErr)
		return
	}

	if oauth.GetCallerId(r) == 0 {
		resp.Error(w, errs.NewAuthenticationErr("unable to retrieve caller id from token"))
		return
	}

	var ir item.Dto
	err := json.NewDecoder(r.Body).Decode(&ir)
	if err != nil {
		resp.Error(w, errs.NewBadRequestErr("invalid request body"))
		return
	}
	defer r.Body.Close()

	ir.Seller = oauth.GetCallerId(r)

	ic, restErr := service.ItemService.Create(ir)
	if restErr != nil {
		resp.Error(w, restErr)
		return
	}

	resp.JSON(w, http.StatusCreated, ic)
}

func (c *defaultItemController) Get(w http.ResponseWriter, r *http.Request) {

}
