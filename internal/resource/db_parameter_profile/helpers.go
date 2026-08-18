package db_parameter_profile

import (
	//"fmt"
	//"time"

	"sort"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"terraform-provider-tessell/internal/model"
)

func setResourceData(d *schema.ResourceData, databaseParameterProfileResponse *model.DatabaseParameterProfileResponse) error {

	if err := d.Set("id", databaseParameterProfileResponse.Id); err != nil {
		return err
	}

	if err := d.Set("version_id", databaseParameterProfileResponse.VersionId); err != nil {
		return err
	}

	if err := d.Set("name", databaseParameterProfileResponse.Name); err != nil {
		return err
	}

	if err := d.Set("description", databaseParameterProfileResponse.Description); err != nil {
		return err
	}

	if err := d.Set("oob", databaseParameterProfileResponse.Oob); err != nil {
		return err
	}

	if err := d.Set("engine_type", databaseParameterProfileResponse.EngineType); err != nil {
		return err
	}

	if err := d.Set("engine_info", parseDatabaseParameterEngineInfoWithResData(databaseParameterProfileResponse.EngineInfo, d)); err != nil {
		return err
	}

	if err := d.Set("factory_parameter_id", databaseParameterProfileResponse.FactoryParameterId); err != nil {
		return err
	}

	if err := d.Set("status", databaseParameterProfileResponse.Status); err != nil {
		return err
	}

	// Map API response values (PUBLISHED, UNPUBLISHED, DRAFT) back to the
	// action verbs accepted by the schema (publish, unpublish, draft) so that
	// state stays consistent with what the user wrote in HCL and no drift occurs.
	maturityActionMap := map[string]string{
		"PUBLISHED":   "publish",
		"UNPUBLISHED": "unpublish",
		"DRAFT":       "draft",
	}
	maturityVal := ""
	if databaseParameterProfileResponse.MaturityStatus != nil {
		if mapped, ok := maturityActionMap[*databaseParameterProfileResponse.MaturityStatus]; ok {
			maturityVal = mapped
		} else {
			maturityVal = *databaseParameterProfileResponse.MaturityStatus
		}
	}
	if err := d.Set("maturity_status", maturityVal); err != nil {
		return err
	}

	if err := d.Set("owner", databaseParameterProfileResponse.Owner); err != nil {
		return err
	}

	if err := d.Set("parameters", parseDatabaseProfileParameterTypeListWithResData(databaseParameterProfileResponse.Parameters)); err != nil {
		return err
	}

	if err := d.Set("metadata", parseDatabaseParameterProfileMetadataWithResData(databaseParameterProfileResponse.Metadata, d)); err != nil {
		return err
	}

	if err := d.Set("driver_info", parseDatabaseParameterProfileDriverInfoWithResData(databaseParameterProfileResponse.DriverInfo, d)); err != nil {
		return err
	}

	if err := d.Set("user_id", databaseParameterProfileResponse.UserId); err != nil {
		return err
	}

	if err := d.Set("db_version", databaseParameterProfileResponse.DBVersion); err != nil {
		return err
	}

	if err := d.Set("date_created", databaseParameterProfileResponse.DateCreated); err != nil {
		return err
	}

	if err := d.Set("date_modified", databaseParameterProfileResponse.DateModified); err != nil {
		return err
	}

	if err := d.Set("infra_type", databaseParameterProfileResponse.InfraType); err != nil {
		return err
	}

	if err := d.Set("is_legacy", databaseParameterProfileResponse.IsLegacy); err != nil {
		return err
	}

	return nil
}

func parseDatabaseParameterEngineInfoWithResData(engineInfo *model.DatabaseParameterEngineInfo, d *schema.ResourceData) []interface{} {
	if engineInfo == nil {
		return nil
	}
	parsedEngineInfo := make(map[string]interface{})
	if d.Get("engine_info") != nil {
		engineInfoResourceData := d.Get("engine_info").([]interface{})
		if len(engineInfoResourceData) > 0 {
			parsedEngineInfo = (engineInfoResourceData[0]).(map[string]interface{})
		}
	}
	parsedEngineInfo["edition"] = engineInfo.Edition

	var oracle *model.DatabaseParameterEngineInfoOracle
	if engineInfo.Oracle != oracle {
		parsedEngineInfo["oracle"] = []interface{}{parseDatabaseParameterEngineInfoOracle(engineInfo.Oracle)}
	}

	return []interface{}{parsedEngineInfo}
}

func parseDatabaseParameterEngineInfo(engineInfo *model.DatabaseParameterEngineInfo) interface{} {
	if engineInfo == nil {
		return nil
	}
	parsedEngineInfo := make(map[string]interface{})
	parsedEngineInfo["edition"] = engineInfo.Edition

	var oracle *model.DatabaseParameterEngineInfoOracle
	if engineInfo.Oracle != oracle {
		parsedEngineInfo["oracle"] = []interface{}{parseDatabaseParameterEngineInfoOracle(engineInfo.Oracle)}
	}

	return parsedEngineInfo
}

