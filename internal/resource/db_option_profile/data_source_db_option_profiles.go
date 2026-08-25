package db_option_profile

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	apiClient "terraform-provider-tessell/internal/client"
	"terraform-provider-tessell/internal/model"
)

func DataSourceDBOptionProfiles() *schema.Resource {
	return &schema.Resource{

		ReadContext: dataSourceDBOptionProfilesRead,

		Schema: map[string]*schema.Schema{
			"status": {
				Type:        schema.TypeString,
				Description: "status",
				Optional:    true,
			},
			"engine_type": {
				Type:        schema.TypeString,
				Description: "Option Profile's engine-type",
				Optional:    true,
			},
			"version": {
				Type:        schema.TypeString,
				Description: "Option Profile's version",
				Optional:    true,
			},
			"db_option_profiles": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"description": {
							Type:        schema.TypeString,
							Description: "",
							Computed:    true,
						},
						"driver_info": {
							Type:        schema.TypeList,
							Description: "",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"data": {
										Type:        schema.TypeMap,
										Description: "",
										Computed:    true,
									},
								},
							},
						},
						"engine_type": {
							Type:        schema.TypeString,
							Description: "",
							Computed:    true,
						},
						"id": {
							Type:        schema.TypeString,
							Description: "Tessell generated UUID for the entity",
							Computed:    true,
						},
						"option_type_id": {
							Type:        schema.TypeString,
							Description: "Tessell option type UUID for the entity",
							Computed:    true,
						},
						"metadata": {
							Type:        schema.TypeList,
							Description: "",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"data": {
										Type:        schema.TypeMap,
										Description: "",
										Computed:    true,
									},
								},
							},
						},
						"name": {
							Type:        schema.TypeString,
							Description: "Name of the entity",
							Computed:    true,
						},
						"maturity_status": {
							Type:        schema.TypeString,
							Description: "",
							Computed:    true,
						},
						"options": {
							Type:        schema.TypeList,
							Description: "Database Option Profile's associated options",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"apply_now": {
										Type:        schema.TypeBool,
										Description: "",
										Computed:    true,
									},
									"name": {
										Type:        schema.TypeString,
										Description: "Option name",
										Computed:    true,
									},
									"option_settings": {
										Type:        schema.TypeList,
										Description: "",
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": {
													Type:        schema.TypeString,
													Description: "",
													Computed:    true,
												},
												"value": {
													Type:        schema.TypeString,
													Description: "",
													Computed:    true,
												},
												"is_global": {
													Type:        schema.TypeBool,
													Description: "",
													Computed:    true,
												},
											},
										},
									},
								},
							},
						},
						"owner": {
							Type:        schema.TypeString,
							Description: "",
							Computed:    true,
						},
						"tenant_id": {
							Type:        schema.TypeString,
							Description: "",
							Computed:    true,
						},
						"version": {
							Type:        schema.TypeString,
							Description: "Database Option Profile's version",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceDBOptionProfilesRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*apiClient.Client)

	var diags diag.Diagnostics

	var engineType, version, status string
	if !d.GetRawConfig().GetAttr("engine_type").IsNull() {
		engineType = d.Get("engine_type").(string)
	}
	if !d.GetRawConfig().GetAttr("version").IsNull() {
		version = d.Get("version").(string)
	}
	if !d.GetRawConfig().GetAttr("status").IsNull() {
		status = d.Get("status").(string)
	}

	response, _, err := client.GetDatabaseOptionProfilesForConsumption(status, engineType, version)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := setDataSourceValues(d, response.Response); err != nil {
		return diag.FromErr(err)
	}

	d.SetId("DBOptionProfileList_" + engineType + "_" + version + "_" + status)

	return diags
}

func setDataSourceValues(d *schema.ResourceData, DBOptionProfileList *[]model.TessellDatabaseOptionProfileConsumptionDTO) error {
	parsedDBOptionProfileList := make([]interface{}, 0)

	if DBOptionProfileList != nil {
		parsedDBOptionProfileList = make([]interface{}, len(*DBOptionProfileList))
		for i, DBOptionProfile := range *DBOptionProfileList {
			parsedDBOptionProfileList[i] = map[string]interface{}{
				"description":     DBOptionProfile.Description,
				"driver_info":     parseDatabaseOptionProfileDriverInfo(DBOptionProfile.DriverInfo),
				"engine_type":     DBOptionProfile.EngineType,
				"id":              DBOptionProfile.Id,
				"option_type_id":  DBOptionProfile.OptionTypeId,
				"metadata":        parseDatabaseOptionProfileMetadata(DBOptionProfile.Metadata),
				"name":            DBOptionProfile.Name,
				"maturity_status": DBOptionProfile.MaturityStatus,
				"options":         parseDatabaseProfileOptionTypeList(DBOptionProfile.Options),
				"owner":           DBOptionProfile.Owner,
				"tenant_id":       DBOptionProfile.TenantId,
				"version":         DBOptionProfile.Version,
			}
		}
	}

	if err := d.Set("db_option_profiles", parsedDBOptionProfileList); err != nil {
		return err
	}
	return nil
}
