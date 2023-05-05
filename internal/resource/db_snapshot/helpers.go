package db_snapshot

import (
	//"fmt"
	//"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"terraform-provider-tessell/internal/helper"
	"terraform-provider-tessell/internal/model"
)

func setResourceData(d *schema.ResourceData, databaseSnapshot *model.DatabaseSnapshot) error {

	if err := d.Set("id", databaseSnapshot.Id); err != nil {
		return err
	}

	if err := d.Set("name", databaseSnapshot.Name); err != nil {
		return err
	}

	if err := d.Set("description", databaseSnapshot.Description); err != nil {
		return err
	}

	if err := d.Set("snapshot_time", databaseSnapshot.SnapshotTime); err != nil {
		return err
	}

	if err := d.Set("status", databaseSnapshot.Status); err != nil {
		return err
	}

	if err := d.Set("size", databaseSnapshot.Size); err != nil {
		return err
	}

	if err := d.Set("manual", databaseSnapshot.Manual); err != nil {
		return err
	}

	if err := d.Set("cloud_availability", parseCloudRegionInfoListWithResData(databaseSnapshot.CloudAvailability, d)); err != nil {
		return err
	}

	if err := d.Set("availability_config", parseSnapshotAvailabilityConfigListWithResData(databaseSnapshot.AvailabilityConfig, d)); err != nil {
		return err
	}

	if err := d.Set("databases", parseBackupDatabaseInfoListWithResData(databaseSnapshot.Databases, d)); err != nil {
		return err
	}

	if err := d.Set("shared_with", parseEntityAclSharingSummaryInfoWithResData(databaseSnapshot.SharedWith, d)); err != nil {
		return err
	}

	if err := d.Set("backup_status", databaseSnapshot.BackupStatus); err != nil {
		return err
	}

	return nil
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

func parseRegionInfoList(regions *[]model.RegionInfo) []interface{} {
	if regions == nil {
		return nil
	}
	regionInfoList := make([]interface{}, 0)

	if regions != nil {
		regionInfoList = make([]interface{}, len(*regions))
		for i, regionInfoItem := range *regions {
			regionInfoList[i] = parseRegionInfo(&regionInfoItem)
		}
	}

	return regionInfoList
}

func parseRegionInfo(regions *model.RegionInfo) interface{} {
	if regions == nil {
		return nil
	}
	parsedRegions := make(map[string]interface{})
	parsedRegions["region"] = regions.Region
	parsedRegions["availability_zones"] = regions.AvailabilityZones

	return parsedRegions
}

func parseSnapshotAvailabilityConfigListWithResData(availabilityConfig *[]model.SnapshotAvailabilityConfig, d *schema.ResourceData) []interface{} {
	if availabilityConfig == nil {
		return nil
	}
	snapshotAvailabilityConfigList := make([]interface{}, 0)

	if availabilityConfig != nil {
		snapshotAvailabilityConfigList = make([]interface{}, len(*availabilityConfig))
		for i, snapshotAvailabilityConfigItem := range *availabilityConfig {
			snapshotAvailabilityConfigList[i] = parseSnapshotAvailabilityConfig(&snapshotAvailabilityConfigItem)
		}
	}

	return snapshotAvailabilityConfigList
}

func parseSnapshotAvailabilityConfig(availabilityConfig *model.SnapshotAvailabilityConfig) interface{} {
	if availabilityConfig == nil {
		return nil
	}
	parsedAvailabilityConfig := make(map[string]interface{})
	parsedAvailabilityConfig["availability_configured_manually"] = availabilityConfig.AvailabilityConfiguredManually
	parsedAvailabilityConfig["dap_id"] = availabilityConfig.DAPId

	var cloudAvailabilityConfig *[]model.SnapshotCloudAvailabilityInfo
	if availabilityConfig.CloudAvailabilityConfig != cloudAvailabilityConfig {
		parsedAvailabilityConfig["cloud_availability_config"] = parseSnapshotCloudAvailabilityInfoList(availabilityConfig.CloudAvailabilityConfig)
	}

	return parsedAvailabilityConfig
}

func parseSnapshotCloudAvailabilityInfoList(cloudAvailabilityConfig *[]model.SnapshotCloudAvailabilityInfo) []interface{} {
	if cloudAvailabilityConfig == nil {
		return nil
	}
	snapshotCloudAvailabilityInfoList := make([]interface{}, 0)

	if cloudAvailabilityConfig != nil {
		snapshotCloudAvailabilityInfoList = make([]interface{}, len(*cloudAvailabilityConfig))
		for i, snapshotCloudAvailabilityInfoItem := range *cloudAvailabilityConfig {
			snapshotCloudAvailabilityInfoList[i] = parseSnapshotCloudAvailabilityInfo(&snapshotCloudAvailabilityInfoItem)
		}
	}

	return snapshotCloudAvailabilityInfoList
}

func parseSnapshotCloudAvailabilityInfo(cloudAvailabilityConfig *model.SnapshotCloudAvailabilityInfo) interface{} {
	if cloudAvailabilityConfig == nil {
		return nil
	}
	parsedCloudAvailabilityConfig := make(map[string]interface{})
	parsedCloudAvailabilityConfig["cloud"] = cloudAvailabilityConfig.Cloud

	var regions *[]model.SnapshotRegionAvailability
	if cloudAvailabilityConfig.Regions != regions {
		parsedCloudAvailabilityConfig["regions"] = parseSnapshotRegionAvailabilityList(cloudAvailabilityConfig.Regions)
	}

	return parsedCloudAvailabilityConfig
}

func parseSnapshotRegionAvailabilityList(regions *[]model.SnapshotRegionAvailability) []interface{} {
	if regions == nil {
		return nil
	}
	snapshotRegionAvailabilityList := make([]interface{}, 0)

	if regions != nil {
		snapshotRegionAvailabilityList = make([]interface{}, len(*regions))
		for i, snapshotRegionAvailabilityItem := range *regions {
			snapshotRegionAvailabilityList[i] = parseSnapshotRegionAvailability(&snapshotRegionAvailabilityItem)
		}
	}

	return snapshotRegionAvailabilityList
}

func parseSnapshotRegionAvailability(regions *model.SnapshotRegionAvailability) interface{} {
	if regions == nil {
		return nil
	}
	parsedRegions := make(map[string]interface{})
	parsedRegions["region"] = regions.Region
	parsedRegions["status"] = regions.Status

	return parsedRegions
}

func parseBackupDatabaseInfoListWithResData(databases *[]model.BackupDatabaseInfo, d *schema.ResourceData) []interface{} {
	if databases == nil {
		return nil
	}
	backupDatabaseInfoList := make([]interface{}, 0)

	if databases != nil {
		backupDatabaseInfoList = make([]interface{}, len(*databases))
		for i, backupDatabaseInfoItem := range *databases {
			backupDatabaseInfoList[i] = parseBackupDatabaseInfo(&backupDatabaseInfoItem)
		}
	}

	return backupDatabaseInfoList
}

func parseBackupDatabaseInfo(databases *model.BackupDatabaseInfo) interface{} {
	if databases == nil {
		return nil
	}
	parsedDatabases := make(map[string]interface{})
	parsedDatabases["id"] = databases.Id
	parsedDatabases["name"] = databases.Name
	parsedDatabases["status"] = databases.Status

	return parsedDatabases
}

func parseEntityAclSharingSummaryInfoWithResData(sharedWith *model.EntityAclSharingSummaryInfo, d *schema.ResourceData) []interface{} {
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
	parsedSharedWith["users"] = sharedWith.Users

	return []interface{}{parsedSharedWith}
}

func formPayloadForCreateDatabaseSnapshotRequest(d *schema.ResourceData) model.CreateDatabaseSnapshotTaskPayload {
	createDatabaseSnapshotTaskPayloadFormed := model.CreateDatabaseSnapshotTaskPayload{
		Name:        helper.GetStringPointer(d.Get("name")),
		Description: helper.GetStringPointer(d.Get("description")),
	}

	return createDatabaseSnapshotTaskPayloadFormed
}
