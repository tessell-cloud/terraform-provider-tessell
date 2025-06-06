package model

type DBServiceTopology struct {
	Type              *string   `json:"type,omitempty"`
	CloudType         *string   `json:"cloudType,omitempty"`
	Region            *string   `json:"region,omitempty"`
	AvailabilityZones *[]string `json:"availabilityZones,omitempty"`
}

type TessellDAPServiceDTO struct {
	Id                                *string                              `json:"id,omitempty"`                    // ID of the Access Policy
	Name                              *string                              `json:"name,omitempty"`                  // Name of the Access Policy
	AvailabilityMachineId             *string                              `json:"availabilityMachineId,omitempty"` // ID of the Availability Machine
	TessellServiceId                  *string                              `json:"tessellServiceId,omitempty"`      // ID of the associated DB Service
	ServiceName                       *string                              `json:"serviceName,omitempty"`           // Name of the associated DB Service
	EngineType                        *string                              `json:"engineType,omitempty"`            // Database engine type of the associated DB Service
	ContentType                       *string                              `json:"contentType,omitempty"`           // Content Type for the Data Access Policy
	Status                            *string                              `json:"status,omitempty"`                // Database Access Policy Status
	ContentInfo                       *DAPContentInfo                      `json:"contentInfo,omitempty"`
	DataAccessConfig                  *DAPRetentionInfo                    `json:"dataAccessConfig,omitempty"`
	Owner                             *string                              `json:"owner,omitempty"`                             // Owner of the Access Policy
	LoggedInUserRole                  *string                              `json:"loggedInUserRole,omitempty"`                  // The role of the logged in user for accessing the Availability Machine
	SubscriptionsCloudLocationsAndKey *[]SubscriptionsCloudLocationsAndKey `json:"subscriptionsCloudLocationsAndKey,omitempty"` // The subscription, cloud and region information along with encryption key and user info for DAP
	DateCreated                       *string                              `json:"dateCreated,omitempty"`                       // Timestamp when this Access Policy was created at
	DateModified                      *string                              `json:"dateModified,omitempty"`                      // Timestamp when this Access Policy was last updated at
}

type DAPContentInfo struct {
	AsIsContent      *AsIsDAPContent         `json:"asIsContent,omitempty"`
	SanitizedContent *SanitizationDAPContent `json:"sanitizedContent,omitempty"`
	BackupContent    *BackupDAPContent       `json:"backupContent,omitempty"`
}

type AsIsDAPContent struct {
	Automated *bool            `json:"automated,omitempty"` // Share the automated as-is snapshots. This is exclusive with manual specification.
	Manual    *[]DAPManualInfo `json:"manual,omitempty"`    // The list of snapshots that are to be shared as part of this access policy
}

type DAPManualInfo struct {
	Id           *string `json:"id,omitempty"`           // The DB Service snapshot id
	Name         *string `json:"name,omitempty"`         // The DB Service snapshot name
	CreationTime *string `json:"creationTime,omitempty"` // DB Service snapshot capture time
	SharedAt     *string `json:"sharedAt,omitempty"`     // The timestamp when the snapshot was added to DAP for sharing
}

type SanitizationDAPContent struct {
	Automated *SanitizationDAPContentAutomated `json:"automated,omitempty"`
	Manual    *[]DAPManualInfo                 `json:"manual,omitempty"` // The list of sanitized snapshots that are to be shared as part of this access policy
}

type SanitizationDAPContentAutomated struct {
	SanitizationScheduleId *string `json:"sanitizationScheduleId"` // Id of the sanitization schedule to process automated backups, required only if contentType = Sanitized.
}

type BackupDAPContent struct {
	Automated *bool            `json:"automated,omitempty"` // Share the automated backups. This is exclusive with manual specification.
	Manual    *[]DAPManualInfo `json:"manual,omitempty"`    // The list of backups that are to be shared as part of this access policy
}

type DAPRetentionInfo struct {
	PITR         *int `json:"pitr,omitempty"`         // Retention time (in days) for Point-In-Time recoverability
	DailyBackups *int `json:"dailyBackups,omitempty"` // Retention time (in days) to retain daily snapshots
}

