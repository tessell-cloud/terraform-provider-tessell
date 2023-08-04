package model

type TessellServiceContextInfo struct {
	SubStatus   *string `json:"subStatus,omitempty"`
	Description *string `json:"description,omitempty"`
}

type TessellServiceClonedFromInfo struct {
	TessellServiceId      *string `json:"tessellServiceId,omitempty"`      // The DB Service Id using which this DB Service clone is created
	AvailabilityMachineId *string `json:"availabilityMachineId,omitempty"` // The Availability Machine Id using which this DB Service clone is created
	TessellService        *string `json:"tessellService,omitempty"`        // The DB Service name using which this DB Service clone is created
	AvailabilityMachine   *string `json:"availabilityMachine,omitempty"`   // The Availaility Machine name using which this DB Service clone is created
	SnapshotName          *string `json:"snapshotName,omitempty"`          // The snapshot using which this DB Service clone is created
	SnapshotId            *string `json:"snapshotId,omitempty"`            // The snapshot Id using which this DB Service clone is created
	SnapshotTime          *string `json:"snapshotTime,omitempty"`          // DB Service snapshot capture time
	PITRTime              *string `json:"pitrTime,omitempty"`              // If the database was created using a Point-In-Time mechanism, it specifies the timestamp in UTC
	MaximumRecoverability *bool   `json:"maximumRecoverability,omitempty"` // If the service was created using a maximum recoverablity from the parent service
}

type TessellServiceInfrastructureInfo struct {
	Cloud             *string            `json:"cloud,omitempty"`            // The cloud-type in which the DB Service is provisioned (ex. aws, azure)
	Region            *string            `json:"region,omitempty"`           // The region in which the DB Service provisioned
	AvailabilityZone  *string            `json:"availabilityZone,omitempty"` // The availability-zone in which the DB Service is provisioned
	CloudAvailability *[]CloudRegionInfo `json:"cloudAvailability,omitempty"`
	VPC               *string            `json:"vpc,omitempty"` // The VPC to be used for provisioning the DB Service
	EnableEncryption  *bool              `json:"enableEncryption,omitempty"`
	EncryptionKey     *string            `json:"encryptionKey,omitempty"`     // The encryption key name which is used to encrypt the data at rest
	ComputeType       *string            `json:"computeType,omitempty"`       // The compute-type to be used for provisioning the DB Service
	Storage           *int               `json:"storage,omitempty"`           // The storage (in bytes) that has been provisioned for the DB Service
	AdditionalStorage *int               `json:"additionalStorage,omitempty"` // Size in GB. This is maintained for backward compatibility and would be deprecated soon.
}

type TessellServiceConnectivityInfo struct {
	EnableSsl            *bool                                           `json:"enableSSL,omitempty"`
	CaCertId             *string                                         `json:"caCertId,omitempty"`
	DNSPrefix            *string                                         `json:"dnsPrefix,omitempty"`
	ServicePort          *int                                            `json:"servicePort,omitempty"`        // The connection port for the DB Service
	EnablePublicAccess   *bool                                           `json:"enablePublicAccess,omitempty"` // Specify whether to enable public access to the DB Service, default false
	AllowedIpAddresses   *[]string                                       `json:"allowedIpAddresses,omitempty"` // The list of allowed ipv4 addresses that can connect to the DB Service
	ConnectStrings       *[]TessellServiceConnectString                  `json:"connectStrings,omitempty"`     // The list of connect strings for the DB Service
	PrivateLink          *ServiceConnectivityPrivateLink                 `json:"privateLink,omitempty"`
	UpdateInProgressInfo *TessellServiceConnectivityUpdateInProgressInfo `json:"updateInProgressInfo,omitempty"`
}

type TessellServiceConnectString struct {
	Type              *string `json:"type,omitempty"`
	UsageType         *string `json:"usageType,omitempty"`
	ConnectDescriptor *string `json:"connectDescriptor,omitempty"`
	Endpoint          *string `json:"endpoint,omitempty"`
	MasterUser        *string `json:"masterUser,omitempty"`
	ServicePort       *int    `json:"servicePort,omitempty"` // The connection port for the DB Service
}

type ServiceConnectivityPrivateLink struct {
	Status                     *string   `json:"status,omitempty"`
	ServicePrincipals          *[]string `json:"servicePrincipals,omitempty"`          // The list of AWS account principals that are currently enabled
	EndpointServiceName        *string   `json:"endpointServiceName,omitempty"`        // The configured endpoint as a result of configuring the service-pricipals
	ClientAzureSubscriptionIds *[]string `json:"clientAzureSubscriptionIds,omitempty"` // The list of Azure subscription Ids
	PrivateLinkServiceAlias    *string   `json:"privateLinkServiceAlias,omitempty"`    // The Azure private link service alias
}

