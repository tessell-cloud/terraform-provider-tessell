package db_parameter_profile

import (
	"context"
	"fmt"
	"sort"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	apiClient "terraform-provider-tessell/internal/client"
	"terraform-provider-tessell/internal/model"
)

func ResourceDBParameterProfile() *schema.Resource {
	return &schema.Resource{

		CreateContext: resourceDBParameterProfileCreate,
		ReadContext:   resourceDBParameterProfileRead,
		UpdateContext: resourceDBParameterProfileUpdate,
		DeleteContext: resourceDBParameterProfileDelete,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		CustomizeDiff: suppressUnmanagedParametersDiff,

		Schema: map[string]*schema.Schema{
			"id": {
				Type:        schema.TypeString,
				Description: "Tessell generated UUID for the entity",
				Computed:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "Name of the Parameter Profile",
				Required:    true,
			},
			"description": {
				Type:        schema.TypeString,
				Description: "Database Parameter Profile description",
				Optional:    true,
			},
			"engine_type": {
				Type:        schema.TypeString,
				Description: "Database Engine type (e.g., POSTGRESQL, MYSQL, ORACLE, etc.)",
				Required:    true,
			},
			"db_version": {
				Type:        schema.TypeString,
				Description: "Database Engine version",
				Required:    true,
			},
			"parameters": {
				Type:        schema.TypeList,
				Description: "Parameters to manage in this profile. State stores all parameters returned by the API (for accurate import). Only changes to parameters declared here are planned.",
				Optional:    true,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Description: "Parameter name",
							Required:    true,
						},
						"value": {
							Type:        schema.TypeString,
							Description: "Parameter value",
							Required:    true,
						},
						"data_type": {
							Type:        schema.TypeString,
							Description: "Parameter data type",
							Computed:    true,
						},
						"default_value": {
							Type:        schema.TypeString,
							Description: "Default value of the parameter",
							Computed:    true,
						},
						"apply_type": {
							Type:        schema.TypeString,
							Description: "Apply type (static/dynamic)",
							Computed:    true,
						},
						"description": {
							Type:        schema.TypeString,
							Description: "Parameter description",
							Computed:    true,
						},
						"allowed_values": {
							Type:        schema.TypeString,
							Description: "Allowed values for the parameter",
							Computed:    true,
						},
						"is_modified": {
							Type:        schema.TypeBool,
							Description: "Whether the parameter value is modified from default",
							Computed:    true,
						},
						"is_formula_type": {
							Type:        schema.TypeBool,
							Description: "Whether the parameter is a formula type",
							Computed:    true,
						},
						"source": {
							Type:        schema.TypeString,
							Description: "Source of the parameter",
							Computed:    true,
						},
						"top_parameter": {
							Type:        schema.TypeBool,
							Description: "Boolean variable indicating a parameter is a most modified / key parameter",
							Computed:    true,
						},
						"is_modifiable": {
							Type:        schema.TypeBool,
							Description: "Whether the parameter is modifiable",
							Computed:    true,
						},
						"usage_type": {
							Type:        schema.TypeString,
							Description: "Usage type of the parameter",
							Computed:    true,
						},
					},
				},
			},
			"infra_type": {
				Type:        schema.TypeString,
				Description: "Infrastructure type (e.g., AWS, AZURE, GCP)",
				Optional:    true,
				Computed:    true,
			},
			"source_parameter_profile_id": {
				Type:        schema.TypeString,
				Description: "ID of the parameter profile to duplicate from",
				Optional:    true,
			},
			"propagation_policy": {
				Type:        schema.TypeList,
				Description: "Update propagation policy block. If not provided, the API will handle defaults.",
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"strategy": {
							Type:        schema.TypeString,
							Description: "Update propagation strategy: IMMEDIATELY, MAINTENANCE_WINDOW, CUSTOM_DATE_TIME, DO_NOT_APPLY.",
							Required:    true,
						},
						"time": {
							Type:        schema.TypeString,
							Description: "Propagation time (required when strategy is CUSTOM_DATE_TIME, e.g. 2026-05-15T10:00:00Z).",
							Optional:    true,
						},
					},
				},
			},
			// Read-only fields
			"version_id": {
				Type:        schema.TypeString,
				Description: "Tessell generated UUID for the entity version",
				Computed:    true,
			},
			"oob": {
				Type:        schema.TypeBool,
				Description: "Out of box parameter profile",
				Computed:    true,
			},
			"factory_parameter_id": {
				Type:        schema.TypeString,
				Description: "Tessell parameter type UUID for the entity",
				Computed:    true,
			},
			"status": {
				Type:        schema.TypeString,
				Description: "Status of the Parameter Profile",
				Computed:    true,
			},
			"maturity_status": {
				Type:        schema.TypeString,
				Description: "Action to change the maturity status of the Parameter Profile. Allowed values: \"draft\", \"publish\", \"unpublish\".",
				Optional:    true,
				Computed:    true,
				ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
					v := val.(string)
					allowed := map[string]bool{"draft": true, "publish": true, "unpublish": true}
					if !allowed[v] {
						errs = append(errs, fmt.Errorf("%q must be one of: draft, publish, unpublish; got: %s", key, v))
					}
					return
				},
			},
			"owner": {
				Type:        schema.TypeString,
				Description: "Owner of the Parameter Profile",
				Computed:    true,
			},
			"user_id": {
				Type:        schema.TypeString,
				Description: "Database Parameter Profile's user id",
				Computed:    true,
			},
			"date_created": {
				Type:        schema.TypeString,
				Description: "Timestamp when the entity was created",
				Computed:    true,
			},
			"date_modified": {
				Type:        schema.TypeString,
				Description: "Timestamp when the entity was last modified",
				Computed:    true,
			},
			"is_legacy": {
				Type:        schema.TypeBool,
				Description: "Whether this Parameter Profile is Legacy or not",
				Computed:    true,
			},
			"engine_info": {
				Type:        schema.TypeList,
				Description: "Database engine specific info (required for some engine types like PostgreSQL)",
				Optional:    true,
				Computed:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"edition": {
							Type:        schema.TypeString,
							Description: "Engine edition (e.g., COMMUNITY for PostgreSQL)",
							Optional:    true,
							Computed:    true,
						},
						"oracle": {
							Type:        schema.TypeList,
							Description: "Oracle specific engine info parameters",
							Optional:    true,
							Computed:    true,
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"multi_tenancy": {
										Type:        schema.TypeString,
										Description: "Oracle multi-tenancy mode (CDB or NON_CDB)",
										Optional:    true,
										Computed:    true,
									},
								},
							},
						},
					},
				},
			},
			"metadata": {
				Type:        schema.TypeList,
				Description: "Parameter profile metadata",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"data": {
							Type:        schema.TypeMap,
							Description: "Metadata key-value pairs",
							Computed:    true,
						},
					},
				},
			},
			"driver_info": {
				Type:        schema.TypeList,
				Description: "Driver information",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"data": {
							Type:        schema.TypeMap,
							Description: "Driver info key-value pairs",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func resourceDBParameterProfileCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*apiClient.Client)

	var diags diag.Diagnostics

	payload := formPayloadForCreateDatabaseParameterProfile(d)

	response, _, err := client.CreateDatabaseParameterProfile(payload)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(*response.Id)

	resourceDBParameterProfileRead(ctx, d, meta)

	return diags
}

