package kanbandb

import (
	"fmt"
	"github.com/Sharktheone/sharkedule/types"
)

func GetUser(users []types.User, uuid string) (types.User, error) {
	for _, u := range users {
		if u.GetUUID() == uuid {
			return u, nil
		}
	}

	return nil, fmt.Errorf("user with uuid %s does not exist", uuid)
}

func GetUserByMail(users []types.User, mail string) (types.User, error) {
	for _, u := range users {
		if u.GetEmail() == mail {
			return u, nil
		}
	}

	return nil, fmt.Errorf("user with mail %s does not exist", mail)
}

func UpdateUserUsername(users []types.User, uuid string, username string) error {
	for _, u := range users {
		if u.GetUUID() == uuid {
			u.SetUsername(username)
			return nil
		}
	}

	return fmt.Errorf("user with uuid %s does not exist", uuid)
}

func UpdateUserEmail(users []types.User, uuid string, email string) error {
	for _, u := range users {
		if u.GetUUID() == uuid {
			u.SetEmail(email)
			return nil
		}
	}

	return fmt.Errorf("user with uuid %s does not exist", uuid)
}

func UpdateUserPassword(users []types.User, uuid string, password string) error {
	for _, u := range users {
		if u.GetUUID() == uuid {
			u.SetPassword(password)
			return nil
		}
	}

	return fmt.Errorf("user with uuid %s does not exist", uuid)
}

func AddUserWorkspaceAccess(users []types.User, uuid string, workspace string) error {
	for _, u := range users {
		if u.GetUUID() == uuid {
			return u.GetAccess().AddWorkspaceAccess(workspace)
		}
	}

	return fmt.Errorf("user with uuid %s does not exist", uuid)
}

func RemoveUserWorkspaceAccess(users []types.User, user string, workspace string) error {
	for _, u := range users {
		if u.GetUUID() == user {
			return u.GetAccess().RemoveWorkspaceAccess(workspace)
		}
	}

	return fmt.Errorf("user with uuid %s does not exist", workspace)
}

func UpdateUserSettings(users []types.User, uuid string, settings types.Settings) error {
	for _, u := range users {
		if u.GetUUID() == uuid {
			u.SetSettings(settings)
			return nil
		}
	}

	return fmt.Errorf("user with uuid %s does not exist", uuid)
}

//func EnableUser2FA(users []*types.User, uuid string) error { //TODO: this needs some more parameters and return types (e.g. 2FA Token, recovery codes, ...)
//	for _, u := range users {
//		if u.GetUUID() == uuid {
//			u.MFA.TFA = true
//			return nil
//		}
//	}
//
//	return fmt.Errorf("user with uuid %s does not exist", uuid)
//}

//TODO: other methods for 2FA and MFA, OAUTH...

func DeleteUser(users []types.User, uuid string) error {
	for i, u := range users {
		if u.GetUUID() == uuid {
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("user with uuid %s does not exist", uuid)
}
