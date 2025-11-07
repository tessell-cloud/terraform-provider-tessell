package availability_machine

import (
	//"fmt"
	//"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"terraform-provider-tessell/internal/model"
)

func setResourceData(d *schema.ResourceData, dmmConsumerView *model.DMMConsumerView) error {

	if err := d.Set("id", dmmConsumerView.Id); err != nil {
		return err
	}

	if err := d.Set("tessell_service_id", dmmConsumerView.TessellServiceId); err != nil {
		return err
	}

	if err := d.Set("service_name", dmmConsumerView.ServiceName); err != nil {
		return err
	}

	if err := d.Set("tenant", dmmConsumerView.Tenant); err != nil {
		return err
	}

	if err := d.Set("subscription", dmmConsumerView.Subscription); err != nil {
		return err
	}

	if err := d.Set("engine_type", dmmConsumerView.EngineType); err != nil {
		return err
	}

	if err := d.Set("data_ingestion_status", dmmConsumerView.DataIngestionStatus); err != nil {
		return err
	}

	if err := d.Set("user_id", dmmConsumerView.UserId); err != nil {
		return err
	}

	if err := d.Set("owner", dmmConsumerView.Owner); err != nil {
		return err
	}

	if err := d.Set("logged_in_user_role", dmmConsumerView.LoggedInUserRole); err != nil {
		return err
	}

	if err := d.Set("shared_with", parseEntityAclSharingInfoWithResData(dmmConsumerView.SharedWith, d)); err != nil {
		return err
	}

	if err := d.Set("cloud_availability", parseCloudRegionInfoListWithResData(dmmConsumerView.CloudAvailability, d)); err != nil {
		return err
	}

	if err := d.Set("topology", parseDBServiceTopologyListWithResData(dmmConsumerView.Topology, d)); err != nil {
		return err
	}

	if err := d.Set("rpo_policy", parseRPOPolicyConfigWithResData(dmmConsumerView.RPOPolicy, d)); err != nil {
		return err
	}

	if err := d.Set("daps", parseTessellDAPServiceDTOListWithResData(dmmConsumerView.DAPs, d)); err != nil {
		return err
	}

	if err := d.Set("clones", parseTessellCloneSummaryInfoListWithResData(dmmConsumerView.Clones, d)); err != nil {
		return err
	}

	if err := d.Set("date_created", dmmConsumerView.DateCreated); err != nil {
		return err
	}

	if err := d.Set("date_modified", dmmConsumerView.DateModified); err != nil {
		return err
	}

	if err := d.Set("tsm", dmmConsumerView.Tsm); err != nil {
		return err
	}

	if err := d.Set("backup_download_config", parseBackupDownloadConfigWithResData(dmmConsumerView.BackupDownloadConfig, d)); err != nil {
		return err
	}

	if err := d.Set("storage_config", parseStorageConfigPayloadWithResData(dmmConsumerView.StorageConfig, d)); err != nil {
		return err
	}

	return nil
}

func parseEntityAclSharingInfoWithResData(sharedWith *model.EntityAclSharingInfo, d *schema.ResourceData) []interface{} {
	if sharedWith == nil {
		return nil
	}
	parsedSharedWith := make(map[string]interface{})
	if d.Get("shared_with") != nil {
		sharedWithResourceData := d.Get("shared_with").([]interface{})
		if len(sharedWithResourceData) > 0 {
			parsedSharedWith = (sharedWithResourceData[0]).(map[string]interface{})
		}
	}

	var users *[]model.EntityUserAclSharingInfo
	if sharedWith.Users != users {
		parsedSharedWith["users"] = parseEntityUserAclSharingInfoList(sharedWith.Users)
	}

	return []interface{}{parsedSharedWith}
}

func parseEntityAclSharingInfo(sharedWith *model.EntityAclSharingInfo) interface{} {
	if sharedWith == nil {
		return nil
	}
	parsedSharedWith := make(map[string]interface{})

	var users *[]model.EntityUserAclSharingInfo
	if sharedWith.Users != users {
		parsedSharedWith["users"] = parseEntityUserAclSharingInfoList(sharedWith.Users)
	}

	return parsedSharedWith
}

func parseEntityUserAclSharingInfoList(entityUserAclSharingInfo *[]model.EntityUserAclSharingInfo) []interface{} {
	if entityUserAclSharingInfo == nil {
		return nil
	}
	entityUserAclSharingInfoList := make([]interface{}, 0)

	if entityUserAclSharingInfo != nil {
		entityUserAclSharingInfoList = make([]interface{}, len(*entityUserAclSharingInfo))
		for i, entityUserAclSharingInfoItem := range *entityUserAclSharingInfo {
			entityUserAclSharingInfoList[i] = parseEntityUserAclSharingInfo(&entityUserAclSharingInfoItem)
		}
	}

	return entityUserAclSharingInfoList
}

func parseEntityUserAclSharingInfo(entityUserAclSharingInfo *model.EntityUserAclSharingInfo) interface{} {
	if entityUserAclSharingInfo == nil {
		return nil
	}
	parsedEntityUserAclSharingInfo := make(map[string]interface{})
	parsedEntityUserAclSharingInfo["email_id"] = entityUserAclSharingInfo.EmailId
	parsedEntityUserAclSharingInfo["role"] = entityUserAclSharingInfo.Role

	return parsedEntityUserAclSharingInfo
}

func parseCloudRegionInfoListWithResData(cloudAvailability *[]model.CloudRegionInfo, d *schema.ResourceData) []interface{} {
	if cloudAvailability == nil {
		return nil
	}
	cloudRegionInfoList := make([]interface{}, 0)

	if cloudAvailability != nil {
		cloudRegionInfoList = make([]interface{}, len(*cloudAvailability))
		for i, cloudRegionInfoItem := range *cloudAvailability {
			cloudRegionInfoList[i] = parseCloudRegionInfo(&cloudRegionInfoItem)
		}
	}

	return cloudRegionInfoList
}

func parseCloudRegionInfoList(cloudAvailability *[]model.CloudRegionInfo) []interface{} {
	if cloudAvailability == nil {
		return nil
	}
	cloudRegionInfoList := make([]interface{}, 0)

	if cloudAvailability != nil {
		cloudRegionInfoList = make([]interface{}, len(*cloudAvailability))
		for i, cloudRegionInfoItem := range *cloudAvailability {
			cloudRegionInfoList[i] = parseCloudRegionInfo(&cloudRegionInfoItem)
		}
	}

	return cloudRegionInfoList
}

func parseCloudRegionInfo(cloudAvailability *model.CloudRegionInfo) interface{} {
	if cloudAvailability == nil {
		return nil
	}
	parsedCloudAvailability := make(map[string]interface{})
	parsedCloudAvailability["cloud"] = cloudAvailability.Cloud

	var regions *[]model.RegionInfo
	if cloudAvailability.Regions != regions {
		parsedCloudAvailability["regions"] = parseRegionInfoList(cloudAvailability.Regions)
	}

	return parsedCloudAvailability
}

func parseRegionInfoList(regionInfo *[]model.RegionInfo) []interface{} {
	if regionInfo == nil {
		return nil
	}
	regionInfoList := make([]interface{}, 0)

	if regionInfo != nil {
		regionInfoList = make([]interface{}, len(*regionInfo))
		for i, regionInfoItem := range *regionInfo {
			regionInfoList[i] = parseRegionInfo(&regionInfoItem)
		}
	}

	return regionInfoList
}

func parseRegionInfo(regionInfo *model.RegionInfo) interface{} {
	if regionInfo == nil {
		return nil
	}
	parsedRegionInfo := make(map[string]interface{})
	parsedRegionInfo["region"] = regionInfo.Region
	parsedRegionInfo["availability_zones"] = regionInfo.AvailabilityZones

	return parsedRegionInfo
}