func resourceDBParameterProfileRead(_ context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*apiClient.Client)

	var diags diag.Diagnostics

	id := d.Get("id").(string)
	if id == "" {
		id = d.Id()
	}

	response, statusCode, err := client.GetDatabaseParameterProfileById(id)
	if err != nil {
		if statusCode == 404 {
			d.SetId("")
			return diags
		}
		return diag.FromErr(err)
	}

	if err := setResourceData(d, response); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(*response.Id)

	return diags
}

func resourceDBParameterProfileUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*apiClient.Client)

	var diags diag.Diagnostics

	id := d.Get("id").(string)
	if id == "" {
		id = d.Id()
	}

	// Update description and/or parameters
	if d.HasChanges("description", "parameters") {
		// Fetch current usages to include in the PATCH request
		usageResponse, _, err := client.GetParameterProfileUsages(id)
		if err != nil {
			return diag.FromErr(fmt.Errorf("failed to fetch parameter profile usages: %w", err))
		}

		payload, configParamsEmpty := formPayloadForUpdateDatabaseParameterProfile(d, usageResponse)

		if configParamsEmpty && d.HasChange("parameters") {
			// Scenario C: user cleared all parameters (parameters = [] or removed all blocks).
			// API rejects PATCH with empty parameters, so call /reset instead.
			_, _, err := client.ResetDatabaseParameterProfile(id)
			if err != nil {
				return diag.FromErr(fmt.Errorf("failed to reset parameter profile to defaults: %w", err))
			}
		} else {
			// Scenario A (partial removal) or B (value change): send PATCH.
			// API replace-all semantics handle resets for removed params.
			_, _, err = client.UpdateDatabaseParameterProfile(id, payload)
			if err != nil {
				return diag.FromErr(err)
			}
		}
	}

	// Change maturity status (publish/unpublish/draft)
	// Value is validated at plan time by ValidateFunc — passed directly as the API action.
	if d.HasChange("maturity_status") {
		name := d.Get("name").(string)
		action := d.Get("maturity_status").(string)

		_, _, err := client.UpdateDatabaseParameterProfileMaturityStatus(name, action)
		if err != nil {
			return diag.FromErr(fmt.Errorf("failed to change maturity status to %s: %w", action, err))
		}
	}

	resourceDBParameterProfileRead(ctx, d, meta)

	return diags
}