type SubscriptionsCloudLocationsAndKey struct {
	SubscriptionName  *string                             `json:"subscriptionName"`
	CloudRegionAndKey *map[string][]RegionToEncryptionKey `json:"cloudRegionAndKey"`
	Users             *[]string                           `json:"users,omitempty"` // List of users email id who have access to the data/content managed by this Access Policy
}

type RegionToEncryptionKey struct {
	Region            *string `json:"region,omitempty"`
	EncryptionKeyName *string `json:"encryptionKeyName,omitempty"`
}

type TessellCloneSummaryInfo struct {
	Id                *string                         `json:"id,omitempty"`
	Name              *string                         `json:"name"`                   // Name of the clone database
	Subscription      *string                         `json:"subscription,omitempty"` // Clone&#39;s subscription name
	ComputeType       *string                         `json:"computeType,omitempty"`  // Clone&#39;s compute type
	Status            *string                         `json:"status,omitempty"`       // Status of the clone database
	CloudAvailability *[]CloudRegionInfo              `json:"cloudAvailability,omitempty"`
	CloneInfo         *map[string]string              `json:"cloneInfo,omitempty"`   // Miscellaneous information
	Owner             *string                         `json:"owner,omitempty"`       // The user who created database clone
	Instances         *[]TessellServiceInstanceOpsDTO `json:"instances,omitempty"`   // Instances associated with this DB Service
	DateCreated       *string                         `json:"dateCreated,omitempty"` // Timestamp when the entity was created
}

type TessellServiceInstanceOpsDTO struct {
	Id                    *string                                        `json:"id,omitempty"`                 // Tessell generated UUID for the DB Service Instance
	Name                  *string                                        `json:"name,omitempty"`               // Name of the DB Service Instance
	ComputeName           *string                                        `json:"computeName,omitempty"`        // compute-name of the DB Service Instance on Cloud
	Description           *string                                        `json:"description,omitempty"`        // DB Service Instance description
	TessellServiceId      *string                                        `json:"tessellServiceId,omitempty"`   // DB Service Instance&#39;s associated DB Service ID
	ComputeResourceId     *string                                        `json:"computeResourceId,omitempty"`  // Associated compute resource ID
	CloudLocationId       *string                                        `json:"cloudLocationId,omitempty"`    // DB Service Instance&#39;s cloud location
	ParameterProfileId    *string                                        `json:"parameterProfileId,omitempty"` // Parameter Profile linked with the DB service instance
	CloudAccountId        *string                                        `json:"cloudAccountId,omitempty"`     // The cloud account on which the instance is hosted
	InstanceGroupId       *string                                        `json:"instanceGroupId,omitempty"`    // The instance group Id
	Type                  *string                                        `json:"type,omitempty"`
	Role                  *string                                        `json:"role,omitempty"`
	UserVisibleRole       *string                                        `json:"userVisibleRole,omitempty"`
	Status                *string                                        `json:"status,omitempty"`
	PluginStatus          *string                                        `json:"pluginStatus,omitempty"`
	ConnectionInfo        *TessellServiceInstanceConnectionInfo          `json:"connectionInfo,omitempty"`
	GenericInfo           *ServiceInstanceGenericInfo                    `json:"genericInfo,omitempty"`
	LicenseInfo           *DBLicenseInfo                                 `json:"licenseInfo,omitempty"`
	MonitoringConfig      *MonitoringConfig                              `json:"monitoringConfig,omitempty"`
	DateCreated           *string                                        `json:"dateCreated,omitempty"`        // Timestamp when the entity was created
	DateModified          *string                                        `json:"dateModified,omitempty"`       // Timestamp when the entity was last modified, either by system or by user
	DateModifieDByUser    *string                                        `json:"dateModifiedByUser,omitempty"` // Timestamp when the entity was last modified by the user
	Metadata              *TessellServiceInstanceMetadata                `json:"metadata,omitempty"`
	DriverInfo            *TessellServiceInstanceDriverInfo              `json:"driverInfo,omitempty"`
	UpdatesInProgressInfo *TessellServiceInstanceInProgressUpdateInfoOps `json:"updatesInProgressInfo,omitempty"`
	LastStartedAt         *string                                        `json:"lastStartedAt,omitempty"`    // Timestamp when the service instance was last started at
	LastStoppedAt         *string                                        `json:"lastStoppedAt,omitempty"`    // Timestamp when the Service Instance was last stopped at
	LastDegradedAt        *string                                        `json:"lastDegradedAt,omitempty"`   // Timestamp when the Service Instance was DEGRADED
	DeletedForUserAt      *string                                        `json:"deletedForUserAt,omitempty"` // Timestamp when the service instance was marked &#39;delete for user&#39;.
	IsConsumable          *bool                                          `json:"isConsumable,omitempty"`     // Whether the service instance is consumable for purposes like billing
	TessellAgentLcmInfo   *TessellAgentLcmInfo                           `json:"tessellAgentLcmInfo,omitempty"`
	ComputeResource       *TessellComputeResourceOpsDTO                  `json:"computeResource,omitempty"`
}