func parseDBServiceTopologyListWithResData(topology *[]model.DBServiceTopology, d *schema.ResourceData) []interface{} {
	if topology == nil {
		return nil
	}
	dbServiceTopologyList := make([]interface{}, 0)

	if topology != nil {
		dbServiceTopologyList = make([]interface{}, len(*topology))
		for i, dbServiceTopologyItem := range *topology {
			dbServiceTopologyList[i] = parseDBServiceTopology(&dbServiceTopologyItem)
		}
	}

	return dbServiceTopologyList
}

func parseDBServiceTopologyList(topology *[]model.DBServiceTopology) []interface{} {
	if topology == nil {
		return nil
	}
	dbServiceTopologyList := make([]interface{}, 0)

	if topology != nil {
		dbServiceTopologyList = make([]interface{}, len(*topology))
		for i, dbServiceTopologyItem := range *topology {
			dbServiceTopologyList[i] = parseDBServiceTopology(&dbServiceTopologyItem)
		}
	}

	return dbServiceTopologyList
}

func parseDBServiceTopology(topology *model.DBServiceTopology) interface{} {
	if topology == nil {
		return nil
	}
	parsedTopology := make(map[string]interface{})
	parsedTopology["type"] = topology.Type
	parsedTopology["cloud_type"] = topology.CloudType
	parsedTopology["region"] = topology.Region
	parsedTopology["availability_zones"] = topology.AvailabilityZones

	return parsedTopology
}

func parseRPOPolicyConfigWithResData(rpoPolicy *model.RPOPolicyConfig, d *schema.ResourceData) []interface{} {
	if rpoPolicy == nil {
		return nil
	}
	parsedRpoPolicy := make(map[string]interface{})
	if d.Get("rpo_policy") != nil {
		rpoPolicyResourceData := d.Get("rpo_policy").([]interface{})
		if len(rpoPolicyResourceData) > 0 {
			parsedRpoPolicy = (rpoPolicyResourceData[0]).(map[string]interface{})
		}
	}
	parsedRpoPolicy["include_transaction_logs"] = rpoPolicy.IncludeTransactionLogs
	parsedRpoPolicy["enable_auto_snapshot"] = rpoPolicy.EnableAutoSnapshot

	parsedRpoPolicy["enable_auto_backup"] = rpoPolicy.EnableAutoBackup

	var standardPolicy *model.StandardRPOPolicy
	if rpoPolicy.StandardPolicy != standardPolicy {
		parsedRpoPolicy["standard_policy"] = []interface{}{parseStandardRPOPolicy(rpoPolicy.StandardPolicy)}
	}

	var customPolicy *model.CustomRPOPolicy
	if rpoPolicy.CustomPolicy != customPolicy {
		parsedRpoPolicy["custom_policy"] = []interface{}{parseCustomRPOPolicy(rpoPolicy.CustomPolicy)}
	}

	var fullBackupSchedule *model.FullBackupSchedule
	if rpoPolicy.FullBackupSchedule != fullBackupSchedule {
		parsedRpoPolicy["full_backup_schedule"] = []interface{}{parseFullBackupSchedule(rpoPolicy.FullBackupSchedule)}
	}

	var backupRPOConfig *model.RPOPolicyConfigBackupRPOConfig
	if rpoPolicy.BackupRPOConfig != backupRPOConfig {
		parsedRpoPolicy["backup_rpo_config"] = []interface{}{parseRPOPolicyConfigBackupRPOConfig(rpoPolicy.BackupRPOConfig)}
	}

	return []interface{}{parsedRpoPolicy}
}

func parseRPOPolicyConfig(rpoPolicy *model.RPOPolicyConfig) interface{} {
	if rpoPolicy == nil {
		return nil
	}
	parsedRpoPolicy := make(map[string]interface{})
	parsedRpoPolicy["include_transaction_logs"] = rpoPolicy.IncludeTransactionLogs
	parsedRpoPolicy["enable_auto_snapshot"] = rpoPolicy.EnableAutoSnapshot

	parsedRpoPolicy["enable_auto_backup"] = rpoPolicy.EnableAutoBackup

	var standardPolicy *model.StandardRPOPolicy
	if rpoPolicy.StandardPolicy != standardPolicy {
		parsedRpoPolicy["standard_policy"] = []interface{}{parseStandardRPOPolicy(rpoPolicy.StandardPolicy)}
	}

	var customPolicy *model.CustomRPOPolicy
	if rpoPolicy.CustomPolicy != customPolicy {
		parsedRpoPolicy["custom_policy"] = []interface{}{parseCustomRPOPolicy(rpoPolicy.CustomPolicy)}
	}

	var fullBackupSchedule *model.FullBackupSchedule
	if rpoPolicy.FullBackupSchedule != fullBackupSchedule {
		parsedRpoPolicy["full_backup_schedule"] = []interface{}{parseFullBackupSchedule(rpoPolicy.FullBackupSchedule)}
	}

	var backupRPOConfig *model.RPOPolicyConfigBackupRPOConfig
	if rpoPolicy.BackupRPOConfig != backupRPOConfig {
		parsedRpoPolicy["backup_rpo_config"] = []interface{}{parseRPOPolicyConfigBackupRPOConfig(rpoPolicy.BackupRPOConfig)}
	}

	return parsedRpoPolicy
}

func parseStandardRPOPolicy(standardRpoPolicy *model.StandardRPOPolicy) interface{} {
	if standardRpoPolicy == nil {
		return nil
	}
	parsedStandardRpoPolicy := make(map[string]interface{})
	parsedStandardRpoPolicy["retention_days"] = standardRpoPolicy.RetentionDays
	parsedStandardRpoPolicy["include_transaction_logs"] = standardRpoPolicy.IncludeTransactionLogs

	var snapshotStartTime *model.TimeFormat
	if standardRpoPolicy.SnapshotStartTime != snapshotStartTime {
		parsedStandardRpoPolicy["snapshot_start_time"] = []interface{}{parseTimeFormat(standardRpoPolicy.SnapshotStartTime)}
	}

	return parsedStandardRpoPolicy
}

func parseTimeFormat(timeFormat *model.TimeFormat) interface{} {
	if timeFormat == nil {
		return nil
	}
	parsedTimeFormat := make(map[string]interface{})
	parsedTimeFormat["hour"] = timeFormat.Hour
	parsedTimeFormat["minute"] = timeFormat.Minute

	return parsedTimeFormat
}

func parseCustomRPOPolicy(customRpoPolicy *model.CustomRPOPolicy) interface{} {
	if customRpoPolicy == nil {
		return nil
	}
	parsedCustomRpoPolicy := make(map[string]interface{})
	parsedCustomRpoPolicy["name"] = customRpoPolicy.Name

	var schedule *model.ScheduleInfo
	if customRpoPolicy.Schedule != schedule {
		parsedCustomRpoPolicy["schedule"] = []interface{}{parseScheduleInfo(customRpoPolicy.Schedule)}
	}

	return parsedCustomRpoPolicy
}

func parseScheduleInfo(scheduleInfo *model.ScheduleInfo) interface{} {
	if scheduleInfo == nil {
		return nil
	}
	parsedScheduleInfo := make(map[string]interface{})

	var backupStartTime *model.TimeFormat
	if scheduleInfo.BackupStartTime != backupStartTime {
		parsedScheduleInfo["backup_start_time"] = []interface{}{parseTimeFormat(scheduleInfo.BackupStartTime)}
	}

	var dailySchedule *model.DailySchedule
	if scheduleInfo.DailySchedule != dailySchedule {
		parsedScheduleInfo["daily_schedule"] = []interface{}{parseDailySchedule(scheduleInfo.DailySchedule)}
	}

	var weeklySchedule *model.WeeklySchedule
	if scheduleInfo.WeeklySchedule != weeklySchedule {
		parsedScheduleInfo["weekly_schedule"] = []interface{}{parseWeeklySchedule(scheduleInfo.WeeklySchedule)}
	}

	var monthlySchedule *model.MonthlySchedule
	if scheduleInfo.MonthlySchedule != monthlySchedule {
		parsedScheduleInfo["monthly_schedule"] = []interface{}{parseMonthlySchedule(scheduleInfo.MonthlySchedule)}
	}

	var yearlySchedule *model.YearlySchedule
	if scheduleInfo.YearlySchedule != yearlySchedule {
		parsedScheduleInfo["yearly_schedule"] = []interface{}{parseYearlySchedule(scheduleInfo.YearlySchedule)}
	}

	return parsedScheduleInfo
}