type TessellServiceConnectivityUpdateInProgressInfo struct {
	DNSPrefix          *string                                  `json:"dnsPrefix,omitempty"`
	EnablePublicAccess *bool                                    `json:"enablePublicAccess,omitempty"` // Specify whether to enable public access to the DB Service, default false
	AllowedIpAddresses *[]string                                `json:"allowedIpAddresses,omitempty"` // The list of allowed ipv4 addresses that can connect to the DB Service
	PrivateLink        *ServiceConnectivityUpdateInProgressInfo `json:"privateLink,omitempty"`
}

type ServiceConnectivityUpdateInProgressInfo struct {
	ServicePrincipals          *[]string `json:"servicePrincipals,omitempty"`          // The list of AWS account principals that are currently enabled
	ClientAzureSubscriptionIds *[]string `json:"clientAzureSubscriptionIds,omitempty"` // The list of Azure subscription Ids
}

type TessellServiceCredsPayload struct {
	MasterUser     *string `json:"masterUser"`     // DB Service&#39;s master username
	MasterPassword *string `json:"masterPassword"` // DB Service&#39;s master password
}

type TessellServiceMaintenanceWindow struct {
	Day      *string `json:"day"`
	Time     *string `json:"time"` // Time value in (hh:mm) format. ex. &#39;02:00&#39;
	Duration *int    `json:"duration"`
}

type TessellServiceBackupConfigurationPayload struct {
	AutoSnapshot   *bool                                                   `json:"autoSnapshot,omitempty"` // Specify whether to capture automated snapshots for the DB Service, default true.
	SLA            *string                                                 `json:"sla,omitempty"`          // The snapshot SLA for the DB Service. If not specified, a default SLA would be associated with the DB Service
	SnapshotWindow *TessellServiceBackupConfigurationPayloadSnapshotWindow `json:"snapshotWindow,omitempty"`
}

type TessellServiceBackupConfigurationPayloadSnapshotWindow struct {
	Time     *string `json:"time,omitempty"`     // Time value in (hh:mm) format. ex. &#39;02:00&#39;
	Duration *int    `json:"duration,omitempty"` // The allowed duration for capturing the DB Service backup
}

type TessellServiceEngineInfo struct {
	OracleConfig      *TessellServiceOracleEngineConfig      `json:"oracleConfig,omitempty"`
	PostgresqlConfig  *TessellServicePostgresqlEngineConfig  `json:"postgresqlConfig,omitempty"`
	MysqlConfig       *TessellServiceMysqlEngineConfig       `json:"mysqlConfig,omitempty"`
	SqlServerConfig   *TessellServiceSqlServerEngineConfig   `json:"sqlServerConfig,omitempty"`
	ApacheKafkaConfig *TessellServiceApacheKafkaEngineConfig `json:"apacheKafkaConfig,omitempty"`
	PreScriptInfo     *ScriptInfo                            `json:"preScriptInfo,omitempty"`
	PostScriptInfo    *ScriptInfo                            `json:"postScriptInfo,omitempty"`
}

type TessellServiceOracleEngineConfig struct {
	MultiTenant          *bool   `json:"multiTenant,omitempty"`          // Specify whether the DB Service is multi-tenant.
	ParameterProfileId   *string `json:"parameterProfileId,omitempty"`   // The parameter profile id for the database
	OptionsProfile       *string `json:"optionsProfile,omitempty"`       // The options profile for the database
	CharacterSet         *string `json:"characterSet,omitempty"`         // The character-set for the database
	NationalCharacterSet *string `json:"nationalCharacterSet,omitempty"` // The national-character-set for the database
}

type TessellServicePostgresqlEngineConfig struct {
	ParameterProfileId *string `json:"parameterProfileId,omitempty"` // The parameter profile id for the database
}

type TessellServiceMysqlEngineConfig struct {
	ParameterProfileId *string `json:"parameterProfileId,omitempty"` // The parameter profile id for the database
}

type TessellServiceSqlServerEngineConfig struct {
	ParameterProfileId *string `json:"parameterProfileId,omitempty"` // The parameter profile id for the database
	AdDomainId         *string `json:"adDomainId,omitempty"`         // Active Directory Domain id
}

type TessellServiceApacheKafkaEngineConfig struct {
	ParameterProfileId *string `json:"parameterProfileId,omitempty"` // The parameter profile id for the database
}

type ScriptInfo struct {
	ScriptId      *string `json:"scriptId,omitempty"`      // The Tessell Script Id
	ScriptVersion *string `json:"scriptVersion,omitempty"` // The Tessell Script version
}

