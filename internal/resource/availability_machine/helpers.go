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

	var instances *[]model.TessellServiceInstanceOpsDTO
	if clones.Instances != instances {
		parsedClones["instances"] = parseTessellServiceInstanceOpsDTOList(clones.Instances)
	}

	return parsedClones
}

func parseTessellServiceInstanceOpsDTOList(tessellServiceInstanceOpsDTO *[]model.TessellServiceInstanceOpsDTO) []interface{} {
	if tessellServiceInstanceOpsDTO == nil {
		return nil
	}
	tessellServiceInstanceOpsDTOList := make([]interface{}, 0)

	if tessellServiceInstanceOpsDTO != nil {
		tessellServiceInstanceOpsDTOList = make([]interface{}, len(*tessellServiceInstanceOpsDTO))
		for i, tessellServiceInstanceOpsDTOItem := range *tessellServiceInstanceOpsDTO {
			tessellServiceInstanceOpsDTOList[i] = parseTessellServiceInstanceOpsDTO(&tessellServiceInstanceOpsDTOItem)
		}
	}

	return tessellServiceInstanceOpsDTOList
}

func parseTessellServiceInstanceOpsDTO(tessellServiceInstanceOpsDTO *model.TessellServiceInstanceOpsDTO) interface{} {
	if tessellServiceInstanceOpsDTO == nil {
		return nil
	}
	parsedTessellServiceInstanceOpsDTO := make(map[string]interface{})
	parsedTessellServiceInstanceOpsDTO["id"] = tessellServiceInstanceOpsDTO.Id
	parsedTessellServiceInstanceOpsDTO["name"] = tessellServiceInstanceOpsDTO.Name
	parsedTessellServiceInstanceOpsDTO["compute_name"] = tessellServiceInstanceOpsDTO.ComputeName
	parsedTessellServiceInstanceOpsDTO["description"] = tessellServiceInstanceOpsDTO.Description
	parsedTessellServiceInstanceOpsDTO["tessell_service_id"] = tessellServiceInstanceOpsDTO.TessellServiceId
	parsedTessellServiceInstanceOpsDTO["compute_resource_id"] = tessellServiceInstanceOpsDTO.ComputeResourceId
	parsedTessellServiceInstanceOpsDTO["cloud_location_id"] = tessellServiceInstanceOpsDTO.CloudLocationId
	parsedTessellServiceInstanceOpsDTO["parameter_profile_id"] = tessellServiceInstanceOpsDTO.ParameterProfileId
	parsedTessellServiceInstanceOpsDTO["cloud_account_id"] = tessellServiceInstanceOpsDTO.CloudAccountId
	parsedTessellServiceInstanceOpsDTO["instance_group_id"] = tessellServiceInstanceOpsDTO.InstanceGroupId
	parsedTessellServiceInstanceOpsDTO["type"] = tessellServiceInstanceOpsDTO.Type
	parsedTessellServiceInstanceOpsDTO["role"] = tessellServiceInstanceOpsDTO.Role
	parsedTessellServiceInstanceOpsDTO["user_visible_role"] = tessellServiceInstanceOpsDTO.UserVisibleRole
	parsedTessellServiceInstanceOpsDTO["status"] = tessellServiceInstanceOpsDTO.Status
	parsedTessellServiceInstanceOpsDTO["plugin_status"] = tessellServiceInstanceOpsDTO.PluginStatus

	parsedTessellServiceInstanceOpsDTO["date_created"] = tessellServiceInstanceOpsDTO.DateCreated
	parsedTessellServiceInstanceOpsDTO["date_modified"] = tessellServiceInstanceOpsDTO.DateModified
	parsedTessellServiceInstanceOpsDTO["date_modified_by_user"] = tessellServiceInstanceOpsDTO.DateModifieDByUser

	parsedTessellServiceInstanceOpsDTO["last_started_at"] = tessellServiceInstanceOpsDTO.LastStartedAt
	parsedTessellServiceInstanceOpsDTO["last_stopped_at"] = tessellServiceInstanceOpsDTO.LastStoppedAt
	parsedTessellServiceInstanceOpsDTO["last_degraded_at"] = tessellServiceInstanceOpsDTO.LastDegradedAt
	parsedTessellServiceInstanceOpsDTO["deleted_for_user_at"] = tessellServiceInstanceOpsDTO.DeletedForUserAt
	parsedTessellServiceInstanceOpsDTO["is_consumable"] = tessellServiceInstanceOpsDTO.IsConsumable

	var connectionInfo *model.TessellServiceInstanceConnectionInfo
	if tessellServiceInstanceOpsDTO.ConnectionInfo != connectionInfo {
		parsedTessellServiceInstanceOpsDTO["connection_info"] = []interface{}{parseTessellServiceInstanceConnectionInfo(tessellServiceInstanceOpsDTO.ConnectionInfo)}
	}

	var genericInfo *model.ServiceInstanceGenericInfo
	if tessellServiceInstanceOpsDTO.GenericInfo != genericInfo {
		parsedTessellServiceInstanceOpsDTO["generic_info"] = []interface{}{parseServiceInstanceGenericInfo(tessellServiceInstanceOpsDTO.GenericInfo)}
	}

	var licenseInfo *model.DBLicenseInfo
	if tessellServiceInstanceOpsDTO.LicenseInfo != licenseInfo {
		parsedTessellServiceInstanceOpsDTO["license_info"] = []interface{}{parseDBLicenseInfo(tessellServiceInstanceOpsDTO.LicenseInfo)}
	}

	var monitoringConfig *model.MonitoringConfig
	if tessellServiceInstanceOpsDTO.MonitoringConfig != monitoringConfig {
		parsedTessellServiceInstanceOpsDTO["monitoring_config"] = []interface{}{parseMonitoringConfig(tessellServiceInstanceOpsDTO.MonitoringConfig)}
	}

	var metadata *model.TessellServiceInstanceMetadata
	if tessellServiceInstanceOpsDTO.Metadata != metadata {
		parsedTessellServiceInstanceOpsDTO["metadata"] = []interface{}{parseTessellServiceInstanceMetadata(tessellServiceInstanceOpsDTO.Metadata)}
	}

	var driverInfo *model.TessellServiceInstanceDriverInfo
	if tessellServiceInstanceOpsDTO.DriverInfo != driverInfo {
		parsedTessellServiceInstanceOpsDTO["driver_info"] = []interface{}{parseTessellServiceInstanceDriverInfo(tessellServiceInstanceOpsDTO.DriverInfo)}
	}

	var updatesInProgressInfo *model.TessellServiceInstanceInProgressUpdateInfoOps
	if tessellServiceInstanceOpsDTO.UpdatesInProgressInfo != updatesInProgressInfo {
		parsedTessellServiceInstanceOpsDTO["updates_in_progress_info"] = []interface{}{parseTessellServiceInstanceInProgressUpdateInfoOps(tessellServiceInstanceOpsDTO.UpdatesInProgressInfo)}
	}

	var tessellAgentLcmInfo *model.TessellAgentLcmInfo
	if tessellServiceInstanceOpsDTO.TessellAgentLcmInfo != tessellAgentLcmInfo {
		parsedTessellServiceInstanceOpsDTO["tessell_agent_lcm_info"] = []interface{}{parseTessellAgentLcmInfo(tessellServiceInstanceOpsDTO.TessellAgentLcmInfo)}
	}

	var computeResource *model.TessellComputeResourceOpsDTO
	if tessellServiceInstanceOpsDTO.ComputeResource != computeResource {
		parsedTessellServiceInstanceOpsDTO["compute_resource"] = []interface{}{parseTessellComputeResourceOpsDTO(tessellServiceInstanceOpsDTO.ComputeResource)}
	}

	return parsedTessellServiceInstanceOpsDTO
}