type TessellServiceInstanceConnectionInfo struct {
	ConnectString *TessellServiceInstanceConnectString             `json:"connectString,omitempty"`
	EndPoints     *[]TessellServiceInstanceConnectionInfoEndPoints `json:"endPoints,omitempty"`
	Data          *map[string]interface{}                          `json:"data,omitempty"`
}

type TessellServiceInstanceConnectionInfoEndPoints struct {
	Endpoint *string                 `json:"endpoint,omitempty"`
	Labels   *[]string               `json:"labels,omitempty"`
	Data     *map[string]interface{} `json:"data,omitempty"`
}

type ServiceInstanceGenericInfo struct {
	PartOfInitialPrimarySet *bool                      `json:"partOfInitialPrimarySet,omitempty"`
	EncryptionKey           *string                    `json:"encryptionKey,omitempty"`    // The encryption key name which is used to encrypt the data at rest
	EncryptionKeyId         *string                    `json:"encryptionKeyId,omitempty"`  // The encryption key id which is used to encrypt the data at rest
	ServerCertId            *string                    `json:"serverCertId,omitempty"`     // The CA certificate id which is configured for this instance
	VPC                     *string                    `json:"vpc,omitempty"`              // The VPC to be used for provisioning the instance
	VPCId                   *string                    `json:"vpcId,omitempty"`            // The VPC Id which is used for provisioning the instance
	PublicSubnet            *string                    `json:"publicSubnet,omitempty"`     // The public subnet used for provisioning the instance
	PublicSubnetId          *string                    `json:"publicSubnetId,omitempty"`   // The public subnet Id which is used for provisioning the instance
	PrivateSubnet           *string                    `json:"privateSubnet,omitempty"`    // The private subnet used for provisioning the instance
	PrivateSubnetId         *string                    `json:"privateSubnetId,omitempty"`  // The private subnet Id which is used for provisioning the instance
	NetworkProfileId        *string                    `json:"networkProfileId,omitempty"` // The network-profile-id which is used for provisioning the instance
	ComputeType             *string                    `json:"computeType,omitempty"`
	ComputeId               *string                    `json:"computeId,omitempty"`
	AwsInfraConfig          *AwsInfraConfig            `json:"awsInfraConfig,omitempty"`
	BaseStorage             *int                       `json:"baseStorage,omitempty"`            // The base storage (in bytes) that has been provisioned for the DB Service instance.
	AdditionalStorage       *int                       `json:"additionalStorage,omitempty"`      // The additional storage (in bytes) to be provisioned for the DB Service instance. This is in addition to what is specified in the compute type.
	AllocatedStorage        *int                       `json:"allocatedStorage,omitempty"`       // The actual storage (in bytes) that has been provisioned for the DB Service instance.
	MaxMemory               *int                       `json:"maxMemory,omitempty"`              // The allocated max memory (in bytes) for this instance
	SoftwareImage           *string                    `json:"softwareImage,omitempty"`          // The software-image-name which is used for provisioning this instance
	SoftwareImageVersion    *string                    `json:"softwareImageVersion,omitempty"`   // The software-image-version-name which is used for provisioning this instance
	SoftwareImageId         *string                    `json:"softwareImageId,omitempty"`        // The software-image-id which is used for provisioning this instance
	SoftwareImageVersionId  *string                    `json:"softwareImageVersionId,omitempty"` // The software-image-version-id which is used for provisioning this instance
	DataVolumeIops          *int                       `json:"dataVolumeIops,omitempty"`
	Throughput              *int                       `json:"throughput,omitempty"` // Throughput requested for this DB Service instance
	MultiDisk               *bool                      `json:"multiDisk,omitempty"`  // Specify whether the DB service uses multiple data disks
	ParameterProfile        *ParameterProfile          `json:"parameterProfile,omitempty"`
	OptionProfile           *OptionProfile             `json:"optionProfile,omitempty"`
	ParameterProfileId      *string                    `json:"parameterProfileId,omitempty"`
	SyncMode                *string                    `json:"syncMode,omitempty"`
	EngineConfiguration     *ServiceInstanceEngineInfo `json:"engineConfiguration,omitempty"`
	ComputeConfig           *InstanceComputeConfig     `json:"computeConfig,omitempty"`
	StorageConfig           *InstanceStorageConfig     `json:"storageConfig,omitempty"`
	ArchiveStorageConfig    *InstanceStorageConfig     `json:"archiveStorageConfig,omitempty"`
}