func parseDailySchedule(dailySchedule *model.DailySchedule) interface{} {
	if dailySchedule == nil {
		return nil
	}
	parsedDailySchedule := make(map[string]interface{})
	parsedDailySchedule["backups_per_day"] = dailySchedule.BackupsPerDay

	return parsedDailySchedule
}

func parseWeeklySchedule(weeklySchedule *model.WeeklySchedule) interface{} {
	if weeklySchedule == nil {
		return nil
	}
	parsedWeeklySchedule := make(map[string]interface{})
	parsedWeeklySchedule["days"] = weeklySchedule.Days

	return parsedWeeklySchedule
}

func parseMonthlySchedule(monthlySchedule *model.MonthlySchedule) interface{} {
	if monthlySchedule == nil {
		return nil
	}
	parsedMonthlySchedule := make(map[string]interface{})

	var commonSchedule *model.DatesForEachMonth
	if monthlySchedule.CommonSchedule != commonSchedule {
		parsedMonthlySchedule["common_schedule"] = []interface{}{parseDatesForEachMonth(monthlySchedule.CommonSchedule)}
	}

	return parsedMonthlySchedule
}

func parseDatesForEachMonth(datesForEachMonth *model.DatesForEachMonth) interface{} {
	if datesForEachMonth == nil {
		return nil
	}
	parsedDatesForEachMonth := make(map[string]interface{})
	parsedDatesForEachMonth["dates"] = datesForEachMonth.Dates
	parsedDatesForEachMonth["last_day_of_month"] = datesForEachMonth.LastDayOfMonth

	return parsedDatesForEachMonth
}

func parseYearlySchedule(yearlySchedule *model.YearlySchedule) interface{} {
	if yearlySchedule == nil {
		return nil
	}
	parsedYearlySchedule := make(map[string]interface{})

	var commonSchedule *model.CommonYearlySchedule
	if yearlySchedule.CommonSchedule != commonSchedule {
		parsedYearlySchedule["common_schedule"] = []interface{}{parseCommonYearlySchedule(yearlySchedule.CommonSchedule)}
	}

	var monthSpecificSchedule *[]model.MonthWiseDates
	if yearlySchedule.MonthSpecificSchedule != monthSpecificSchedule {
		parsedYearlySchedule["month_specific_schedule"] = parseMonthWiseDatesList(yearlySchedule.MonthSpecificSchedule)
	}

	return parsedYearlySchedule
}

func parseCommonYearlySchedule(commonYearlySchedule *model.CommonYearlySchedule) interface{} {
	if commonYearlySchedule == nil {
		return nil
	}
	parsedCommonYearlySchedule := make(map[string]interface{})
	parsedCommonYearlySchedule["dates"] = commonYearlySchedule.Dates
	parsedCommonYearlySchedule["last_day_of_month"] = commonYearlySchedule.LastDayOfMonth
	parsedCommonYearlySchedule["months"] = commonYearlySchedule.Months

	return parsedCommonYearlySchedule
}

func parseMonthWiseDatesList(monthWiseDates *[]model.MonthWiseDates) []interface{} {
	if monthWiseDates == nil {
		return nil
	}
	monthWiseDatesList := make([]interface{}, 0)

	if monthWiseDates != nil {
		monthWiseDatesList = make([]interface{}, len(*monthWiseDates))
		for i, monthWiseDatesItem := range *monthWiseDates {
			monthWiseDatesList[i] = parseMonthWiseDates(&monthWiseDatesItem)
		}
	}

	return monthWiseDatesList
}

func parseMonthWiseDates(monthWiseDates *model.MonthWiseDates) interface{} {
	if monthWiseDates == nil {
		return nil
	}
	parsedMonthWiseDates := make(map[string]interface{})
	parsedMonthWiseDates["month"] = monthWiseDates.Month
	parsedMonthWiseDates["dates"] = monthWiseDates.Dates

	return parsedMonthWiseDates
}

func parseFullBackupSchedule(fullBackupSchedule *model.FullBackupSchedule) interface{} {
	if fullBackupSchedule == nil {
		return nil
	}
	parsedFullBackupSchedule := make(map[string]interface{})

	var startTime *model.TimeFormat
	if fullBackupSchedule.StartTime != startTime {
		parsedFullBackupSchedule["start_time"] = []interface{}{parseTimeFormat(fullBackupSchedule.StartTime)}
	}

	var weeklySchedule *model.WeeklySchedule
	if fullBackupSchedule.WeeklySchedule != weeklySchedule {
		parsedFullBackupSchedule["weekly_schedule"] = []interface{}{parseWeeklySchedule(fullBackupSchedule.WeeklySchedule)}
	}

	return parsedFullBackupSchedule
}

func parseRPOPolicyConfigBackupRPOConfig(rpoPolicyConfig_backupRpoConfig *model.RPOPolicyConfigBackupRPOConfig) interface{} {
	if rpoPolicyConfig_backupRpoConfig == nil {
		return nil
	}
	parsedRpoPolicyConfig_backupRpoConfig := make(map[string]interface{})

	var fullBackupSchedule *model.FullBackupSchedule
	if rpoPolicyConfig_backupRpoConfig.FullBackupSchedule != fullBackupSchedule {
		parsedRpoPolicyConfig_backupRpoConfig["full_backup_schedule"] = []interface{}{parseFullBackupSchedule(rpoPolicyConfig_backupRpoConfig.FullBackupSchedule)}
	}

	var standardPolicy *model.BackupStandardRPOPolicy
	if rpoPolicyConfig_backupRpoConfig.StandardPolicy != standardPolicy {
		parsedRpoPolicyConfig_backupRpoConfig["standard_policy"] = []interface{}{parseBackupStandardRPOPolicy(rpoPolicyConfig_backupRpoConfig.StandardPolicy)}
	}

	var customPolicy *model.BackupCustomRPOPolicy
	if rpoPolicyConfig_backupRpoConfig.CustomPolicy != customPolicy {
		parsedRpoPolicyConfig_backupRpoConfig["custom_policy"] = []interface{}{parseBackupCustomRPOPolicy(rpoPolicyConfig_backupRpoConfig.CustomPolicy)}
	}

	return parsedRpoPolicyConfig_backupRpoConfig
}

func parseBackupStandardRPOPolicy(backupStandardRpoPolicy *model.BackupStandardRPOPolicy) interface{} {
	if backupStandardRpoPolicy == nil {
		return nil
	}
	parsedBackupStandardRpoPolicy := make(map[string]interface{})
	parsedBackupStandardRpoPolicy["retention_days"] = backupStandardRpoPolicy.RetentionDays

	var backupStartTime *model.TimeFormat
	if backupStandardRpoPolicy.BackupStartTime != backupStartTime {
		parsedBackupStandardRpoPolicy["backup_start_time"] = []interface{}{parseTimeFormat(backupStandardRpoPolicy.BackupStartTime)}
	}

	return parsedBackupStandardRpoPolicy
}

