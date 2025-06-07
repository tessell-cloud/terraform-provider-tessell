package model

type AzureNetAppConfigPayloadConfigurations struct {
	NetworkFeatures *string `json:"networkFeatures,omitempty"`
}

type WeeklySchedule struct {
	Days *[]string `json:"days,omitempty"` // Days in a week to retain weekly backups for
}

type InstanceFsxNetAppConfig struct {
	FileSystemName *string `json:"fileSystemName,omitempty"`
	SvmName        *string `json:"svmName,omitempty"`
	VolumeName     *string `json:"volumeName,omitempty"`
	FileSystemId   *string `json:"fileSystemId,omitempty"` // File System Id of the FSx NetApp registered with Tessell
	SvmId          *string `json:"svmId,omitempty"`        // Storage Virtual Machine Id of the FSx NetApp registered with Tessell
}

type BackupCustomRPOPolicy struct {
	Name     *string       `json:"name"` // Custom RPO policy name
	Schedule *ScheduleInfo `json:"schedule"`
}

type StorageConfigPayload struct {
	Provider          *string                   `json:"provider"`
	FsxNetAppConfig   *FsxNetAppConfigPayload   `json:"fsxNetAppConfig,omitempty"`
	AzureNetAppConfig *AzureNetAppConfigPayload `json:"azureNetAppConfig,omitempty"`
}

type BackupStandardRPOPolicy struct {
	RetentionDays   *int        `json:"retentionDays"` // Number of days for which the backup of DB Service would be retained
	BackupStartTime *TimeFormat `json:"backupStartTime,omitempty"`
}

type ServiceInstanceOracleEngineConfig struct {
	AccessMode *string `json:"accessMode,omitempty"`
}

type RPOPolicyConfigBackupRPOConfig struct {
	FullBackupSchedule *FullBackupSchedule      `json:"fullBackupSchedule,omitempty"`
	StandardPolicy     *BackupStandardRPOPolicy `json:"standardPolicy,omitempty"`
	CustomPolicy       *BackupCustomRPOPolicy   `json:"customPolicy,omitempty"`
}

type FullBackupSchedule struct {
	StartTime      *TimeFormat     `json:"startTime,omitempty"`
	WeeklySchedule *WeeklySchedule `json:"weeklySchedule,omitempty"`
}

type InstanceExadataComputeConfig struct {
	InfrastructureId   *string `json:"infrastructureId"`
	InfrastructureName *string `json:"infrastructureName"`
	VmClusterId        *string `json:"vmClusterId"`
	VmClusterName      *string `json:"vmClusterName"`
	Vcpus              *int    `json:"vcpus"`
	Memory             *int    `json:"memory"`
}

type CommonYearlySchedule struct {
	Dates          *[]int    `json:"dates,omitempty"` // Dates in a month to retain monthly backups
	LastDayOfMonth *bool     `json:"lastDayOfMonth,omitempty"`
	Months         *[]string `json:"months,omitempty"`
}

type APIError struct {
	Code             *string `json:"code,omitempty"`    // Status code for the error response
	Message          *string `json:"message,omitempty"` // Error message for API response
	Resolution       *string `json:"resolution,omitempty"`
	Timestamp        *string `json:"timestamp,omitempty"`
	ContextId        *string `json:"contextId,omitempty"`        // ContextId of API request
	SessionId        *string `json:"sessionId,omitempty"`        // SessionId of API request
	TessellErrorCode *string `json:"tessellErrorCode,omitempty"` // Unique error code specific to Tessell
}

type TessellDataflixFromTimeInfo struct {
	FromTime   *string                      `json:"fromTime,omitempty"` // Recoverability start timestamp
	ToTime     *string                      `json:"toTime,omitempty"`   // Recoverability end timestamp
	SharedWith *EntityAclSharingSummaryInfo `json:"sharedWith,omitempty"`
}

type SnapshotRegionAvailability struct {
	Region *string `json:"region"`
	Status *string `json:"status"` // Database Backup Status
}

