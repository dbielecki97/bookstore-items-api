package controller

import (
	"encoding/json"
	"github.com/dbielecki97/bookstore-items-api/domain/item"
	"github.com/dbielecki97/bookstore-items-api/domain/query"
	"github.com/dbielecki97/bookstore-items-api/service"
	"github.com/dbielecki97/bookstore-items-api/utils/resp"
	"github.com/dbielecki97/bookstore-oauth-go/oauth"
	"github.com/dbielecki97/bookstore-utils-go/errs"
	"github.com/gorilla/mux"
	"net/http"
)

var (
	ItemController itemController = &defaultItemController{}
)

type itemController interface {
	Create(http.ResponseWriter, *http.Request)
	Get(http.ResponseWriter, *http.Request)
	Search(http.ResponseWriter, *http.Request)
	Update(http.ResponseWriter, *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
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
	vars := mux.Vars(r)
	itemId := vars["item_id"]
	it, err := service.ItemService.Get(itemId)
	if err != nil {
		resp.Error(w, err)
		return
	}

	resp.JSON(w, http.StatusOK, it)
}

func (c *defaultItemController) Search(w http.ResponseWriter, r *http.Request) {
	var q query.EsQuery
	err := json.NewDecoder(r.Body).Decode(&q)
	if err != nil {
		resp.Error(w, errs.NewBadRequestErr("invalid json body"))
		return
	}
	defer r.Body.Close()

	results, restErr := service.ItemService.Search(q)
	if restErr != nil {
		resp.Error(w, restErr)
		return
	}

	resp.JSON(w, http.StatusOK, results)
}

func (c *defaultItemController) Update(w http.ResponseWriter, r *http.Request) {
	if restErr := oauth.AuthenticateRequest(r); restErr != nil {
		resp.Error(w, restErr)
		return
	}

	if oauth.GetCallerId(r) == 0 {
		resp.Error(w, errs.NewAuthenticationErr("unable to retrieve caller id from token"))
		return
	}

	vars := mux.Vars(r)
	itemId := vars["item_id"]

	var ir item.Dto
	err := json.NewDecoder(r.Body).Decode(&ir)
	if err != nil {
		resp.Error(w, errs.NewBadRequestErr("invalid request body"))
		return
	}
	defer r.Body.Close()

	ir.ID = itemId

	ui, restErr := service.ItemService.Update(ir)
	if restErr != nil {
		resp.Error(w, restErr)
		return
	}

	resp.JSON(w, http.StatusOK, ui)
}

func (c defaultItemController) Delete(w http.ResponseWriter, r *http.Request) {
	if restErr := oauth.AuthenticateRequest(r); restErr != nil {
		resp.Error(w, restErr)
		return
	}

	if oauth.GetCallerId(r) == 0 {
		resp.Error(w, errs.NewAuthenticationErr("unable to retrieve caller id from token"))
		return
	}

	vars := mux.Vars(r)
	itemId := vars["item_id"]
	err := service.ItemService.Delete(itemId)
	if err != nil {
		resp.Error(w, err)
		return
	}

	resp.JSON(w, http.StatusNoContent, nil)
}
