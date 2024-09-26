package sanitized_db_snapshot

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	apiClient "terraform-provider-tessell/internal/client"
	"terraform-provider-tessell/internal/model"
)

func DataSourceSanitizedDBSnapshots() *schema.Resource {
	return &schema.Resource{

		ReadContext: dataSourceSanitizedDBSnapshotsRead,

		Schema: map[string]*schema.Schema{
			"availability_machine_id": {
				Type:        schema.TypeString,
				Description: "Id of the parent AvailabilityMachine",
				Required:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "Name of the DB Snapshot to filter with",
				Optional:    true,
			},
			"manual": {
				Type:        schema.TypeBool,
				Description: "Specifies whether the backup is captured manually",
				Optional:    true,
			},
			"sanitized_db_snapshots": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Description: "ID of the sanitized snapshot",
							Computed:    true,
						},
						"name": {
							Type:        schema.TypeString,
							Description: "Name of the sanitized snapshot",
							Computed:    true,
						},
						"snapshot_time": {
							Type:        schema.TypeString,
							Description: "Capture time of the source snapshot from which this sanitized snapshot is created",
							Computed:    true,
						},
						"status": {
							Type:        schema.TypeString,
							Description: "Database Backup Status",
							Computed:    true,
						},
						"size": {
							Type:        schema.TypeInt,
							Description: "Size of this snapshot (in bytes)",
							Computed:    true,
						},
						"manual": {
							Type:        schema.TypeBool,
							Description: "Specifies whether this snapshot is created based on a manual user request or through an automated schedule",
							Computed:    true,
						},
						"cloud_availability": {
							Type:        schema.TypeList,
							Description: "The cloud and region information where this snapshot has been made available at",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"cloud": {
										Type:        schema.TypeString,
										Description: "",
										Computed:    true,
									},
									"regions": {
										Type:        schema.TypeList,
										Description: "Region specific availability details for the snapshot",
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"region": {
													Type:        schema.TypeString,
													Description: "The region name",
													Computed:    true,
												},
												"status": {
													Type:        schema.TypeString,
													Description: "The current status of the snapshot in the respective region",
													Computed:    true,
												},
											},
										},
									},
								},
							},
						},
						"availability_config": {
							Type:        schema.TypeList,
							Description: "The config information for cloud and region availability for this snapshot",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"availability_configured_manually": {
										Type:        schema.TypeBool,
										Description: "",
										Computed:    true,
									},
									"dap_id": {
										Type:        schema.TypeString,
										Description: "",
										Computed:    true,
									},
									"cloud_availability_config": {
										Type:        schema.TypeList,
										Description: "",
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"cloud": {
													Type:        schema.TypeString,
													Description: "",
													Computed:    true,
												},
												"regions": {
													Type:        schema.TypeList,
													Description: "The list of regions and respective availability status",
													Computed:    true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"region": {
																Type:        schema.TypeString,
																Description: "",
																Computed:    true,
															},
															"status": {
																Type:        schema.TypeString,
																Description: "Database Backup Status",
																Computed:    true,
															},
														},
													},
												},
											},
										},
									},
								},
							},
						},
						"sanitization_info": {
							Type:        schema.TypeList,
							Description: "",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"source_snapshot_id": {
										Type:        schema.TypeString,
										Description: "ID of the as-is snapshot from which it was sanitized",
										Computed:    true,
									},
									"sanitization_schedule_id": {
										Type:        schema.TypeString,
										Description: "ID of the Sanitization Schedule which has created this snapshot",
										Computed:    true,
									},
									"sanitization_schedule": {
										Type:        schema.TypeString,
										Description: "Name of the Sanitization Schedule which has created this snapshot",
										Computed:    true,
									},
									"sanitization_script_id": {
										Type:        schema.TypeString,
										Description: "ID of the script which was used to create this snapshot",
										Computed:    true,
									},
									"sanitization_script": {
										Type:        schema.TypeString,
										Description: "Name of the script which was used to create this snapshot",
										Computed:    true,
									},
									"script_version": {
										Type:        schema.TypeString,
										Description: "Version of the script which was used to create this snapshot",
										Computed:    true,
									},
								},
							},
						},
						"databases": {
							Type:        schema.TypeList,
							Description: "The databases that are captured as part of the snapshot",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:        schema.TypeString,
										Description: "ID of the database",
										Computed:    true,
									},
									"name": {
										Type:        schema.TypeString,
										Description: "Name of the database",
										Computed:    true,
									},
									"status": {
										Type:        schema.TypeString,
										Description: "Status of the database as of capture of this snapshot",
										Computed:    true,
									},
								},
							},
						},
						"shared_with": {
							Type:        schema.TypeList,
							Description: "Tessell Entity ACL Sharing Summary Info",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"users": {
										Type:        schema.TypeList,
										Description: "",
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
						"backup_status": {
							Type:        schema.TypeString,
							Description: "",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceSanitizedDBSnapshotsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*apiClient.Client)

	var diags diag.Diagnostics

	availabilityMachineId := d.Get("availability_machine_id").(string)
	var name string
	if !d.GetRawConfig().GetAttr("name").IsNull() {
		name = d.Get("name").(string)
	}
	var manual bool
	if !d.GetRawConfig().GetAttr("manual").IsNull() {
		manual = d.Get("manual").(bool)
	}

	response, _, err := client.GetSanitizedDatabaseSnapshots(availabilityMachineId, &name, &manual)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := setDataSourceValues(d, response.Snapshots); err != nil {
		return diag.FromErr(err)
	}

	d.SetId("SanitizedDBSnapshotList")

	return diags
}

func setDataSourceValues(d *schema.ResourceData, SanitizedDBSnapshotList *[]model.SanitizedDatabaseSnapshot) error {
	parsedSanitizedDBSnapshotList := make([]interface{}, 0)

	if SanitizedDBSnapshotList != nil {
		parsedSanitizedDBSnapshotList = make([]interface{}, len(*SanitizedDBSnapshotList))
		for i, SanitizedDBSnapshot := range *SanitizedDBSnapshotList {
			parsedSanitizedDBSnapshotList[i] = map[string]interface{}{
				"id":                  SanitizedDBSnapshot.Id,
				"name":                SanitizedDBSnapshot.Name,
				"snapshot_time":       SanitizedDBSnapshot.SnapshotTime,
				"status":              SanitizedDBSnapshot.Status,
				"size":                SanitizedDBSnapshot.Size,
				"manual":              SanitizedDBSnapshot.Manual,
				"cloud_availability":  parseDatabaseSnapshotCloudRegionInfoList(SanitizedDBSnapshot.CloudAvailability),
				"availability_config": parseSnapshotAvailabilityConfigList(SanitizedDBSnapshot.AvailabilityConfig),
				"sanitization_info":   []interface{}{parseSanitizedDatabaseSnapshotSanitizationInfo(SanitizedDBSnapshot.SanitizationInfo)},
				"databases":           parseBackupDatabaseInfoList(SanitizedDBSnapshot.Databases),
				"shared_with":         []interface{}{parseEntityAclSharingSummaryInfo(SanitizedDBSnapshot.SharedWith)},
				"backup_status":       SanitizedDBSnapshot.BackupStatus,
			}
		}
	}

	if err := d.Set("sanitized_db_snapshots", parsedSanitizedDBSnapshotList); err != nil {
		return err
	}
	return nil
}
