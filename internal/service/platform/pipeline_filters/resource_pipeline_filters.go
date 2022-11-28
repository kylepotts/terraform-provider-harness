package pipeline_filters

import (
	"context"
	"net/http"

	"github.com/harness/harness-go-sdk/harness/nextgen"
	"github.com/harness/terraform-provider-harness/helpers"
	"github.com/harness/terraform-provider-harness/internal"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func ResourcePipelineFilters() *schema.Resource {
	resource := &schema.Resource{
		Description: "Resource for creating Harness Pipeline Filters.",

		ReadContext:   resourcePipelineFiltersRead,
		UpdateContext: resourcePipelineFiltersCreateOrUpdate,
		DeleteContext: resourcePipelineFiltersDelete,
		CreateContext: resourcePipelineFiltersCreateOrUpdate,
		Importer:      helpers.MultiLevelFilterImporter,

		Schema: map[string]*schema.Schema{
			"identifier": {
				Description: "Unique identifier of the resource.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"name": {
				Description: "Name of the pipeline filters.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"type": {
				Description:  "Type of pipeline filters. Currently supported types are {PipelineSetup, PipelineExecution, Deployment, Template, EnvironmentGroup, Environment}",
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice([]string{"PipelineSetup", "PipelineExecution", "Deployment", "Template", "EnvironmentGroup", "Environment"}, false),
			},
			"org_id": {
				Description: "Organization Identifier for the Entity.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"project_id": {
				Description: "Project Identifier for the Entity.",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"filter_properties": {
				Description: "Properties of the filters entity defined in Harness.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Required:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"filter_type": {
							Description:  "Corresponding Entity of the filters. Currently supported types are {Connector, DelegateProfile, Delegate, PipelineSetup, PipelineExecution, Deployment, Audit, Template, EnvironmentGroup, FileStore, CCMRecommendation, Anomaly, Environment}.",
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validation.StringInSlice([]string{"Connector", "DelegateProfile", "Delegate", "PipelineSetup", "PipelineExecution", "Deployment", "Audit", "Template", "EnvironmentGroup", "FileStore", "CCMRecommendation", "Anomaly", "Environment"}, false),
						},
						"tags": {
							Description: "Tags to associate with the resource. Tags should be in the form `name:value`.",
							Type:        schema.TypeSet,
							Optional:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			"filter_visibility": {
				Description:  "This indicates visibility of filters. By default, everyone can view this filter.",
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"EveryOne", "OnlyCreator"}, false),
			},
		},
	}

	return resource
}

func resourcePipelineFiltersRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	id := d.Id()

	type_ := d.Get("type").(string)
	resp, httpResp, err := c.FilterApi.PipelinegetFilter(ctx, c.AccountId, id, type_, &nextgen.FilterApiPipelinegetFilterOpts{
		OrgIdentifier:     helpers.BuildField(d, "org_id"),
		ProjectIdentifier: helpers.BuildField(d, "project_id"),
	})

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	if resp.Data == nil {
		d.SetId("")
		d.MarkNewResource()
		return nil
	}

	readPipelineFilter(d, resp.Data)

	return nil
}

func resourcePipelineFiltersCreateOrUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	var err error
	var resp nextgen.ResponseDtoFilter
	var httpResp *http.Response

	id := d.Id()
	filter := buildPipelineFilter(d)

	if id == "" {
		resp, httpResp, err = c.FilterApi.PipelinepostFilter(ctx, *filter, c.AccountId)
	} else {
		resp, httpResp, err = c.FilterApi.PipelineupdateFilter(ctx, *filter, c.AccountId)
	}

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	readPipelineFilter(d, resp.Data)

	return nil
}

func resourcePipelineFiltersDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c, ctx := meta.(*internal.Session).GetPlatformClientWithContext(ctx)

	type_ := d.Get("type").(string)

	_, httpResp, err := c.FilterApi.PipelinedeleteFilter(ctx, c.AccountId, d.Id(), type_, &nextgen.FilterApiPipelinedeleteFilterOpts{
		OrgIdentifier:     helpers.BuildField(d, "org_id"),
		ProjectIdentifier: helpers.BuildField(d, "project_id"),
	})

	if err != nil {
		return helpers.HandleApiError(err, d, httpResp)
	}

	return nil
}

func buildPipelineFilter(d *schema.ResourceData) *nextgen.Filter {
	filter := &nextgen.Filter{
		FilterProperties: &nextgen.FilterProperties{},
	}

	if attr, ok := d.GetOk("org_id"); ok {
		filter.OrgIdentifier = attr.(string)
	}

	if attr, ok := d.GetOk("project_id"); ok {
		filter.ProjectIdentifier = attr.(string)
	}

	if attr, ok := d.GetOk("filter_visibility"); ok {
		filter.FilterVisibility = attr.(string)
	}

	if attr, ok := d.GetOk("name"); ok {
		filter.Name = attr.(string)
	}

	if attr, ok := d.GetOk("identifier"); ok {
		filter.Identifier = attr.(string)
	}

	if attr, ok := d.GetOk("filter_properties"); ok {
		config := attr.([]interface{})[0].(map[string]interface{})
		if attr, ok := config["filter_type"]; ok {
			filter.FilterProperties.FilterType = attr.(string)
		}

		if attr := config["tags"].(*schema.Set).List(); len(attr) > 0 {
			filter.FilterProperties.Tags = helpers.ExpandTags(attr)
		}
	}

	return filter
}

func readPipelineFilter(d *schema.ResourceData, filter *nextgen.Filter) {
	d.SetId(filter.Identifier)
	d.Set("identifier", filter.Identifier)
	d.Set("org_id", filter.OrgIdentifier)
	d.Set("project_id", filter.ProjectIdentifier)
	d.Set("name", filter.Name)
	d.Set("type", filter.FilterProperties.FilterType)
	d.Set("filter_visibility", filter.FilterVisibility)
	d.Set("filter_properties", []interface{}{
		map[string]interface{}{
			"filter_type": filter.FilterProperties.FilterType,
			"tags":        helpers.FlattenTags(filter.FilterProperties.Tags),
		},
	})
}
