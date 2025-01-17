package enums

type UserRoles string

const (
	UserRoleAdmin    UserRoles = "admin"
	UserRoleGuest    UserRoles = "user_guest"
	UserRoleBusiness UserRoles = "user_business"
)

type UserStatus string

const (
	UserStatusPendingActivation UserStatus = "pending_activation"
	UserStatusEnabled           UserStatus = "enabled"
	UserStatusDisabled          UserStatus = "disabled"
)
