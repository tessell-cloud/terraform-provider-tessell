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
							Description: "DB Service backup Id",
							Computed:    true,
						},
						"name": {
							Type:        schema.TypeString,
							Description: "DB Service backup name",
							Computed:    true,
						},
						"backup_time": {
							Type:        schema.TypeString,
							Description: "DB Service backup capture time",
							Computed:    true,
						},
						"status": {
							Type:        schema.TypeString,
							Description: "Database Backup Status",
							Computed:    true,
						},
						"size": {
							Type:        schema.TypeInt,
							Description: "Backup size in bytes",
							Computed:    true,
						},
						"manual": {
							Type:        schema.TypeBool,
							Description: "Specifies whether the backup is captured manually",
							Computed:    true,
						},
						"cloud_availability": {
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
							Description: "",
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
							Description: "The databases that are captured as part of the backup",
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:        schema.TypeString,
										Description: "Databases Id",
										Computed:    true,
									},
									"name": {
										Type:        schema.TypeString,
										Description: "Databases name",
										Computed:    true,
									},
									"status": {
										Type:        schema.TypeString,
										Description: "Databases status",
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
										Description: "snapshot from which backup was created",
										Computed:    true,
									},
									"snapshot_name": {
										Type:        schema.TypeString,
										Description: "",
										Computed:    true,
									},
									"snapshot_time": {
										Type:        schema.TypeString,
										Description: "snapshot creation time",
										Computed:    true,
									},
								},
							},
						},
						"shared_with": {
							Type:        schema.TypeList,
							Description: "Users having shared access to the Database Backup",
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
																Description: "time till backup will be live",
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