func resourceDBParameterProfileDelete(_ context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*apiClient.Client)

	var diags diag.Diagnostics

	id := d.Get("id").(string)
	if id == "" {
		id = d.Id()
	}

	response, statusCode, err := client.DeleteDatabaseParameterProfile(id)
	if err != nil {
		return diag.FromErr(err)
	}

	if statusCode != 200 {
		return diag.FromErr(fmt.Errorf("deletion failed for tessell_parameter_profile with id %s. Received response: %+v", id, response))
	}

	return diags
}

// suppressUnmanagedParametersDiff controls the planned diff for parameters.
// State stores ALL parameters returned by the API (for import and state show).
// The diff logic handles three scenarios:
//
// Scenario A (remove param): param was isModified in old state but absent from
// new config → show it resetting to default_value so the user sees the reset
// in the plan. The PATCH payload explicitly sends default_value for the removed
// param (including formula-type params whose default_value is a formula string —
// the API accepts formula strings, resets the param and marks isFormulaType=true).
//
// Scenario B (change value): param in both config and state → show new value.
// PATCH includes the param with the new config value.
//
// Scenario C (clear all): config has zero parameter blocks → show all modified
// params resetting to defaults. Update function routes this to /reset endpoint
// ({} body) instead of PATCH, which fully restores all params to system defaults.
//
// Non-modified params not in config are pinned to old state (no drift shown).
func suppressUnmanagedParametersDiff(_ context.Context, d *schema.ResourceDiff, _ interface{}) error {
	// Nothing to suppress on create — no old state exists yet
	if d.Id() == "" {
		return nil
	}

	// Build a map of name → config-value from the user's raw .tf config
	configRaw := d.GetRawConfig()
	configParams := configRaw.GetAttr("parameters")
	if configParams.IsNull() || !configParams.IsKnown() {
		return nil
	}

	configByName := map[string]string{}
	for it := configParams.ElementIterator(); it.Next(); {
		_, paramVal := it.Element()
		if paramVal.IsNull() || !paramVal.IsKnown() {
			continue
		}
		nameVal := paramVal.GetAttr("name")
		valueVal := paramVal.GetAttr("value")
		if !nameVal.IsNull() && nameVal.IsKnown() && !valueVal.IsNull() && valueVal.IsKnown() {
			configByName[nameVal.AsString()] = valueVal.AsString()
		}
	}

	// Get old state (all ~100 params with computed fields already populated)
	oldParams, _ := d.GetChange("parameters")
	oldParamList, ok := oldParams.([]interface{})
	if !ok || len(oldParamList) == 0 {
		return nil
	}

	// Build merged list from old state:
	// - Config params (Scenario B): clone old state entry, update "value" to config value
	// - Removed modified params (Scenario A/C): clone entry, reset "value" to default_value
	// - Non-modified non-config params: pin to old state (no drift)
	merged := make([]interface{}, 0, len(oldParamList))
	for _, p := range oldParamList {
		pMap, ok := p.(map[string]interface{})
		if !ok {
			continue
		}
		name, _ := pMap["name"].(string)

		if configValue, inConfig := configByName[name]; inConfig {
			// Scenario B: param in config — use the new value
			updated := make(map[string]interface{}, len(pMap))
			for k, v := range pMap {
				updated[k] = v
			}
			updated["value"] = configValue
			merged = append(merged, updated)
		} else {
			isModified, _ := pMap["is_modified"].(bool)
			if isModified {
				// Scenario A/C: was modified, removed from config — show reset to default
				updated := make(map[string]interface{}, len(pMap))
				for k, v := range pMap {
					updated[k] = v
				}
				if defaultVal, ok := pMap["default_value"].(string); ok && defaultVal != "" {
					updated["value"] = defaultVal
				}
				updated["is_modified"] = false
				merged = append(merged, updated)
			} else {
				// Not modified, not in config — pin to old state (no drift)
				merged = append(merged, p)
			}
		}
	}

	// Sort by name — must match the stable order produced by Read (helpers.go)
	sort.Slice(merged, func(i, j int) bool {
		iMap, _ := merged[i].(map[string]interface{})
		jMap, _ := merged[j].(map[string]interface{})
		nameI, _ := iMap["name"].(string)
		nameJ, _ := jMap["name"].(string)
		return nameI < nameJ
	})

	return d.SetNew("parameters", merged)
}

