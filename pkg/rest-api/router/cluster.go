package router

import (
	"net/http"

	"github.com/cloud-barista/poc-cb-ladybug/pkg/core/common"
	"github.com/cloud-barista/poc-cb-ladybug/pkg/core/model"
	"github.com/cloud-barista/poc-cb-ladybug/pkg/core/service"
	"github.com/labstack/echo/v4"
)

func ListCluster(c echo.Context) error {
	common.CBLog.Debugf("[CALLED]")

	if err := Validate(c, []string{"namespace"}); err != nil {
		common.CBLog.Error(err)
		return SendMessage(c, http.StatusBadRequest, err.Error())
	}

	clusterList, err := service.ListCluster(c.Param("namespace"))
	if err != nil {
		common.CBLog.Error(err)
		return SendMessage(c, http.StatusBadRequest, err.Error())
	}

	return Send(c, http.StatusOK, clusterList)
}

func GetCluster(c echo.Context) error {
	common.CBLog.Debugf("[CALLED]")

	if err := Validate(c, []string{"namespace", "cluster"}); err != nil {
		common.CBLog.Error(err)
		return SendMessage(c, http.StatusBadRequest, err.Error())
	}

	cluster, err := service.GetCluster(c.Param("namespace"), c.Param("cluster"))
	if err != nil {
		common.CBLog.Error(err)
		return SendMessage(c, http.StatusBadRequest, err.Error())
	}

	return Send(c, http.StatusOK, cluster)
}

func CreateCluster(c echo.Context) error {
	common.CBLog.Debugf("[CALLED]")

	clusterReq := &model.ClusterReq{}
	if err := c.Bind(clusterReq); err != nil {
		common.CBLog.Error(err)
		return SendMessage(c, http.StatusBadRequest, err.Error())
	}

	model.ClusterReqConfKubeDef(clusterReq)

	err := model.ClusterReqValidate(clusterReq)
	if err != nil {
		common.CBLog.Error(err)
		return SendMessage(c, http.StatusBadRequest, err.Error())
	}

	cluster, err := service.CreateCluster(c.Param("namespace"), clusterReq)
	if err != nil {
		common.CBLog.Error(err)
		return SendMessage(c, http.StatusInternalServerError, err.Error())
	}

	return Send(c, http.StatusOK, cluster)
}

func DeleteCluster(c echo.Context) error {
	common.CBLog.Debugf("[CALLED]")

	if err := Validate(c, []string{"namespace", "cluster"}); err != nil {
		common.CBLog.Error(err)
		return SendMessage(c, http.StatusBadRequest, err.Error())
	}

	status, err := service.DeleteCluster(c.Param("namespace"), c.Param("cluster"))
	if err != nil {
		common.CBLog.Error(err)
		return SendMessage(c, http.StatusInternalServerError, err.Error())
	}

	return Send(c, http.StatusOK, status)
}