type OptionProfile struct {
	Id     *string `json:"id,omitempty"`   // Tessell generated UUID for the the option profile
	Name   *string `json:"name,omitempty"` // The name used to identify the option profile
	Status *string `json:"status,omitempty"`
}

type DBLicenseInfo struct {
	Licenses *[]LicenseInfo `json:"licenses,omitempty"`
}

type LicenseInfo struct {
	LicenseId *string  `json:"licenseId,omitempty"`
	LockHash  *string  `json:"lockHash,omitempty"` // Acquired licenses lock-hash
	Quantity  *float64 `json:"quantity,omitempty"` // quantity of acquired license
}

type TessellServiceInstanceMetadata struct {
	InstanceGroupName             *string                 `json:"instanceGroupName,omitempty"`
	DeletionAttempts              *int                    `json:"deletionAttempts,omitempty"`
	LastDeletionDispatchTime      *string                 `json:"lastDeletionDispatchTime,omitempty"`
	LastResizeDispatchTime        *string                 `json:"lastResizeDispatchTime,omitempty"`
	LastStorageResizeDispatchTime *string                 `json:"lastStorageResizeDispatchTime,omitempty"`
	AddReplicaContextId           *string                 `json:"addReplicaContextId,omitempty"`
	Data                          *map[string]interface{} `json:"data,omitempty"`
	IsDcrEnabled                  *bool                   `json:"isDcrEnabled,omitempty"`
}

type TessellServiceInstanceDriverInfo struct {
	Data *map[string]interface{} `json:"data,omitempty"`
}

type TessellServiceInstanceInProgressUpdateInfoOps struct {
	Infra *TessellServiceInstanceInProgressUpdateInfoOpsInfra `json:"infra,omitempty"`
}

type TessellServiceInstanceInProgressUpdateInfoOpsInfra struct {
	ResourceUpdateInfo *TessellResourceUpdateInfo             `json:"resourceUpdateInfo,omitempty"`
	InfraUpdateInfo    *TessellServiceInstanceInfraUpdateInfo `json:"infraUpdateInfo,omitempty"`
}

type TessellServiceInstanceInfraUpdateInfo struct {
	ComputeType *string `json:"computeType,omitempty"`
}

type TessellAgentLcmInfo struct {
	ComputeResourceInfo *map[string]interface{} `json:"computeResourceInfo,omitempty"`
	ServiceInfo         *map[string]interface{} `json:"serviceInfo,omitempty"`
	InstanceInfo        *map[string]interface{} `json:"instanceInfo,omitempty"`
	Data                *map[string]interface{} `json:"data,omitempty"`
}

