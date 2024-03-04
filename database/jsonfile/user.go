package jsonfile

import (
	kanbandb "github.com/Sharktheone/sharkedule/kanban/database"
	"github.com/Sharktheone/sharkedule/types"
)

func (J *JSONFile) GetUser(uuid string) (types.User, error) {
	return kanbandb.GetUser(J.db.Users, uuid)
}

func (J *JSONFile) GetUserByMail(mail string) (types.User, error) {
	return kanbandb.GetUserByMail(J.db.Users, mail)
}

func (J *JSONFile) UpdateUserUsername(uuid string, username string) error {
	return kanbandb.UpdateUserUsername(J.db.Users, uuid, username)
}

func (J *JSONFile) UpdateUserEmail(uuid string, email string) error {
	return kanbandb.UpdateUserEmail(J.db.Users, uuid, email)
}

func (J *JSONFile) UpdateUserPassword(uuid string, password string) error {
	return kanbandb.UpdateUserPassword(J.db.Users, uuid, password)
}

func (J *JSONFile) AddUserWorkspaceAccess(uuid, workspace string) error {
	return kanbandb.AddUserWorkspaceAccess(J.db.Users, uuid, workspace)
}

func (J *JSONFile) RemoveUserWorkspaceAccess(user, workspace string) error {
	return kanbandb.RemoveUserWorkspaceAccess(J.db.Users, user, workspace)
}

func (J *JSONFile) UpdateUserSettings(uuid string, settings types.Settings) error {
	return kanbandb.UpdateUserSettings(J.db.Users, uuid, settings)
}

func (J *JSONFile) DeleteUser(uuid string) error {
	return kanbandb.DeleteUser(J.db.Users, uuid)
}