func formPayloadForCreateDatabaseParameterProfile(d *schema.ResourceData) model.DatabaseParameterProfileRequest {
	name := d.Get("name").(string)
	engineType := d.Get("engine_type").(string)
	dbVersion := d.Get("db_version").(string)

	payload := model.DatabaseParameterProfileRequest{
		Name:       &name,
		EngineType: &engineType,
		DBVersion:  &dbVersion,
	}

	if v, ok := d.GetOk("description"); ok {
		description := v.(string)
		payload.Description = &description
	}

	if v, ok := d.GetOk("infra_type"); ok {
		infraType := v.(string)
		payload.InfraType = &infraType
	}

	if v, ok := d.GetOk("source_parameter_profile_id"); ok {
		sourceId := v.(string)
		payload.SourceParameterProfileId = &sourceId
	}

	// Parse engine_info if provided
	if v, ok := d.GetOk("engine_info"); ok {
		engineInfoList := v.([]interface{})
		if len(engineInfoList) > 0 {
			engineInfoMap := engineInfoList[0].(map[string]interface{})
			engineInfo := model.DatabaseParameterEngineInfo{}

			if edition, ok := engineInfoMap["edition"].(string); ok && edition != "" {
				engineInfo.Edition = &edition
			}

			if oracleList, ok := engineInfoMap["oracle"].([]interface{}); ok && len(oracleList) > 0 {
				oracleMap := oracleList[0].(map[string]interface{})
				oracle := model.DatabaseParameterEngineInfoOracle{}
				if multiTenancy, ok := oracleMap["multi_tenancy"].(string); ok && multiTenancy != "" {
					oracle.MultiTenancy = &multiTenancy
				}
				engineInfo.Oracle = &oracle
			}

			payload.EngineInfo = &engineInfo
		}
	}

	// Always send a non-null parameters array. API rejects null even when source_parameter_profile_id is set.
	// If no parameters block is declared in HCL, send an empty slice so the API receives [] instead of null.
	parameters := []model.DatabaseProfileParameterRequest{}
	if v, ok := d.GetOk("parameters"); ok {
		paramList := v.([]interface{})
		parameters = make([]model.DatabaseProfileParameterRequest, len(paramList))
		for i, param := range paramList {
			paramMap := param.(map[string]interface{})
			name := paramMap["name"].(string)
			value := paramMap["value"].(string)
			parameters[i] = model.DatabaseProfileParameterRequest{
				Name:  &name,
				Value: &value,
			}
		}
	}
	payload.Parameters = &parameters

	return payload
}

