package filters

import (
	"context"
	"errors"
	"net/http"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func DataSourceFilters() *schema.Resource {
	resource := &schema.Resource{
		Description: "Data source for retrieving a Harness Filter. This data source allows you to fetch filters of the following types: {Connector, DelegateProfile, Delegate, EnvironmentGroup, FileStore, Environment}",

		ReadContext: dataSourceFiltersRead,

		Schema: map[string]*schema.Schema{
			"identifier": {
				Description: "Unique identifier of the resource",
				Type:        schema.TypeString,
				Required:    true,
			},
			"name": {
				Description: "Name of the Filter",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"type": {
				Description:  "Type of filter. Currently supported types are {Connector, DelegateProfile, Delegate, EnvironmentGroup, FileStore, Environment}.",
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice([]string{"Connector", "DelegateProfile", "Delegate", "EnvironmentGroup", "FileStore", "Environment"}, false),
			},
			"org_id": {
				Description: "organization Identifier for the Entity",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"project_id": {
				Description: "project Identifier for the Entity",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"filter_properties": {
				Description: "Properties of the filter entity defined in Harness.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"filter_type": {
							Description: "Corresponding Entity of the filter. Currently supported types are {Connector, DelegateProfile, Delegate, EnvironmentGroup, FileStore, Environment}.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"tags": {
							Description: "Tags to associate with the resource. Tags should be in the form `name:value`.",
							Type:        schema.TypeSet,
							Computed:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			"filter_visibility": {
				Description: "This indicates visibility of filter. By default, everyone can view this filter.",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}

	return resource
}

func dataSourceFiltersRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	var filter *nextgen.Filter
	var err error
	var httpResp *http.Response

	id := d.Get("identifier").(string)
	type_ := d.Get("type").(string)

	if id != "" {
		var resp nextgen.ResponseDtoFilter
		resp, httpResp, err = c.FilterApi.GetFilter(ctx, c.AccountId, id, type_, &nextgen.FilterApiGetFilterOpts{
			OrgIdentifier:     helpers.BuildField(d, "org_id"),
			ProjectIdentifier: helpers.BuildField(d, "project_id"),
		})
		filter = resp.Data
	} else {
		return diag.FromErr(errors.New(" identifier  must be specified"))
	}

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	if filter == nil {
		d.SetId("")
		d.MarkNewResource()
		return nil
	}

	readFilter(d, filter)

	return nil
}