func parseBackupCustomRPOPolicy(backupCustomRpoPolicy *model.BackupCustomRPOPolicy) interface{} {
	if backupCustomRpoPolicy == nil {
		return nil
	}
	parsedBackupCustomRpoPolicy := make(map[string]interface{})
	parsedBackupCustomRpoPolicy["name"] = backupCustomRpoPolicy.Name

	var schedule *model.ScheduleInfo
	if backupCustomRpoPolicy.Schedule != schedule {
		parsedBackupCustomRpoPolicy["schedule"] = []interface{}{parseScheduleInfo(backupCustomRpoPolicy.Schedule)}
	}

	return parsedBackupCustomRpoPolicy
}

func parseTessellDAPServiceDTOListWithResData(daps *[]model.TessellDAPServiceDTO, d *schema.ResourceData) []interface{} {
	if daps == nil {
		return nil
	}
	tessellDAPServiceDTOList := make([]interface{}, 0)

	if daps != nil {
		tessellDAPServiceDTOList = make([]interface{}, len(*daps))
		for i, tessellDAPServiceDTOItem := range *daps {
			tessellDAPServiceDTOList[i] = parseTessellDAPServiceDTO(&tessellDAPServiceDTOItem)
		}
	}

	return tessellDAPServiceDTOList
}

func parseTessellDAPServiceDTOList(daps *[]model.TessellDAPServiceDTO) []interface{} {
	if daps == nil {
		return nil
	}
	tessellDAPServiceDTOList := make([]interface{}, 0)

	if daps != nil {
		tessellDAPServiceDTOList = make([]interface{}, len(*daps))
		for i, tessellDAPServiceDTOItem := range *daps {
			tessellDAPServiceDTOList[i] = parseTessellDAPServiceDTO(&tessellDAPServiceDTOItem)
		}
	}

	return tessellDAPServiceDTOList
}

func parseTessellDAPServiceDTO(daps *model.TessellDAPServiceDTO) interface{} {
	if daps == nil {
		return nil
	}
	parsedDaps := make(map[string]interface{})
	parsedDaps["id"] = daps.Id
	parsedDaps["name"] = daps.Name
	parsedDaps["availability_machine_id"] = daps.AvailabilityMachineId
	parsedDaps["tessell_service_id"] = daps.TessellServiceId
	parsedDaps["service_name"] = daps.ServiceName
	parsedDaps["engine_type"] = daps.EngineType
	parsedDaps["content_type"] = daps.ContentType
	parsedDaps["status"] = daps.Status

	parsedDaps["owner"] = daps.Owner
	parsedDaps["logged_in_user_role"] = daps.LoggedInUserRole

	parsedDaps["date_created"] = daps.DateCreated
	parsedDaps["date_modified"] = daps.DateModified

	var contentInfo *model.DAPContentInfo
	if daps.ContentInfo != contentInfo {
		parsedDaps["content_info"] = []interface{}{parseDAPContentInfo(daps.ContentInfo)}
	}

	var dataAccessConfig *model.DAPRetentionInfo
	if daps.DataAccessConfig != dataAccessConfig {
		parsedDaps["data_access_config"] = []interface{}{parseDAPRetentionInfo(daps.DataAccessConfig)}
	}

	var subscriptionsCloudLocationsAndKey *[]model.SubscriptionsCloudLocationsAndKey
	if daps.SubscriptionsCloudLocationsAndKey != subscriptionsCloudLocationsAndKey {
		parsedDaps["subscriptions_cloud_locations_and_key"] = parseSubscriptionsCloudLocationsAndKeyList(daps.SubscriptionsCloudLocationsAndKey)
	}

	return parsedDaps
}

func parseDAPContentInfo(dapContentInfo *model.DAPContentInfo) interface{} {
	if dapContentInfo == nil {
		return nil
	}
	parsedDapContentInfo := make(map[string]interface{})

	var asIsContent *model.AsIsDAPContent
	if dapContentInfo.AsIsContent != asIsContent {
		parsedDapContentInfo["as_is_content"] = []interface{}{parseAsIsDAPContent(dapContentInfo.AsIsContent)}
	}

	var sanitizedContent *model.SanitizationDAPContent
	if dapContentInfo.SanitizedContent != sanitizedContent {
		parsedDapContentInfo["sanitized_content"] = []interface{}{parseSanitizationDAPContent(dapContentInfo.SanitizedContent)}
	}

	var backupContent *model.BackupDAPContent
	if dapContentInfo.BackupContent != backupContent {
		parsedDapContentInfo["backup_content"] = []interface{}{parseBackupDAPContent(dapContentInfo.BackupContent)}
	}

	return parsedDapContentInfo
}

func parseAsIsDAPContent(asIsDapContent *model.AsIsDAPContent) interface{} {
	if asIsDapContent == nil {
		return nil
	}
	parsedAsIsDapContent := make(map[string]interface{})
	parsedAsIsDapContent["automated"] = asIsDapContent.Automated

	var manual *[]model.DAPManualInfo
	if asIsDapContent.Manual != manual {
		parsedAsIsDapContent["manual"] = parseDAPManualInfoList(asIsDapContent.Manual)
	}

	return parsedAsIsDapContent
}

func parseDAPManualInfoList(dapManualInfo *[]model.DAPManualInfo) []interface{} {
	if dapManualInfo == nil {
		return nil
	}
	dapManualInfoList := make([]interface{}, 0)

	if dapManualInfo != nil {
		dapManualInfoList = make([]interface{}, len(*dapManualInfo))
		for i, dapManualInfoItem := range *dapManualInfo {
			dapManualInfoList[i] = parseDAPManualInfo(&dapManualInfoItem)
		}
	}

	return dapManualInfoList
}

func parseDAPManualInfo(dapManualInfo *model.DAPManualInfo) interface{} {
	if dapManualInfo == nil {
		return nil
	}
	parsedDapManualInfo := make(map[string]interface{})
	parsedDapManualInfo["id"] = dapManualInfo.Id
	parsedDapManualInfo["name"] = dapManualInfo.Name
	parsedDapManualInfo["creation_time"] = dapManualInfo.CreationTime
	parsedDapManualInfo["shared_at"] = dapManualInfo.SharedAt

	return parsedDapManualInfo
}

func parseSanitizationDAPContent(sanitizationDapContent *model.SanitizationDAPContent) interface{} {
	if sanitizationDapContent == nil {
		return nil
	}
	parsedSanitizationDapContent := make(map[string]interface{})

	var automated *model.SanitizationDAPContentAutomated
	if sanitizationDapContent.Automated != automated {
		parsedSanitizationDapContent["automated"] = []interface{}{parseSanitizationDAPContentAutomated(sanitizationDapContent.Automated)}
	}

	var manual *[]model.DAPManualInfo
	if sanitizationDapContent.Manual != manual {
		parsedSanitizationDapContent["manual"] = parseDAPManualInfoList(sanitizationDapContent.Manual)
	}

	return parsedSanitizationDapContent
}

func parseSanitizationDAPContentAutomated(sanitizationDapContent_automated *model.SanitizationDAPContentAutomated) interface{} {
	if sanitizationDapContent_automated == nil {
		return nil
	}
	parsedSanitizationDapContent_automated := make(map[string]interface{})
	parsedSanitizationDapContent_automated["sanitization_schedule_id"] = sanitizationDapContent_automated.SanitizationScheduleId

	return parsedSanitizationDapContent_automated
}

func parseBackupDAPContent(backupDapContent *model.BackupDAPContent) interface{} {
	if backupDapContent == nil {
		return nil
	}
	parsedBackupDapContent := make(map[string]interface{})
	parsedBackupDapContent["automated"] = backupDapContent.Automated

	var manual *[]model.DAPManualInfo
	if backupDapContent.Manual != manual {
		parsedBackupDapContent["manual"] = parseDAPManualInfoList(backupDapContent.Manual)
	}

	return parsedBackupDapContent
}

