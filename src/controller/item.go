package controller

import (
	"encoding/json"
	item2 "github.com/dbielecki97/bookstore-items-api/src/domain/item"
	query2 "github.com/dbielecki97/bookstore-items-api/src/domain/query"
	service2 "github.com/dbielecki97/bookstore-items-api/src/service"
	resp2 "github.com/dbielecki97/bookstore-items-api/src/utils/resp"
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
		resp2.Error(w, restErr)
		return
	}

	if oauth.GetCallerId(r) == 0 {
		resp2.Error(w, errs.NewAuthenticationErr("unable to retrieve caller id from token"))
		return
	}

	var ir item2.Dto
	err := json.NewDecoder(r.Body).Decode(&ir)
	if err != nil {
		resp2.Error(w, errs.NewBadRequestErr("invalid request body"))
		return
	}
	defer r.Body.Close()

	ir.Seller = oauth.GetCallerId(r)

	ic, restErr := service2.ItemService.Create(ir)
	if restErr != nil {
		resp2.Error(w, restErr)
		return
	}

	resp2.JSON(w, http.StatusCreated, ic)
}

func (c *defaultItemController) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	itemId := vars["item_id"]
	it, err := service2.ItemService.Get(itemId)
	if err != nil {
		resp2.Error(w, err)
		return
	}

	resp2.JSON(w, http.StatusOK, it)
}

func (c *defaultItemController) Search(w http.ResponseWriter, r *http.Request) {
	var q query2.EsQuery
	err := json.NewDecoder(r.Body).Decode(&q)
	if err != nil {
		resp2.Error(w, errs.NewBadRequestErr("invalid json body"))
		return
	}
	defer r.Body.Close()

	results, restErr := service2.ItemService.Search(q)
	if restErr != nil {
		resp2.Error(w, restErr)
		return
	}

	resp2.JSON(w, http.StatusOK, results)
}

func (c *defaultItemController) Update(w http.ResponseWriter, r *http.Request) {
	if restErr := oauth.AuthenticateRequest(r); restErr != nil {
		resp2.Error(w, restErr)
		return
	}

	if oauth.GetCallerId(r) == 0 {
		resp2.Error(w, errs.NewAuthenticationErr("unable to retrieve caller id from token"))
		return
	}

	vars := mux.Vars(r)
	itemId := vars["item_id"]

	var ir item2.Dto
	err := json.NewDecoder(r.Body).Decode(&ir)
	if err != nil {
		resp2.Error(w, errs.NewBadRequestErr("invalid request body"))
		return
	}
	defer r.Body.Close()

	ir.ID = itemId

	ui, restErr := service2.ItemService.Update(ir)
	if restErr != nil {
		resp2.Error(w, restErr)
		return
	}

	resp2.JSON(w, http.StatusOK, ui)
}

func (c defaultItemController) Delete(w http.ResponseWriter, r *http.Request) {
	if restErr := oauth.AuthenticateRequest(r); restErr != nil {
		resp2.Error(w, restErr)
		return
	}

	if oauth.GetCallerId(r) == 0 {
		resp2.Error(w, errs.NewAuthenticationErr("unable to retrieve caller id from token"))
		return
	}

	vars := mux.Vars(r)
	itemId := vars["item_id"]
	err := service2.ItemService.Delete(itemId)
	if err != nil {
		resp2.Error(w, err)
		return
	}

	resp2.JSON(w, http.StatusNoContent, nil)
}