type TessellResourceUpdateInfo struct {
	UpdateType  *string                 `json:"updateType,omitempty"`  // Type of the update
	ReferenceId *string                 `json:"referenceId,omitempty"` // The reference-id of the update request
	SubmittedAt *string                 `json:"submittedAt,omitempty"` // Timestamp when the resource update was requested
	UpdateInfo  *map[string]interface{} `json:"updateInfo,omitempty"`  // The specific details for a Tessell resource that are being updated
}

type SnapshotAvailabilityConfig struct {
	AvailabilityConfiguredManually *bool                            `json:"availabilityConfiguredManually,omitempty"`
	DAPId                          *string                          `json:"dapId,omitempty"`
	CloudAvailabilityConfig        *[]SnapshotCloudAvailabilityInfo `json:"cloudAvailabilityConfig,omitempty"`
}

type RPOPolicyConfig struct {
	IncludeTransactionLogs *bool                           `json:"includeTransactionLogs,omitempty"` // Determines whether transaction logs should be retained to enable Point-In-Time Recovery (PITR) functionality
	EnableAutoSnapshot     *bool                           `json:"enableAutoSnapshot"`               // Specify whether system will take automatic snapshots
	StandardPolicy         *StandardRPOPolicy              `json:"standardPolicy,omitempty"`
	CustomPolicy           *CustomRPOPolicy                `json:"customPolicy,omitempty"`
	FullBackupSchedule     *FullBackupSchedule             `json:"fullBackupSchedule,omitempty"`
	EnableAutoBackup       *bool                           `json:"enableAutoBackup,omitempty"` // Specify whether system will take automatic backups
	BackupRPOConfig        *RPOPolicyConfigBackupRPOConfig `json:"backupRpoConfig,omitempty"`
}

type EntityAclSharingSummaryInfo struct {
	Users *[]string `json:"users,omitempty"`
}

type MonthlySchedule struct {
	CommonSchedule *DatesForEachMonth `json:"commonSchedule,omitempty"`
}

type APIMetadata struct {
	TimeZone   *string            `json:"timeZone,omitempty"`
	Records    *int               `json:"records,omitempty"`
	Pagination *APIPaginationInfo `json:"pagination,omitempty"`
}

type EntityAclSharingInfo struct {
	Users *[]EntityUserAclSharingInfo `json:"users,omitempty"`
}

type TimeFormat struct {
	Hour   *int `json:"hour"`
	Minute *int `json:"minute"`
}

type InstanceStorageConfig struct {
	Provider          *string                    `json:"provider,omitempty"`
	FsxNetAppConfig   *InstanceFsxNetAppConfig   `json:"fsxNetAppConfig,omitempty"`
	AzureNetAppConfig *InstanceAzureNetAppConfig `json:"azureNetAppConfig,omitempty"`
}

type YearlySchedule struct {
	CommonSchedule        *CommonYearlySchedule `json:"commonSchedule,omitempty"`
	MonthSpecificSchedule *[]MonthWiseDates     `json:"monthSpecificSchedule,omitempty"`
}

type StandardRPOPolicy struct {
	RetentionDays          *int        `json:"retentionDays"`                    // Number of days for which the snapshot of DB Service would be retained
	IncludeTransactionLogs *bool       `json:"includeTransactionLogs,omitempty"` // Determines whether transaction logs should be retained to enable Point-In-Time Recovery (PITR) functionality
	SnapshotStartTime      *TimeFormat `json:"snapshotStartTime"`
}

type ServiceInstanceEngineInfo struct {
	OracleConfig *ServiceInstanceOracleEngineConfig `json:"oracleConfig,omitempty"`
}

