package settings

import (
	"os"
)

//
// Environment variables
const (
	EnvAuthRequired         = "AUTH_REQUIRED"
	EnvKeycloakHost         = "KEYCLOAK_HOST"
	EnvKeycloakRealm        = "KEYCLOAK_REALM"
	EnvKeycloakClientID     = "KEYCLOAK_CLIENT_ID"
	EnvKeycloakClientSecret = "KEYCLOAK_CLIENT_SECRET"
	EnvKeycloakAdminUser    = "KEYCLOAK_ADMIN_USER"
	EnvKeycloakAdminPass    = "KEYCLOAK_ADMIN_PASS"
	EnvKeycloakAdminRealm   = "KEYCLOAK_ADMIN_REALM"
	EnvAddonToken           = "ADDON_TOKEN"
	EnvRolePath             = "ROLE_PATH"
	EnvUserPath             = "USER_PATH"
)

type Auth struct {
	// Auth required
	Required bool
	// Keycloak client config
	Keycloak struct {
		Host         string
		Realm        string
		ClientID     string
		ClientSecret string
		Admin        struct {
			User  string
			Pass  string
			Realm string
		}
	}
	// Addon API access token
	AddonToken string
	// Path to role yaml
	RolePath string
	// Path to user yaml
	UserPath string
}

func (r *Auth) Load() (err error) {
	var found bool
	r.Required = getEnvBool(EnvAuthRequired, false)
	if !r.Required {
		return
	}
	r.Keycloak.Host, found = os.LookupEnv(EnvKeycloakHost)
	if !found {
		r.Keycloak.Host = "https://localhost:8081"
	}
	r.Keycloak.Realm, found = os.LookupEnv(EnvKeycloakRealm)
	if !found {
		r.Keycloak.Realm = "konveyor"
	}
	r.Keycloak.ClientID, found = os.LookupEnv(EnvKeycloakClientID)
	if !found {
		r.Keycloak.ClientID = "konveyor"
	}
	r.Keycloak.ClientSecret, found = os.LookupEnv(EnvKeycloakClientSecret)
	if !found {
		r.Keycloak.ClientSecret = ""
	}
	r.Keycloak.Admin.User, found = os.LookupEnv(EnvKeycloakAdminUser)
	if !found {
		r.Keycloak.Admin.User = "admin"
	}
	r.Keycloak.Admin.Pass, found = os.LookupEnv(EnvKeycloakAdminPass)
	if !found {
		r.Keycloak.Admin.Pass = "admin"
	}
	r.Keycloak.Admin.Realm, found = os.LookupEnv(EnvKeycloakAdminRealm)
	if !found {
		r.Keycloak.Admin.Realm = "master"
	}
	r.AddonToken, found = os.LookupEnv(EnvAddonToken)
	if !found {
		r.AddonToken = "konveyor"
	}
	r.RolePath, found = os.LookupEnv(EnvRolePath)
	if !found {
		r.RolePath = "/tmp/roles.yaml"
	}
	r.UserPath, found = os.LookupEnv(EnvUserPath)
	if !found {
		r.UserPath = "/tmp/users.yaml"
	}
	return
}