type TessellComputeResourceOpsDTO struct {
	Id                           *string                               `json:"id,omitempty"`          // Tessell generated UUID for the entity
	Name                         *string                               `json:"name"`                  // Name of the entity
	Description                  *string                               `json:"description,omitempty"` // Compute Resource description
	TenantId                     *string                               `json:"tenantId,omitempty"`
	SubscriptionId               *string                               `json:"subscriptionId,omitempty"`
	EngineType                   *string                               `json:"engineType,omitempty"`
	Status                       *string                               `json:"status,omitempty"`
	Tsm                          *bool                                 `json:"tsm,omitempty"`
	CloudStatus                  *string                               `json:"cloudStatus,omitempty"`           // Compute Resource&#39;s status in the cloud
	ComputeSharingEnabled        *bool                                 `json:"computeSharingEnabled,omitempty"` // Whether the Compute Resource is shared across multiple DB Services
	CloudAccountId               *string                               `json:"cloudAccountId,omitempty"`        // Compute Resource&#39;s Tessell cloud account identifier
	CloudLocation                *string                               `json:"cloudLocation,omitempty"`         // Compute Resource&#39;s location in the cloud
	CloudResourceId              *string                               `json:"cloudResourceId,omitempty"`       // Compute Resource&#39;s cloud identifier
	Type                         *string                               `json:"type,omitempty"`
	MachineType                  *string                               `json:"machineType,omitempty"`
	OsInfo                       *OsInfo                               `json:"osInfo,omitempty"`
	SoftwareImageId              *string                               `json:"softwareImageId,omitempty"`        // Compute Resource&#39;s Software Image Id
	SoftwareImageVersionId       *string                               `json:"softwareImageVersionId,omitempty"` // Compute Resource&#39;s Software Image Version Id
	NetworkProfileId             *string                               `json:"networkProfileId,omitempty"`       // Compute Resource&#39;s Network Profile Id
	ComputeTypeId                *string                               `json:"computeTypeId,omitempty"`          // Compute Resource&#39;s compute type Id
	UserId                       *string                               `json:"userId,omitempty"`                 // Compute Resource&#39;s user id
	Owner                        *string                               `json:"owner,omitempty"`                  // Compute resource&#39;s owner email address
	DateCreated                  *string                               `json:"dateCreated,omitempty"`            // Timestamp when the entity was created
	DateModified                 *string                               `json:"dateModified,omitempty"`           // Timestamp when the entity was last modified
	Timezone                     *string                               `json:"timezone,omitempty"`               // The timezone detail
	MachineFqdnInfo              *ComputeResourceMachineFqdnInfo       `json:"machineFqdnInfo,omitempty"`
	IpAddressInfo                *ComputeResourceIpAddressInfo         `json:"ipAddressInfo,omitempty"`
	Metadata                     *ComputeResourceMetadata              `json:"metadata,omitempty"`
	DriverInfo                   *ComputeResourceDriverInfo            `json:"driverInfo,omitempty"`
	TessellStackInfo             *ComputeResourceTessellStackInfo      `json:"tessellStackInfo,omitempty"`
	TessellAgentLcmInfo          *TessellAgentLcmInfo                  `json:"tessellAgentLcmInfo,omitempty"`
	TessellEndpointMigrationInfo *ComputeResourceEndpointMigrationInfo `json:"tessellEndpointMigrationInfo,omitempty"`
	Internal                     *bool                                 `json:"internal,omitempty"` // Whether the Compute Resource is created for internal usage
	ContextInfo                  *ComputeResourceContextInfo           `json:"contextInfo,omitempty"`
	ActionMetadata               *ComputeResourceActionMetadata        `json:"actionMetadata,omitempty"`
}

type OsInfo struct {
	Type  *string `json:"type,omitempty"`
	Image *string `json:"image,omitempty"`
}

type ComputeResourceMachineFqdnInfo struct {
	Data *map[string]interface{} `json:"data,omitempty"`
}

type ComputeResourceIpAddressInfo struct {
	Data        *map[string]interface{} `json:"data,omitempty"`
	IpAddresses *map[string]interface{} `json:"ipAddresses,omitempty"`
}