type ScheduleInfo struct {
	BackupStartTime *TimeFormat      `json:"backupStartTime,omitempty"`
	DailySchedule   *DailySchedule   `json:"dailySchedule,omitempty"`
	WeeklySchedule  *WeeklySchedule  `json:"weeklySchedule,omitempty"`
	MonthlySchedule *MonthlySchedule `json:"monthlySchedule,omitempty"`
	YearlySchedule  *YearlySchedule  `json:"yearlySchedule,omitempty"`
}

type APIPaginationInfo struct {
	PageSize   *int `json:"pageSize,omitempty"`
	PageOffset *int `json:"pageOffset,omitempty"`
}

type TessellServiceInstanceDTO struct {
	Id                   *string                              `json:"id,omitempty"`                // Tessell generated UUID for the DB Service Instance
	Name                 *string                              `json:"name"`                        // Name of the DB Service Instance
	InstanceGroupName    *string                              `json:"instanceGroupName,omitempty"` // Name of the instance group
	Type                 *string                              `json:"type,omitempty"`              // DB Service instance type
	Role                 *string                              `json:"role,omitempty"`              // DB Service instance role
	Status               *string                              `json:"status,omitempty"`            // DB Service instance status
	TessellServiceId     *string                              `json:"tessellServiceId,omitempty"`  // DB Service Instance&#39;s associated DB Service id
	Cloud                *string                              `json:"cloud,omitempty"`             // DB Service Instance&#39;s cloud type
	Region               *string                              `json:"region,omitempty"`            // DB Service Instance&#39;s cloud region
	AvailabilityZone     *string                              `json:"availabilityZone,omitempty"`  // DB Service Instance&#39;s cloud availability zone
	InstanceGroupId      *string                              `json:"instanceGroupId,omitempty"`   // The instance groupd Id
	ComputeType          *string                              `json:"computeType,omitempty"`       // The compute used for creation of the Tessell Service Instance
	AwsInfraConfig       *AwsInfraConfig                      `json:"awsInfraConfig,omitempty"`
	ComputeId            *string                              `json:"computeId,omitempty"`   // The associated compute identifier
	ComputeName          *string                              `json:"computeName,omitempty"` // The associated compute name
	Storage              *int                                 `json:"storage,omitempty"`     // The storage (in bytes) that has been provisioned for the DB Service instance.
	DataVolumeIops       *int                                 `json:"dataVolumeIops,omitempty"`
	Throughput           *int                                 `json:"throughput,omitempty"` // Throughput requested for this DB Service instance
	EnablePerfInsights   *bool                                `json:"enablePerfInsights,omitempty"`
	ParameterProfile     *ParameterProfile                    `json:"parameterProfile,omitempty"`
	MonitoringConfig     *MonitoringConfig                    `json:"monitoringConfig,omitempty"`
	VPC                  *string                              `json:"vpc,omitempty"`                  // The VPC used for creation of the DB Service Instance
	PublicSubnet         *string                              `json:"publicSubnet,omitempty"`         // The public subnet used for creation of the DB Service Instance
	PrivateSubnet        *string                              `json:"privateSubnet,omitempty"`        // The private subnet used for creation of the DB Service Instance
	EncryptionKey        *string                              `json:"encryptionKey,omitempty"`        // The encryption key name which is used to encrypt the data at rest
	SoftwareImage        *string                              `json:"softwareImage,omitempty"`        // Software Image to be used to create the instance
	SoftwareImageVersion *string                              `json:"softwareImageVersion,omitempty"` // Software Image Version to be used to create the instance
	DateCreated          *string                              `json:"dateCreated,omitempty"`          // Timestamp when the entity was created
	ConnectString        *TessellServiceInstanceConnectString `json:"connectString,omitempty"`
	UpdatesInProgress    *[]TessellResourceUpdateInfo         `json:"updatesInProgress,omitempty"` // The updates that are in progress for this resource
	LastStartedAt        *string                              `json:"lastStartedAt,omitempty"`     // Timestamp when the service instance was last started at
	LastStoppedAt        *string                              `json:"lastStoppedAt,omitempty"`     // Timestamp when the Service Instance was last stopped at
	SyncMode             *string                              `json:"syncMode,omitempty"`
	EngineConfiguration  *ServiceInstanceEngineInfo           `json:"engineConfiguration,omitempty"`
	ComputeConfig        *InstanceComputeConfig               `json:"computeConfig,omitempty"`
	StorageConfig        *InstanceStorageConfig               `json:"storageConfig,omitempty"`
	ArchiveStorageConfig *InstanceStorageConfig               `json:"archiveStorageConfig,omitempty"`
}

