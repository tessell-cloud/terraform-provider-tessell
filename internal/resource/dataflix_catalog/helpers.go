package dataflix_catalog

import (
	//"fmt"
	//"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"terraform-provider-tessell/internal/model"
)

func setResourceData(d *schema.ResourceData, tessellDMMDataflixServiceView *model.TessellDMMDataflixServiceView) error {

	if err := d.Set("availability_machine_id", tessellDMMDataflixServiceView.AvailabilityMachineId); err != nil {
		return err
	}

	if err := d.Set("tessell_service_id", tessellDMMDataflixServiceView.TessellServiceId); err != nil {
		return err
	}

	if err := d.Set("service_name", tessellDMMDataflixServiceView.ServiceName); err != nil {
		return err
	}

	if err := d.Set("engine_type", tessellDMMDataflixServiceView.EngineType); err != nil {
		return err
	}

	if err := d.Set("time_zone", tessellDMMDataflixServiceView.TimeZone); err != nil {
		return err
	}

	if err := d.Set("owner", tessellDMMDataflixServiceView.Owner); err != nil {
		return err
	}

	if err := d.Set("pitr_catalog", parseTessellDataflixPITRInfoListWithResData(tessellDMMDataflixServiceView.PITRCatalog, d)); err != nil {
		return err
	}

	if err := d.Set("snapshot_catalog", parseTessellDMMDataflixBackupDTOListWithResData(tessellDMMDataflixServiceView.SnapshotCatalog, d)); err != nil {
		return err
	}

	return nil
}

func parseTessellDataflixPITRInfoListWithResData(pitrCatalog *[]model.TessellDataflixPITRInfo, d *schema.ResourceData) []interface{} {
	if pitrCatalog == nil {
		return nil
	}
	tessellDataflixPITRInfoList := make([]interface{}, 0)

	if pitrCatalog != nil {
		tessellDataflixPITRInfoList = make([]interface{}, len(*pitrCatalog))
		for i, tessellDataflixPITRInfoItem := range *pitrCatalog {
			tessellDataflixPITRInfoList[i] = parseTessellDataflixPITRInfo(&tessellDataflixPITRInfoItem)
		}
	}

	return tessellDataflixPITRInfoList
}

func parseTessellDataflixPITRInfo(pitrCatalog *model.TessellDataflixPITRInfo) interface{} {
	if pitrCatalog == nil {
		return nil
	}
	parsedPitrCatalog := make(map[string]interface{})
	parsedPitrCatalog["cloud"] = pitrCatalog.Cloud

	var regions *[]model.TessellDataflixPITRInfoForRegion
	if pitrCatalog.Regions != regions {
		parsedPitrCatalog["regions"] = parseTessellDataflixPITRInfoForRegionList(pitrCatalog.Regions)
	}

	return parsedPitrCatalog
}

func parseTessellDataflixPITRInfoForRegionList(regions *[]model.TessellDataflixPITRInfoForRegion) []interface{} {
	if regions == nil {
		return nil
	}
	tessellDataflixPITRInfoForRegionList := make([]interface{}, 0)

	if regions != nil {
		tessellDataflixPITRInfoForRegionList = make([]interface{}, len(*regions))
		for i, tessellDataflixPITRInfoForRegionItem := range *regions {
			tessellDataflixPITRInfoForRegionList[i] = parseTessellDataflixPITRInfoForRegion(&tessellDataflixPITRInfoForRegionItem)
		}
	}

	return tessellDataflixPITRInfoForRegionList
}

func parseTessellDataflixPITRInfoForRegion(regions *model.TessellDataflixPITRInfoForRegion) interface{} {
	if regions == nil {
		return nil
	}
	parsedRegions := make(map[string]interface{})
	parsedRegions["region"] = regions.Region

	var timeRanges *[]model.TessellDataflixFromTimeInfo
	if regions.TimeRanges != timeRanges {
		parsedRegions["time_ranges"] = parseTessellDataflixFromTimeInfoList(regions.TimeRanges)
	}

	return parsedRegions
}