type TerraformTessellDatabaseDTO struct {
	SourceDatabaseId      *string                        `json:"sourceDatabaseId,omitempty"` // Required while creating a clone. It specifies the Id of the source database from which the clone is being created.
	Id                    *string                        `json:"id,omitempty"`
	DatabaseName          *string                        `json:"databaseName,omitempty"`     // Database name
	Description           *string                        `json:"description,omitempty"`      // Database description
	TessellServiceId      *string                        `json:"tessellServiceId,omitempty"` // Associated DB Service Id
	EngineType            *string                        `json:"engineType,omitempty"`       // Database Engine Type
	Status                *string                        `json:"status,omitempty"`           // Database status
	DateCreated           *string                        `json:"dateCreated,omitempty"`      // Timestamp when the entity was created
	ClonedFromInfo        *TessellDatabaseClonedFromInfo `json:"clonedFromInfo,omitempty"`
	DatabaseConfiguration *DatabaseConfiguration         `json:"databaseConfiguration,omitempty"`
}

type TessellDatabaseClonedFromInfo struct {
	DatabaseId *string `json:"databaseId,omitempty"` // The original database Id using which this database clone is created
}

type DatabaseConfiguration struct {
	OracleConfig     *OracleDatabaseConfig     `json:"oracleConfig,omitempty"`
	PostgresqlConfig *PostgresqlDatabaseConfig `json:"postgresqlConfig,omitempty"`
	MysqlConfig      *MysqlDatabaseConfig      `json:"mySqlConfig,omitempty"`
	SqlServerConfig  *SqlServerDatabaseConfig  `json:"sqlServerConfig,omitempty"`
}

type OracleDatabaseConfig struct {
	ParameterProfileId *string `json:"parameterProfileId,omitempty"` // The parameter profile id for the database
	OptionsProfile     *string `json:"optionsProfile,omitempty"`     // The options profile for the database
}

type PostgresqlDatabaseConfig struct {
	ParameterProfileId *string `json:"parameterProfileId,omitempty"` // The parameter profile id for the database
}

type MysqlDatabaseConfig struct {
	ParameterProfileId *string `json:"parameterProfileId,omitempty"` // The parameter profile id for the database
}

type SqlServerDatabaseConfig struct {
	ParameterProfileId *string `json:"parameterProfileId,omitempty"` // The parameter profile id for the database
}

type TessellServiceIntegrationsPayload struct {
	Integrations *[]string `json:"integrations,omitempty"`
}

type TessellTag struct {
	Name  *string `json:"name,omitempty"`  // Case sensitive, tag name
	Value *string `json:"value,omitempty"` // Case sensitive, tag value
}

type TessellServiceInstanceDTO struct {
	Id                   *string                              `json:"id,omitempty"`   // Tessell generated UUID for the DB Service Instance
	Name                 *string                              `json:"name,omitempty"` // Name of the DB Service Instance
	Type                 *string                              `json:"type,omitempty"`
	Role                 *string                              `json:"role,omitempty"` // DB Service Topology
	Status               *string                              `json:"status,omitempty"`
	TessellServiceId     *string                              `json:"tessellServiceId,omitempty"` // DB Service Instance&#39;s associated DB Service id
	Cloud                *string                              `json:"cloud,omitempty"`            // DB Service Instance&#39;s cloud type
	Region               *string                              `json:"region,omitempty"`           // DB Service Instance&#39;s cloud region
	AvailabilityZone     *string                              `json:"availabilityZone,omitempty"` // DB Service Instance&#39;s cloud availability zone
	InstanceGroupId      *string                              `json:"instanceGroupId,omitempty"`  // The instance groupd Id
	ComputeType          *string                              `json:"computeType,omitempty"`      // The compute used for creation of the Tessell Service Instance
	Storage              *int                                 `json:"storage,omitempty"`          // The storage (in bytes) that has been provisioned for the DB Service instance.
	DataVolumeIops       *int                                 `json:"dataVolumeIops,omitempty"`
	ParameterProfile     *ParameterProfile                    `json:"parameterProfile,omitempty"`
	VPC                  *string                              `json:"vpc,omitempty"`                  // The VPC used for creation of the DB Service Instance
	EncryptionKey        *string                              `json:"encryptionKey,omitempty"`        // The encryption key name which is used to encrypt the data at rest
	SoftwareImage        *string                              `json:"softwareImage,omitempty"`        // Software Image to be used to create the instance
	SoftwareImageVersion *string                              `json:"softwareImageVersion,omitempty"` // Software Image Version to be used to create the instance
	DateCreated          *string                              `json:"dateCreated,omitempty"`          // Timestamp when the entity was created
	ConnectString        *TessellServiceInstanceConnectString `json:"connectString,omitempty"`
	UpdatesInProgress    *[]TessellResourceUpdateInfo         `json:"updatesInProgress,omitempty"` // The updates that are in progress for this resource
	LastStartedAt        *string                              `json:"lastStartedAt,omitempty"`     // Timestamp when the service instance was last started at
	LastStoppedAt        *string                              `json:"lastStoppedAt,omitempty"`     // Timestamp when the Service Instance was last stopped at
}

