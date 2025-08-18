package db_parameter_profile

import (
	//"fmt"
	//"time"

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

	if err := d.Set("factory_parameter_id", databaseParameterProfileResponse.FactoryParameterId); err != nil {
		return err
	}

	if err := d.Set("status", databaseParameterProfileResponse.Status); err != nil {
		return err
	}

	if err := d.Set("maturity_status", databaseParameterProfileResponse.MaturityStatus); err != nil {
		return err
	}

	if err := d.Set("owner", databaseParameterProfileResponse.Owner); err != nil {
		return err
	}

	if err := d.Set("tenant_id", databaseParameterProfileResponse.TenantId); err != nil {
		return err
	}

	if err := d.Set("logged_in_user_role", databaseParameterProfileResponse.LoggedInUserRole); err != nil {
		return err
	}

	if err := d.Set("parameters", parseDatabaseProfileParameterTypeListWithResData(databaseParameterProfileResponse.Parameters, d)); err != nil {
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

	if err := d.Set("shared_with", parseEntityAclSharingInfoWithResData(databaseParameterProfileResponse.SharedWith, d)); err != nil {
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

	return nil
}

func parseDatabaseProfileParameterTypeListWithResData(parameters *[]model.DatabaseProfileParameterType, d *schema.ResourceData) []interface{} {
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
	parsedParameters["value"] = parameters.Value
	parsedParameters["allowed_values"] = parameters.AllowedValues
	parsedParameters["is_modified"] = parameters.IsModified
	parsedParameters["is_formula_type"] = parameters.IsFormulaType

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
