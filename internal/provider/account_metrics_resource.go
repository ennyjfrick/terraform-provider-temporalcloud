package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/temporalio/tcld/protogen/api/account/v1"
	"github.com/temporalio/tcld/protogen/api/accountservice/v1"
	"github.com/temporalio/tcld/protogen/api/requestservice/v1"

	"github.com/temporalio/terraform-provider-temporalcloud/internal/client"
)

type (
	accountMetricsResource struct {
		accountClient accountservice.AccountServiceClient
		requestClient requestservice.RequestServiceClient
	}

	accountMetricsResourceModel struct {
		ID               types.String `tfsdk:"id"`
		Enabled          types.Bool   `tfsdk:"enabled"`
		AcceptedClientCA types.String `tfsdk:"accepted_client_ca"`
		Endpoint         types.String `tfsdk:"endpoint"`

		Timeouts timeouts.Value `tfsdk:"timeouts"`
	}
)

var (
	_ resource.Resource                   = (*accountMetricsResource)(nil)
	_ resource.ResourceWithConfigure      = (*accountMetricsResource)(nil)
	_ resource.ResourceWithImportState    = (*accountMetricsResource)(nil)
	_ resource.ResourceWithValidateConfig = (*accountMetricsResource)(nil)
)

func NewAccountMetricsResource() resource.Resource {
	return &accountMetricsResource{}
}

func (r *accountMetricsResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	clientStore, ok := req.ProviderData.(client.ClientStore)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected client.CloudClient, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.accountClient = clientStore.AccountServiceClient()
	r.requestClient = clientStore.RequestServiceClient()
}

// Metadata returns the resource type name
func (r *accountMetricsResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "account_metrics"
}

// Schema defines the schema for the resource
func (r *accountMetricsResource) Schema(ctx context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Configures a Temporal Cloud account's metrics",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "A unique identifier for the account's metrics configuration",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"enabled": schema.BoolAttribute{
				Description: "If true, enables metrics for the account",
				Required:    true,
			},
			"accepted_client_ca": schema.StringAttribute{
				Description: "The Base64-encoded CA cert in PEM format used to authenticate clients connecting to the metrics endpoint. Required when enabled is true.",
			},
			"endpoint": schema.StringAttribute{
				Description: "The Prometheus metrics endpoint URI",
				Computed:    true,
			},
		},
		Blocks: map[string]schema.Block{
			"timeouts": timeouts.Block(
				ctx, timeouts.Opts{
					Create: true,
					Delete: true,
					Update: true,
				},
			),
		},
	}
}

// ValidateConfig validates the resource's config
func (r *accountMetricsResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data accountMetricsResourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if data.Enabled.ValueBool() && (data.AcceptedClientCA.IsNull() || data.AcceptedClientCA.IsUnknown() || data.AcceptedClientCA.ValueString() == "") {
		resp.Diagnostics.AddAttributeError(
			path.Root("accepted_client_ca"),
			"Missing Attribute Configuration",
			"Expected accepted_client_ca to be configured when enabled is true",
		)

		return
	}
}