func parseDatabaseParameterEngineInfoOracle(databaseParameterEngineInfo_oracle *model.DatabaseParameterEngineInfoOracle) interface{} {
	if databaseParameterEngineInfo_oracle == nil {
		return nil
	}
	parsedDatabaseParameterEngineInfo_oracle := make(map[string]interface{})
	parsedDatabaseParameterEngineInfo_oracle["multi_tenancy"] = databaseParameterEngineInfo_oracle.MultiTenancy

	return parsedDatabaseParameterEngineInfo_oracle
}

func parseDatabaseProfileParameterTypeListWithResData(parameters *[]model.DatabaseProfileParameterType) []interface{} {
	if parameters == nil {
		return nil
	}

	// Store ALL parameters from the API so that:
	// 1. State accurately reflects reality
	// 2. terraform import captures the full profile
	// 3. terraform state show displays all params
	//
	// Sort by name for a stable TypeList ordering — this prevents phantom diffs
	// when the API returns parameters in different orders across reads.
	// CustomizeDiff handles suppressing drift for non-config params.
	sorted := make([]model.DatabaseProfileParameterType, len(*parameters))
	copy(sorted, *parameters)
	sort.Slice(sorted, func(i, j int) bool {
		nameI, nameJ := "", ""
		if sorted[i].Name != nil {
			nameI = *sorted[i].Name
		}
		if sorted[j].Name != nil {
			nameJ = *sorted[j].Name
		}
		return nameI < nameJ
	})

	databaseProfileParameterTypeList := make([]interface{}, len(sorted))
	for i, param := range sorted {
		databaseProfileParameterTypeList[i] = parseDatabaseProfileParameterType(&param)
	}

	return databaseProfileParameterTypeList
}

func parseDatabaseProfileParameterTypeList(parameters *[]model.DatabaseProfileParameterType) []interface{} {
	if parameters == nil {
		return nil
	}
	databaseProfileParameterTypeList := make([]interface{}, 0)

	if parameters != nil {
		databaseProfileParameterTypeList = make([]interface{}, len(*parameters))
		for i, databaseProfileParameterTypeItem := range *parameters {
			databaseProfileParameterTypeList[i] = parseDatabaseProfileParameterType(&databaseProfileParameterTypeItem)
		}
	}

	return databaseProfileParameterTypeList
}

func parseDatabaseProfileParameterType(parameters *model.DatabaseProfileParameterType) interface{} {
	if parameters == nil {
		return nil
	}
	parsedParameters := make(map[string]interface{})
	parsedParameters["data_type"] = parameters.DataType
	parsedParameters["default_value"] = parameters.DefaultValue
	parsedParameters["apply_type"] = parameters.ApplyType
	parsedParameters["name"] = parameters.Name
	parsedParameters["description"] = parameters.Description
	parsedParameters["value"] = parameters.Value
	parsedParameters["allowed_values"] = parameters.AllowedValues
	parsedParameters["is_modified"] = parameters.IsModified
	parsedParameters["is_formula_type"] = parameters.IsFormulaType
	parsedParameters["source"] = parameters.Source
	parsedParameters["top_parameter"] = parameters.TopParameter
	parsedParameters["is_modifiable"] = parameters.IsModifiable
	parsedParameters["usage_type"] = parameters.UsageType

	return parsedParameters
}

func parseDatabaseParameterProfileMetadataWithResData(metadata *model.DatabaseParameterProfileMetadata, d *schema.ResourceData) []interface{} {
	if metadata == nil {
		return nil
	}
	parsedMetadata := make(map[string]interface{})
	if d.Get("metadata") != nil {
		metadataResourceData := d.Get("metadata").([]interface{})
		if len(metadataResourceData) > 0 {
			parsedMetadata = (metadataResourceData[0]).(map[string]interface{})
		}
	}
	parsedMetadata["data"] = metadata.Data

	return []interface{}{parsedMetadata}
}

func parseDatabaseParameterProfileMetadata(metadata *model.DatabaseParameterProfileMetadata) interface{} {
	if metadata == nil {
		return nil
	}
	parsedMetadata := make(map[string]interface{})
	parsedMetadata["data"] = metadata.Data

	return parsedMetadata
}

func parseDatabaseParameterProfileDriverInfoWithResData(driverInfo *model.DatabaseParameterProfileDriverInfo, d *schema.ResourceData) []interface{} {
	if driverInfo == nil {
		return nil
	}
	parsedDriverInfo := make(map[string]interface{})
	if d.Get("driver_info") != nil {
		driverInfoResourceData := d.Get("driver_info").([]interface{})
		if len(driverInfoResourceData) > 0 {
			parsedDriverInfo = (driverInfoResourceData[0]).(map[string]interface{})
		}
	}
	parsedDriverInfo["data"] = driverInfo.Data

	return []interface{}{parsedDriverInfo}
}

func parseDatabaseParameterProfileDriverInfo(driverInfo *model.DatabaseParameterProfileDriverInfo) interface{} {
	if driverInfo == nil {
		return nil
	}
	parsedDriverInfo := make(map[string]interface{})
	parsedDriverInfo["data"] = driverInfo.Data

	return parsedDriverInfo
}
