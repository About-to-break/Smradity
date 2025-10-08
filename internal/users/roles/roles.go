package roles

type Role string

const (
	RoleDefault   Role = "default"
	RoleModerator Role = "moderator"
	RoleUSer      Role = "user"
	RoleAdmin     Role = "admin"
)

var AllRoles = []Role{
	RoleDefault,
	RoleModerator,
	RoleUSer,
	RoleAdmin,
}

func IsValidRole(role Role) bool {
	for _, r := range AllRoles {
		if r == role {
			return true
		}
	}
	return false
}
