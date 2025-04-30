package common

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strings"

	commonproto "github.com/cloudfresco/sc-dcsa/internal/protogen/common/v1"
	"go.uber.org/zap"
)

// CreateRoleResp - For Creating Role
func CreateRoleResp(ctx context.Context, in *commonproto.CreateRole) (*commonproto.Role, error) {
	url := "https://" + in.Auth0Domain + "/api/v2/roles"
	payload := strings.NewReader(`{"name":"` + in.Name + `","description":"` + in.Description + `"}`)

	respBody, err := SendRequest("POST", url, payload, "Bearer "+in.Auth0MgmtToken)
	if err != nil {
		log.Error("Error", zap.String("user", in.UserEmail), zap.String("reqid", in.RequestId), zap.Error(err))
		return nil, err
	}

	roleResp, err := DecodeRespBody(respBody)
	if err != nil {
		log.Error("Error", zap.String("user", in.UserEmail), zap.String("reqid", in.RequestId), zap.Error(err))
		return nil, err
	}
	role := commonproto.Role{}
	role.Id = roleResp["id"].(string)
	role.Name = roleResp["name"].(string)
	role.Description = roleResp["description"].(string)

	return &role, nil
}

// GetRoleResp - to get Role
func GetRoleResp(ctx context.Context, in *commonproto.GetRole) (*commonproto.Role, error) {
	url := "https://" + in.Auth0Domain + "/api/v2/roles/" + in.RoleId
	respBody, err := SendRequest("GET", url, nil, "Bearer "+in.Auth0MgmtToken)
	if err != nil {
		log.Error("Error", zap.String("user", in.UserEmail), zap.String("reqid", in.RequestId), zap.Error(err))
		return nil, err
	}

	roleResp, err := DecodeRespBody(respBody)
	if err != nil {
		log.Error("Error", zap.String("user", in.UserEmail), zap.String("reqid", in.RequestId), zap.Error(err))
		return nil, err
	}
	role := commonproto.Role{}
	role.Id = roleResp["id"].(string)
	role.Name = roleResp["name"].(string)
	role.Description = roleResp["description"].(string)

	return &role, nil
}

// GetRolesResp - to get Roles
func GetRolesResp(ctx context.Context, in *commonproto.GetRoles) ([]*commonproto.Role, error) {
	url := "https://" + in.Auth0Domain + "/api/v2/roles"
	respBody, err := SendRequest("GET", url, nil, "Bearer "+in.Auth0MgmtToken)
	if err != nil {
		log.Error("Error", zap.String("user", in.UserEmail), zap.String("reqid", in.RequestId), zap.Error(err))
		return nil, err
	}

	roleResp, err := DecodeListRespBody(respBody)
	if err != nil {
		log.Error("Error", zap.String("user", in.UserEmail), zap.String("reqid", in.RequestId), zap.Error(err))
		return nil, err
	}
	roles := []*commonproto.Role{}
	for _, rl := range roleResp {
		role := commonproto.Role{}
		role.Id = rl["id"].(string)
		role.Name = rl["name"].(string)
		role.Description = rl["description"].(string)
		roles = append(roles, &role)
	}

	return roles, nil
}

// UpdateRoleResp - For Updating Role
func UpdateRoleResp(ctx context.Context, in *commonproto.UpdateRole) (*commonproto.Role, error) {
	url := "https://" + in.Auth0Domain + "/api/v2/roles" + in.RoleId
	payload := strings.NewReader(`{"name":"` + in.Name + `","description":"` + in.Description + `"}`)

	respBody, err := SendRequest("PATCH", url, payload, "Bearer "+in.Auth0MgmtToken)
	if err != nil {
		log.Error("Error", zap.String("user", in.UserEmail), zap.String("reqid", in.RequestId), zap.Error(err))
		return nil, err
	}

	roleResp, err := DecodeRespBody(respBody)
	if err != nil {
		log.Error("Error", zap.String("user", in.UserEmail), zap.String("reqid", in.RequestId), zap.Error(err))
		return nil, err
	}
	role := commonproto.Role{}
	role.Id = roleResp["id"].(string)
	role.Name = roleResp["name"].(string)
	role.Description = roleResp["description"].(string)

	return &role, nil
}

// DeleteRoleResp - to delete Role
func DeleteRoleResp(ctx context.Context, in *commonproto.DeleteRole) error {
	url := "https://" + in.Auth0Domain + "/api/v2/roles/" + in.RoleId
	_, err := SendRequest("DELETE", url, nil, "Bearer "+in.Auth0MgmtToken)
	if err != nil {
		log.Error("Error", zap.String("user", in.UserEmail), zap.String("reqid", in.RequestId), zap.Error(err))
		return err
	}
	return nil
}

// AddPermisionsToRolesResp -  Add Permisions To Roles
func AddPermisionsToRolesResp(ctx context.Context, in *commonproto.AddPermisionsToRoles) error {
	url := "https://" + in.Auth0Domain + "/api/v2/roles/" + in.RoleId + "/permissions"

	payload := strings.NewReader(`{"permissions": [{"resource_server_identifier":"` + in.ResourceServerIdentifier + `","permission_name":"` + in.PermissionName + `"}]}`)

	_, err := SendRequest("POST", url, payload, "Bearer "+in.Auth0MgmtToken)
	if err != nil {
		log.Error("Error", zap.String("user", in.UserEmail), zap.String("reqid", in.RequestId), zap.Error(err))
		return err
	}
	return nil
}