type ParameterProfile struct {
	Id      *string `json:"id,omitempty"`
	Name    *string `json:"name,omitempty"`
	Version *string `json:"version,omitempty"`
	Status  *string `json:"status,omitempty"`
}

type TessellServiceInstanceConnectString struct {
	ConnectDescriptor *string `json:"connectDescriptor,omitempty"`
	MasterUser        *string `json:"masterUser,omitempty"`
	Endpoint          *string `json:"endpoint,omitempty"`
	ServicePort       *string `json:"servicePort,omitempty"`
}

type TessellResourceUpdateInfo struct {
	UpdateType  *string                 `json:"updateType,omitempty"`  // Type of the update
	ReferenceId *string                 `json:"referenceId,omitempty"` // The reference-id of the update request
	SubmittedAt *string                 `json:"submittedAt,omitempty"` // Timestamp when the resource update was requested
	UpdateInfo  *map[string]interface{} `json:"updateInfo,omitempty"`  // The specific details for a Tessell resource that are being updated
}

type ServiceUpcomingScheduledActions struct {
	StartStop *ServiceUpcomingScheduledActionsStartStop `json:"startStop,omitempty"`
	Patch     *ServiceUpcomingScheduledActionsPatch     `json:"patch,omitempty"`
	Delete    *ServiceUpcomingScheduledActionsDelete    `json:"delete,omitempty"`
}

type ServiceUpcomingScheduledActionsStartStop struct {
	Action *string `json:"action,omitempty"` // Action which can be either start/stop
	At     *string `json:"at,omitempty"`
}

type ServiceUpcomingScheduledActionsPatch struct {
	At      *string `json:"at,omitempty"`
	Message *string `json:"message,omitempty"`
}

type ServiceUpcomingScheduledActionsDelete struct {
	At *string `json:"at,omitempty"`
}

type TessellServiceDeletionConfig struct {
	RetainAvailabilityMachine *bool `json:"retainAvailabilityMachine,omitempty"` // If specified as true, the associated Availability Machine (snapshots, sanitized-snapshots, logs) would be retained
}

type DeletionScheduleDTO struct {
	DeleteAt       *string                       `json:"deleteAt"` // DB Service deletion Time
	DeletionConfig *TessellServiceDeletionConfig `json:"deletionConfig,omitempty"`
}

