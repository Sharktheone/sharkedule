package permissions

type UserPerms struct {
	CreateWorkspaces bool `json:"create_workspaces" yaml:"create_workspaces" bson:"create_workspaces"`
}

type WorkspacePerms struct {
	DeleteWorkspace bool `json:"delete_workspaces" yaml:"delete_workspaces" bson:"delete_workspaces"`
}