// RemoveRolePermissionResp -  Remove Role Permission
func RemoveRolePermissionResp(ctx context.Context, in *commonproto.RemoveRolePermission) error {
	url := "https://" + in.Auth0Domain + "/api/v2/roles/" + in.RoleId + "/permissions"

	payload := strings.NewReader(`{"permissions": [{"resource_server_identifier":"` + in.ResourceServerIdentifier + `","permission_name":"` + in.PermissionName + `"}]}`)

	_, err := SendRequest("DELETE", url, payload, "Bearer "+in.Auth0MgmtToken)
	if err != nil {
		log.Error("Error", zap.String("user", in.UserEmail), zap.String("reqid", in.RequestId), zap.Error(err))
		return err
	}
	return nil
}

// GetRolePermissionsResp - to get RolePermissions
func GetRolePermissionsResp(ctx context.Context, in *commonproto.GetRolePermissions) ([]*commonproto.RolePermission, error) {
	url := "https://" + in.Auth0Domain + "/api/v2/roles/" + in.RoleId + "/permissions"
	respBody, err := SendRequest("GET", url, nil, "Bearer "+in.Auth0MgmtToken)
	if err != nil {
		log.Error("Error", zap.String("user", in.UserEmail), zap.String("reqid", in.RequestId), zap.Error(err))
		return nil, err
	}

	rolePermissionsResp, err := DecodeListRespBody(respBody)
	if err != nil {
		log.Error("Error", zap.String("user", in.UserEmail), zap.String("reqid", in.RequestId), zap.Error(err))
		return nil, err
	}
	rolePermissions := []*commonproto.RolePermission{}
	for _, rl := range rolePermissionsResp {
		rolePermission := commonproto.RolePermission{}
		rolePermission.PermissionName = rl["permission_name"].(string)
		rolePermission.Description = rl["description"].(string)
		rolePermission.ResourceServerName = rl["resource_server_name"].(string)
		rolePermission.ResourceServerIdentifier = rl["resource_server_identifier"].(string)
		rolePermissions = append(rolePermissions, &rolePermission)
	}
	return rolePermissions, nil
}

// AssignRolesToUsersResp -  Assign Roles To Users
func AssignRolesToUsersResp(ctx context.Context, in *commonproto.AssignRolesToUsers) error {
	url := "https://" + in.Auth0Domain + "/api/v2/users/" + in.AssignToUserId + "/roles"

	payload := strings.NewReader(`{ "roles": [ "` + in.RoleId + `"] }`)

	_, err := SendRequest("POST", url, payload, "Bearer "+in.Auth0MgmtToken)
	if err != nil {
		log.Error("Error", zap.String("user", in.UserEmail), zap.String("reqid", in.RequestId), zap.Error(err))
		return err
	}
	return nil
}

// ViewUserRolesResp - view user roles
func ViewUserRolesResp(ctx context.Context, in *commonproto.ViewUserRoles) ([]*commonproto.Role, error) {
	url := "https://" + in.Auth0Domain + "/api/v2/users/" + in.UserId + "/roles"
	respBody, err := SendRequest("GET", url, nil, "Bearer "+in.Auth0MgmtToken)
	if err != nil {
		log.Error("Error", zap.String("user", in.UserEmail), zap.String("reqid", in.RequestId), zap.Error(err))
		return nil, err
	}

	roleResp, err := DecodeListRespBody(respBody)
	if err != nil {
		log.Error("Error", zap.String("user", in.UserEmail), zap.String("reqid", in.RequestId), zap.Error(err))
		return nil, err
	}
	roles := []*commonproto.Role{}
	for _, rl := range roleResp {
		role := commonproto.Role{}
		role.Id = rl["id"].(string)
		role.Name = rl["name"].(string)
		role.Description = rl["description"].(string)
		roles = append(roles, &role)
	}

	return roles, nil
}

type APIPermission struct {
	Value       string `json:"value"`
	Description string `json:"description"`
}

// AddAPIPermissionResp - add api permissions
func AddAPIPermissionResp(ctx context.Context, in *commonproto.AddAPIPermission) error {
	APIPermissions := []APIPermission{}
	for _, permission := range in.Permissions {
		ap := APIPermission{Value: permission.PermissionName, Description: permission.PermissionDescription}
		APIPermissions = append(APIPermissions, ap)
	}

	permissionData, _ := json.Marshal(APIPermissions)

	url := "https://" + in.Auth0Domain + "/api/v2/resource-servers/" + in.Auth0ApiId

	payload := strings.NewReader(`{"scopes":` + string(permissionData) + `}`)

	_, err := SendRequest("PATCH", url, payload, "Bearer "+in.Auth0MgmtToken)
	if err != nil {
		log.Error("Error", zap.String("user", in.UserEmail), zap.String("reqid", in.RequestId), zap.Error(err))
		return err
	}

	return nil
}

// SendRequest - common code for all requests
func SendRequest(method, url string, payload io.Reader, mgmtToken string) ([]byte, error) {
	req, _ := http.NewRequest(method, url, payload)

	req.Header.Add("content-type", "application/json")
	req.Header.Add("authorization", mgmtToken)
	req.Header.Add("cache-control", "no-cache")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// DecodeRespBody - decode Response
func DecodeRespBody(respBody []byte) (map[string]interface{}, error) {
	jsonDataReader := strings.NewReader(string(respBody))
	decoder := json.NewDecoder(jsonDataReader)
	var roleResp map[string]interface{}
	err := decoder.Decode(&roleResp)
	if err != nil {
		return nil, err
	}
	return roleResp, nil
}

// DecodeListRespBody - decode Response
func DecodeListRespBody(respBody []byte) ([]map[string]interface{}, error) {
	jsonDataReader := strings.NewReader(string(respBody))
	decoder := json.NewDecoder(jsonDataReader)
	var roleResp []map[string]interface{}
	err := decoder.Decode(&roleResp)
	if err != nil {
		return nil, err
	}
	return roleResp, nil
}
