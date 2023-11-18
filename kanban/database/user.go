package kanbandb

import (
	"fmt"
	"github.com/Sharktheone/sharkedule/kanban/types"
	"github.com/Sharktheone/sharkedule/user/access/workspaceaccess"
	"github.com/Sharktheone/sharkedule/user/settings"
)

func GetUser(users []*types.User, uuid string) (*types.User, error) {
	for _, u := range users {
		if u.UUID == uuid {
			return u, nil
		}
	}

	return nil, fmt.Errorf("user with uuid %s does not exist", uuid)
}

func GetUserByMail(users []*types.User, mail string) (*types.User, error) {
	for _, u := range users {
		if u.Email == mail {
			return u, nil
		}
	}

	return nil, fmt.Errorf("user with mail %s does not exist", mail)
}

func UpdateUserUsername(users []*types.User, uuid string, username string) error {
	for _, u := range users {
		if u.UUID == uuid {
			u.Username = username
			return nil
		}
	}

	return fmt.Errorf("user with uuid %s does not exist", uuid)
}

func UpdateUserEmail(users []*types.User, uuid string, email string) error {
	for _, u := range users {
		if u.UUID == uuid {
			u.Email = email
			return nil
		}
	}

	return fmt.Errorf("user with uuid %s does not exist", uuid)
}

func UpdateUserPassword(users []*types.User, uuid string, password string) error {
	for _, u := range users {
		if u.UUID == uuid {
			u.Password = password
			return nil
		}
	}

	return fmt.Errorf("user with uuid %s does not exist", uuid)
}

func AddUserWorkspaceAccess(users []*types.User, uuid string, workspace string) error {
	for _, u := range users {
		if u.UUID == uuid {
			u.Access.Workspaces = append(u.Access.Workspaces, workspaceaccess.WorkspaceAccess{
				UUID: workspace,
			})
			return nil
		}
	}

	return fmt.Errorf("user with uuid %s does not exist", uuid)
}

func RemoveUserWorkspaceAccess(users []*types.User, uuid string, access string) error {
	for _, u := range users {
		if u.UUID == uuid {
			for i, a := range u.Access.Workspaces {
				if a.UUID == access {
					u.Access.Workspaces = append(u.Access.Workspaces[:i], u.Access.Workspaces[i+1:]...)
					return nil
				}
			}
		}
	}

	return fmt.Errorf("user with uuid %s does not exist", uuid)
}

func UpdateUserSettings(users []*types.User, uuid string, settings settings.Settings) error {
	for _, u := range users {
		if u.UUID == uuid {
			u.Settings = settings
			return nil
		}
	}

	return fmt.Errorf("user with uuid %s does not exist", uuid)
}

//func EnableUser2FA(users []*types.User, uuid string) error { //TODO: this needs some more parameters and return types (e.g. 2FA Token, recovery codes, ...)
//	for _, u := range users {
//		if u.UUID == uuid {
//			u.MFA.TFA = true
//			return nil
//		}
//	}
//
//	return fmt.Errorf("user with uuid %s does not exist", uuid)
//}

//TODO: other methods for 2FA and MFA, OAUTH...

func DeleteUser(users []*types.User, uuid string) error {
	for i, u := range users {
		if u.UUID == uuid {
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("user with uuid %s does not exist", uuid)
}