type ComputeResourceMetadata struct {
	Subscription                 *string                                  `json:"subscription,omitempty"`
	ComputeType                  *string                                  `json:"computeType,omitempty"`
	VPC                          *string                                  `json:"vpc,omitempty"`
	PrivateSubnetId              *string                                  `json:"privateSubnetId,omitempty"`
	PrivateSubnet                *string                                  `json:"privateSubnet,omitempty"`
	PublicSubnetId               *string                                  `json:"publicSubnetId,omitempty"`
	PublicSubnet                 *string                                  `json:"publicSubnet,omitempty"`
	EncryptionKeyId              *string                                  `json:"encryptionKeyId,omitempty"`
	EncryptionKey                *string                                  `json:"encryptionKey,omitempty"`
	ConnectivityInfo             *ComputeResourceMetadataConnectivityInfo `json:"connectivityInfo,omitempty"`
	OsTimezone                   *string                                  `json:"osTimezone,omitempty"` // The timezone detail
	DBserverInfo                 *ComputeResourceDBserverInfo             `json:"dbserverInfo,omitempty"`
	AwsInfraConfig               *AwsInfraConfig                          `json:"awsInfraConfig,omitempty"`
	LicenseInfo                  *ComputeLicenseInfo                      `json:"licenseInfo,omitempty"`
	Data                         *map[string]interface{}                  `json:"data,omitempty"`
	IsAzureMonitorAgentInstalled *bool                                    `json:"isAzureMonitorAgentInstalled,omitempty"`
	NodeIndex                    *int                                     `json:"nodeIndex,omitempty"`
	UseAzureMonitorAgent         *bool                                    `json:"useAzureMonitorAgent,omitempty"`
	ComputeConfig                *ComputeConfig                           `json:"computeConfig,omitempty"`
}

type ComputeResourceMetadataConnectivityInfo struct {
	ServicePort *int `json:"servicePort,omitempty"` // The connection port for the DB Service
}

type ComputeResourceDBserverInfo struct {
	FirstProvisionedDBserviceId *string                    `json:"firstProvisionedDbserviceId,omitempty"`
	DBServiceIds                *[]string                  `json:"dbServiceIds,omitempty"` // The list of DB Service ids that are hosted on this compute resource
	EnablePublicAccess          *bool                      `json:"enablePublicAccess,omitempty"`
	EnableSSL                   *bool                      `json:"enableSSL,omitempty"`
	SoftwareImageInfo           *DBserverSoftwareImageInfo `json:"softwareImageInfo,omitempty"`
}

type DBserverSoftwareImageInfo struct {
	SoftwareImage         *string                             `json:"softwareImage,omitempty"`
	SoftwareImageId       *string                             `json:"softwareImageId,omitempty"`
	SoftwareImageVersions *[]DBserverSoftwareImageVersionInfo `json:"softwareImageVersions,omitempty"`
}

type DBserverSoftwareImageVersionInfo struct {
	SoftwareImageVersion   *string `json:"softwareImageVersion,omitempty"`
	SoftwareImageVersionId *string `json:"softwareImageVersionId,omitempty"`
	Supported              *bool   `json:"supported,omitempty"`
}

type ComputeLicenseInfo struct {
	AcquirerId *string  `json:"acquirerId,omitempty"`
	LicenseId  *string  `json:"licenseId,omitempty"`
	LockHash   *string  `json:"lockHash,omitempty"` // Acquired licenses lock-hash
	Quantity   *float64 `json:"quantity,omitempty"` // quantity of acquired license
}

type ComputeConfig struct {
	Provider      *string               `json:"provider,omitempty"`
	ExadataConfig *ExadataComputeConfig `json:"exadataConfig,omitempty"`
}

type ExadataComputeConfig struct {
	InfrastructureId   *string `json:"infrastructureId"`
	InfrastructureName *string `json:"infrastructureName"`
	VmClusterId        *string `json:"vmClusterId"`
	VmClusterName      *string `json:"vmClusterName"`
	Ocpu               *int    `json:"ocpu,omitempty"`
	MemoryInGbs        *int    `json:"memoryInGbs,omitempty"`
	FloatingIpAddress  *string `json:"floatingIpAddress,omitempty"`
	DBServer           *string `json:"dbServer,omitempty"`
	PrivateIpAddress   *string `json:"privateIpAddress,omitempty"`
	LocalStorageInGbs  *int    `json:"localStorageInGbs,omitempty"`
	DNSName            *string `json:"dnsName,omitempty"`
}

type ComputeResourceDriverInfo struct {
	Data *map[string]interface{} `json:"data,omitempty"`
}

type ComputeResourceTessellStackInfo struct {
	Data *map[string]interface{} `json:"data,omitempty"`
}

