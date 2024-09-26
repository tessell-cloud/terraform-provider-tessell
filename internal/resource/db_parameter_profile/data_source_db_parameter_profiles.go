package db_parameter_profile

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	apiClient "terraform-provider-tessell/internal/client"
	"terraform-provider-tessell/internal/model"
)

func DataSourceDBParameterProfiles() *schema.Resource {
	return &schema.Resource{

		ReadContext: dataSourceDBParameterProfilesRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Description: "Name of the Parameter Profile",
				Optional:    true,
			},
			"status": {
				Type:        schema.TypeString,
				Description: "status",
				Optional:    true,
			},
			"engine_type": {
				Type:        schema.TypeString,
				Description: "Availaility Machine's engine-types",
				Optional:    true,
			},
			"db_parameter_profiles": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Description: "Tessell generated UUID for the entity",
							Computed:    true,
						},
						"version_id": {
							Type:        schema.TypeString,
							Description: "Tessell generated UUID for the entity",
							Computed:    true,
						},
						"name": {
							Type:        schema.TypeString,
							Description: "Name of the entity",
							Computed:    true,
						},
						"description": {
							Type:        schema.TypeString,
							Description: "Database Parameter Profile description",
							Computed:    true,
						},
						"oob": {
							Type:        schema.TypeBool,
							Description: "",
							Computed:    true,
						},
						"engine_type": {
							Type:        schema.TypeString,
							Description: "",
							Computed:    true,
						},
						"factory_parameter_id": {
							Type:        schema.TypeString,
							Description: "Tessell parameter type UUID for the entity",
							Computed:    true,
						},
						"status": {
							Type:        schema.TypeString,
							Description: "",
							Computed:    true,
						},
						"maturity_status": {
							Type:        schema.TypeString,
							Description: "",
							Computed:    true,
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
						"logged_in_user_role": {
							Type:        schema.TypeString,
							Description: "The role of the logged in user for accessing the db profile",
							Computed:    true,
						},
						"parameters": {
							Type:        schema.TypeList,
							Description: "Parameter Profile's associated parameters",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"data_type": {
										Type:        schema.TypeString,
										Description: "",
										Computed:    true,
									},
									"default_value": {
										Type:        schema.TypeString,
										Description: "",
										Computed:    true,
									},
									"apply_type": {
										Type:        schema.TypeString,
										Description: "",
										Computed:    true,
									},
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
									"allowed_values": {
										Type:        schema.TypeString,
										Description: "",
										Computed:    true,
									},
									"is_modified": {
										Type:        schema.TypeBool,
										Description: "",
										Computed:    true,
									},
									"is_formula_type": {
										Type:        schema.TypeBool,
										Description: "",
										Computed:    true,
									},
								},
							},
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
						"user_id": {
							Type:        schema.TypeString,
							Description: "Database Parameter Profile's user id",
							Computed:    true,
						},
						"shared_with": {
							Type:        schema.TypeList,
							Description: "Tessell Entity ACL Sharing Info",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"users": {
										Type:        schema.TypeList,
										Description: "",
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"email_id": {
													Type:        schema.TypeString,
													Description: "",
													Computed:    true,
												},
												"role": {
													Type:        schema.TypeString,
													Description: "",
													Computed:    true,
												},
											},
										},
									},
								},
							},
						},
						"db_version": {
							Type:        schema.TypeString,
							Description: "Database Parameter Profile's version",
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
					},
				},
			},
		},
	}
}

func dataSourceDBParameterProfilesRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*apiClient.Client)

	var diags diag.Diagnostics

	var name string
	if !d.GetRawConfig().GetAttr("name").IsNull() {
		name = d.Get("name").(string)
	}
	var engineType string
	if !d.GetRawConfig().GetAttr("engine_type").IsNull() {
		engineType = d.Get("engine_type").(string)
	}
	var status string
	if !d.GetRawConfig().GetAttr("status").IsNull() {
		status = d.Get("status").(string)
	}

	response, _, err := client.GetDatabaseParameterProfilesForConsumers(&status, &engineType, &name)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := setDataSourceValues(d, response.Response); err != nil {
		return diag.FromErr(err)
	}

	d.SetId("DBParameterProfileList_" + strings.ReplaceAll(name, " ", "_"))

	return diags
}

func setDataSourceValues(d *schema.ResourceData, DBParameterProfileList *[]model.DatabaseParameterProfileResponse) error {
	parsedDBParameterProfileList := make([]interface{}, 0)

	if DBParameterProfileList != nil {
		parsedDBParameterProfileList = make([]interface{}, len(*DBParameterProfileList))
		for i, DBParameterProfile := range *DBParameterProfileList {
			parsedDBParameterProfileList[i] = map[string]interface{}{
				"id":                   DBParameterProfile.Id,
				"version_id":           DBParameterProfile.VersionId,
				"name":                 DBParameterProfile.Name,
				"description":          DBParameterProfile.Description,
				"oob":                  DBParameterProfile.Oob,
				"engine_type":          DBParameterProfile.EngineType,
				"factory_parameter_id": DBParameterProfile.FactoryParameterId,
				"status":               DBParameterProfile.Status,
				"maturity_status":      DBParameterProfile.MaturityStatus,
				"owner":                DBParameterProfile.Owner,
				"tenant_id":            DBParameterProfile.TenantId,
				"logged_in_user_role":  DBParameterProfile.LoggedInUserRole,
				"parameters":           parseDatabaseProfileParameterTypeList(DBParameterProfile.Parameters),
				"metadata":             []interface{}{parseDatabaseParameterProfileMetadata(DBParameterProfile.Metadata)},
				"driver_info":          []interface{}{parseDatabaseParameterProfileDriverInfo(DBParameterProfile.DriverInfo)},
				"user_id":              DBParameterProfile.UserId,
				"shared_with":          []interface{}{parseEntityAclSharingInfo(DBParameterProfile.SharedWith)},
				"db_version":           DBParameterProfile.DBVersion,
				"date_created":         DBParameterProfile.DateCreated,
				"date_modified":        DBParameterProfile.DateModified,
			}
		}
	}

	if err := d.Set("db_parameter_profiles", parsedDBParameterProfileList); err != nil {
		return err
	}
	return nil
}