type APIStatus struct {
	Status  *string `json:"status,omitempty"`
	Message *string `json:"message,omitempty"`
}

type AzureNetAppEncryptionKeyInfo struct {
	Id                      *string `json:"id,omitempty"`                      // Id of the encryption key
	Name                    *string `json:"name,omitempty"`                    // name of the encryption key
	KeyVaultCloudResourceId *string `json:"keyVaultCloudResourceId,omitempty"` // name of the encryption key vault in cloud
	KeySource               *string `json:"keySource,omitempty"`
}

type AwsInfraConfig struct {
	AwsCpuOptions *AwsCpuOptions `json:"awsCpuOptions,omitempty"`
}

type InstanceAzureNetAppConfig struct {
	AzureNetAppName     *string                       `json:"azureNetAppName,omitempty"`
	CapacityPoolName    *string                       `json:"capacityPoolName,omitempty"`
	VolumeName          *string                       `json:"volumeName,omitempty"`
	AzureNetAppId       *string                       `json:"azureNetAppId,omitempty"`       // Azure NetApp Id registered with Tessell
	CapacityPoolId      *string                       `json:"capacityPoolId,omitempty"`      // Capacity Pool Id of the Azure NetApp registered with Tessell
	DelegatedSubnetId   *string                       `json:"delegatedSubnetId,omitempty"`   // Delegated Subnet name registered with Tessell for the Azure NetApp volume
	DelegatedSubnetName *string                       `json:"delegatedSubnetName,omitempty"` // Delegated Subnet Id registered with Tessell for the Azure NetApp volume
	EncryptionKeyInfo   *AzureNetAppEncryptionKeyInfo `json:"encryptionKeyInfo,omitempty"`
	NetworkFeatures     *string                       `json:"networkFeatures,omitempty"`
	ServiceLevel        *string                       `json:"serviceLevel,omitempty"`
}

type EntityUserAclSharingInfo struct {
	EmailId *string `json:"emailId,omitempty"`
	Role    *string `json:"role,omitempty"`
}

type TessellServiceInstanceConnectString struct {
	ConnectDescriptor *string `json:"connectDescriptor,omitempty"`
	MasterUser        *string `json:"masterUser,omitempty"`
	Endpoint          *string `json:"endpoint,omitempty"`
	ServicePort       *string `json:"servicePort,omitempty"`
}

type CustomRPOPolicy struct {
	Name     *string       `json:"name"` // Custom RPO policy name
	Schedule *ScheduleInfo `json:"schedule"`
}

type MonitoringConfig struct {
	PerfInsights *PerfInsightsConfig `json:"perfInsights,omitempty"`
}

type MonthWiseDates struct {
	Month *string `json:"month"` // Name of a month
	Dates *[]int  `json:"dates"`
}

type DatesForEachMonth struct {
	Dates          *[]int `json:"dates,omitempty"` // Dates in a month to retain monthly backups
	LastDayOfMonth *bool  `json:"lastDayOfMonth,omitempty"`
}

type InstanceComputeConfig struct {
	Provider      *string                       `json:"provider,omitempty"`
	ExadataConfig *InstanceExadataComputeConfig `json:"exadataConfig,omitempty"`
}

type AzureNetAppConfigPayload struct {
	AzureNetAppId  *string                                 `json:"azureNetAppId,omitempty"`  // Azure NetApp Id registered with Tessell
	CapacityPoolId *string                                 `json:"capacityPoolId,omitempty"` // Capacity pool Id of the Azure NetApp registered with Tessell
	Configurations *AzureNetAppConfigPayloadConfigurations `json:"configurations,omitempty"`
}