type TerraformTessellServiceDTO struct {
	Id                         *string                                   `json:"id,omitempty"`                    // Tessell generated UUID for the DB Service. This is the unique identifier for the DB Service.
	AvailabilityMachineId      *string                                   `json:"availabilityMachineId,omitempty"` // Associated Availability Machine Id
	SnapshotId                 *string                                   `json:"snapshotId,omitempty"`            // Tessell service snapshot Id, using which the clone is to be created
	PITR                       *string                                   `json:"pitr,omitempty"`                  // PITR Timestamp, using which the clone is to be created
	Name                       *string                                   `json:"name"`                            // Name of the DB Service
	Description                *string                                   `json:"description,omitempty"`           // DB Service&#39;s description
	TenantId                   *string                                   `json:"tenantId,omitempty"`              // The tenant-id for the DB Service
	Subscription               *string                                   `json:"subscription"`                    // Tessell Subscription in which the DB Service is to be created
	EngineType                 *string                                   `json:"engineType"`
	Topology                   *string                                   `json:"topology"`
	NumOfInstances             *int                                      `json:"numOfInstances,omitempty"` // Number of instance (nodes) to be created for the DB Service. This is a required input for Apache Kafka. For all other engines, this input would be ignored even if specified.
	Status                     *string                                   `json:"status,omitempty"`
	ContextInfo                *TessellServiceContextInfo                `json:"contextInfo,omitempty"`
	LicenseType                *string                                   `json:"licenseType,omitempty"` // DB Service License Type
	Edition                    *string                                   `json:"edition,omitempty"`
	SoftwareImage              *string                                   `json:"softwareImage"`                        // Software Image to be used to create the DB Service
	SoftwareImageVersion       *string                                   `json:"softwareImageVersion"`                 // Software Image Version to be used to create the DB Service
	SoftwareImageVersionFamily *string                                   `json:"softwareImageVersionFamily,omitempty"` // Software Image Family DB Service belongs to
	AutoMinorVersionUpdate     *bool                                     `json:"autoMinorVersionUpdate,omitempty"`     // Specify whether to automatically update minor version for DB Service
	EnableDeletionProtection   *bool                                     `json:"enableDeletionProtection,omitempty"`   // Specify whether to enable deletion protection for the DB Service
	EnableStopProtection       *bool                                     `json:"enableStopProtection,omitempty"`       // This field specifies whether to enable stop protection for the DB Service. If this is enabled, the stop for the DB Service would be disallowed until this setting is disabled.
	Owner                      *string                                   `json:"owner,omitempty"`                      // DB Service owner email address
	LoggedInUserRole           *string                                   `json:"loggedInUserRole,omitempty"`           // Access role for the currently logged in user
	DateCreated                *string                                   `json:"dateCreated,omitempty"`                // Timestamp when the DB Service was created at
	StartedAt                  *string                                   `json:"startedAt,omitempty"`                  // Timestamp when the DB Service was last started at
	StoppedAt                  *string                                   `json:"stoppedAt,omitempty"`                  // Timestamp when the DB Service was last stopped at
	ClonedFromInfo             *TessellServiceClonedFromInfo             `json:"clonedFromInfo,omitempty"`
	Infrastructure             *TessellServiceInfrastructureInfo         `json:"infrastructure"`
	ServiceConnectivity        *TessellServiceConnectivityInfo           `json:"serviceConnectivity"`
	TessellGenieStatus         *string                                   `json:"tessellGenieStatus,omitempty"` // DB Service&#39;s Genie status
	Creds                      *TessellServiceCredsPayload               `json:"creds"`
	MaintenanceWindow          *TessellServiceMaintenanceWindow          `json:"maintenanceWindow,omitempty"`
	SnapshotConfiguration      *TessellServiceBackupConfigurationPayload `json:"snapshotConfiguration,omitempty"`
	EngineConfiguration        *TessellServiceEngineInfo                 `json:"engineConfiguration"`
	Databases                  *[]TerraformTessellDatabaseDTO            `json:"databases,omitempty"` // Databases that are part of this DB Service
	IntegrationsConfig         *TessellServiceIntegrationsPayload        `json:"integrationsConfig,omitempty"`
	Tags                       *[]TessellTag                             `json:"tags,omitempty"`      // The tags to be associated with the DB Service
	Instances                  *[]TessellServiceInstanceDTO              `json:"instances,omitempty"` // Instances associated with this DB Service
	SharedWith                 *EntityAclSharingInfo                     `json:"sharedWith,omitempty"`
	UpcomingScheduledActions   *ServiceUpcomingScheduledActions          `json:"upcomingScheduledActions,omitempty"`
	DeletionConfig             *TessellServiceDeletionConfig             `json:"deletionConfig,omitempty"`
	DeletionSchedule           *DeletionScheduleDTO                      `json:"deletionSchedule,omitempty"`
	UpdatesInProgress          *[]TessellResourceUpdateInfo              `json:"updatesInProgress,omitempty"` // The updates that are in progress for this resource
}

type CloneTessellServicePayload struct {
	SnapshotId               *string                                   `json:"snapshotId,omitempty"`  // Tessell service snapshot Id, using which the clone is to be created
	PITR                     *string                                   `json:"pitr,omitempty"`        // PITR Timestamp, using which the clone is to be created
	Name                     *string                                   `json:"name"`                  // DB Service name
	Description              *string                                   `json:"description,omitempty"` // DB Service&#39;s description
	Subscription             *string                                   `json:"subscription"`          // Tessell Subscription in which the DB Service is to be created
	Edition                  *string                                   `json:"edition,omitempty"`
	EngineType               *string                                   `json:"engineType"`
	Topology                 *string                                   `json:"topology"`
	NumOfInstances           *int                                      `json:"numOfInstances,omitempty"`           // Number of instance (nodes) to be created for the DB Service. This is a required input for Apache Kafka. For all other engines, this input would be ignored even if specified.
	SoftwareImage            *string                                   `json:"softwareImage"`                      // Software Image to be used to create the DB Service
	SoftwareImageVersion     *string                                   `json:"softwareImageVersion"`               // Software Image Version to be used to create the DB Service
	AutoMinorVersionUpdate   *bool                                     `json:"autoMinorVersionUpdate,omitempty"`   // Specify whether to automatically update minor version for DB Service
	EnableDeletionProtection *bool                                     `json:"enableDeletionProtection,omitempty"` // Specify whether to enable deletion protection for the DB Service
	EnableStopProtection     *bool                                     `json:"enableStopProtection,omitempty"`     // Specify whether to enable stop protection for the DB Service
	Infrastructure           *TessellServiceInfrastructurePayload      `json:"infrastructure"`
	ServiceConnectivity      *TessellServiceConnectivityInfoPayload    `json:"serviceConnectivity"`
	Creds                    *TessellServiceCredsPayload               `json:"creds"`
	MaintenanceWindow        *TessellServiceMaintenanceWindow          `json:"maintenanceWindow,omitempty"`
	DeletionConfig           *TessellServiceDeletionConfig             `json:"deletionConfig,omitempty"`
	SnapshotConfiguration    *TessellServiceBackupConfigurationPayload `json:"snapshotConfiguration,omitempty"`
	EngineConfiguration      *TessellServiceEngineConfigurationPayload `json:"engineConfiguration"`
	Databases                *[]CreateDatabasePayload                  `json:"databases,omitempty"` // Specify the databases to be created in the DB Service
	IntegrationsConfig       *TessellServiceIntegrationsPayload        `json:"integrationsConfig,omitempty"`
	Tags                     *[]TessellTag                             `json:"tags,omitempty"` // The tags to be associated with the DB Service
}

