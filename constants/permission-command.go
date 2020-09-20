package constants

// Khai báo các permission command
const (
	SuperAdmin = "super_admin" // Quyền cao nhất hệ thống TCNB, có thể thêm bất kỳ nhóm quyền, kho cho nhân viên và không kiểm tra cấp bậc

	// Thao tac voi ma loi - Error Code
	AddErrorCodePermission    = "add_error_code"
	RemoveErrorCodePermission = "remove_error_code"
	ViewErrorCodePermission   = "view_error_code"
	UpdateErrorCodePermission = "update_error_code"
	SearchErrorCodePermission = "search_error_code"

	// Thao tac voi master data
	// Ngan hang - Bank
	AddBankPermission    = "add_bank"
	RemoveBankPermission = "remove_bank"
	ViewBankPermission   = "view_bank"
	UpdateBankPermission = "update_bank"

	// Tinh thanh - Province
	AddProvincePermission    = "add_province"
	RemoveProvincePermission = "remove_province"
	ViewProvincePermission   = "view_province"
	UpdateProvincePermission = "update_province"

	// Quan huyen - District
	AddDistrictPermission    = "add_district"
	RemoveDistrictPermission = "remove_district"
	ViewDistrictPermission   = "view_district"
	UpdateDistrictPermission = "update_district"

	// Phuong xa - Ward
	AddWardPermission    = "add_ward"
	RemoveWardPermission = "remove_ward"
	ViewWardPermission   = "view_ward"
	UpdateWardPermission = "update_ward"

	// Truy cap va quan ly ung dung
	AppAccessPermission = "app_info_access"
)

// PermissionCheckList ..
var PermissionCheckList map[string]bool = map[string]bool{
	SuperAdmin:                false,
	AddErrorCodePermission:    false,
	RemoveErrorCodePermission: false,
	ViewErrorCodePermission:   false,
	UpdateErrorCodePermission: false,
	SearchErrorCodePermission: false,
	AddBankPermission:         false,
	RemoveBankPermission:      false,
	ViewBankPermission:        false,
	UpdateBankPermission:      false,
	AddProvincePermission:     false,
	RemoveProvincePermission:  false,
	ViewProvincePermission:    false,
	UpdateProvincePermission:  false,
	AddDistrictPermission:     false,
	RemoveDistrictPermission:  false,
	ViewDistrictPermission:    false,
	UpdateDistrictPermission:  false,
	AddWardPermission:         false,
	RemoveWardPermission:      false,
	ViewWardPermission:        false,
	UpdateWardPermission:      false,
}