// Create creates the resource and sets the initial Terraform state
func (r *accountMetricsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan accountMetricsResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	createTimeout, diags := plan.Timeouts.Create(ctx, defaultCreateTimeout)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	accResp, err := r.accountClient.GetAccount(ctx, &accountservice.GetAccountRequest{})
	if err != nil {
		resp.Diagnostics.AddError("Failed to get account information", err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(ctx, createTimeout)
	defer cancel()

	// account metrics config always exists, "create" only really sets the state
	metricsReq := &accountservice.UpdateAccountRequest{
		ResourceVersion: accResp.GetAccount().GetResourceVersion(),
		Spec: &account.AccountSpec{
			Metrics: &account.MetricsSpec{
				Enabled:          plan.Enabled.ValueBool(),
				AcceptedClientCa: plan.AcceptedClientCA.ValueString(),
			},
			OutputSinks: accResp.GetAccount().GetSpec().GetOutputSinks(),
		},
	}

	metricsResp, err := r.accountClient.UpdateAccount(ctx, metricsReq)
	if err != nil {
		resp.Diagnostics.AddError("Failed to create metrics resource", err.Error())
		return
	}

	if err := client.AwaitRequestStatus(ctx, r.requestClient, metricsResp.GetRequestStatus()); err != nil {
		resp.Diagnostics.AddError("Failed to create metrics resource", err.Error())
		return
	}

	accResp, err = r.accountClient.GetAccount(ctx, &accountservice.GetAccountRequest{})
	if err != nil {
		resp.Diagnostics.AddError("Failed to get resource state", err.Error())
		return
	}

	updateAccountMetricsModelFromSpec(&plan, accResp.GetAccount())
	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

// Read refreshes the Terraform state with the latest data
func (r *accountMetricsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state accountMetricsResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	accResp, err := r.accountClient.GetAccount(ctx, &accountservice.GetAccountRequest{})
	if err != nil {
		resp.Diagnostics.AddError("Failed to get account metrics", err.Error())
		return
	}

	updateAccountMetricsModelFromSpec(&state, accResp.GetAccount())
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}

// Update updates the resource and sets the updated Terraform state on success
func (r *accountMetricsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan accountMetricsResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	updateTimeout, diags := plan.Timeouts.Create(ctx, defaultCreateTimeout)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	accResp, err := r.accountClient.GetAccount(ctx, &accountservice.GetAccountRequest{})
	if err != nil {
		resp.Diagnostics.AddError("Failed to get current account metrics status", err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(ctx, updateTimeout)
	defer cancel()

	metricsReq := &accountservice.UpdateAccountRequest{
		ResourceVersion: accResp.GetAccount().GetResourceVersion(),
		Spec: &account.AccountSpec{
			Metrics: &account.MetricsSpec{
				Enabled:          plan.Enabled.ValueBool(),
				AcceptedClientCa: plan.AcceptedClientCA.ValueString(),
			},
			OutputSinks: accResp.GetAccount().GetSpec().GetOutputSinks(),
		},
	}

	metricsResp, err := r.accountClient.UpdateAccount(ctx, metricsReq)
	if err != nil {
		resp.Diagnostics.AddError("Failed to update account metrics", err.Error())
		return
	}

	if err := client.AwaitRequestStatus(ctx, r.requestClient, metricsResp.GetRequestStatus()); err != nil {
		resp.Diagnostics.AddError("Failed to update account metrics", err.Error())
		return
	}

	accResp, err = r.accountClient.GetAccount(ctx, &accountservice.GetAccountRequest{})
	if err != nil {
		resp.Diagnostics.AddError("Failed to get account metrics after update", err.Error())
		return
	}

	updateAccountMetricsModelFromSpec(&plan, accResp.GetAccount())
	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

// Delete deletes the resource and removes the Terraform state on success
func (r *accountMetricsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state accountMetricsResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	deleteTimeout, diags := state.Timeouts.Delete(ctx, defaultDeleteTimeout)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	accResp, err := r.accountClient.GetAccount(ctx, &accountservice.GetAccountRequest{})
	if err != nil {
		resp.Diagnostics.AddError("Failed to get current account metrics status", err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(ctx, deleteTimeout)
	defer cancel()

	// can't actually "delete" account metrics config, setting to disabled implicitly is the best equivalent
	metricsReq := &accountservice.UpdateAccountRequest{
		ResourceVersion: accResp.GetAccount().GetResourceVersion(),
		Spec: &account.AccountSpec{
			Metrics: &account.MetricsSpec{
				Enabled: false,
			},
			OutputSinks: accResp.GetAccount().GetSpec().GetOutputSinks(),
		},
	}

	metricsResp, err := r.accountClient.UpdateAccount(ctx, metricsReq)
	if err != nil {
		resp.Diagnostics.AddError("Failed to delete metrics resource", err.Error())
		return
	}

	if err := client.AwaitRequestStatus(ctx, r.requestClient, metricsResp.GetRequestStatus()); err != nil {
		resp.Diagnostics.AddError("Failed to delete metrics resource", err.Error())
		return
	}
}

func (r *accountMetricsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func updateAccountMetricsModelFromSpec(state *accountMetricsResourceModel, spec *account.Account) {
	state.Enabled = types.BoolValue(spec.GetSpec().GetMetrics().GetEnabled())
	state.AcceptedClientCA = types.StringValue(spec.GetSpec().GetMetrics().GetAcceptedClientCa())
	state.Endpoint = types.StringValue(spec.GetMetrics().GetUri())
	state.ID = types.StringValue("account-metrics") // no real ID to key off of here other than account ID, which is hard to get via the API
}
