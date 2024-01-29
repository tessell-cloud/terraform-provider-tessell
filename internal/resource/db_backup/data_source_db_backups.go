package db_backup

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	apiClient "terraform-provider-tessell/internal/client"
	"terraform-provider-tessell/internal/model"
)

func DataSourceDBBackups() *schema.Resource {
	return &schema.Resource{

		ReadContext: dataSourceDBBackupsRead,

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
			"db_backups": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Description: "ID of the backup",
							Computed:    true,
						},
						"name": {
							Type:        schema.TypeString,
							Description: "Name of the backup",
							Computed:    true,
						},
						"backup_time": {
							Type:        schema.TypeString,
							Description: "Backup capture time",
							Computed:    true,
						},
						"status": {
							Type:        schema.TypeString,
							Description: "Database Backup Status",
							Computed:    true,
						},
						"size": {
							Type:        schema.TypeInt,
							Description: "Size of this backup (in bytes)",
							Computed:    true,
						},
						"manual": {
							Type:        schema.TypeBool,
							Description: "Specifies whether the backup is captured as per manual user request or per automated schedule",
							Computed:    true,
						},
						"cloud_availability": {
							Type:        schema.TypeList,
							Description: "The cloud and region information where this backup has been made available at",
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
										Description: "The regions details",
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"region": {
													Type:        schema.TypeString,
													Description: "The cloud region name",
													Computed:    true,
												},
												"availability_zones": {
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
								},
							},
						},
						"availability_config": {
							Type:        schema.TypeList,
							Description: "The config information for cloud and region availability for this backup",
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
													Description: "The list of regions and respective avaoilability status",
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
						"databases": {
							Type:        schema.TypeList,
							Description: "The databases that are captured as part of this backup",
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
						"backup_info": {
							Type:        schema.TypeList,
							Description: "",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"source_snapshot_id": {
										Type:        schema.TypeString,
										Description: "ID of snapshot from which this backup was created",
										Computed:    true,
									},
									"snapshot_name": {
										Type:        schema.TypeString,
										Description: "Name of snapshot from which this backup was created",
										Computed:    true,
									},
									"snapshot_time": {
										Type:        schema.TypeString,
										Description: "Capture time of snapshot from which this backup was created",
										Computed:    true,
									},
								},
							},
						},
						"shared_with": {
							Type:        schema.TypeList,
							Description: "List of users who have access to this backup",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"users": {
										Type:        schema.TypeList,
										Description: "",
										Computed:    true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"user_email": {
													Type:        schema.TypeString,
													Description: "email of the user",
													Computed:    true,
												},
												"download_url_status": {
													Type:        schema.TypeString,
													Description: "",
													Computed:    true,
												},
												"expiry_config": {
													Type:        schema.TypeList,
													Description: "The backup expiry config",
													Computed:    true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"expire_at": {
																Type:        schema.TypeString,
																Description: "time-to-live for the Pre auth url",
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
						"download_url_status": {
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

func dataSourceDBBackupsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*apiClient.Client)

	var diags diag.Diagnostics

	name := d.Get("name").(string)
	availabilityMachineId := d.Get("availability_machine_id").(string)
	manual := d.Get("manual").(bool)

	response, _, err := client.GetDatabaseBackups(availabilityMachineId, name, manual)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := setDataSourceValues(d, response.Backups); err != nil {
		return diag.FromErr(err)
	}

	d.SetId("DBBackupList")

	return diags
}

func setDataSourceValues(d *schema.ResourceData, DBBackupList *[]model.DatabaseBackup) error {
	parsedDBBackupList := make([]interface{}, 0)

	if DBBackupList != nil {
		parsedDBBackupList = make([]interface{}, len(*DBBackupList))
		for i, DBBackup := range *DBBackupList {
			parsedDBBackupList[i] = map[string]interface{}{
				"id":                  DBBackup.Id,
				"name":                DBBackup.Name,
				"backup_time":         DBBackup.BackupTime,
				"status":              DBBackup.Status,
				"size":                DBBackup.Size,
				"manual":              DBBackup.Manual,
				"cloud_availability":  parseCloudRegionInfoList(DBBackup.CloudAvailability),
				"availability_config": parseSnapshotAvailabilityConfigList(DBBackup.AvailabilityConfig),
				"databases":           parseBackupDatabaseInfoList(DBBackup.Databases),
				"backup_info":         []interface{}{parseBackupSourceInfo(DBBackup.BackupInfo)},
				"shared_with":         []interface{}{parseDatabaseBackupSharedWith(DBBackup.SharedWith)},
				"download_url_status": DBBackup.DownloadUrlStatus,
			}
		}
	}

	if err := d.Set("db_backups", parsedDBBackupList); err != nil {
		return err
	}
	return nil
}
