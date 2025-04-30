package main

import (
	"context"

	"github.com/cloudfresco/sc-dcsa/internal/common"
	commonproto "github.com/cloudfresco/sc-dcsa/internal/protogen/common/v1"
	"github.com/spf13/viper"
)

// RoleOptions - for Role
type RoleOptions struct {
	Roles []Role `mapstructure:"roles"`
}

type Role struct {
	RoleName                 string   `mapstructure:"role_name"`                  //- used for create role
	RoleDescription          string   `mapstructure:"role_description"`           //- used for role description
	ResourceServerIdentifier string   `mapstructure:"resource_server_identifier"` //- used for AddPermisionsToRoles
	PermissionName           string   `mapstructure:"permission_name"`            // - used for AddPermisionsToRoles
	PermissionDescription    string   `mapstructure:"permission_description"`     // - used for AddPermisionsToRoles
	UserIds                  []string `mapstructure:"user_ids"`                   //  - used to assign roles
}

func main() {
	v := viper.New()
	v.AutomaticEnv()

	v.SetConfigName("roles")
	v.SetConfigType("json")
	v.AddConfigPath("./test/roles")
	if err := v.ReadInConfig(); err != nil {
		return
	}

	rolesData := RoleOptions{}
	if err := v.Unmarshal(&rolesData); err != nil {
		return
	}

	mgmtToken := v.GetString("SC_DCSA_AUTH0_MGMTTOKEN")

	API_ID := v.GetString("SC_DCSA_AUTH0_API_ID")

	domain := v.GetString("SC_DCSA_AUTH0_DOMAIN")

	userEmail := v.GetString("SC_DCSA_EMAIL_TEST")

	requestId := v.GetString("SC_DCSA_REQUESTID_TEST")

	ctx := context.Background()

	// get roles
	getRoleReq := commonproto.GetRoles{}
	getRoleReq.Auth0Domain = domain
	getRoleReq.Auth0MgmtToken = mgmtToken
	getRoleReq.UserEmail = userEmail
	getRoleReq.RequestId = requestId

	roleResponse, err := common.GetRolesResp(ctx, &getRoleReq)
	if err != nil {
		return
	}

	// remove permissions from role and its roles
	if len(roleResponse) != 0 {
		for _, role := range roleResponse {
			// get role permissions
			getRolePermissionReq := commonproto.GetRolePermissions{}
			getRolePermissionReq.RoleId = role.Id
			getRolePermissionReq.Auth0Domain = domain
			getRolePermissionReq.Auth0MgmtToken = mgmtToken
			getRolePermissionReq.UserEmail = userEmail
			getRolePermissionReq.RequestId = requestId
			rolePermissionResponse, err := common.GetRolePermissionsResp(ctx, &getRolePermissionReq)
			if err != nil {
				return
			}

			for _, rolePermission := range rolePermissionResponse {
				// remove Role Permission
				removeRolePermissionReq := commonproto.RemoveRolePermission{}
				removeRolePermissionReq.ResourceServerIdentifier = rolePermission.ResourceServerIdentifier
				removeRolePermissionReq.PermissionName = rolePermission.PermissionName
				removeRolePermissionReq.RoleId = role.Id
				removeRolePermissionReq.Auth0Domain = domain
				removeRolePermissionReq.Auth0MgmtToken = mgmtToken
				removeRolePermissionReq.UserEmail = userEmail
				removeRolePermissionReq.RequestId = requestId
				err = common.RemoveRolePermissionResp(ctx, &removeRolePermissionReq)
				if err != nil {
					return
				}
			}
			// Delete Role - delete role
			deleteRoleReq := commonproto.DeleteRole{}
			deleteRoleReq.RoleId = role.Id
			deleteRoleReq.Auth0Domain = domain
			deleteRoleReq.Auth0MgmtToken = mgmtToken
			deleteRoleReq.UserEmail = userEmail
			deleteRoleReq.RequestId = requestId
			err = common.DeleteRoleResp(ctx, &deleteRoleReq)
			if err != nil {
				return
			}
		}
	}

	// add all permissions
	permissions := []*commonproto.Permission{}
	for _, rl := range rolesData.Roles {
		p := commonproto.Permission{}
		p.PermissionName = rl.PermissionName
		p.PermissionDescription = rl.PermissionDescription
		permissions = append(permissions, &p)
	}

	apiPermissionReq := commonproto.AddAPIPermission{}
	apiPermissionReq.Permissions = permissions
	apiPermissionReq.Auth0ApiId = API_ID
	apiPermissionReq.Auth0Domain = domain
	apiPermissionReq.Auth0MgmtToken = mgmtToken
	apiPermissionReq.UserEmail = userEmail
	apiPermissionReq.RequestId = requestId

	err = common.AddAPIPermissionResp(ctx, &apiPermissionReq)
	if err != nil {
		return
	}

	// create role with permissions and assign to users
	for _, rl := range rolesData.Roles {
		createRoleReq := commonproto.CreateRole{}
		createRoleReq.Name = rl.RoleName
		createRoleReq.Description = rl.RoleDescription
		createRoleReq.Auth0Domain = domain
		createRoleReq.Auth0MgmtToken = mgmtToken
		createRoleReq.UserEmail = userEmail
		createRoleReq.RequestId = requestId
		role, err := common.CreateRoleResp(ctx, &createRoleReq)
		if err != nil {
			return
		}
		addPermisionsToRolesReq := commonproto.AddPermisionsToRoles{}
		addPermisionsToRolesReq.ResourceServerIdentifier = rl.ResourceServerIdentifier
		addPermisionsToRolesReq.PermissionName = rl.PermissionName
		addPermisionsToRolesReq.RoleId = role.Id
		addPermisionsToRolesReq.Auth0Domain = domain
		addPermisionsToRolesReq.Auth0MgmtToken = mgmtToken
		addPermisionsToRolesReq.UserEmail = userEmail
		addPermisionsToRolesReq.RequestId = requestId

		err = common.AddPermisionsToRolesResp(ctx, &addPermisionsToRolesReq)
		if err != nil {
			return
		}

		for _, userId := range rl.UserIds {
			assignRolesToUsersReq := commonproto.AssignRolesToUsers{}
			assignRolesToUsersReq.RoleId = role.Id
			assignRolesToUsersReq.AssignToUserId = userId
			assignRolesToUsersReq.Auth0Domain = domain
			assignRolesToUsersReq.Auth0MgmtToken = mgmtToken
			assignRolesToUsersReq.UserEmail = userEmail
			assignRolesToUsersReq.RequestId = requestId
			err = common.AssignRolesToUsersResp(ctx, &assignRolesToUsersReq)
			if err != nil {
				return
			}
		}
	}
}
