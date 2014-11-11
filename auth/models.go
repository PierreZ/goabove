package auth

// keyring stores all auth parameters
type keyring struct {
	raCk         string // runabove consumer key
	osUsername   string // OpenStack username
	osPassword   string // Openstack passsword
	osTenantName string // Openstack tenant name
	osAuthUrl    string // Openstack auth url
	osTenantId   string // Openstack tenant id
	osRegionName string // Openstack region
}

// raApiCredentials are credentials returned by auth/credential
type raApiCredentials struct {
	ValidationUrl string
	ConsumerKey   string
	State         string
}
