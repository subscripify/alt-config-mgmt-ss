package tenantconfigmodel

type TenantConfigType string

const (
	LordConfig   TenantConfigType = "lord"   //can be used by lord tenant only
	SuperConfig  TenantConfigType = "super"  //can be used by super and lor tenant only
	PublicConfig TenantConfigType = "public" //can be used by any tenant
)