func parseDAPRetentionInfo(dapRetentionInfo *model.DAPRetentionInfo) interface{} {
	if dapRetentionInfo == nil {
		return nil
	}
	parsedDapRetentionInfo := make(map[string]interface{})
	parsedDapRetentionInfo["pitr"] = dapRetentionInfo.PITR
	parsedDapRetentionInfo["daily_backups"] = dapRetentionInfo.DailyBackups

	return parsedDapRetentionInfo
}

func parseSubscriptionsCloudLocationsAndKeyList(subscriptionsCloudLocationsAndKey *[]model.SubscriptionsCloudLocationsAndKey) []interface{} {
	if subscriptionsCloudLocationsAndKey == nil {
		return nil
	}
	subscriptionsCloudLocationsAndKeyList := make([]interface{}, 0)

	if subscriptionsCloudLocationsAndKey != nil {
		subscriptionsCloudLocationsAndKeyList = make([]interface{}, len(*subscriptionsCloudLocationsAndKey))
		for i, subscriptionsCloudLocationsAndKeyItem := range *subscriptionsCloudLocationsAndKey {
			subscriptionsCloudLocationsAndKeyList[i] = parseSubscriptionsCloudLocationsAndKey(&subscriptionsCloudLocationsAndKeyItem)
		}
	}

	return subscriptionsCloudLocationsAndKeyList
}

func parseSubscriptionsCloudLocationsAndKey(subscriptionsCloudLocationsAndKey *model.SubscriptionsCloudLocationsAndKey) interface{} {
	if subscriptionsCloudLocationsAndKey == nil {
		return nil
	}
	parsedSubscriptionsCloudLocationsAndKey := make(map[string]interface{})
	parsedSubscriptionsCloudLocationsAndKey["subscription_name"] = subscriptionsCloudLocationsAndKey.SubscriptionName
	parsedSubscriptionsCloudLocationsAndKey["cloud_region_and_key"] = subscriptionsCloudLocationsAndKey.CloudRegionAndKey
	parsedSubscriptionsCloudLocationsAndKey["users"] = subscriptionsCloudLocationsAndKey.Users

	return parsedSubscriptionsCloudLocationsAndKey
}

func parseTessellCloneSummaryInfoListWithResData(clones *[]model.TessellCloneSummaryInfo, d *schema.ResourceData) []interface{} {
	if clones == nil {
		return nil
	}
	tessellCloneSummaryInfoList := make([]interface{}, 0)

	if clones != nil {
		tessellCloneSummaryInfoList = make([]interface{}, len(*clones))
		for i, tessellCloneSummaryInfoItem := range *clones {
			tessellCloneSummaryInfoList[i] = parseTessellCloneSummaryInfo(&tessellCloneSummaryInfoItem)
		}
	}

	return tessellCloneSummaryInfoList
}

func parseTessellCloneSummaryInfoList(clones *[]model.TessellCloneSummaryInfo) []interface{} {
	if clones == nil {
		return nil
	}
	tessellCloneSummaryInfoList := make([]interface{}, 0)

	if clones != nil {
		tessellCloneSummaryInfoList = make([]interface{}, len(*clones))
		for i, tessellCloneSummaryInfoItem := range *clones {
			tessellCloneSummaryInfoList[i] = parseTessellCloneSummaryInfo(&tessellCloneSummaryInfoItem)
		}
	}

	return tessellCloneSummaryInfoList
}

func parseTessellCloneSummaryInfo(clones *model.TessellCloneSummaryInfo) interface{} {
	if clones == nil {
		return nil
	}
	parsedClones := make(map[string]interface{})
	parsedClones["id"] = clones.Id
	parsedClones["name"] = clones.Name
	parsedClones["subscription"] = clones.Subscription
	parsedClones["compute_type"] = clones.ComputeType
	parsedClones["status"] = clones.Status

	parsedClones["clone_info"] = clones.CloneInfo
	parsedClones["owner"] = clones.Owner
	parsedClones["date_created"] = clones.DateCreated

	var cloudAvailability *[]model.CloudRegionInfo
	if clones.CloudAvailability != cloudAvailability {
		parsedClones["cloud_availability"] = parseCloudRegionInfoList(clones.CloudAvailability)
	}

	var instances *[]model.TessellServiceInstanceDTO
	if clones.Instances != instances {
		parsedClones["instances"] = parseTessellServiceInstanceDTOList(clones.Instances)
	}

	return parsedClones
}

func parseTessellServiceInstanceDTOList(tessellServiceInstanceDTO *[]model.TessellServiceInstanceDTO) []interface{} {
	if tessellServiceInstanceDTO == nil {
		return nil
	}
	tessellServiceInstanceDTOList := make([]interface{}, 0)

	if tessellServiceInstanceDTO != nil {
		tessellServiceInstanceDTOList = make([]interface{}, len(*tessellServiceInstanceDTO))
		for i, tessellServiceInstanceDTOItem := range *tessellServiceInstanceDTO {
			tessellServiceInstanceDTOList[i] = parseTessellServiceInstanceDTO(&tessellServiceInstanceDTOItem)
		}
	}

	return tessellServiceInstanceDTOList
}