func parseTessellDataflixFromTimeInfoList(timeRanges *[]model.TessellDataflixFromTimeInfo) []interface{} {
	if timeRanges == nil {
		return nil
	}
	tessellDataflixFromTimeInfoList := make([]interface{}, 0)

	if timeRanges != nil {
		tessellDataflixFromTimeInfoList = make([]interface{}, len(*timeRanges))
		for i, tessellDataflixFromTimeInfoItem := range *timeRanges {
			tessellDataflixFromTimeInfoList[i] = parseTessellDataflixFromTimeInfo(&tessellDataflixFromTimeInfoItem)
		}
	}

	return tessellDataflixFromTimeInfoList
}

func parseTessellDataflixFromTimeInfo(timeRanges *model.TessellDataflixFromTimeInfo) interface{} {
	if timeRanges == nil {
		return nil
	}
	parsedTimeRanges := make(map[string]interface{})
	parsedTimeRanges["from_time"] = timeRanges.FromTime
	parsedTimeRanges["to_time"] = timeRanges.ToTime

	var sharedWith *model.EntityAclSharingSummaryInfo
	if timeRanges.SharedWith != sharedWith {
		parsedTimeRanges["shared_with"] = []interface{}{parseEntityAclSharingSummaryInfo(timeRanges.SharedWith)}
	}

	return parsedTimeRanges
}

func parseEntityAclSharingSummaryInfo(entityAclSharingSummaryInfo *model.EntityAclSharingSummaryInfo) interface{} {
	if entityAclSharingSummaryInfo == nil {
		return nil
	}
	parsedEntityAclSharingSummaryInfo := make(map[string]interface{})
	parsedEntityAclSharingSummaryInfo["users"] = entityAclSharingSummaryInfo.Users

	return parsedEntityAclSharingSummaryInfo
}

func parseTessellDMMDataflixBackupDTOListWithResData(snapshotCatalog *[]model.TessellDMMDataflixBackupDTO, d *schema.ResourceData) []interface{} {
	if snapshotCatalog == nil {
		return nil
	}
	tessellDMMDataflixBackupDTOList := make([]interface{}, 0)

	if snapshotCatalog != nil {
		tessellDMMDataflixBackupDTOList = make([]interface{}, len(*snapshotCatalog))
		for i, tessellDMMDataflixBackupDTOItem := range *snapshotCatalog {
			tessellDMMDataflixBackupDTOList[i] = parseTessellDMMDataflixBackupDTO(&tessellDMMDataflixBackupDTOItem)
		}
	}

	return tessellDMMDataflixBackupDTOList
}

func parseTessellDMMDataflixBackupDTO(snapshotCatalog *model.TessellDMMDataflixBackupDTO) interface{} {
	if snapshotCatalog == nil {
		return nil
	}
	parsedSnapshotCatalog := make(map[string]interface{})
	parsedSnapshotCatalog["id"] = snapshotCatalog.Id
	parsedSnapshotCatalog["name"] = snapshotCatalog.Name
	parsedSnapshotCatalog["description"] = snapshotCatalog.Description
	parsedSnapshotCatalog["snapshot_time"] = snapshotCatalog.SnapshotTime
	parsedSnapshotCatalog["status"] = snapshotCatalog.Status
	parsedSnapshotCatalog["size"] = snapshotCatalog.Size
	parsedSnapshotCatalog["manual"] = snapshotCatalog.Manual

	var cloudAvailability *[]model.CloudRegionInfo
	if snapshotCatalog.CloudAvailability != cloudAvailability {
		parsedSnapshotCatalog["cloud_availability"] = parseCloudRegionInfoList(snapshotCatalog.CloudAvailability)
	}

	var databases *[]model.BackupDatabaseInfo
	if snapshotCatalog.Databases != databases {
		parsedSnapshotCatalog["databases"] = parseBackupDatabaseInfoList(snapshotCatalog.Databases)
	}

	var sharedWith *model.EntityAclSharingSummaryInfo
	if snapshotCatalog.SharedWith != sharedWith {
		parsedSnapshotCatalog["shared_with"] = []interface{}{parseEntityAclSharingSummaryInfo(snapshotCatalog.SharedWith)}
	}

	return parsedSnapshotCatalog
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

func parseBackupDatabaseInfoList(databases *[]model.BackupDatabaseInfo) []interface{} {
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
