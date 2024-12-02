package main

import (
	//"encoding/json"
	"context"
	"fmt"

	"github.com/cloudfresco/sc-dcsa/internal/common"
	commonproto "github.com/cloudfresco/sc-dcsa/internal/protogen/common/v1"
	"github.com/spf13/viper"
	//"io"
	//"io/ioutil"
	//"net/http"
	//"strings"
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

/*func GetRoleData(roleName string, roleDescription string, resourceServerIdentifier string, permissionName string, permissionDescription string, userIds []string) *RoleData {
	role := new(RoleData)
	role.RoleName = roleName
	role.RoleDescription = roleDescription
	role.ResourceServerIdentifier = resourceServerIdentifier
	role.PermissionName = permissionName
	role.PermissionDescription = permissionDescription
	role.UserIds = userIds
	return role
}*/

// common code for all requests
/*func SendRequest(method, url string, payload io.Reader, mgmtToken string) ([]byte, error) {
	req, _ := http.NewRequest(method, url, payload)

	req.Header.Add("content-type", "application/json")
	req.Header.Add("authorization", mgmtToken)
	req.Header.Add("cache-control", "no-cache")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	fmt.Println("SendRequest res", res)

	return body, nil
}

// https://auth0.com/docs/get-started/apis/add-api-permissions
func AddAPIPermission(API_ID string, mgmtToken string, permissionData string) error {
	fmt.Println("AddAPIPermission started")
	url := "https://dev-8l3ua0t8fhmh08vq.us.auth0.com/api/v2/resource-servers/" + API_ID

	// payload := strings.NewReader("{ \"scopes\": [ { \"value\": \"PERMISSION_NAME\", \"description\": \"PERMISSION_DESC\" }, { \"value\": \"PERMISSION_NAME\", \"description\": \"PERMISSION_DESC\" } ] }")

	// req, _ := http.NewRequest("PATCH", url, payload)

	// payload := strings.NewReader(`{"scopes": [ {"value":"` +permissionName+ `", "description":"` + permissionDescription + `" }] }`)

	// payload := strings.NewReader("{ \"scopes\": [ { \"value\": \"cud:usersdata\", \"description\": \"create or update or delete users\" }, { \"value\": \"read:usersdata\", \"description\": \"read users\" }, { \"value\": \"cud:partiesdata\", \"description\": \"create or update or delete parties\" }, { \"value\": \"read:partiesdata\", \"description\": \"read parties\" } ] }")

	payload := strings.NewReader(`{"scopes":` + permissionData + `}`)

	fmt.Println("payload", payload)

	respBody, err := SendRequest("PATCH", url, payload, mgmtToken)
	if err != nil {
		fmt.Println("err", err)
		return err
	}

	fmt.Println("AddAPIPermission string(respBody) is", string(respBody))
	fmt.Println("AddAPIPermission ended")
	return nil
}

// https://auth0.com/docs/secure/tokens/access-tokens/management-api-access-tokens/get-management-api-access-tokens-for-testing
// https://auth0.com/docs/manage-users/access-control/configure-core-rbac/roles/create-roles
func CreateRole(name string, description string, mgmtToken string) (*Role, error) {
	fmt.Println("CreateRole started")
	url := "https://dev-8l3ua0t8fhmh08vq.us.auth0.com/api/v2/roles"

	// payload := strings.NewReader("{ \"name\": \"read-msg\", \"description\": \"Access messaging features\" }")
	payload := strings.NewReader(`{"name":"` + name + `","description":"` + description + `"}`)

	respBody, err := SendRequest("POST", url, payload, mgmtToken)
	if err != nil {
		fmt.Println("err", err)
		return nil, err
	}
	role := &Role{}
	err = json.Unmarshal(respBody, role)
	if err != nil {
		fmt.Println("err", err)
		return nil, err
	}
	fmt.Println("CreateRole role.Id is", role.Id)
	fmt.Println("CreateRole ended")

	return role, nil
}

// https://auth0.com/docs/api/management/v2/roles/get-roles-by-id
func GetRole(roleID string, mgmtToken string) error {
	url := "https://dev-8l3ua0t8fhmh08vq.us.auth0.com/api/v2/roles/" + roleID

	respBody, err := SendRequest("GET", url, nil, mgmtToken)
	if err != nil {
		fmt.Println("err", err)
		return err
	}

	fmt.Println("GetRole string(respBody) is", string(respBody))
	return nil
}

// https://auth0.com/docs/api/management/v2/roles/get-roles
func GetRoles(mgmtToken string) (*GetRolesResponse, error) {
	fmt.Println("GetRoles started")
	url := "https://dev-8l3ua0t8fhmh08vq.us.auth0.com/api/v2/roles"

	respBody, err := SendRequest("GET", url, nil, mgmtToken)
	if err != nil {
		fmt.Println("err", err)
		return nil, err
	}

	fmt.Println("GetRoles string(respBody) is", string(respBody))
	jsonDataReader := strings.NewReader(string(respBody))
	decoder := json.NewDecoder(jsonDataReader)
	var roleResp []map[string]interface{}
	err = decoder.Decode(&roleResp)
	if err != nil {
		fmt.Println("err", err)
		return nil, err
	}
	roles := []*Role{}
	for _, rl := range roleResp {
		role := Role{}
		role.Id = rl["id"].(string)
		role.Name = rl["name"].(string)
		role.Description = rl["description"].(string)
		roles = append(roles, &role)
	}
	roleResponse := GetRolesResponse{}
	roleResponse.Roles = roles
	fmt.Println("GetRoles ended")
	return &roleResponse, nil
}

// https://auth0.com/docs/api/management/v2/roles/delete-roles-by-id
func DeleteRole(roleID string, mgmtToken string) error {
	fmt.Println("DeleteRole started")
	url := "https://dev-8l3ua0t8fhmh08vq.us.auth0.com/api/v2/roles/" + roleID
	respBody, err := SendRequest("DELETE", url, nil, mgmtToken)
	if err != nil {
		fmt.Println("err", err)
		return err
	}

	fmt.Println("DeleteRole string(respBody) is", string(respBody))
	fmt.Println("DeleteRole ended")
	return nil
}

// https://auth0.com/docs/api/management/v2/roles/patch-roles-by-id
func UpdateRole(roleID string, mgmtToken string) error {
	url := "https://dev-8l3ua0t8fhmh08vq.us.auth0.com/api/v2/roles/" + roleID
	payload := strings.NewReader(`{"name":"read-all-msg","description":"Access all messaging features"}`)

	respBody, err := SendRequest("PATCH", url, payload, mgmtToken)
	if err != nil {
		fmt.Println("err", err)
		return err
	}
	fmt.Println("UpdateRole string(body) is", string(respBody))
	return nil
}

// https://auth0.com/docs/manage-users/access-control/configure-core-rbac/roles/add-permissions-to-roles
func AddPermisionsToRoles(roleID string, resourceServerIdentifier string, permissionName string, mgmtToken string) error {
	fmt.Println("AddPermisionsToRoles started")
	url := "https://dev-8l3ua0t8fhmh08vq.us.auth0.com/api/v2/roles/" + roleID + "/permissions"

	// payload := strings.NewReader("{ \"permissions\": [ { \"resource_server_identifier\": \"https://authjwt.com/\", \"permission_name\": \"read:messages\" }] }")

	payload := strings.NewReader(`{"permissions": [{"resource_server_identifier":"` + resourceServerIdentifier + `","permission_name":"` + permissionName + `"}]}`)

	respBody, err := SendRequest("POST", url, payload, mgmtToken)
	if err != nil {
		fmt.Println("err", err)
		return err
	}
	fmt.Println("AddPermisionsToRoles string(body) is", string(respBody))
	fmt.Println("AddPermisionsToRoles ended")
	return nil
}

// https://auth0.com/docs/api/management/v2/roles/get-role-permission
func GetRolePermission(roleID string, mgmtToken string) (*GetRolePermissionResponse, error) {
	fmt.Println("GetRolePermission started")
	url := "https://dev-8l3ua0t8fhmh08vq.us.auth0.com/api/v2/roles/" + roleID + "/permissions"

	respBody, err := SendRequest("GET", url, nil, mgmtToken)
	if err != nil {
		fmt.Println("err", err)
		return nil, err
	}

	fmt.Println("GetRolePermission string(respBody) is", string(respBody))
	jsonDataReader := strings.NewReader(string(respBody))
	decoder := json.NewDecoder(jsonDataReader)
	var rolePermissionResp []map[string]interface{}
	err = decoder.Decode(&rolePermissionResp)
	if err != nil {
		fmt.Println("err", err)
		return nil, err
	}
	rolePermissions := []*RolePermission{}
	for _, rl := range rolePermissionResp {
		rolePermission := RolePermission{}
		rolePermission.PermissionName = rl["permission_name"].(string)
		rolePermission.Description = rl["description"].(string)
		rolePermission.ResourceServerName = rl["resource_server_name"].(string)
		rolePermission.ResourceServerIdentifier = rl["resource_server_identifier"].(string)
		rolePermissions = append(rolePermissions, &rolePermission)
	}
	rolePermissionResponse := GetRolePermissionResponse{}
	rolePermissionResponse.RolePermissions = rolePermissions
	fmt.Println("GetRolePermission ended")
	return &rolePermissionResponse, nil
}

// https://auth0.com/docs/api/management/v2/roles/delete-role-permission-assignment
func RemoveRolePermission(roleID string, resourceServerIdentifier string, permissionName string, mgmtToken string) error {
	fmt.Println("RemoveRolePermission started")
	url := "https://dev-8l3ua0t8fhmh08vq.us.auth0.com/api/v2/roles/" + roleID + "/permissions"

	payload := strings.NewReader(`{"permissions": [{"resource_server_identifier":"` + resourceServerIdentifier + `","permission_name":"` + permissionName + `"}]}`)

	// payload := strings.NewReader("{ \"permissions\": [ { \"resource_server_identifier\": \"https://authjwt.com/\", \"permission_name\": \"read:messages\" }] }")

	respBody, err := SendRequest("DELETE", url, payload, mgmtToken)
	if err != nil {
		fmt.Println("err", err)
		return err
	}

	fmt.Println("RemoveRolePermission string(respBody) is", string(respBody))
	fmt.Println("RemoveRolePermission ended")
	return nil
}

// https://auth0.com/docs/manage-users/access-control/configure-core-rbac/rbac-users/assign-roles-to-users

func AssignRolesToUsers(userID string, roleID string, mgmtToken string) error {
	fmt.Println("CreateRole started")
	url := "https://dev-8l3ua0t8fhmh08vq.us.auth0.com/api/v2/users/" + userID + "/roles"

	payload := strings.NewReader(`{ "roles": [ "` + roleID + `"] }`)

	respBody, err := SendRequest("POST", url, payload, mgmtToken)
	if err != nil {
		fmt.Println("err", err)
		return err
	}

	fmt.Println("AssignRolesToUsers string(respBody) is", string(respBody))
	return nil
}

// https://auth0.com/docs/api/management/v2/roles/get-role-user
func GetRoleUsers(roleID string, mgmtToken string) error {
	url := "https://dev-8l3ua0t8fhmh08vq.us.auth0.com/api/v2/roles/" + roleID + "/users"

	respBody, err := SendRequest("GET", url, nil, mgmtToken)
	if err != nil {
		fmt.Println("err", err)
		return err
	}

	fmt.Println("GetRoleUsers string(respBody) is", string(respBody))
	return nil
}

// https://auth0.com/docs/manage-users/access-control/configure-core-rbac/rbac-users/view-user-roles

func ViewUserRoles(userID string, mgmtToken string) error {
	url := "https://dev-8l3ua0t8fhmh08vq.us.auth0.com/api/v2/users/" + userID + "/roles"

	respBody, err := SendRequest("GET", url, nil, mgmtToken)
	if err != nil {
		fmt.Println("err", err)
		return err
	}

	fmt.Println("ViewUserRoles string(respBody)", string(respBody))
	return nil
}

type APIPermission struct {
	Value       string `json:"value"`
	Description string `json:"description"`
}*/

func main() {
	v := viper.New()
	v.AutomaticEnv()

	v.SetConfigName("roles")
	v.SetConfigType("json")
	v.AddConfigPath("./test/roles")
	if err := v.ReadInConfig(); err != nil {
		fmt.Println(err)
	}

	rolesData := RoleOptions{}
	if err := v.Unmarshal(&rolesData); err != nil {
		fmt.Println(err)
	}

	fmt.Println("roles", rolesData)

	mgmtToken := v.GetString("SC_DCSA_AUTH0_MGMTTOKEN")
	fmt.Println("mgmtToken", mgmtToken)

	API_ID := v.GetString("SC_DCSA_AUTH0_API_ID")
	fmt.Println("API_ID", API_ID)

	domain := v.GetString("SC_DCSA_AUTH0_DOMAIN")
	fmt.Println("domain", domain)

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
		fmt.Println("err", err)
		return
	}
	// fmt.Println("roleResponse", roleResponse)
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
				fmt.Println("err", err)
				return
			}
			fmt.Println("rolePermissionResponse", rolePermissionResponse)
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
					fmt.Println("err", err)
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
				fmt.Println("err", err)
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
		fmt.Println("err", err)
		return
	}

	// create role with permissions and assign to users
	for _, rl := range rolesData.Roles {
		fmt.Println("rl.RoleName", rl.RoleName)
		createRoleReq := commonproto.CreateRole{}
		createRoleReq.Name = rl.RoleName
		createRoleReq.Description = rl.RoleDescription
		createRoleReq.Auth0Domain = domain
		createRoleReq.Auth0MgmtToken = mgmtToken
		createRoleReq.UserEmail = userEmail
		createRoleReq.RequestId = requestId
		role, err := common.CreateRoleResp(ctx, &createRoleReq)
		if err != nil {
			fmt.Println("err", err)
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
			fmt.Println("err", err)
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
				fmt.Println("err", err)
				return
			}
		}
	}

	/*mgmtToken := "Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6Ik1TU09vRW41dU5JbDAzeVN3d3FJbSJ9.eyJpc3MiOiJodHRwczovL2Rldi04bDN1YTB0OGZobWgwOHZxLnVzLmF1dGgwLmNvbS8iLCJzdWIiOiJmckVMd01CWFM0WDFhR3B5WEhGQkNQQWNzYkxtYjROVUBjbGllbnRzIiwiYXVkIjoiaHR0cHM6Ly9kZXYtOGwzdWEwdDhmaG1oMDh2cS51cy5hdXRoMC5jb20vYXBpL3YyLyIsImlhdCI6MTcyOTIzMjgxMywiZXhwIjoxNzMxODI0ODEzLCJzY29wZSI6InJlYWQ6Y2xpZW50X2dyYW50cyBjcmVhdGU6Y2xpZW50X2dyYW50cyBkZWxldGU6Y2xpZW50X2dyYW50cyB1cGRhdGU6Y2xpZW50X2dyYW50cyByZWFkOnVzZXJzIHVwZGF0ZTp1c2VycyBkZWxldGU6dXNlcnMgY3JlYXRlOnVzZXJzIHJlYWQ6dXNlcnNfYXBwX21ldGFkYXRhIHVwZGF0ZTp1c2Vyc19hcHBfbWV0YWRhdGEgZGVsZXRlOnVzZXJzX2FwcF9tZXRhZGF0YSBjcmVhdGU6dXNlcnNfYXBwX21ldGFkYXRhIHJlYWQ6dXNlcl9jdXN0b21fYmxvY2tzIGNyZWF0ZTp1c2VyX2N1c3RvbV9ibG9ja3MgZGVsZXRlOnVzZXJfY3VzdG9tX2Jsb2NrcyBjcmVhdGU6dXNlcl90aWNrZXRzIHJlYWQ6Y2xpZW50cyB1cGRhdGU6Y2xpZW50cyBkZWxldGU6Y2xpZW50cyBjcmVhdGU6Y2xpZW50cyByZWFkOmNsaWVudF9rZXlzIHVwZGF0ZTpjbGllbnRfa2V5cyBkZWxldGU6Y2xpZW50X2tleXMgY3JlYXRlOmNsaWVudF9rZXlzIHJlYWQ6Y29ubmVjdGlvbnMgdXBkYXRlOmNvbm5lY3Rpb25zIGRlbGV0ZTpjb25uZWN0aW9ucyBjcmVhdGU6Y29ubmVjdGlvbnMgcmVhZDpyZXNvdXJjZV9zZXJ2ZXJzIHVwZGF0ZTpyZXNvdXJjZV9zZXJ2ZXJzIGRlbGV0ZTpyZXNvdXJjZV9zZXJ2ZXJzIGNyZWF0ZTpyZXNvdXJjZV9zZXJ2ZXJzIHJlYWQ6ZGV2aWNlX2NyZWRlbnRpYWxzIHVwZGF0ZTpkZXZpY2VfY3JlZGVudGlhbHMgZGVsZXRlOmRldmljZV9jcmVkZW50aWFscyBjcmVhdGU6ZGV2aWNlX2NyZWRlbnRpYWxzIHJlYWQ6cnVsZXMgdXBkYXRlOnJ1bGVzIGRlbGV0ZTpydWxlcyBjcmVhdGU6cnVsZXMgcmVhZDpydWxlc19jb25maWdzIHVwZGF0ZTpydWxlc19jb25maWdzIGRlbGV0ZTpydWxlc19jb25maWdzIHJlYWQ6aG9va3MgdXBkYXRlOmhvb2tzIGRlbGV0ZTpob29rcyBjcmVhdGU6aG9va3MgcmVhZDphY3Rpb25zIHVwZGF0ZTphY3Rpb25zIGRlbGV0ZTphY3Rpb25zIGNyZWF0ZTphY3Rpb25zIHJlYWQ6ZW1haWxfcHJvdmlkZXIgdXBkYXRlOmVtYWlsX3Byb3ZpZGVyIGRlbGV0ZTplbWFpbF9wcm92aWRlciBjcmVhdGU6ZW1haWxfcHJvdmlkZXIgYmxhY2tsaXN0OnRva2VucyByZWFkOnN0YXRzIHJlYWQ6aW5zaWdodHMgcmVhZDp0ZW5hbnRfc2V0dGluZ3MgdXBkYXRlOnRlbmFudF9zZXR0aW5ncyByZWFkOmxvZ3MgcmVhZDpsb2dzX3VzZXJzIHJlYWQ6c2hpZWxkcyBjcmVhdGU6c2hpZWxkcyB1cGRhdGU6c2hpZWxkcyBkZWxldGU6c2hpZWxkcyByZWFkOmFub21hbHlfYmxvY2tzIGRlbGV0ZTphbm9tYWx5X2Jsb2NrcyB1cGRhdGU6dHJpZ2dlcnMgcmVhZDp0cmlnZ2VycyByZWFkOmdyYW50cyBkZWxldGU6Z3JhbnRzIHJlYWQ6Z3VhcmRpYW5fZmFjdG9ycyB1cGRhdGU6Z3VhcmRpYW5fZmFjdG9ycyByZWFkOmd1YXJkaWFuX2Vucm9sbG1lbnRzIGRlbGV0ZTpndWFyZGlhbl9lbnJvbGxtZW50cyBjcmVhdGU6Z3VhcmRpYW5fZW5yb2xsbWVudF90aWNrZXRzIHJlYWQ6dXNlcl9pZHBfdG9rZW5zIGNyZWF0ZTpwYXNzd29yZHNfY2hlY2tpbmdfam9iIGRlbGV0ZTpwYXNzd29yZHNfY2hlY2tpbmdfam9iIHJlYWQ6Y3VzdG9tX2RvbWFpbnMgZGVsZXRlOmN1c3RvbV9kb21haW5zIGNyZWF0ZTpjdXN0b21fZG9tYWlucyB1cGRhdGU6Y3VzdG9tX2RvbWFpbnMgcmVhZDplbWFpbF90ZW1wbGF0ZXMgY3JlYXRlOmVtYWlsX3RlbXBsYXRlcyB1cGRhdGU6ZW1haWxfdGVtcGxhdGVzIHJlYWQ6bWZhX3BvbGljaWVzIHVwZGF0ZTptZmFfcG9saWNpZXMgcmVhZDpyb2xlcyBjcmVhdGU6cm9sZXMgZGVsZXRlOnJvbGVzIHVwZGF0ZTpyb2xlcyByZWFkOnByb21wdHMgdXBkYXRlOnByb21wdHMgcmVhZDpicmFuZGluZyB1cGRhdGU6YnJhbmRpbmcgZGVsZXRlOmJyYW5kaW5nIHJlYWQ6bG9nX3N0cmVhbXMgY3JlYXRlOmxvZ19zdHJlYW1zIGRlbGV0ZTpsb2dfc3RyZWFtcyB1cGRhdGU6bG9nX3N0cmVhbXMgY3JlYXRlOnNpZ25pbmdfa2V5cyByZWFkOnNpZ25pbmdfa2V5cyB1cGRhdGU6c2lnbmluZ19rZXlzIHJlYWQ6bGltaXRzIHVwZGF0ZTpsaW1pdHMgY3JlYXRlOnJvbGVfbWVtYmVycyByZWFkOnJvbGVfbWVtYmVycyBkZWxldGU6cm9sZV9tZW1iZXJzIHJlYWQ6ZW50aXRsZW1lbnRzIHJlYWQ6YXR0YWNrX3Byb3RlY3Rpb24gdXBkYXRlOmF0dGFja19wcm90ZWN0aW9uIHJlYWQ6b3JnYW5pemF0aW9uc19zdW1tYXJ5IGNyZWF0ZTphdXRoZW50aWNhdGlvbl9tZXRob2RzIHJlYWQ6YXV0aGVudGljYXRpb25fbWV0aG9kcyB1cGRhdGU6YXV0aGVudGljYXRpb25fbWV0aG9kcyBkZWxldGU6YXV0aGVudGljYXRpb25fbWV0aG9kcyByZWFkOm9yZ2FuaXphdGlvbnMgdXBkYXRlOm9yZ2FuaXphdGlvbnMgY3JlYXRlOm9yZ2FuaXphdGlvbnMgZGVsZXRlOm9yZ2FuaXphdGlvbnMgY3JlYXRlOm9yZ2FuaXphdGlvbl9tZW1iZXJzIHJlYWQ6b3JnYW5pemF0aW9uX21lbWJlcnMgZGVsZXRlOm9yZ2FuaXphdGlvbl9tZW1iZXJzIGNyZWF0ZTpvcmdhbml6YXRpb25fY29ubmVjdGlvbnMgcmVhZDpvcmdhbml6YXRpb25fY29ubmVjdGlvbnMgdXBkYXRlOm9yZ2FuaXphdGlvbl9jb25uZWN0aW9ucyBkZWxldGU6b3JnYW5pemF0aW9uX2Nvbm5lY3Rpb25zIGNyZWF0ZTpvcmdhbml6YXRpb25fbWVtYmVyX3JvbGVzIHJlYWQ6b3JnYW5pemF0aW9uX21lbWJlcl9yb2xlcyBkZWxldGU6b3JnYW5pemF0aW9uX21lbWJlcl9yb2xlcyBjcmVhdGU6b3JnYW5pemF0aW9uX2ludml0YXRpb25zIHJlYWQ6b3JnYW5pemF0aW9uX2ludml0YXRpb25zIGRlbGV0ZTpvcmdhbml6YXRpb25faW52aXRhdGlvbnMgcmVhZDpzY2ltX2NvbmZpZyBjcmVhdGU6c2NpbV9jb25maWcgdXBkYXRlOnNjaW1fY29uZmlnIGRlbGV0ZTpzY2ltX2NvbmZpZyBjcmVhdGU6c2NpbV90b2tlbiByZWFkOnNjaW1fdG9rZW4gZGVsZXRlOnNjaW1fdG9rZW4gZGVsZXRlOnBob25lX3Byb3ZpZGVycyBjcmVhdGU6cGhvbmVfcHJvdmlkZXJzIHJlYWQ6cGhvbmVfcHJvdmlkZXJzIHVwZGF0ZTpwaG9uZV9wcm92aWRlcnMgZGVsZXRlOnBob25lX3RlbXBsYXRlcyBjcmVhdGU6cGhvbmVfdGVtcGxhdGVzIHJlYWQ6cGhvbmVfdGVtcGxhdGVzIHVwZGF0ZTpwaG9uZV90ZW1wbGF0ZXMgY3JlYXRlOmVuY3J5cHRpb25fa2V5cyByZWFkOmVuY3J5cHRpb25fa2V5cyB1cGRhdGU6ZW5jcnlwdGlvbl9rZXlzIGRlbGV0ZTplbmNyeXB0aW9uX2tleXMgcmVhZDpzZXNzaW9ucyBkZWxldGU6c2Vzc2lvbnMgcmVhZDpyZWZyZXNoX3Rva2VucyBkZWxldGU6cmVmcmVzaF90b2tlbnMgY3JlYXRlOnNlbGZfc2VydmljZV9wcm9maWxlcyByZWFkOnNlbGZfc2VydmljZV9wcm9maWxlcyB1cGRhdGU6c2VsZl9zZXJ2aWNlX3Byb2ZpbGVzIGRlbGV0ZTpzZWxmX3NlcnZpY2VfcHJvZmlsZXMgY3JlYXRlOnNzb19hY2Nlc3NfdGlja2V0cyByZWFkOmZvcm1zIHVwZGF0ZTpmb3JtcyBkZWxldGU6Zm9ybXMgY3JlYXRlOmZvcm1zIHJlYWQ6Zmxvd3MgdXBkYXRlOmZsb3dzIGRlbGV0ZTpmbG93cyBjcmVhdGU6Zmxvd3MgcmVhZDpmbG93c192YXVsdCByZWFkOmZsb3dzX3ZhdWx0X2Nvbm5lY3Rpb25zIHVwZGF0ZTpmbG93c192YXVsdF9jb25uZWN0aW9ucyBkZWxldGU6Zmxvd3NfdmF1bHRfY29ubmVjdGlvbnMgY3JlYXRlOmZsb3dzX3ZhdWx0X2Nvbm5lY3Rpb25zIHJlYWQ6Zmxvd3NfZXhlY3V0aW9ucyBkZWxldGU6Zmxvd3NfZXhlY3V0aW9ucyByZWFkOmNvbm5lY3Rpb25zX29wdGlvbnMgdXBkYXRlOmNvbm5lY3Rpb25zX29wdGlvbnMgcmVhZDpjbGllbnRfY3JlZGVudGlhbHMgY3JlYXRlOmNsaWVudF9jcmVkZW50aWFscyB1cGRhdGU6Y2xpZW50X2NyZWRlbnRpYWxzIGRlbGV0ZTpjbGllbnRfY3JlZGVudGlhbHMiLCJndHkiOiJjbGllbnQtY3JlZGVudGlhbHMiLCJhenAiOiJmckVMd01CWFM0WDFhR3B5WEhGQkNQQWNzYkxtYjROVSJ9.md42jX6gMEZ5biPO2c1A5q34y0i7q-SLbUEqOanZohijkfbddcytsag9gKr_zMhFuYqd94z3STHrg4tR2CtU5LrGwgXy_OI2VezaRdCBYGonwuWpkajjScFCuTt6ITrRskJYy_SwMoWZLPuYZ4LYa1f7RnzUgoxk3O11t2WLmRPvCg8kMTJ2X4hioE559VFCJauf4fg12iA6Yavk_BkQreh8GVMZ2Lf0nFFaVxWcNcXWM_bf7_UTG5AT_plvVpFv3_aMJ1v_HLlwU4DjnYahbYIxYmIOxWmEAn4wzB1zbqUePkS43aZzQcZ9IGFxzTZube9OOAAPjeX-nNEgTfcINQ"

	API_ID := "67121dd2e4e2c893f694b684"

	roles := []*RoleData{}
	role1 := GetRoleData("cud:users", "create or update or delete users", "https://hello-world.example.com", "cud:usersdata", "create or update or delete users", []string{"auth0|66fd06d0bfea78a82bb42459"})

	role2 := GetRoleData("read:users", "read users", "https://hello-world.example.com", "read:usersdata", "read users", []string{"auth0|66fd06d0bfea78a82bb42459"})

	role3 := GetRoleData("cud:parties", "create or update or delete parties", "https://hello-world.example.com", "cud:partiesdata", "create or update or delete parties", []string{"auth0|66fd06d0bfea78a82bb42459", "auth0|66fcdfb6d20dcb68e3fcbc3b"})

	role4 := GetRoleData("read:parties", "read parties", "https://hello-world.example.com", "read:partiesdata", "read parties", []string{"auth0|66fd06d0bfea78a82bb42459", "auth0|66fcdfb6d20dcb68e3fcbc3b"})

	roles = append(roles, role1, role2, role3, role4)

	// get roles
	roleResponse, err := GetRoles(mgmtToken)
	if err != nil {
		fmt.Println("err", err)
		return
	}
	// fmt.Println("roleResponse", roleResponse)
	// remove permissions from role and its roles
	if len(roleResponse.Roles) != 0 {
		for _, role := range roleResponse.Roles {
			// get role permissions
			rolePermissionResponse, err := GetRolePermission(role.Id, mgmtToken)
			if err != nil {
				fmt.Println("err", err)
				return
			}
			fmt.Println("rolePermissionResponse", rolePermissionResponse)
			for _, rolePermission := range rolePermissionResponse.RolePermissions {
				// remove Role Permission
				err = RemoveRolePermission(role.Id, rolePermission.ResourceServerIdentifier, rolePermission.PermissionName, mgmtToken)
				if err != nil {
					fmt.Println("err", err)
					return
				}
			}
			// Delete Role - delete role
			err = DeleteRole(role.Id, mgmtToken)
			if err != nil {
				fmt.Println("err", err)
				return
			}
		}
	}

	// add all permissions
	APIPermissions := []APIPermission{}
	for _, rl := range roles {
		ap := APIPermission{Value: rl.PermissionName, Description: rl.PermissionDescription}
		APIPermissions = append(APIPermissions, ap)
	}
	j, _ := json.Marshal(APIPermissions)
	// fmt.Println("j", j)
	// fmt.Println("j", string(j))

	err = AddAPIPermission(API_ID, mgmtToken, string(j))
	if err != nil {
		fmt.Println("err", err)
		return
	}

	// create role with permissions and assign to users
	for _, rl := range roles {
		role, err := CreateRole(rl.RoleName, rl.RoleDescription, mgmtToken)
		if err != nil {
			fmt.Println("err", err)
			return
		}
		err = AddPermisionsToRoles(role.Id, rl.ResourceServerIdentifier, rl.PermissionName, mgmtToken)
		if err != nil {
			fmt.Println("err", err)
			return
		}
		for _, userId := range rl.UserIds {
			err = AssignRolesToUsers(userId, role.Id, mgmtToken)
			if err != nil {
				fmt.Println("err", err)
				return
			}
		}
	}
	*/
}