func parseTessellServiceInstanceDTO(tessellServiceInstanceDTO *model.TessellServiceInstanceDTO) interface{} {
	if tessellServiceInstanceDTO == nil {
		return nil
	}
	parsedTessellServiceInstanceDTO := make(map[string]interface{})
	parsedTessellServiceInstanceDTO["id"] = tessellServiceInstanceDTO.Id
	parsedTessellServiceInstanceDTO["name"] = tessellServiceInstanceDTO.Name
	parsedTessellServiceInstanceDTO["instance_group_name"] = tessellServiceInstanceDTO.InstanceGroupName
	parsedTessellServiceInstanceDTO["type"] = tessellServiceInstanceDTO.Type
	parsedTessellServiceInstanceDTO["role"] = tessellServiceInstanceDTO.Role
	parsedTessellServiceInstanceDTO["status"] = tessellServiceInstanceDTO.Status
	parsedTessellServiceInstanceDTO["tessell_service_id"] = tessellServiceInstanceDTO.TessellServiceId
	parsedTessellServiceInstanceDTO["cloud"] = tessellServiceInstanceDTO.Cloud
	parsedTessellServiceInstanceDTO["region"] = tessellServiceInstanceDTO.Region
	parsedTessellServiceInstanceDTO["availability_zone"] = tessellServiceInstanceDTO.AvailabilityZone
	parsedTessellServiceInstanceDTO["instance_group_id"] = tessellServiceInstanceDTO.InstanceGroupId
	parsedTessellServiceInstanceDTO["compute_type"] = tessellServiceInstanceDTO.ComputeType

	parsedTessellServiceInstanceDTO["compute_id"] = tessellServiceInstanceDTO.ComputeId
	parsedTessellServiceInstanceDTO["compute_name"] = tessellServiceInstanceDTO.ComputeName
	parsedTessellServiceInstanceDTO["storage"] = tessellServiceInstanceDTO.Storage
	parsedTessellServiceInstanceDTO["data_volume_iops"] = tessellServiceInstanceDTO.DataVolumeIops
	parsedTessellServiceInstanceDTO["throughput"] = tessellServiceInstanceDTO.Throughput
	parsedTessellServiceInstanceDTO["enable_perf_insights"] = tessellServiceInstanceDTO.EnablePerfInsights

	parsedTessellServiceInstanceDTO["vpc"] = tessellServiceInstanceDTO.VPC
	parsedTessellServiceInstanceDTO["public_subnet"] = tessellServiceInstanceDTO.PublicSubnet
	parsedTessellServiceInstanceDTO["private_subnet"] = tessellServiceInstanceDTO.PrivateSubnet
	parsedTessellServiceInstanceDTO["encryption_key"] = tessellServiceInstanceDTO.EncryptionKey
	parsedTessellServiceInstanceDTO["software_image"] = tessellServiceInstanceDTO.SoftwareImage
	parsedTessellServiceInstanceDTO["software_image_version"] = tessellServiceInstanceDTO.SoftwareImageVersion
	parsedTessellServiceInstanceDTO["date_created"] = tessellServiceInstanceDTO.DateCreated

	parsedTessellServiceInstanceDTO["last_started_at"] = tessellServiceInstanceDTO.LastStartedAt
	parsedTessellServiceInstanceDTO["last_stopped_at"] = tessellServiceInstanceDTO.LastStoppedAt
	parsedTessellServiceInstanceDTO["sync_mode"] = tessellServiceInstanceDTO.SyncMode

	var awsInfraConfig *model.AwsInfraConfig
	if tessellServiceInstanceDTO.AwsInfraConfig != awsInfraConfig {
		parsedTessellServiceInstanceDTO["aws_infra_config"] = []interface{}{parseAwsInfraConfig(tessellServiceInstanceDTO.AwsInfraConfig)}
	}

	var parameterProfile *model.ParameterProfile
	if tessellServiceInstanceDTO.ParameterProfile != parameterProfile {
		parsedTessellServiceInstanceDTO["parameter_profile"] = []interface{}{parseParameterProfile(tessellServiceInstanceDTO.ParameterProfile)}
	}

	var optionProfile *model.OptionProfile
	if tessellServiceInstanceDTO.OptionProfile != optionProfile {
		parsedTessellServiceInstanceDTO["option_profile"] = []interface{}{parseOptionProfile(tessellServiceInstanceDTO.OptionProfile)}
	}

	var monitoringConfig *model.MonitoringConfig
	if tessellServiceInstanceDTO.MonitoringConfig != monitoringConfig {
		parsedTessellServiceInstanceDTO["monitoring_config"] = []interface{}{parseMonitoringConfig(tessellServiceInstanceDTO.MonitoringConfig)}
	}

	var connectString *model.TessellServiceInstanceConnectString
	if tessellServiceInstanceDTO.ConnectString != connectString {
		parsedTessellServiceInstanceDTO["connect_string"] = []interface{}{parseTessellServiceInstanceConnectString(tessellServiceInstanceDTO.ConnectString)}
	}

	var updatesInProgress *[]model.TessellResourceUpdateInfo
	if tessellServiceInstanceDTO.UpdatesInProgress != updatesInProgress {
		parsedTessellServiceInstanceDTO["updates_in_progress"] = parseTessellResourceUpdateInfoList(tessellServiceInstanceDTO.UpdatesInProgress)
	}

	var engineConfiguration *model.ServiceInstanceEngineInfo
	if tessellServiceInstanceDTO.EngineConfiguration != engineConfiguration {
		parsedTessellServiceInstanceDTO["engine_configuration"] = []interface{}{parseServiceInstanceEngineInfo(tessellServiceInstanceDTO.EngineConfiguration)}
	}

	var computeConfig *model.InstanceComputeConfig
	if tessellServiceInstanceDTO.ComputeConfig != computeConfig {
		parsedTessellServiceInstanceDTO["compute_config"] = []interface{}{parseInstanceComputeConfig(tessellServiceInstanceDTO.ComputeConfig)}
	}

	var storageConfig *model.InstanceStorageConfig
	if tessellServiceInstanceDTO.StorageConfig != storageConfig {
		parsedTessellServiceInstanceDTO["storage_config"] = []interface{}{parseInstanceStorageConfig(tessellServiceInstanceDTO.StorageConfig)}
	}

	var archiveStorageConfig *model.InstanceStorageConfig
	if tessellServiceInstanceDTO.ArchiveStorageConfig != archiveStorageConfig {
		parsedTessellServiceInstanceDTO["archive_storage_config"] = []interface{}{parseInstanceStorageConfig(tessellServiceInstanceDTO.ArchiveStorageConfig)}
	}

	var privateLinkInfo *model.PrivateLinkInfo
	if tessellServiceInstanceDTO.PrivateLinkInfo != privateLinkInfo {
		parsedTessellServiceInstanceDTO["private_link_info"] = []interface{}{parsePrivateLinkInfo(tessellServiceInstanceDTO.PrivateLinkInfo)}
	}

	return parsedTessellServiceInstanceDTO
}

func parseAwsInfraConfig(awsInfraConfig *model.AwsInfraConfig) interface{} {
	if awsInfraConfig == nil {
		return nil
	}
	parsedAwsInfraConfig := make(map[string]interface{})

	var awsCpuOptions *model.AwsCpuOptions
	if awsInfraConfig.AwsCpuOptions != awsCpuOptions {
		parsedAwsInfraConfig["aws_cpu_options"] = []interface{}{parseAwsCpuOptions(awsInfraConfig.AwsCpuOptions)}
	}

	return parsedAwsInfraConfig
}

func parseAwsCpuOptions(awsCpuOptions *model.AwsCpuOptions) interface{} {
	if awsCpuOptions == nil {
		return nil
	}
	parsedAwsCpuOptions := make(map[string]interface{})
	parsedAwsCpuOptions["vcpus"] = awsCpuOptions.Vcpus

	return parsedAwsCpuOptions
}

func parseParameterProfile(parameterProfile *model.ParameterProfile) interface{} {
	if parameterProfile == nil {
		return nil
	}
	parsedParameterProfile := make(map[string]interface{})
	parsedParameterProfile["id"] = parameterProfile.Id
	parsedParameterProfile["name"] = parameterProfile.Name
	parsedParameterProfile["version"] = parameterProfile.Version
	parsedParameterProfile["status"] = parameterProfile.Status

	return parsedParameterProfile
}

func parseOptionProfile(optionProfile *model.OptionProfile) interface{} {
	if optionProfile == nil {
		return nil
	}
	parsedOptionProfile := make(map[string]interface{})
	parsedOptionProfile["id"] = optionProfile.Id
	parsedOptionProfile["name"] = optionProfile.Name
	parsedOptionProfile["version"] = optionProfile.Version
	parsedOptionProfile["status"] = optionProfile.Status

	return parsedOptionProfile
}

func parseMonitoringConfig(monitoringConfig *model.MonitoringConfig) interface{} {
	if monitoringConfig == nil {
		return nil
	}
	parsedMonitoringConfig := make(map[string]interface{})

	var perfInsights *model.PerfInsightsConfig
	if monitoringConfig.PerfInsights != perfInsights {
		parsedMonitoringConfig["perf_insights"] = []interface{}{parsePerfInsightsConfig(monitoringConfig.PerfInsights)}
	}

	return parsedMonitoringConfig
}

func parsePerfInsightsConfig(perfInsightsConfig *model.PerfInsightsConfig) interface{} {
	if perfInsightsConfig == nil {
		return nil
	}
	parsedPerfInsightsConfig := make(map[string]interface{})
	parsedPerfInsightsConfig["perf_insights_enabled"] = perfInsightsConfig.PerfInsightsEnabled
	parsedPerfInsightsConfig["monitoring_deployment_id"] = perfInsightsConfig.MonitoringDeploymentId
	parsedPerfInsightsConfig["status"] = perfInsightsConfig.Status

	return parsedPerfInsightsConfig
}

func parseTessellServiceInstanceConnectString(tessellServiceInstanceConnectString *model.TessellServiceInstanceConnectString) interface{} {
	if tessellServiceInstanceConnectString == nil {
		return nil
	}
	parsedTessellServiceInstanceConnectString := make(map[string]interface{})
	parsedTessellServiceInstanceConnectString["connect_descriptor"] = tessellServiceInstanceConnectString.ConnectDescriptor
	parsedTessellServiceInstanceConnectString["master_user"] = tessellServiceInstanceConnectString.MasterUser
	parsedTessellServiceInstanceConnectString["endpoint"] = tessellServiceInstanceConnectString.Endpoint
	parsedTessellServiceInstanceConnectString["service_port"] = tessellServiceInstanceConnectString.ServicePort

	return parsedTessellServiceInstanceConnectString
}

