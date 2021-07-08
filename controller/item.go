package controller

import (
	"encoding/json"
	"github.com/dbielecki97/bookstore-items-api/domain"
	"github.com/dbielecki97/bookstore-items-api/service"
	"github.com/dbielecki97/bookstore-items-api/utils/resp"
	"github.com/dbielecki97/bookstore-oauth-go/oauth"
	"github.com/dbielecki97/bookstore-utils-go/errs"
	"github.com/dbielecki97/bookstore-utils-go/logger"
	"go.uber.org/zap"
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

	var ir domain.Item
	err := json.NewDecoder(r.Body).Decode(&ir)
	if err != nil {
		resp.Error(w, errs.NewBadRequestErr("invalid request body"))
		return
	}
	defer r.Body.Close()

	ir.Seller = oauth.GetCallerId(r)

	item, restErr := service.ItemService.Create(ir)
	if restErr != nil {
		resp.Error(w, restErr)
		return
	}

	logger.Info("", zap.Reflect("item", item))
	resp.JSON(w, http.StatusCreated, item)
}

func (c *defaultItemController) Get(w http.ResponseWriter, r *http.Request) {

}