type TessellServiceInfrastructurePayload struct {
	Cloud             *string `json:"cloud"`                       // The cloud-type in which the DB Service is to be provisioned (ex. aws, azure)
	Region            *string `json:"region"`                      // The region in which the DB Service is to be provisioned
	AvailabilityZone  *string `json:"availabilityZone,omitempty"`  // The availability-zone in which the DB Service is to be provisioned
	VPC               *string `json:"vpc"`                         // The VPC to be used for provisioning the DB Service
	EnableEncryption  *bool   `json:"enableEncryption,omitempty"`  // Specify whether to enable the encryption at rest for the DB Service.
	EncryptionKey     *string `json:"encryptionKey,omitempty"`     // The encryption key name which is to be used to encrypt the data at rest. This is honoured only if &#39;enableEncryption&#39; is true. If this is not specified, Tessell will use a default out-of-the-box encryption key.
	ComputeType       *string `json:"computeType"`                 // The compute-type to be used for provisioning the DB Service
	AdditionalStorage *int    `json:"additionalStorage,omitempty"` // The additional storage (in bytes) to be provisioned for the DB Service. This is in addition to what is specified in the compute type.
}

type TessellServiceConnectivityInfoPayload struct {
	EnableSsl          *bool     `json:"enableSSL,omitempty"`
	DNSPrefix          *string   `json:"dnsPrefix,omitempty"`          // If not specified, Tessell will use a randomly generated prefix
	ServicePort        *int      `json:"servicePort"`                  // The connection port for the DB Service
	EnablePublicAccess *bool     `json:"enablePublicAccess,omitempty"` // Specify whether to enable public access to the DB Service, default false
	AllowedIpAddresses *[]string `json:"allowedIpAddresses,omitempty"` // The list of allowed ipv4 addresses that can connect to the DB Service
}

type TessellServiceEngineConfigurationPayload struct {
	PreScriptInfo     *ScriptInfo                     `json:"preScriptInfo,omitempty"`
	PostScriptInfo    *ScriptInfo                     `json:"postScriptInfo,omitempty"`
	OracleConfig      *OracleEngineConfigPayload      `json:"oracleConfig,omitempty"`
	PostgresqlConfig  *PostgresqlEngineConfigPayload  `json:"postgresqlConfig,omitempty"`
	MysqlConfig       *MysqlEngineConfigPayload       `json:"mysqlConfig,omitempty"`
	SqlServerConfig   *SqlServerEngineConfigPayload   `json:"sqlServerConfig,omitempty"`
	ApacheKafkaConfig *ApacheKafkaEngineConfigPayload `json:"apacheKafkaConfig,omitempty"`
}

type OracleEngineConfigPayload struct {
	MultiTenant          *bool   `json:"multiTenant,omitempty"` // Specify whether the DB Service is multi-tenant.
	ParameterProfileId   *string `json:"parameterProfileId"`    // The parameter profile id for the database
	OptionsProfile       *string `json:"optionsProfile"`        // The options profile for the database
	CharacterSet         *string `json:"characterSet"`          // The character-set for the database
	NationalCharacterSet *string `json:"nationalCharacterSet"`  // The national-character-set for the database
}

type PostgresqlEngineConfigPayload struct {
	ParameterProfileId *string `json:"parameterProfileId"` // The parameter profile id for the database
}

type MysqlEngineConfigPayload struct {
	ParameterProfileId *string `json:"parameterProfileId"` // The parameter profile id for the database
}

type SqlServerEngineConfigPayload struct {
	ParameterProfileId *string `json:"parameterProfileId"`   // The parameter profile id for the database
	AdDomainId         *string `json:"adDomainId,omitempty"` // Active Directory Domain id
}

type ApacheKafkaEngineConfigPayload struct {
	ParameterProfileId *string `json:"parameterProfileId,omitempty"` // The parameter profile id for the database
}