type ComputeResourceEndpointMigrationInfo struct {
	ValidateCommonEndpoint                 *bool   `json:"validateCommonEndpoint,omitempty"`
	MigrateToCommonEndpoint                *bool   `json:"migrateToCommonEndpoint,omitempty"`
	CommonEndpointSuccessfulValidationTime *string `json:"commonEndpointSuccessfulValidationTime,omitempty"`
	CommonEndpointSuccessfulMigrationTime  *string `json:"commonEndpointSuccessfulMigrationTime,omitempty"`
}

type ComputeResourceContextInfo struct {
	SubStatus   *string `json:"subStatus,omitempty"`
	Description *string `json:"description,omitempty"`
}

type ComputeResourceActionMetadata struct {
	LastActionMetadata     *CrLastActionMetadata `json:"lastActionMetadata,omitempty"`
	LastResizeDispatchTime *string               `json:"lastResizeDispatchTime,omitempty"`
}

type CrLastActionMetadata struct {
	ReferenceId *string `json:"referenceId,omitempty"`
	ContextId   *string `json:"contextId,omitempty"`
	ActionType  *string `json:"actionType,omitempty"`
}

type BackupDownloadConfig struct {
	AllowBackupDownloadsForAllUsers *bool `json:"allowBackupDownloadsForAllUsers,omitempty"` // Allow all users to download the backup, if false only owner/co-owner(s) will be allowed
	AllowBackupDownloads            *bool `json:"allowBackupDownloads,omitempty"`            // Allow download of the backup for owner/co-owner of the AM
}

type DMMConsumerView struct {
	Id                   *string                    `json:"id,omitempty"`                  // ID of the Availability Machine
	TessellServiceId     *string                    `json:"tessellServiceId,omitempty"`    // ID of the DB Service that is associated with the Availability Machine
	ServiceName          *string                    `json:"serviceName,omitempty"`         // Name of the DB Service that is associated with the Availability Machine
	Tenant               *string                    `json:"tenant,omitempty"`              // ID of the tenant under which this Availability Machine is effective
	Subscription         *string                    `json:"subscription,omitempty"`        // Name of the subscription under which the associated DB Service is hosted
	EngineType           *string                    `json:"engineType,omitempty"`          // Database Engine Type
	DataIngestionStatus  *string                    `json:"dataIngestionStatus,omitempty"` // Availability Machine&#39;s data ingestion status
	UserId               *string                    `json:"userId,omitempty"`              // User details representing the owner for the Availability Machine
	Owner                *string                    `json:"owner,omitempty"`               // User details representing the owner for the Availability Machine
	LoggedInUserRole     *string                    `json:"loggedInUserRole,omitempty"`    // The role of the logged in user for accessing this Availability Machine
	SharedWith           *EntityAclSharingInfo      `json:"sharedWith,omitempty"`
	CloudAvailability    *[]CloudRegionInfo         `json:"cloudAvailability,omitempty"` // Availability Machine manages data across multiple regions within a cloud. This sections provides information about the cloud and regions where this Availability Machine is managing the data.
	Topology             *[]DBServiceTopology       `json:"topology,omitempty"`          // The availability location details: cloudAccount to region
	RPOPolicy            *RPOPolicyConfig           `json:"rpoPolicy,omitempty"`
	DAPs                 *[]TessellDAPServiceDTO    `json:"daps,omitempty"`         // The Access Policies (DAP) that have configured for this Availability Machine
	Clones               *[]TessellCloneSummaryInfo `json:"clones,omitempty"`       // The clone DB Services that have been created using contents (snapshots, Sanitized Snapshots, PITR, backups) from this Availability Machine
	DateCreated          *string                    `json:"dateCreated,omitempty"`  // The timestamp when the Availability Machine was incarnated
	DateModified         *string                    `json:"dateModified,omitempty"` // The timestamp when the Availability Machine was last updated
	Tsm                  *bool                      `json:"tsm,omitempty"`          // Specify whether the associated DB Service is created using TSM compute type
	BackupDownloadConfig *BackupDownloadConfig      `json:"backupDownloadConfig,omitempty"`
	StorageConfig        *StorageConfigPayload      `json:"storageConfig,omitempty"`
}

type GetDMMsServiceView struct {
	Metadata *APIMetadata       `json:"metadata,omitempty"`
	Response *[]DMMConsumerView `json:"response,omitempty"`
}
