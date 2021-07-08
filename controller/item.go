package controller

import (
	"github.com/dbielecki97/bookstore-items-api/domain"
	"github.com/dbielecki97/bookstore-items-api/service"
	"github.com/dbielecki97/bookstore-oauth-go/oauth"
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
	if err := oauth.AuthenticateRequest(r); err != nil {
		//TODO: Return error to the caller
		return
	}

	i := domain.Item{Seller: oauth.GetCallerId(r)}

	item, err := service.ItemService.Create(i)
	if err != nil {
		//TODO: Return error json to the user
	}

	logger.Info("", zap.Reflect("item", item))
	//TODO: Return created ite as json with HTTP status 201 - Created
}

func (c *defaultItemController) Get(w http.ResponseWriter, r *http.Request) {

}