func parseTessellResourceUpdateInfoList(tessellResourceUpdateInfo *[]model.TessellResourceUpdateInfo) []interface{} {
	if tessellResourceUpdateInfo == nil {
		return nil
	}
	tessellResourceUpdateInfoList := make([]interface{}, 0)

	if tessellResourceUpdateInfo != nil {
		tessellResourceUpdateInfoList = make([]interface{}, len(*tessellResourceUpdateInfo))
		for i, tessellResourceUpdateInfoItem := range *tessellResourceUpdateInfo {
			tessellResourceUpdateInfoList[i] = parseTessellResourceUpdateInfo(&tessellResourceUpdateInfoItem)
		}
	}

	return tessellResourceUpdateInfoList
}

func parseTessellResourceUpdateInfo(tessellResourceUpdateInfo *model.TessellResourceUpdateInfo) interface{} {
	if tessellResourceUpdateInfo == nil {
		return nil
	}
	parsedTessellResourceUpdateInfo := make(map[string]interface{})
	parsedTessellResourceUpdateInfo["update_type"] = tessellResourceUpdateInfo.UpdateType
	parsedTessellResourceUpdateInfo["reference_id"] = tessellResourceUpdateInfo.ReferenceId
	parsedTessellResourceUpdateInfo["submitted_at"] = tessellResourceUpdateInfo.SubmittedAt
	parsedTessellResourceUpdateInfo["update_info"] = tessellResourceUpdateInfo.UpdateInfo

	return parsedTessellResourceUpdateInfo
}

func parseServiceInstanceEngineInfo(serviceInstanceEngineInfo *model.ServiceInstanceEngineInfo) interface{} {
	if serviceInstanceEngineInfo == nil {
		return nil
	}
	parsedServiceInstanceEngineInfo := make(map[string]interface{})

	var oracleConfig *model.ServiceInstanceOracleEngineConfig
	if serviceInstanceEngineInfo.OracleConfig != oracleConfig {
		parsedServiceInstanceEngineInfo["oracle_config"] = []interface{}{parseServiceInstanceOracleEngineConfig(serviceInstanceEngineInfo.OracleConfig)}
	}

	return parsedServiceInstanceEngineInfo
}

func parseServiceInstanceOracleEngineConfig(serviceInstanceOracleEngineConfig *model.ServiceInstanceOracleEngineConfig) interface{} {
	if serviceInstanceOracleEngineConfig == nil {
		return nil
	}
	parsedServiceInstanceOracleEngineConfig := make(map[string]interface{})
	parsedServiceInstanceOracleEngineConfig["access_mode"] = serviceInstanceOracleEngineConfig.AccessMode

	return parsedServiceInstanceOracleEngineConfig
}

func parseInstanceComputeConfig(instanceComputeConfig *model.InstanceComputeConfig) interface{} {
	if instanceComputeConfig == nil {
		return nil
	}
	parsedInstanceComputeConfig := make(map[string]interface{})
	parsedInstanceComputeConfig["provider"] = instanceComputeConfig.Provider

	var exadataConfig *model.InstanceExadataComputeConfig
	if instanceComputeConfig.ExadataConfig != exadataConfig {
		parsedInstanceComputeConfig["exadata_config"] = []interface{}{parseInstanceExadataComputeConfig(instanceComputeConfig.ExadataConfig)}
	}

	return parsedInstanceComputeConfig
}

func parseInstanceExadataComputeConfig(instanceExadataComputeConfig *model.InstanceExadataComputeConfig) interface{} {
	if instanceExadataComputeConfig == nil {
		return nil
	}
	parsedInstanceExadataComputeConfig := make(map[string]interface{})
	parsedInstanceExadataComputeConfig["infrastructure_id"] = instanceExadataComputeConfig.InfrastructureId
	parsedInstanceExadataComputeConfig["infrastructure_name"] = instanceExadataComputeConfig.InfrastructureName
	parsedInstanceExadataComputeConfig["vm_cluster_id"] = instanceExadataComputeConfig.VmClusterId
	parsedInstanceExadataComputeConfig["vm_cluster_name"] = instanceExadataComputeConfig.VmClusterName
	parsedInstanceExadataComputeConfig["vcpus"] = instanceExadataComputeConfig.Vcpus
	parsedInstanceExadataComputeConfig["memory"] = instanceExadataComputeConfig.Memory

	return parsedInstanceExadataComputeConfig
}

func parseInstanceStorageConfig(storageConfig *model.InstanceStorageConfig) interface{} {
	if storageConfig == nil {
		return nil
	}
	parsedStorageConfig := make(map[string]interface{})
	parsedStorageConfig["provider"] = storageConfig.Provider
	parsedStorageConfig["volume_type"] = storageConfig.VolumeType

	var fsxNetAppConfig *model.InstanceFsxNetAppConfig
	if storageConfig.FsxNetAppConfig != fsxNetAppConfig {
		parsedStorageConfig["fsx_net_app_config"] = []interface{}{parseInstanceFsxNetAppConfig(storageConfig.FsxNetAppConfig)}
	}

	var azureNetAppConfig *model.InstanceAzureNetAppConfig
	if storageConfig.AzureNetAppConfig != azureNetAppConfig {
		parsedStorageConfig["azure_net_app_config"] = []interface{}{parseInstanceAzureNetAppConfig(storageConfig.AzureNetAppConfig)}
	}

	return parsedStorageConfig
}

func parseInstanceFsxNetAppConfig(instanceFsxNetAppConfig *model.InstanceFsxNetAppConfig) interface{} {
	if instanceFsxNetAppConfig == nil {
		return nil
	}
	parsedInstanceFsxNetAppConfig := make(map[string]interface{})
	parsedInstanceFsxNetAppConfig["file_system_name"] = instanceFsxNetAppConfig.FileSystemName
	parsedInstanceFsxNetAppConfig["svm_name"] = instanceFsxNetAppConfig.SvmName
	parsedInstanceFsxNetAppConfig["volume_name"] = instanceFsxNetAppConfig.VolumeName
	parsedInstanceFsxNetAppConfig["file_system_id"] = instanceFsxNetAppConfig.FileSystemId
	parsedInstanceFsxNetAppConfig["svm_id"] = instanceFsxNetAppConfig.SvmId

	return parsedInstanceFsxNetAppConfig
}

func parseInstanceAzureNetAppConfig(instanceAzureNetAppConfig *model.InstanceAzureNetAppConfig) interface{} {
	if instanceAzureNetAppConfig == nil {
		return nil
	}
	parsedInstanceAzureNetAppConfig := make(map[string]interface{})
	parsedInstanceAzureNetAppConfig["azure_net_app_name"] = instanceAzureNetAppConfig.AzureNetAppName
	parsedInstanceAzureNetAppConfig["capacity_pool_name"] = instanceAzureNetAppConfig.CapacityPoolName
	parsedInstanceAzureNetAppConfig["volume_name"] = instanceAzureNetAppConfig.VolumeName
	parsedInstanceAzureNetAppConfig["azure_net_app_id"] = instanceAzureNetAppConfig.AzureNetAppId
	parsedInstanceAzureNetAppConfig["capacity_pool_id"] = instanceAzureNetAppConfig.CapacityPoolId
	parsedInstanceAzureNetAppConfig["delegated_subnet_id"] = instanceAzureNetAppConfig.DelegatedSubnetId
	parsedInstanceAzureNetAppConfig["delegated_subnet_name"] = instanceAzureNetAppConfig.DelegatedSubnetName

	parsedInstanceAzureNetAppConfig["network_features"] = instanceAzureNetAppConfig.NetworkFeatures
	parsedInstanceAzureNetAppConfig["service_level"] = instanceAzureNetAppConfig.ServiceLevel

	var encryptionKeyInfo *model.AzureNetAppEncryptionKeyInfo
	if instanceAzureNetAppConfig.EncryptionKeyInfo != encryptionKeyInfo {
		parsedInstanceAzureNetAppConfig["encryption_key_info"] = []interface{}{parseAzureNetAppEncryptionKeyInfo(instanceAzureNetAppConfig.EncryptionKeyInfo)}
	}

	return parsedInstanceAzureNetAppConfig
}