type DatabaseSnapshotRegionInfo struct {
	Region *string `json:"region"`           // The region name
	Status *string `json:"status,omitempty"` // The current status of the snapshot in the respective region
}

type SnapshotCloudAvailabilityInfo struct {
	Cloud   *string                       `json:"cloud"`
	Regions *[]SnapshotRegionAvailability `json:"regions,omitempty"` // The list of regions and respective availability status
}

type TaskSummary struct {
	TaskId                *string            `json:"taskId,omitempty"`
	TaskType              *string            `json:"taskType,omitempty"`
	ResourceId            *string            `json:"resourceId,omitempty"`
	AssociatedResourceIds *[]string          `json:"associatedResourceIds,omitempty"`
	Details               *map[string]string `json:"details,omitempty"`
}

type TessellDataflixPITRInfo struct {
	Cloud   *string                             `json:"cloud,omitempty"`
	Regions *[]TessellDataflixPITRInfoForRegion `json:"regions,omitempty"`
}

type DeletionScheduleDTO struct {
	Id             *string                       `json:"id,omitempty"`
	DeleteAt       *string                       `json:"deleteAt"` // Time at which the DB Service should be deleted at
	DeletionConfig *TessellServiceDeletionConfig `json:"deletionConfig,omitempty"`
}

type TessellDataflixPITRInfoForRegion struct {
	Region     *string                        `json:"region,omitempty"` // Region name
	TimeRanges *[]TessellDataflixFromTimeInfo `json:"timeRanges,omitempty"`
}

type FsxNetAppConfigPayload struct {
	FileSystemId *string `json:"fileSystemId"` // File System Id of the FSx NetApp registered with Tessell
	SvmId        *string `json:"svmId"`        // Storage Virtual Machine Id of the FSx NetApp registered with Tessell
}

type ParameterProfile struct {
	Id      *string `json:"id,omitempty"`      // Tessell generated UUID for the the parameter profile
	Name    *string `json:"name,omitempty"`    // The name used to identify the parameter profile
	Version *string `json:"version,omitempty"` // The version of the parameter profile associated with the instance
	Status  *string `json:"status,omitempty"`
}

type TessellServiceDeletionConfig struct {
	RetainAvailabilityMachine *bool `json:"retainAvailabilityMachine,omitempty"` // If specified as true, the associated Availability Machine (snapshots, sanitized-snapshots, logs) would be retained
}

type RegionInfo struct {
	Region            *string   `json:"region"` // The cloud region name
	AvailabilityZones *[]string `json:"availabilityZones,omitempty"`
}

type CloudRegionInfo struct {
	Cloud   *string       `json:"cloud"`
	Regions *[]RegionInfo `json:"regions,omitempty"` // The regions details
}

type DailySchedule struct {
	BackupsPerDay *int `json:"backupsPerDay,omitempty"` // The number of backups to be captured per day.
}

type DatabaseSnapshotCloudRegionInfo struct {
	Cloud   *string                       `json:"cloud"`
	Regions *[]DatabaseSnapshotRegionInfo `json:"regions,omitempty"` // Region specific availability details for the snapshot
}

type PerfInsightsConfig struct {
	PerfInsightsEnabled    *bool   `json:"perfInsightsEnabled,omitempty"`
	MonitoringDeploymentId *string `json:"monitoringDeploymentId,omitempty"`
	Status                 *string `json:"status,omitempty"`
}

type AwsCpuOptions struct {
	Vcpus *int `json:"vcpus,omitempty"` // Number of vcpus for aws cpu options
}

type BackupDatabaseInfo struct {
	Id     *string `json:"id,omitempty"`     // ID of the database
	Name   *string `json:"name,omitempty"`   // Name of the database
	Status *string `json:"status,omitempty"` // Status of the database as of capture of this snapshot
}