type CreateDatabasePayload struct {
	DatabaseName          *string                                     `json:"databaseName"`
	SourceDatabaseId      *string                                     `json:"sourceDatabaseId,omitempty"` // Required while creating a clone. It specifies the Id of the source database from which the clone is being created.
	DatabaseConfiguration *CreateDatabasePayloadDatabaseConfiguration `json:"databaseConfiguration,omitempty"`
}

type CreateDatabasePayloadDatabaseConfiguration struct {
	OracleConfig     *OracleDatabaseConfig     `json:"oracleConfig,omitempty"`
	PostgresqlConfig *PostgresqlDatabaseConfig `json:"postgresqlConfig,omitempty"`
	MysqlConfig      *MysqlDatabaseConfig      `json:"mysqlConfig,omitempty"`
	SqlServerConfig  *SqlServerDatabaseConfig  `json:"sqlServerConfig,omitempty"`
}

type DeleteTessellServicePayload struct {
	DeletionConfig  *TessellServiceDeletionConfig `json:"deletionConfig,omitempty"`
	Comment         *string                       `json:"comment,omitempty"` // Comment for the the action
	PublishEventLog *bool                         `json:"publishEventLog,omitempty"`
}

type TessellServiceDTO struct {
	Id                         *string                           `json:"id,omitempty"`                    // Tessell generated UUID for the DB Service. This is the unique identifier for the DB Service.
	AvailabilityMachineId      *string                           `json:"availabilityMachineId,omitempty"` // Unique id of the associated Availability Machine
	Name                       *string                           `json:"name,omitempty"`                  // Name of the DB Service
	Description                *string                           `json:"description,omitempty"`           // User specified description for the DB Service
	EngineType                 *string                           `json:"engineType,omitempty"`
	Topology                   *string                           `json:"topology,omitempty"`
	NumOfInstances             *int                              `json:"numOfInstances,omitempty"` // This specifies the number of instances (nodes) that are created for the DB Service
	Status                     *string                           `json:"status,omitempty"`
	ContextInfo                *TessellServiceContextInfo        `json:"contextInfo,omitempty"`
	LicenseType                *string                           `json:"licenseType,omitempty"`
	AutoMinorVersionUpdate     *bool                             `json:"autoMinorVersionUpdate,omitempty"`   // This field specifies whether to automatically update minor version for the DB Service
	EnableDeletionProtection   *bool                             `json:"enableDeletionProtection,omitempty"` // This field specifies whether to enable deletion protection for the DB Service. If this is enabled, the deletion for the DB Service would be disallowed until this setting is disabled.
	EnableStopProtection       *bool                             `json:"enableStopProtection,omitempty"`     // This field specifies whether to enable stop protection for the DB Service. If this is enabled, the stop for the DB Service would be disallowed until this setting is disabled.
	Edition                    *string                           `json:"edition,omitempty"`
	SoftwareImage              *string                           `json:"softwareImage,omitempty"`              // The software image that has been used to create the DB Service
	SoftwareImageVersion       *string                           `json:"softwareImageVersion,omitempty"`       // The software image version that is used to create the DB Service
	SoftwareImageVersionFamily *string                           `json:"softwareImageVersionFamily,omitempty"` // The software image version family that is used to create the DB Service
	TenantId                   *string                           `json:"tenantId,omitempty"`                   // The tenant identifier under which the DB Service is created
	Subscription               *string                           `json:"subscription,omitempty"`               // The Tessell Subscription under which this DB Service is created
	Owner                      *string                           `json:"owner,omitempty"`                      // This field specifies who is the owner for the DB Service
	LoggedInUserRole           *string                           `json:"loggedInUserRole,omitempty"`           // This field specifies access role on the DB Service for the currently logged in user
	DateCreated                *string                           `json:"dateCreated,omitempty"`                // This field specifies the timestamp when the DB Service was created at
	StartedAt                  *string                           `json:"startedAt,omitempty"`                  // This field specifies the timestamp when the DB Service was last started at
	StoppedAt                  *string                           `json:"stoppedAt,omitempty"`                  // This field specifies the timestamp when the DB Service was last stopped at
	ClonedFromInfo             *TessellServiceClonedFromInfo     `json:"clonedFromInfo,omitempty"`
	ServiceConnectivity        *TessellServiceConnectivityInfo   `json:"serviceConnectivity,omitempty"`
	TessellGenieStatus         *string                           `json:"tessellGenieStatus,omitempty"`
	Infrastructure             *TessellServiceInfrastructureInfo `json:"infrastructure,omitempty"`
	MaintenanceWindow          *TessellServiceMaintenanceWindow  `json:"maintenanceWindow,omitempty"`
	EngineConfiguration        *TessellServiceEngineInfo         `json:"engineConfiguration,omitempty"`
	IntegrationsConfig         *TessellServiceIntegrationsInfo   `json:"integrationsConfig,omitempty"`
	DeletionConfig             *TessellServiceDeletionConfig     `json:"deletionConfig,omitempty"`
	Tags                       *[]TessellTag                     `json:"tags,omitempty"`              // The tags associated with the DB Service
	UpdatesInProgress          *[]TessellResourceUpdateInfo      `json:"updatesInProgress,omitempty"` // The updates that are in progress for this resource
	Instances                  *[]TessellServiceInstanceDTO      `json:"instances,omitempty"`         // The instances (nodes) associated with this DB Service
	Databases                  *[]TessellDatabaseDTO             `json:"databases,omitempty"`         // This field details about the databases that are created under this DB Service
	SharedWith                 *EntityAclSharingInfo             `json:"sharedWith,omitempty"`
	DeletionSchedule           *DeletionScheduleDTO              `json:"deletionSchedule,omitempty"`
	UpcomingScheduledActions   *ServiceUpcomingScheduledActions  `json:"upcomingScheduledActions,omitempty"`
}