// formPayloadForUpdateDatabaseParameterProfile builds the PATCH payload.
// Returns the payload and a boolean indicating whether config params are empty
// (true = no params declared in HCL → caller should route to /reset instead).
func formPayloadForUpdateDatabaseParameterProfile(d *schema.ResourceData, usageResponse *model.ParameterProfileUsageResponse) (model.DatabaseParameterProfilePatchRequest, bool) {
	var propagationPolicy *model.UpdatePropagationPolicy
	if v, ok := d.GetOk("propagation_policy"); ok {
		blockList := v.([]interface{})
		if len(blockList) > 0 && blockList[0] != nil {
			blockMap := blockList[0].(map[string]interface{})
			strategy := blockMap["strategy"].(string)
			propagationPolicy = &model.UpdatePropagationPolicy{
				Strategy: &strategy,
			}
			if t, ok := blockMap["time"].(string); ok && t != "" {
				propagationPolicy.Time = &t
			}
		}
	}

	patchInfo := model.DatabaseParameterProfilePatchInfo{}

	if v, ok := d.GetOk("description"); ok {
		description := v.(string)
		patchInfo.Description = &description
	}

	// Build a map of name→value from what the user explicitly wrote in their .tf config.
	// d.GetOk("parameters") returns ALL params from state (including API defaults),
	// so we use GetRawConfig to get only what the user actually declared.
	configParamsEmpty := true
	configByName := make(map[string]string)
	configRaw := d.GetRawConfig()
	configParams := configRaw.GetAttr("parameters")
	if !configParams.IsNull() && configParams.IsKnown() {
		for it := configParams.ElementIterator(); it.Next(); {
			_, paramVal := it.Element()
			if paramVal.IsNull() || !paramVal.IsKnown() {
				continue
			}
			nameVal := paramVal.GetAttr("name")
			valueVal := paramVal.GetAttr("value")
			if nameVal.IsNull() || !nameVal.IsKnown() || valueVal.IsNull() || !valueVal.IsKnown() {
				continue
			}
			configByName[nameVal.AsString()] = valueVal.AsString()
		}
		configParamsEmpty = len(configByName) == 0
	}

	if !configParamsEmpty {
		// Build the full parameter list from old state, merging in config changes:
		//   - Param in config (Scenario B / new additions): use the new config value.
		//   - Param removed from config, is_modifiable=true (Scenario A): explicitly
		//     send default_value so the API resets it deterministically.
		//   - Param removed from config, is_modifiable=false: skip — API rejects any
		//     modification attempt even when resetting to the same default value.
		//   - Param new in config but absent from old state: use the config value.
		oldParams, _ := d.GetChange("parameters")
		oldParamList, _ := oldParams.([]interface{})

		oldParamNames := make(map[string]bool, len(oldParamList))
		allParams := make([]model.DatabaseProfileParameterRequest, 0, len(oldParamList))

		for _, p := range oldParamList {
			pMap, ok := p.(map[string]interface{})
			if !ok {
				continue
			}
			name, _ := pMap["name"].(string)
			oldParamNames[name] = true

			if newValue, inConfig := configByName[name]; inConfig {
				// Declared in config: send the (possibly updated) config value.
				allParams = append(allParams, model.DatabaseProfileParameterRequest{
					Name:  &name,
					Value: &newValue,
				})
			} else {
				// Removed from config (Scenario A): only reset if the param was actually
				// user-modified (is_modified=true). API-returned defaults that were never
				// touched by the user have is_modified=false and are already at their
				// default — no need to send them, and doing so causes API errors for
				// non-modifiable or enum-constrained params.
				// Also skip non-modifiable params (is_modifiable=false) — the API rejects
				// any attempt to set them even to their own default value.
				// The API accepts formula strings (e.g. "LEAST({instance_memory/9531392},5000)")
				// as values and correctly resets the param (isFormulaType=true, isModified=false).
				isModified, _ := pMap["is_modified"].(bool)
				isModifiable, _ := pMap["is_modifiable"].(bool)
				if !isModified || !isModifiable {
					continue
				}
				defaultVal, _ := pMap["default_value"].(string)
				allParams = append(allParams, model.DatabaseProfileParameterRequest{
					Name:  &name,
					Value: &defaultVal,
				})
			}
		}

		// Include newly-added params that were not present in the old state at all.
		for n, v := range configByName {
			if !oldParamNames[n] {
				nameCopy := n
				valueCopy := v
				allParams = append(allParams, model.DatabaseProfileParameterRequest{
					Name:  &nameCopy,
					Value: &valueCopy,
				})
			}
		}

		patchInfo.Parameters = &allParams
	}

	payload := model.DatabaseParameterProfilePatchRequest{
		ParameterProfileInfo: &patchInfo,
		PropagationPolicy:    propagationPolicy,
	}

	// Convert usage response to service instances for the PATCH request
	if usageResponse != nil && usageResponse.ServiceInstances != nil && len(*usageResponse.ServiceInstances) > 0 {
		serviceInstances := make([]model.ServiceInstancesPatchRequest, len(*usageResponse.ServiceInstances))
		for i, si := range *usageResponse.ServiceInstances {
			patchRequest := model.ServiceInstancesPatchRequest{
				Service: si.Service,
			}
			// Convert GovernanceInstanceInfo to GovernancePatchInstanceRequest
			if si.Instances != nil && len(*si.Instances) > 0 {
				instances := make([]model.GovernancePatchInstanceRequest, len(*si.Instances))
				for j, inst := range *si.Instances {
					instances[j] = model.GovernancePatchInstanceRequest{
						Name: inst.Name,
						Id:   inst.Id,
					}
				}
				patchRequest.Instances = &instances
			}
			serviceInstances[i] = patchRequest
		}
		payload.ServiceInstances = &serviceInstances
	}

	return payload, configParamsEmpty
}