func parseAzureNetAppEncryptionKeyInfo(azureNetAppEncryptionKeyInfo *model.AzureNetAppEncryptionKeyInfo) interface{} {
	if azureNetAppEncryptionKeyInfo == nil {
		return nil
	}
	parsedAzureNetAppEncryptionKeyInfo := make(map[string]interface{})
	parsedAzureNetAppEncryptionKeyInfo["id"] = azureNetAppEncryptionKeyInfo.Id
	parsedAzureNetAppEncryptionKeyInfo["name"] = azureNetAppEncryptionKeyInfo.Name
	parsedAzureNetAppEncryptionKeyInfo["key_vault_cloud_resource_id"] = azureNetAppEncryptionKeyInfo.KeyVaultCloudResourceId
	parsedAzureNetAppEncryptionKeyInfo["key_source"] = azureNetAppEncryptionKeyInfo.KeySource

	return parsedAzureNetAppEncryptionKeyInfo
}

func parsePrivateLinkInfo(privateLinkInfo *model.PrivateLinkInfo) interface{} {
	if privateLinkInfo == nil {
		return nil
	}
	parsedPrivateLinkInfo := make(map[string]interface{})
	parsedPrivateLinkInfo["id"] = privateLinkInfo.Id
	parsedPrivateLinkInfo["status"] = privateLinkInfo.Status
	parsedPrivateLinkInfo["endpoint_service_name"] = privateLinkInfo.EndpointServiceName
	parsedPrivateLinkInfo["private_link_service_alias"] = privateLinkInfo.PrivateLinkServiceAlias
	parsedPrivateLinkInfo["service_principals"] = privateLinkInfo.ServicePrincipals
	parsedPrivateLinkInfo["client_azure_subscription_ids"] = privateLinkInfo.ClientAzureSubscriptionIds

	return parsedPrivateLinkInfo
}

func parseBackupDownloadConfigWithResData(backupDownloadConfig *model.BackupDownloadConfig, d *schema.ResourceData) []interface{} {
	if backupDownloadConfig == nil {
		return nil
	}
	parsedBackupDownloadConfig := make(map[string]interface{})
	if d.Get("backup_download_config") != nil {
		backupDownloadConfigResourceData := d.Get("backup_download_config").([]interface{})
		if len(backupDownloadConfigResourceData) > 0 {
			parsedBackupDownloadConfig = (backupDownloadConfigResourceData[0]).(map[string]interface{})
		}
	}
	parsedBackupDownloadConfig["allow_backup_downloads_for_all_users"] = backupDownloadConfig.AllowBackupDownloadsForAllUsers
	parsedBackupDownloadConfig["allow_backup_downloads"] = backupDownloadConfig.AllowBackupDownloads

	return []interface{}{parsedBackupDownloadConfig}
}

func parseBackupDownloadConfig(backupDownloadConfig *model.BackupDownloadConfig) interface{} {
	if backupDownloadConfig == nil {
		return nil
	}
	parsedBackupDownloadConfig := make(map[string]interface{})
	parsedBackupDownloadConfig["allow_backup_downloads_for_all_users"] = backupDownloadConfig.AllowBackupDownloadsForAllUsers
	parsedBackupDownloadConfig["allow_backup_downloads"] = backupDownloadConfig.AllowBackupDownloads

	return parsedBackupDownloadConfig
}

func parseStorageConfigPayloadWithResData(storageConfig *model.StorageConfigPayload, d *schema.ResourceData) []interface{} {
	if storageConfig == nil {
		return nil
	}
	parsedStorageConfig := make(map[string]interface{})
	if d.Get("storage_config") != nil {
		storageConfigResourceData := d.Get("storage_config").([]interface{})
		if len(storageConfigResourceData) > 0 {
			parsedStorageConfig = (storageConfigResourceData[0]).(map[string]interface{})
		}
	}
	parsedStorageConfig["provider"] = storageConfig.Provider

	var fsxNetAppConfig *model.FsxNetAppConfigPayload
	if storageConfig.FsxNetAppConfig != fsxNetAppConfig {
		parsedStorageConfig["fsx_net_app_config"] = []interface{}{parseFsxNetAppConfigPayload(storageConfig.FsxNetAppConfig)}
	}

	var azureNetAppConfig *model.AzureNetAppConfigPayload
	if storageConfig.AzureNetAppConfig != azureNetAppConfig {
		parsedStorageConfig["azure_net_app_config"] = []interface{}{parseAzureNetAppConfigPayload(storageConfig.AzureNetAppConfig)}
	}

	return []interface{}{parsedStorageConfig}
}

func parseStorageConfigPayload(storageConfig *model.StorageConfigPayload) interface{} {
	if storageConfig == nil {
		return nil
	}
	parsedStorageConfig := make(map[string]interface{})
	parsedStorageConfig["provider"] = storageConfig.Provider

	var fsxNetAppConfig *model.FsxNetAppConfigPayload
	if storageConfig.FsxNetAppConfig != fsxNetAppConfig {
		parsedStorageConfig["fsx_net_app_config"] = []interface{}{parseFsxNetAppConfigPayload(storageConfig.FsxNetAppConfig)}
	}

	var azureNetAppConfig *model.AzureNetAppConfigPayload
	if storageConfig.AzureNetAppConfig != azureNetAppConfig {
		parsedStorageConfig["azure_net_app_config"] = []interface{}{parseAzureNetAppConfigPayload(storageConfig.AzureNetAppConfig)}
	}

	return parsedStorageConfig
}

func parseFsxNetAppConfigPayload(fsxNetAppConfigPayload *model.FsxNetAppConfigPayload) interface{} {
	if fsxNetAppConfigPayload == nil {
		return nil
	}
	parsedFsxNetAppConfigPayload := make(map[string]interface{})
	parsedFsxNetAppConfigPayload["file_system_id"] = fsxNetAppConfigPayload.FileSystemId
	parsedFsxNetAppConfigPayload["svm_id"] = fsxNetAppConfigPayload.SvmId

	return parsedFsxNetAppConfigPayload
}

func parseAzureNetAppConfigPayload(azureNetAppConfigPayload *model.AzureNetAppConfigPayload) interface{} {
	if azureNetAppConfigPayload == nil {
		return nil
	}
	parsedAzureNetAppConfigPayload := make(map[string]interface{})
	parsedAzureNetAppConfigPayload["azure_net_app_id"] = azureNetAppConfigPayload.AzureNetAppId
	parsedAzureNetAppConfigPayload["capacity_pool_id"] = azureNetAppConfigPayload.CapacityPoolId

	var configurations *model.AzureNetAppConfigPayloadConfigurations
	if azureNetAppConfigPayload.Configurations != configurations {
		parsedAzureNetAppConfigPayload["configurations"] = []interface{}{parseAzureNetAppConfigPayloadConfigurations(azureNetAppConfigPayload.Configurations)}
	}

	return parsedAzureNetAppConfigPayload
}

func parseAzureNetAppConfigPayloadConfigurations(azureNetAppConfigPayload_configurations *model.AzureNetAppConfigPayloadConfigurations) interface{} {
	if azureNetAppConfigPayload_configurations == nil {
		return nil
	}
	parsedAzureNetAppConfigPayload_configurations := make(map[string]interface{})
	parsedAzureNetAppConfigPayload_configurations["network_features"] = azureNetAppConfigPayload_configurations.NetworkFeatures

	return parsedAzureNetAppConfigPayload_configurations
}