func parseTessellServiceInstanceConnectionInfo(tessellServiceInstanceConnectionInfo *model.TessellServiceInstanceConnectionInfo) interface{} {
	if tessellServiceInstanceConnectionInfo == nil {
		return nil
	}
	parsedTessellServiceInstanceConnectionInfo := make(map[string]interface{})

	parsedTessellServiceInstanceConnectionInfo["data"] = tessellServiceInstanceConnectionInfo.Data

	var connectString *model.TessellServiceInstanceConnectString
	if tessellServiceInstanceConnectionInfo.ConnectString != connectString {
		parsedTessellServiceInstanceConnectionInfo["connect_string"] = []interface{}{parseTessellServiceInstanceConnectString(tessellServiceInstanceConnectionInfo.ConnectString)}
	}

	var endPoints *[]model.TessellServiceInstanceConnectionInfoEndPoints
	if tessellServiceInstanceConnectionInfo.EndPoints != endPoints {
		parsedTessellServiceInstanceConnectionInfo["end_points"] = parseTessellServiceInstanceConnectionInfoEndPointsList(tessellServiceInstanceConnectionInfo.EndPoints)
	}

	return parsedTessellServiceInstanceConnectionInfo
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

func parseTessellServiceInstanceConnectionInfoEndPointsList(tessellServiceInstanceConnectionInfo_endPoints *[]model.TessellServiceInstanceConnectionInfoEndPoints) []interface{} {
	if tessellServiceInstanceConnectionInfo_endPoints == nil {
		return nil
	}
	tessellServiceInstanceConnectionInfoEndPointsList := make([]interface{}, 0)

	if tessellServiceInstanceConnectionInfo_endPoints != nil {
		tessellServiceInstanceConnectionInfoEndPointsList = make([]interface{}, len(*tessellServiceInstanceConnectionInfo_endPoints))
		for i, tessellServiceInstanceConnectionInfoEndPointsItem := range *tessellServiceInstanceConnectionInfo_endPoints {
			tessellServiceInstanceConnectionInfoEndPointsList[i] = parseTessellServiceInstanceConnectionInfoEndPoints(&tessellServiceInstanceConnectionInfoEndPointsItem)
		}
	}

	return tessellServiceInstanceConnectionInfoEndPointsList
}

func parseTessellServiceInstanceConnectionInfoEndPoints(tessellServiceInstanceConnectionInfo_endPoints *model.TessellServiceInstanceConnectionInfoEndPoints) interface{} {
	if tessellServiceInstanceConnectionInfo_endPoints == nil {
		return nil
	}
	parsedTessellServiceInstanceConnectionInfo_endPoints := make(map[string]interface{})
	parsedTessellServiceInstanceConnectionInfo_endPoints["endpoint"] = tessellServiceInstanceConnectionInfo_endPoints.Endpoint
	parsedTessellServiceInstanceConnectionInfo_endPoints["labels"] = tessellServiceInstanceConnectionInfo_endPoints.Labels
	parsedTessellServiceInstanceConnectionInfo_endPoints["data"] = tessellServiceInstanceConnectionInfo_endPoints.Data

	return parsedTessellServiceInstanceConnectionInfo_endPoints
}

func parseServiceInstanceGenericInfo(serviceInstanceGenericInfo *model.ServiceInstanceGenericInfo) interface{} {
	if serviceInstanceGenericInfo == nil {
		return nil
	}
	parsedServiceInstanceGenericInfo := make(map[string]interface{})
	parsedServiceInstanceGenericInfo["part_of_initial_primary_set"] = serviceInstanceGenericInfo.PartOfInitialPrimarySet
	parsedServiceInstanceGenericInfo["encryption_key"] = serviceInstanceGenericInfo.EncryptionKey
	parsedServiceInstanceGenericInfo["encryption_key_id"] = serviceInstanceGenericInfo.EncryptionKeyId
	parsedServiceInstanceGenericInfo["server_cert_id"] = serviceInstanceGenericInfo.ServerCertId
	parsedServiceInstanceGenericInfo["vpc"] = serviceInstanceGenericInfo.VPC
	parsedServiceInstanceGenericInfo["vpc_id"] = serviceInstanceGenericInfo.VPCId
	parsedServiceInstanceGenericInfo["public_subnet"] = serviceInstanceGenericInfo.PublicSubnet
	parsedServiceInstanceGenericInfo["public_subnet_id"] = serviceInstanceGenericInfo.PublicSubnetId
	parsedServiceInstanceGenericInfo["private_subnet"] = serviceInstanceGenericInfo.PrivateSubnet
	parsedServiceInstanceGenericInfo["private_subnet_id"] = serviceInstanceGenericInfo.PrivateSubnetId
	parsedServiceInstanceGenericInfo["network_profile_id"] = serviceInstanceGenericInfo.NetworkProfileId
	parsedServiceInstanceGenericInfo["compute_type"] = serviceInstanceGenericInfo.ComputeType
	parsedServiceInstanceGenericInfo["compute_id"] = serviceInstanceGenericInfo.ComputeId

	parsedServiceInstanceGenericInfo["base_storage"] = serviceInstanceGenericInfo.BaseStorage
	parsedServiceInstanceGenericInfo["additional_storage"] = serviceInstanceGenericInfo.AdditionalStorage
	parsedServiceInstanceGenericInfo["allocated_storage"] = serviceInstanceGenericInfo.AllocatedStorage
	parsedServiceInstanceGenericInfo["max_memory"] = serviceInstanceGenericInfo.MaxMemory
	parsedServiceInstanceGenericInfo["software_image"] = serviceInstanceGenericInfo.SoftwareImage
	parsedServiceInstanceGenericInfo["software_image_version"] = serviceInstanceGenericInfo.SoftwareImageVersion
	parsedServiceInstanceGenericInfo["software_image_id"] = serviceInstanceGenericInfo.SoftwareImageId
	parsedServiceInstanceGenericInfo["software_image_version_id"] = serviceInstanceGenericInfo.SoftwareImageVersionId
	parsedServiceInstanceGenericInfo["data_volume_iops"] = serviceInstanceGenericInfo.DataVolumeIops
	parsedServiceInstanceGenericInfo["throughput"] = serviceInstanceGenericInfo.Throughput
	parsedServiceInstanceGenericInfo["multi_disk"] = serviceInstanceGenericInfo.MultiDisk

	parsedServiceInstanceGenericInfo["parameter_profile_id"] = serviceInstanceGenericInfo.ParameterProfileId
	parsedServiceInstanceGenericInfo["sync_mode"] = serviceInstanceGenericInfo.SyncMode

	var awsInfraConfig *model.AwsInfraConfig
	if serviceInstanceGenericInfo.AwsInfraConfig != awsInfraConfig {
		parsedServiceInstanceGenericInfo["aws_infra_config"] = []interface{}{parseAwsInfraConfig(serviceInstanceGenericInfo.AwsInfraConfig)}
	}

	var parameterProfile *model.ParameterProfile
	if serviceInstanceGenericInfo.ParameterProfile != parameterProfile {
		parsedServiceInstanceGenericInfo["parameter_profile"] = []interface{}{parseParameterProfile(serviceInstanceGenericInfo.ParameterProfile)}
	}

	var optionProfile *model.OptionProfile
	if serviceInstanceGenericInfo.OptionProfile != optionProfile {
		parsedServiceInstanceGenericInfo["option_profile"] = []interface{}{parseOptionProfile(serviceInstanceGenericInfo.OptionProfile)}
	}

	var engineConfiguration *model.ServiceInstanceEngineInfo
	if serviceInstanceGenericInfo.EngineConfiguration != engineConfiguration {
		parsedServiceInstanceGenericInfo["engine_configuration"] = []interface{}{parseServiceInstanceEngineInfo(serviceInstanceGenericInfo.EngineConfiguration)}
	}

	var computeConfig *model.InstanceComputeConfig
	if serviceInstanceGenericInfo.ComputeConfig != computeConfig {
		parsedServiceInstanceGenericInfo["compute_config"] = []interface{}{parseInstanceComputeConfig(serviceInstanceGenericInfo.ComputeConfig)}
	}

	var storageConfig *model.InstanceStorageConfig
	if serviceInstanceGenericInfo.StorageConfig != storageConfig {
		parsedServiceInstanceGenericInfo["storage_config"] = []interface{}{parseInstanceStorageConfig(serviceInstanceGenericInfo.StorageConfig)}
	}

	var archiveStorageConfig *model.InstanceStorageConfig
	if serviceInstanceGenericInfo.ArchiveStorageConfig != archiveStorageConfig {
		parsedServiceInstanceGenericInfo["archive_storage_config"] = []interface{}{parseInstanceStorageConfig(serviceInstanceGenericInfo.ArchiveStorageConfig)}
	}

	return parsedServiceInstanceGenericInfo
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
	parsedOptionProfile["status"] = optionProfile.Status

	return parsedOptionProfile
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

func parseDBLicenseInfo(dbLicenseInfo *model.DBLicenseInfo) interface{} {
	if dbLicenseInfo == nil {
		return nil
	}
	parsedDbLicenseInfo := make(map[string]interface{})

	var licenses *[]model.LicenseInfo
	if dbLicenseInfo.Licenses != licenses {
		parsedDbLicenseInfo["licenses"] = parseLicenseInfoList(dbLicenseInfo.Licenses)
	}

	return parsedDbLicenseInfo
}

func parseLicenseInfoList(licenseInfo *[]model.LicenseInfo) []interface{} {
	if licenseInfo == nil {
		return nil
	}
	licenseInfoList := make([]interface{}, 0)

	if licenseInfo != nil {
		licenseInfoList = make([]interface{}, len(*licenseInfo))
		for i, licenseInfoItem := range *licenseInfo {
			licenseInfoList[i] = parseLicenseInfo(&licenseInfoItem)
		}
	}

	return licenseInfoList
}

func parseLicenseInfo(licenseInfo *model.LicenseInfo) interface{} {
	if licenseInfo == nil {
		return nil
	}
	parsedLicenseInfo := make(map[string]interface{})
	parsedLicenseInfo["license_id"] = licenseInfo.LicenseId
	parsedLicenseInfo["lock_hash"] = licenseInfo.LockHash
	parsedLicenseInfo["quantity"] = licenseInfo.Quantity

	return parsedLicenseInfo
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

func parseTessellServiceInstanceMetadata(tessellServiceInstanceMetadata *model.TessellServiceInstanceMetadata) interface{} {
	if tessellServiceInstanceMetadata == nil {
		return nil
	}
	parsedTessellServiceInstanceMetadata := make(map[string]interface{})
	parsedTessellServiceInstanceMetadata["instance_group_name"] = tessellServiceInstanceMetadata.InstanceGroupName
	parsedTessellServiceInstanceMetadata["deletion_attempts"] = tessellServiceInstanceMetadata.DeletionAttempts
	parsedTessellServiceInstanceMetadata["last_deletion_dispatch_time"] = tessellServiceInstanceMetadata.LastDeletionDispatchTime
	parsedTessellServiceInstanceMetadata["last_resize_dispatch_time"] = tessellServiceInstanceMetadata.LastResizeDispatchTime
	parsedTessellServiceInstanceMetadata["last_storage_resize_dispatch_time"] = tessellServiceInstanceMetadata.LastStorageResizeDispatchTime
	parsedTessellServiceInstanceMetadata["add_replica_context_id"] = tessellServiceInstanceMetadata.AddReplicaContextId
	parsedTessellServiceInstanceMetadata["data"] = tessellServiceInstanceMetadata.Data
	parsedTessellServiceInstanceMetadata["is_dcr_enabled"] = tessellServiceInstanceMetadata.IsDcrEnabled

	return parsedTessellServiceInstanceMetadata
}

func parseTessellServiceInstanceDriverInfo(tessellServiceInstanceDriverInfo *model.TessellServiceInstanceDriverInfo) interface{} {
	if tessellServiceInstanceDriverInfo == nil {
		return nil
	}
	parsedTessellServiceInstanceDriverInfo := make(map[string]interface{})
	parsedTessellServiceInstanceDriverInfo["data"] = tessellServiceInstanceDriverInfo.Data

	return parsedTessellServiceInstanceDriverInfo
}

func parseTessellServiceInstanceInProgressUpdateInfoOps(tessellServiceInstanceInProgressUpdateInfoOps *model.TessellServiceInstanceInProgressUpdateInfoOps) interface{} {
	if tessellServiceInstanceInProgressUpdateInfoOps == nil {
		return nil
	}
	parsedTessellServiceInstanceInProgressUpdateInfoOps := make(map[string]interface{})

	var infra *model.TessellServiceInstanceInProgressUpdateInfoOpsInfra
	if tessellServiceInstanceInProgressUpdateInfoOps.Infra != infra {
		parsedTessellServiceInstanceInProgressUpdateInfoOps["infra"] = []interface{}{parseTessellServiceInstanceInProgressUpdateInfoOpsInfra(tessellServiceInstanceInProgressUpdateInfoOps.Infra)}
	}

	return parsedTessellServiceInstanceInProgressUpdateInfoOps
}

func parseTessellServiceInstanceInProgressUpdateInfoOpsInfra(tessellServiceInstanceInProgressUpdateInfoOps_infra *model.TessellServiceInstanceInProgressUpdateInfoOpsInfra) interface{} {
	if tessellServiceInstanceInProgressUpdateInfoOps_infra == nil {
		return nil
	}
	parsedTessellServiceInstanceInProgressUpdateInfoOps_infra := make(map[string]interface{})

	var resourceUpdateInfo *model.TessellResourceUpdateInfo
	if tessellServiceInstanceInProgressUpdateInfoOps_infra.ResourceUpdateInfo != resourceUpdateInfo {
		parsedTessellServiceInstanceInProgressUpdateInfoOps_infra["resource_update_info"] = []interface{}{parseTessellResourceUpdateInfo(tessellServiceInstanceInProgressUpdateInfoOps_infra.ResourceUpdateInfo)}
	}

	var infraUpdateInfo *model.TessellServiceInstanceInfraUpdateInfo
	if tessellServiceInstanceInProgressUpdateInfoOps_infra.InfraUpdateInfo != infraUpdateInfo {
		parsedTessellServiceInstanceInProgressUpdateInfoOps_infra["infra_update_info"] = []interface{}{parseTessellServiceInstanceInfraUpdateInfo(tessellServiceInstanceInProgressUpdateInfoOps_infra.InfraUpdateInfo)}
	}

	return parsedTessellServiceInstanceInProgressUpdateInfoOps_infra
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

func parseTessellServiceInstanceInfraUpdateInfo(tessellServiceInstanceInfraUpdateInfo *model.TessellServiceInstanceInfraUpdateInfo) interface{} {
	if tessellServiceInstanceInfraUpdateInfo == nil {
		return nil
	}
	parsedTessellServiceInstanceInfraUpdateInfo := make(map[string]interface{})
	parsedTessellServiceInstanceInfraUpdateInfo["compute_type"] = tessellServiceInstanceInfraUpdateInfo.ComputeType

	return parsedTessellServiceInstanceInfraUpdateInfo
}

func parseTessellAgentLcmInfo(tessellAgentLcmInfo *model.TessellAgentLcmInfo) interface{} {
	if tessellAgentLcmInfo == nil {
		return nil
	}
	parsedTessellAgentLcmInfo := make(map[string]interface{})
	parsedTessellAgentLcmInfo["compute_resource_info"] = tessellAgentLcmInfo.ComputeResourceInfo
	parsedTessellAgentLcmInfo["service_info"] = tessellAgentLcmInfo.ServiceInfo
	parsedTessellAgentLcmInfo["instance_info"] = tessellAgentLcmInfo.InstanceInfo
	parsedTessellAgentLcmInfo["data"] = tessellAgentLcmInfo.Data

	return parsedTessellAgentLcmInfo
}

func parseTessellComputeResourceOpsDTO(tessellComputeResourceOpsDTO *model.TessellComputeResourceOpsDTO) interface{} {
	if tessellComputeResourceOpsDTO == nil {
		return nil
	}
	parsedTessellComputeResourceOpsDTO := make(map[string]interface{})
	parsedTessellComputeResourceOpsDTO["id"] = tessellComputeResourceOpsDTO.Id
	parsedTessellComputeResourceOpsDTO["name"] = tessellComputeResourceOpsDTO.Name
	parsedTessellComputeResourceOpsDTO["description"] = tessellComputeResourceOpsDTO.Description
	parsedTessellComputeResourceOpsDTO["tenant_id"] = tessellComputeResourceOpsDTO.TenantId
	parsedTessellComputeResourceOpsDTO["subscription_id"] = tessellComputeResourceOpsDTO.SubscriptionId
	parsedTessellComputeResourceOpsDTO["engine_type"] = tessellComputeResourceOpsDTO.EngineType
	parsedTessellComputeResourceOpsDTO["status"] = tessellComputeResourceOpsDTO.Status
	parsedTessellComputeResourceOpsDTO["tsm"] = tessellComputeResourceOpsDTO.Tsm
	parsedTessellComputeResourceOpsDTO["cloud_status"] = tessellComputeResourceOpsDTO.CloudStatus
	parsedTessellComputeResourceOpsDTO["compute_sharing_enabled"] = tessellComputeResourceOpsDTO.ComputeSharingEnabled
	parsedTessellComputeResourceOpsDTO["cloud_account_id"] = tessellComputeResourceOpsDTO.CloudAccountId
	parsedTessellComputeResourceOpsDTO["cloud_location"] = tessellComputeResourceOpsDTO.CloudLocation
	parsedTessellComputeResourceOpsDTO["cloud_resource_id"] = tessellComputeResourceOpsDTO.CloudResourceId
	parsedTessellComputeResourceOpsDTO["type"] = tessellComputeResourceOpsDTO.Type
	parsedTessellComputeResourceOpsDTO["machine_type"] = tessellComputeResourceOpsDTO.MachineType

	parsedTessellComputeResourceOpsDTO["software_image_id"] = tessellComputeResourceOpsDTO.SoftwareImageId
	parsedTessellComputeResourceOpsDTO["software_image_version_id"] = tessellComputeResourceOpsDTO.SoftwareImageVersionId
	parsedTessellComputeResourceOpsDTO["network_profile_id"] = tessellComputeResourceOpsDTO.NetworkProfileId
	parsedTessellComputeResourceOpsDTO["compute_type_id"] = tessellComputeResourceOpsDTO.ComputeTypeId
	parsedTessellComputeResourceOpsDTO["user_id"] = tessellComputeResourceOpsDTO.UserId
	parsedTessellComputeResourceOpsDTO["owner"] = tessellComputeResourceOpsDTO.Owner
	parsedTessellComputeResourceOpsDTO["date_created"] = tessellComputeResourceOpsDTO.DateCreated
	parsedTessellComputeResourceOpsDTO["date_modified"] = tessellComputeResourceOpsDTO.DateModified
	parsedTessellComputeResourceOpsDTO["timezone"] = tessellComputeResourceOpsDTO.Timezone

	parsedTessellComputeResourceOpsDTO["internal"] = tessellComputeResourceOpsDTO.Internal

	var osInfo *model.OsInfo
	if tessellComputeResourceOpsDTO.OsInfo != osInfo {
		parsedTessellComputeResourceOpsDTO["os_info"] = []interface{}{parseOsInfo(tessellComputeResourceOpsDTO.OsInfo)}
	}

	var machineFqdnInfo *model.ComputeResourceMachineFqdnInfo
	if tessellComputeResourceOpsDTO.MachineFqdnInfo != machineFqdnInfo {
		parsedTessellComputeResourceOpsDTO["machine_fqdn_info"] = []interface{}{parseComputeResourceMachineFqdnInfo(tessellComputeResourceOpsDTO.MachineFqdnInfo)}
	}

	var ipAddressInfo *model.ComputeResourceIpAddressInfo
	if tessellComputeResourceOpsDTO.IpAddressInfo != ipAddressInfo {
		parsedTessellComputeResourceOpsDTO["ip_address_info"] = []interface{}{parseComputeResourceIpAddressInfo(tessellComputeResourceOpsDTO.IpAddressInfo)}
	}

	var metadata *model.ComputeResourceMetadata
	if tessellComputeResourceOpsDTO.Metadata != metadata {
		parsedTessellComputeResourceOpsDTO["metadata"] = []interface{}{parseComputeResourceMetadata(tessellComputeResourceOpsDTO.Metadata)}
	}

	var driverInfo *model.ComputeResourceDriverInfo
	if tessellComputeResourceOpsDTO.DriverInfo != driverInfo {
		parsedTessellComputeResourceOpsDTO["driver_info"] = []interface{}{parseComputeResourceDriverInfo(tessellComputeResourceOpsDTO.DriverInfo)}
	}

	var tessellStackInfo *model.ComputeResourceTessellStackInfo
	if tessellComputeResourceOpsDTO.TessellStackInfo != tessellStackInfo {
		parsedTessellComputeResourceOpsDTO["tessell_stack_info"] = []interface{}{parseComputeResourceTessellStackInfo(tessellComputeResourceOpsDTO.TessellStackInfo)}
	}

	var tessellAgentLcmInfo *model.TessellAgentLcmInfo
	if tessellComputeResourceOpsDTO.TessellAgentLcmInfo != tessellAgentLcmInfo {
		parsedTessellComputeResourceOpsDTO["tessell_agent_lcm_info"] = []interface{}{parseTessellAgentLcmInfo(tessellComputeResourceOpsDTO.TessellAgentLcmInfo)}
	}

	var tessellEndpointMigrationInfo *model.ComputeResourceEndpointMigrationInfo
	if tessellComputeResourceOpsDTO.TessellEndpointMigrationInfo != tessellEndpointMigrationInfo {
		parsedTessellComputeResourceOpsDTO["tessell_endpoint_migration_info"] = []interface{}{parseComputeResourceEndpointMigrationInfo(tessellComputeResourceOpsDTO.TessellEndpointMigrationInfo)}
	}

	var contextInfo *model.ComputeResourceContextInfo
	if tessellComputeResourceOpsDTO.ContextInfo != contextInfo {
		parsedTessellComputeResourceOpsDTO["context_info"] = []interface{}{parseComputeResourceContextInfo(tessellComputeResourceOpsDTO.ContextInfo)}
	}

	var actionMetadata *model.ComputeResourceActionMetadata
	if tessellComputeResourceOpsDTO.ActionMetadata != actionMetadata {
		parsedTessellComputeResourceOpsDTO["action_metadata"] = []interface{}{parseComputeResourceActionMetadata(tessellComputeResourceOpsDTO.ActionMetadata)}
	}

	return parsedTessellComputeResourceOpsDTO
}

func parseOsInfo(osInfo *model.OsInfo) interface{} {
	if osInfo == nil {
		return nil
	}
	parsedOsInfo := make(map[string]interface{})
	parsedOsInfo["type"] = osInfo.Type
	parsedOsInfo["image"] = osInfo.Image

	return parsedOsInfo
}

func parseComputeResourceMachineFqdnInfo(computeResourceMachineFqdnInfo *model.ComputeResourceMachineFqdnInfo) interface{} {
	if computeResourceMachineFqdnInfo == nil {
		return nil
	}
	parsedComputeResourceMachineFqdnInfo := make(map[string]interface{})
	parsedComputeResourceMachineFqdnInfo["data"] = computeResourceMachineFqdnInfo.Data

	return parsedComputeResourceMachineFqdnInfo
}

func parseComputeResourceIpAddressInfo(computeResourceIpAddressInfo *model.ComputeResourceIpAddressInfo) interface{} {
	if computeResourceIpAddressInfo == nil {
		return nil
	}
	parsedComputeResourceIpAddressInfo := make(map[string]interface{})
	parsedComputeResourceIpAddressInfo["data"] = computeResourceIpAddressInfo.Data
	parsedComputeResourceIpAddressInfo["ip_addresses"] = computeResourceIpAddressInfo.IpAddresses

	return parsedComputeResourceIpAddressInfo
}

func parseComputeResourceMetadata(computeResourceMetadata *model.ComputeResourceMetadata) interface{} {
	if computeResourceMetadata == nil {
		return nil
	}
	parsedComputeResourceMetadata := make(map[string]interface{})
	parsedComputeResourceMetadata["subscription"] = computeResourceMetadata.Subscription
	parsedComputeResourceMetadata["compute_type"] = computeResourceMetadata.ComputeType
	parsedComputeResourceMetadata["vpc"] = computeResourceMetadata.VPC
	parsedComputeResourceMetadata["private_subnet_id"] = computeResourceMetadata.PrivateSubnetId
	parsedComputeResourceMetadata["private_subnet"] = computeResourceMetadata.PrivateSubnet
	parsedComputeResourceMetadata["public_subnet_id"] = computeResourceMetadata.PublicSubnetId
	parsedComputeResourceMetadata["public_subnet"] = computeResourceMetadata.PublicSubnet
	parsedComputeResourceMetadata["encryption_key_id"] = computeResourceMetadata.EncryptionKeyId
	parsedComputeResourceMetadata["encryption_key"] = computeResourceMetadata.EncryptionKey

	parsedComputeResourceMetadata["os_timezone"] = computeResourceMetadata.OsTimezone

	parsedComputeResourceMetadata["data"] = computeResourceMetadata.Data
	parsedComputeResourceMetadata["is_azure_monitor_agent_installed"] = computeResourceMetadata.IsAzureMonitorAgentInstalled
	parsedComputeResourceMetadata["node_index"] = computeResourceMetadata.NodeIndex
	parsedComputeResourceMetadata["use_azure_monitor_agent"] = computeResourceMetadata.UseAzureMonitorAgent

	var connectivityInfo *model.ComputeResourceMetadataConnectivityInfo
	if computeResourceMetadata.ConnectivityInfo != connectivityInfo {
		parsedComputeResourceMetadata["connectivity_info"] = []interface{}{parseComputeResourceMetadataConnectivityInfo(computeResourceMetadata.ConnectivityInfo)}
	}

	var dBserverInfo *model.ComputeResourceDBserverInfo
	if computeResourceMetadata.DBserverInfo != dBserverInfo {
		parsedComputeResourceMetadata["dbserver_info"] = []interface{}{parseComputeResourceDBserverInfo(computeResourceMetadata.DBserverInfo)}
	}

	var awsInfraConfig *model.AwsInfraConfig
	if computeResourceMetadata.AwsInfraConfig != awsInfraConfig {
		parsedComputeResourceMetadata["aws_infra_config"] = []interface{}{parseAwsInfraConfig(computeResourceMetadata.AwsInfraConfig)}
	}

	var licenseInfo *model.ComputeLicenseInfo
	if computeResourceMetadata.LicenseInfo != licenseInfo {
		parsedComputeResourceMetadata["license_info"] = []interface{}{parseComputeLicenseInfo(computeResourceMetadata.LicenseInfo)}
	}

	var computeConfig *model.ComputeConfig
	if computeResourceMetadata.ComputeConfig != computeConfig {
		parsedComputeResourceMetadata["compute_config"] = []interface{}{parseComputeConfig(computeResourceMetadata.ComputeConfig)}
	}

	return parsedComputeResourceMetadata
}

func parseComputeResourceMetadataConnectivityInfo(computeResourceMetadata_connectivityInfo *model.ComputeResourceMetadataConnectivityInfo) interface{} {
	if computeResourceMetadata_connectivityInfo == nil {
		return nil
	}
	parsedComputeResourceMetadata_connectivityInfo := make(map[string]interface{})
	parsedComputeResourceMetadata_connectivityInfo["service_port"] = computeResourceMetadata_connectivityInfo.ServicePort

	return parsedComputeResourceMetadata_connectivityInfo
}

func parseComputeResourceDBserverInfo(computeResourceDbserverInfo *model.ComputeResourceDBserverInfo) interface{} {
	if computeResourceDbserverInfo == nil {
		return nil
	}
	parsedComputeResourceDbserverInfo := make(map[string]interface{})
	parsedComputeResourceDbserverInfo["first_provisioned_dbservice_id"] = computeResourceDbserverInfo.FirstProvisionedDBserviceId
	parsedComputeResourceDbserverInfo["db_service_ids"] = computeResourceDbserverInfo.DBServiceIds
	parsedComputeResourceDbserverInfo["enable_public_access"] = computeResourceDbserverInfo.EnablePublicAccess
	parsedComputeResourceDbserverInfo["enable_ssl"] = computeResourceDbserverInfo.EnableSSL

	var softwareImageInfo *model.DBserverSoftwareImageInfo
	if computeResourceDbserverInfo.SoftwareImageInfo != softwareImageInfo {
		parsedComputeResourceDbserverInfo["software_image_info"] = []interface{}{parseDBserverSoftwareImageInfo(computeResourceDbserverInfo.SoftwareImageInfo)}
	}

	return parsedComputeResourceDbserverInfo
}

func parseDBserverSoftwareImageInfo(dbserverSoftwareImageInfo *model.DBserverSoftwareImageInfo) interface{} {
	if dbserverSoftwareImageInfo == nil {
		return nil
	}
	parsedDbserverSoftwareImageInfo := make(map[string]interface{})
	parsedDbserverSoftwareImageInfo["software_image"] = dbserverSoftwareImageInfo.SoftwareImage
	parsedDbserverSoftwareImageInfo["software_image_id"] = dbserverSoftwareImageInfo.SoftwareImageId

	var softwareImageVersions *[]model.DBserverSoftwareImageVersionInfo
	if dbserverSoftwareImageInfo.SoftwareImageVersions != softwareImageVersions {
		parsedDbserverSoftwareImageInfo["software_image_versions"] = parseDBserverSoftwareImageVersionInfoList(dbserverSoftwareImageInfo.SoftwareImageVersions)
	}

	return parsedDbserverSoftwareImageInfo
}

func parseDBserverSoftwareImageVersionInfoList(dbserverSoftwareImageVersionInfo *[]model.DBserverSoftwareImageVersionInfo) []interface{} {
	if dbserverSoftwareImageVersionInfo == nil {
		return nil
	}
	dBserverSoftwareImageVersionInfoList := make([]interface{}, 0)

	if dbserverSoftwareImageVersionInfo != nil {
		dBserverSoftwareImageVersionInfoList = make([]interface{}, len(*dbserverSoftwareImageVersionInfo))
		for i, dBserverSoftwareImageVersionInfoItem := range *dbserverSoftwareImageVersionInfo {
			dBserverSoftwareImageVersionInfoList[i] = parseDBserverSoftwareImageVersionInfo(&dBserverSoftwareImageVersionInfoItem)
		}
	}

	return dBserverSoftwareImageVersionInfoList
}

func parseDBserverSoftwareImageVersionInfo(dbserverSoftwareImageVersionInfo *model.DBserverSoftwareImageVersionInfo) interface{} {
	if dbserverSoftwareImageVersionInfo == nil {
		return nil
	}
	parsedDbserverSoftwareImageVersionInfo := make(map[string]interface{})
	parsedDbserverSoftwareImageVersionInfo["software_image_version"] = dbserverSoftwareImageVersionInfo.SoftwareImageVersion
	parsedDbserverSoftwareImageVersionInfo["software_image_version_id"] = dbserverSoftwareImageVersionInfo.SoftwareImageVersionId
	parsedDbserverSoftwareImageVersionInfo["supported"] = dbserverSoftwareImageVersionInfo.Supported

	return parsedDbserverSoftwareImageVersionInfo
}

func parseComputeLicenseInfo(computeLicenseInfo *model.ComputeLicenseInfo) interface{} {
	if computeLicenseInfo == nil {
		return nil
	}
	parsedComputeLicenseInfo := make(map[string]interface{})
	parsedComputeLicenseInfo["acquirer_id"] = computeLicenseInfo.AcquirerId
	parsedComputeLicenseInfo["license_id"] = computeLicenseInfo.LicenseId
	parsedComputeLicenseInfo["lock_hash"] = computeLicenseInfo.LockHash
	parsedComputeLicenseInfo["quantity"] = computeLicenseInfo.Quantity

	return parsedComputeLicenseInfo
}

func parseComputeConfig(computeConfig *model.ComputeConfig) interface{} {
	if computeConfig == nil {
		return nil
	}
	parsedComputeConfig := make(map[string]interface{})
	parsedComputeConfig["provider"] = computeConfig.Provider

	var exadataConfig *model.ExadataComputeConfig
	if computeConfig.ExadataConfig != exadataConfig {
		parsedComputeConfig["exadata_config"] = []interface{}{parseExadataComputeConfig(computeConfig.ExadataConfig)}
	}

	return parsedComputeConfig
}

func parseExadataComputeConfig(exadataComputeConfig *model.ExadataComputeConfig) interface{} {
	if exadataComputeConfig == nil {
		return nil
	}
	parsedExadataComputeConfig := make(map[string]interface{})
	parsedExadataComputeConfig["infrastructure_id"] = exadataComputeConfig.InfrastructureId
	parsedExadataComputeConfig["infrastructure_name"] = exadataComputeConfig.InfrastructureName
	parsedExadataComputeConfig["vm_cluster_id"] = exadataComputeConfig.VmClusterId
	parsedExadataComputeConfig["vm_cluster_name"] = exadataComputeConfig.VmClusterName
	parsedExadataComputeConfig["ocpu"] = exadataComputeConfig.Ocpu
	parsedExadataComputeConfig["memory_in_gbs"] = exadataComputeConfig.MemoryInGbs
	parsedExadataComputeConfig["floating_ip_address"] = exadataComputeConfig.FloatingIpAddress
	parsedExadataComputeConfig["db_server"] = exadataComputeConfig.DBServer
	parsedExadataComputeConfig["private_ip_address"] = exadataComputeConfig.PrivateIpAddress
	parsedExadataComputeConfig["local_storage_in_gbs"] = exadataComputeConfig.LocalStorageInGbs
	parsedExadataComputeConfig["dns_name"] = exadataComputeConfig.DNSName

	return parsedExadataComputeConfig
}

func parseComputeResourceDriverInfo(computeResourceDriverInfo *model.ComputeResourceDriverInfo) interface{} {
	if computeResourceDriverInfo == nil {
		return nil
	}
	parsedComputeResourceDriverInfo := make(map[string]interface{})
	parsedComputeResourceDriverInfo["data"] = computeResourceDriverInfo.Data

	return parsedComputeResourceDriverInfo
}

func parseComputeResourceTessellStackInfo(computeResourceTessellStackInfo *model.ComputeResourceTessellStackInfo) interface{} {
	if computeResourceTessellStackInfo == nil {
		return nil
	}
	parsedComputeResourceTessellStackInfo := make(map[string]interface{})
	parsedComputeResourceTessellStackInfo["data"] = computeResourceTessellStackInfo.Data

	return parsedComputeResourceTessellStackInfo
}

func parseComputeResourceEndpointMigrationInfo(computeResourceEndpointMigrationInfo *model.ComputeResourceEndpointMigrationInfo) interface{} {
	if computeResourceEndpointMigrationInfo == nil {
		return nil
	}
	parsedComputeResourceEndpointMigrationInfo := make(map[string]interface{})
	parsedComputeResourceEndpointMigrationInfo["validate_common_endpoint"] = computeResourceEndpointMigrationInfo.ValidateCommonEndpoint
	parsedComputeResourceEndpointMigrationInfo["migrate_to_common_endpoint"] = computeResourceEndpointMigrationInfo.MigrateToCommonEndpoint
	parsedComputeResourceEndpointMigrationInfo["common_endpoint_successful_validation_time"] = computeResourceEndpointMigrationInfo.CommonEndpointSuccessfulValidationTime
	parsedComputeResourceEndpointMigrationInfo["common_endpoint_successful_migration_time"] = computeResourceEndpointMigrationInfo.CommonEndpointSuccessfulMigrationTime

	return parsedComputeResourceEndpointMigrationInfo
}

func parseComputeResourceContextInfo(computeResourceContextInfo *model.ComputeResourceContextInfo) interface{} {
	if computeResourceContextInfo == nil {
		return nil
	}
	parsedComputeResourceContextInfo := make(map[string]interface{})
	parsedComputeResourceContextInfo["sub_status"] = computeResourceContextInfo.SubStatus
	parsedComputeResourceContextInfo["description"] = computeResourceContextInfo.Description

	return parsedComputeResourceContextInfo
}

func parseComputeResourceActionMetadata(computeResourceActionMetadata *model.ComputeResourceActionMetadata) interface{} {
	if computeResourceActionMetadata == nil {
		return nil
	}
	parsedComputeResourceActionMetadata := make(map[string]interface{})

	parsedComputeResourceActionMetadata["last_resize_dispatch_time"] = computeResourceActionMetadata.LastResizeDispatchTime

	var lastActionMetadata *model.CrLastActionMetadata
	if computeResourceActionMetadata.LastActionMetadata != lastActionMetadata {
		parsedComputeResourceActionMetadata["last_action_metadata"] = []interface{}{parseCrLastActionMetadata(computeResourceActionMetadata.LastActionMetadata)}
	}

	return parsedComputeResourceActionMetadata
}

func parseCrLastActionMetadata(crLastActionMetadata *model.CrLastActionMetadata) interface{} {
	if crLastActionMetadata == nil {
		return nil
	}
	parsedCrLastActionMetadata := make(map[string]interface{})
	parsedCrLastActionMetadata["reference_id"] = crLastActionMetadata.ReferenceId
	parsedCrLastActionMetadata["context_id"] = crLastActionMetadata.ContextId
	parsedCrLastActionMetadata["action_type"] = crLastActionMetadata.ActionType

	return parsedCrLastActionMetadata
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
