package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/haproxytech/client-native"

	"github.com/haproxytech/dataplaneapi/haproxy"
	"github.com/haproxytech/dataplaneapi/misc"
	"github.com/haproxytech/dataplaneapi/operations/configuration"
)

//GetRawConfigurationHandlerImpl implementation of the GetHAProxyConfigurationHandler interface
type GetRawConfigurationHandlerImpl struct {
	Client *client_native.HAProxyClient
}

// PostRawConfigurationHandlerImpl implementation of the PostHAProxyConfigurationHandler interface
type PostRawConfigurationHandlerImpl struct {
	Client      *client_native.HAProxyClient
	ReloadAgent *haproxy.ReloadAgent
}

//Handle executing the request and returning a response
func (h *GetRawConfigurationHandlerImpl) Handle(params configuration.GetHAProxyConfigurationParams, principal interface{}) middleware.Responder {
	data, err := h.Client.Configuration.GetRawConfiguration()
	if err != nil {
		e := misc.HandleError(err)
		return configuration.NewGetHAProxyConfigurationDefault(int(*e.Code)).WithPayload(e)
	}
	return configuration.NewGetHAProxyConfigurationOK().WithPayload(data)
}

//Handle executing the request and returning a response
func (h *PostRawConfigurationHandlerImpl) Handle(params configuration.PostHAProxyConfigurationParams, principal interface{}) middleware.Responder {
	v := int64(0)
	if params.Version != nil {
		v = *params.Version
	}

	err := h.Client.Configuration.PostRawConfiguration(params.Data, v)
	if err != nil {
		e := misc.HandleError(err)
		return configuration.NewPostHAProxyConfigurationDefault(int(*e.Code)).WithPayload(e)
	}
	if *params.ForceReload {
		err := h.ReloadAgent.ForceReload()
		if err != nil {
			e := misc.HandleError(err)
			return configuration.NewPostHAProxyConfigurationDefault(int(*e.Code)).WithPayload(e)
		}
		return configuration.NewPostHAProxyConfigurationCreated().WithPayload(*params.Data)
	}
	rID := h.ReloadAgent.Reload()
	return configuration.NewPostHAProxyConfigurationAccepted().WithReloadID(rID).WithPayload(*params.Data)
}
