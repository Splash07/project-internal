package errcode

const (
	EmployeeNotHavePermission              = "EMPLOYEE_NOT_HAVE_PERMISSION"
	GroupPermissionNotFound                = "ERROR_PERMISSION_GROUP_PERMISSION_NOT_FOUND"
	PermissionNotFoundInGroup              = "ERROR_PERMISSION_NOT_FOUND_IN_GROUP"
	UserNotFound                           = "USER_NOT_FOUND"
	UserNotHavePermission                  = "USER_NOT_HAVE_PERMISSION"
	CreateGroupPermissionFail              = "CREATE_GROUP_PERMISSION_FAIL"
	DeleteGroupPermissionFail              = "DELETE_GROUP_PERMISSION_FAIL"
	UpdateGroupPermissionFail              = "UPDATE_GROUP_PERMISSION_FAIL"
	AddPermissionsToGroupPermissionFail    = "ADD_PERMISSIONS_TO_GROUP_PERMISSION_FAIL"
	RemovePermissionsToGroupPermissionFail = "REMOVE_PERMISSIONS_TO_GROUP_PERMISSION_FAIL"
	AddGroupToUserFail                     = "ADD_GROUP_TO_USER_FAIL"
	AddExtraPermissionToUserFail           = "ADD_EXTRA_PERMISSION_TO_USER_FAIL"
	RemoveExtraPermissionInUserFail        = "REMOVE_EXTRA_PERMISSION_IN_USER_FAIL"
)
