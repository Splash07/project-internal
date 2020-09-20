package masterdata

import (
	"github.com/labstack/echo/v4"
	insideAPI "gitlab.ghn.vn/online/common/apicall/v5/inside"
)

type warehouseHandler struct{}

// WarehouseHandler ..
var WarehouseHandler warehouseHandler

// Warehouses Danh sách Warehouses ..
func (warehouseHandler) Warehouses(c echo.Context) (err error) {

	insideAPICall := insideAPI.New(getRequestID(c), cfg.API)
	warehouseList := []insideAPI.Warehouse{}

	offset := 0
	limit := 100
	for {
		warehouses, _, errCall := insideAPICall.GetAllHub(offset, limit)
		if errCall != nil {
			return err
		}
		warehouseList = append(warehouseList, warehouses...)
		if len(warehouses) < limit {
			break
		}
		offset += limit
	}

	return c.JSON(success(warehouseList))
}

// Chi tiết Warehouse ..
func (warehouseHandler) WarehouseDetail(c echo.Context) (err error) {

	insideAPICall := insideAPI.New(getRequestID(c), cfg.API)

	type myRequest struct {
		WarehouseID int `json:"warehouse_id" query:"warehouse_id" validate:"required"`
	}

	request := new(myRequest)
	if err = c.Bind(request); err != nil {
		return
	}
	if err = c.Validate(request); err != nil {
		return
	}

	waehouse, err := insideAPICall.GetWarehouseDetail(request.WarehouseID)

	if err != nil {
		return
	}

	return (c.JSON(success(waehouse)))
}