type TessellServiceIntegrationsInfo struct {
	Integrations *[]string `json:"integrations,omitempty"`
}

type TessellDatabaseDTO struct {
	Id                    *string                        `json:"id,omitempty"`
	DatabaseName          *string                        `json:"databaseName,omitempty"`     // Database name
	Description           *string                        `json:"description,omitempty"`      // Database description
	TessellServiceId      *string                        `json:"tessellServiceId,omitempty"` // Associated DB Service Id
	EngineType            *string                        `json:"engineType,omitempty"`       // Database Engine Type
	Status                *string                        `json:"status,omitempty"`           // Database status
	DateCreated           *string                        `json:"dateCreated,omitempty"`      // Timestamp when the entity was created
	ClonedFromInfo        *TessellDatabaseClonedFromInfo `json:"clonedFromInfo,omitempty"`
	DatabaseConfiguration *DatabaseConfiguration         `json:"databaseConfiguration,omitempty"`
}

type TessellServicesResponse struct {
	Metadata *APIMetadata         `json:"metadata,omitempty"`
	Response *[]TessellServiceDTO `json:"response,omitempty"`
}

type ProvisionTessellServicePayload struct {
	Name                     *string                                   `json:"name"`                  // DB Service name
	Description              *string                                   `json:"description,omitempty"` // DB Service&#39;s description
	Subscription             *string                                   `json:"subscription"`          // Tessell Subscription in which the DB Service is to be created
	Edition                  *string                                   `json:"edition,omitempty"`
	EngineType               *string                                   `json:"engineType"`
	Topology                 *string                                   `json:"topology"`
	NumOfInstances           *int                                      `json:"numOfInstances,omitempty"`           // Number of instance (nodes) to be created for the DB Service. This is a required input for Apache Kafka. For all other engines, this input would be ignored even if specified.
	SoftwareImage            *string                                   `json:"softwareImage"`                      // Software Image to be used to create the DB Service
	SoftwareImageVersion     *string                                   `json:"softwareImageVersion"`               // Software Image Version to be used to create the DB Service
	AutoMinorVersionUpdate   *bool                                     `json:"autoMinorVersionUpdate,omitempty"`   // Specify whether to automatically update minor version for DB Service
	EnableDeletionProtection *bool                                     `json:"enableDeletionProtection,omitempty"` // Specify whether to enable deletion protection for the DB Service
	EnableStopProtection     *bool                                     `json:"enableStopProtection,omitempty"`     // Specify whether to enable stop protection for the DB Service
	Infrastructure           *TessellServiceInfrastructurePayload      `json:"infrastructure"`
	ServiceConnectivity      *TessellServiceConnectivityInfoPayload    `json:"serviceConnectivity"`
	Creds                    *TessellServiceCredsPayload               `json:"creds"`
	MaintenanceWindow        *TessellServiceMaintenanceWindow          `json:"maintenanceWindow,omitempty"`
	DeletionConfig           *TessellServiceDeletionConfig             `json:"deletionConfig,omitempty"`
	SnapshotConfiguration    *TessellServiceBackupConfigurationPayload `json:"snapshotConfiguration,omitempty"`
	EngineConfiguration      *TessellServiceEngineConfigurationPayload `json:"engineConfiguration"`
	Databases                *[]CreateDatabasePayload                  `json:"databases,omitempty"` // Specify the databases to be created in the DB Service
	IntegrationsConfig       *TessellServiceIntegrationsPayload        `json:"integrationsConfig,omitempty"`
	Tags                     *[]TessellTag                             `json:"tags,omitempty"` // The tags to be associated with the DB Service
}

type StartTessellServicePayload struct {
	Comment *string `json:"comment,omitempty"` // Comment for the the action
}

type StopTessellServicePayload struct {
	Comment *string `json:"comment,omitempty"` // Comment for the the action
}
