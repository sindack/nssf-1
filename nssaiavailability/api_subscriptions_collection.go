// Copyright 2019 free5GC.org
//
// SPDX-License-Identifier: Apache-2.0
//

/*
 * NSSF NSSAI Availability
 *
 * NSSF NSSAI Availability Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package nssaiavailability

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/free5gc/http_wrapper"
	"github.com/free5gc/nssf/logger"
	"github.com/free5gc/nssf/producer"
	"github.com/free5gc/openapi"
	. "github.com/free5gc/openapi/models"
)

func HTTPNSSAIAvailabilityPost(c *gin.Context) {
	var createData NssfEventSubscriptionCreateData

	requestBody, err := c.GetRawData()
	if err != nil {
		problemDetail := ProblemDetails{
			Title:  "System failure",
			Status: http.StatusInternalServerError,
			Detail: err.Error(),
			Cause:  "SYSTEM_FAILURE",
		}
		logger.HandlerLog.Errorf("Get Request Body error: %+v", err)
		c.JSON(http.StatusInternalServerError, problemDetail)
		return
	}

	err = openapi.Deserialize(&createData, requestBody, "application/json")
	if err != nil {
		problemDetail := "[Request Body] " + err.Error()
		rsp := ProblemDetails{
			Title:  "Malformed request syntax",
			Status: http.StatusBadRequest,
			Detail: problemDetail,
		}
		logger.HandlerLog.Errorln(problemDetail)
		c.JSON(http.StatusBadRequest, rsp)
		return
	}

	req := http_wrapper.NewRequest(c.Request, createData)

	rsp := producer.HandleNSSAIAvailabilityPost(req)

	// TODO: Based on TS 29.531 5.3.2.3.1, add location header

	responseBody, err := openapi.Serialize(rsp.Body, "application/json")
	if err != nil {
		logger.HandlerLog.Errorln(err)
		problemDetails := ProblemDetails{
			Status: http.StatusInternalServerError,
			Cause:  "SYSTEM_FAILURE",
			Detail: err.Error(),
		}
		c.JSON(http.StatusInternalServerError, problemDetails)
	} else {
		c.Data(rsp.Status, "application/json", responseBody)
	}
}
