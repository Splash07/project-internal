package route

import (
	"github.com/labstack/echo/v4"
	"gitlab.com/Splash07/project-internal/web/handler/masterdata"
)

func setUpMasterDataRoutes(g *echo.Group) {
	//////////////////// master-data ////////////////////
	//////banks
	g.Any("/master-data/bank", masterdata.BankHandler.GetOne)
	g.Any("/master-data/banks", masterdata.BankHandler.GetAll)
	g.Any("/master-data/bank/add", masterdata.BankHandler.AddNewBank)
	g.Any("/master-data/bank/remove", masterdata.BankHandler.RemoveBank)
	g.Any("/master-data/bank/update", masterdata.BankHandler.UpdateBank)
	g.Any("/master-data/bank/switch-status", masterdata.BankHandler.SwitchStatus)

	//////provinces
	g.Any("/v2/master-data/province", masterdata.ProvinceHandler.GetByIDV2)
	g.Any("/v2/master-data/provinces", masterdata.ProvinceHandler.GetAllV2)
	g.Any("/v2/master-data/province/remove", masterdata.ProvinceHandler.RemoveV2)
	g.Any("/v2/master-data/province/switch-status", masterdata.ProvinceHandler.SwitchStatus)
	g.Any("/v2/master-data/province/add", masterdata.ProvinceHandler.AddV2)
	g.Any("/v2/master-data/province/update", masterdata.ProvinceHandler.UpdateV2)

	//////districts
	g.Any("/master-data/districts", masterdata.DistrictHandler.Districts)            // might be removed ..
	g.Any("/master-data/district/detail", masterdata.DistrictHandler.DistrictDetail) // might be removed ..
	g.Any("/v2/master-data/district", masterdata.DistrictHandler.GetByIDV2)
	g.Any("/v2/master-data/districts", masterdata.DistrictHandler.DistrictsByProvinceIDV2)
	g.Any("/v2/master-data/district/remove", masterdata.DistrictHandler.RemoveV2)
	g.Any("/v2/master-data/district/switch-status", masterdata.DistrictHandler.SwitchStatus)
	g.Any("/v2/master-data/district/add", masterdata.DistrictHandler.AddV2)
	g.Any("/v2/master-data/district/update", masterdata.DistrictHandler.UpdateV2)

	g.Any("/v2/master-data/ward", masterdata.WardHandler.GetByCodeV2)
	g.Any("/v2/master-data/wards", masterdata.WardHandler.WardsByDistrictIDV2)
	g.Any("/v2/master-data/ward/remove", masterdata.WardHandler.RemoveV2)
	g.Any("/v2/master-data/ward/switch-status", masterdata.WardHandler.SwitchStatus)
	g.Any("/v2/master-data/ward/add", masterdata.WardHandler.AddV2)
	g.Any("/v2/master-data/ward/update", masterdata.WardHandler.UpdateV2)
}
