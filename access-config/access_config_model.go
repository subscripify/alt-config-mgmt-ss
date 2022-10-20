package accessconfig

import (
	"fmt"
	"regexp"
	"time"

	goaway "github.com/TwiN/go-away"
	"github.com/google/uuid"
)

type tenantConfig struct {
	accessConfigUUID  uuid.UUID
	configAlias       string
	configLocation    string
	configOwnerTenant uuid.UUID
	createdBy         string
	accessType        AccessConfigType
	createTimestamp   time.Time
}

type AccessConfigType string

const (
	PrivateConfig AccessConfigType = "private" // can be used by super tenants to access super services as assigned by the lord tenants
	CustomConfig  AccessConfigType = "custom"  // can be used by the main tenants to access public services and custom services developed by the super tenants

)

// creates a brand new UUID for a config
func (tc *tenantConfig) setNewAccessConfigUUID() {
	tc.accessConfigUUID = uuid.New()
}

// this will set the config uuid to the value in the string. the accessConfigUUID sent in must be a UUID in the form ########-####-####-####-############ and must parse as a valid UUID
func (tc *tenantConfig) setAccessConfigUUID(accessConfigUUID string) error {
	loadedUUID, err := uuid.Parse(accessConfigUUID)
	if err != nil {
		return fmt.Errorf("the tenant UUID failed to parse: %s", err)
	}
	tc.accessConfigUUID = loadedUUID
	return nil
}

// returns the current access config uuid
func (tc *tenantConfig) getAccessConfigUUID() uuid.UUID {
	return tc.accessConfigUUID
}

// the alias name does not need to be unique and is alias used for quick reference when searching. no starting spaces and watch your language
func (tc *tenantConfig) setAlias(alias string) error {
	// no spaces, special characters, or swear words
	profanityDetector := goaway.NewProfanityDetector().WithSanitizeLeetSpeak(false).WithSanitizeSpecialCharacters(true).WithSanitizeAccents(false)
	if profanityDetector.IsProfane(alias) {
		err := fmt.Errorf("this is not a valid alias name, only a-z A-Z 1-9 - and spaces allowed")
		return err
	}
	pattern := `^([a-zA-Z0-9]|(?:[a-zA-Z0-9]+[a-zA-Z0-9.\s\-]*[a-zA-Z0-9]+))$`
	r := regexp.MustCompile(pattern)
	if !r.MatchString(alias) {
		err := fmt.Errorf(`this is not a valid alias name `)
		return err
	}

	tc.configAlias = alias
	return nil
}

// returns the config alias
func (tc *tenantConfig) getAlias() string {
	return tc.configAlias
}

// this value will constrain the apply of this config to the owner and any of its vassal tenants. a valid uuid will be allowed here but if it does not match a valid tenant it will fail on INSERT or UPDATE to the database
func (tc *tenantConfig) setConfigOwnerUUID(ownerUUID string) error {
	loadedUUID, err := uuid.Parse(ownerUUID)
	if err != nil {
		return fmt.Errorf("the tenant UUID failed to parse: %s", err)
	}
	tc.configOwnerTenant = loadedUUID
	return nil
}

// returns the config owner uuid
func (tc *tenantConfig) getConfigOwnerUUID() uuid.UUID {
	return tc.configOwnerTenant
}

// id from identity system - must validate as email of current user session
func (tc *tenantConfig) setCreatedBy(userIdentifier string) error {
	tc.createdBy = userIdentifier
	return nil
}

// returns the creator of the config
func (tc *tenantConfig) getCreator() string {
	return tc.createdBy
}

// sets the access type for the config - custom or private
func (tc *tenantConfig) setTenantType(accessConfigType AccessConfigType) error {

	switch accessConfigType {
	case CustomConfig:
		tc.accessType = CustomConfig
	case PrivateConfig:
		tc.accessType = PrivateConfig
	default:
		err := fmt.Errorf("invalid access type")
		return err
	}
	return nil
}

// returns a string indicating the access type
func (tc *tenantConfig) getAccessType() string {
	return string(tc.accessType)
}

// returns the creation time of the tenant as recorded in the database - this value is empty for tenants yet to be inserted into the database
func (tc *tenantConfig) getCreateTime() time.Time {
	return tc.createTimestamp
}
