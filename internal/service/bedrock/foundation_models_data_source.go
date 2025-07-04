// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bedrock

import (
	"context"

	"github.com/YakDriver/regexache"
	"github.com/aws/aws-sdk-go-v2/service/bedrock"
	awstypes "github.com/aws/aws-sdk-go-v2/service/bedrock/types"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-aws/internal/framework"
	fwflex "github.com/hashicorp/terraform-provider-aws/internal/framework/flex"
	fwtypes "github.com/hashicorp/terraform-provider-aws/internal/framework/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

// @FrameworkDataSource("aws_bedrock_foundation_models", name="Foundation Models")
func newFoundationModelsDataSource(context.Context) (datasource.DataSourceWithConfigure, error) {
	return &foundationModelsDataSource{}, nil
}

type foundationModelsDataSource struct {
	framework.DataSourceWithModel[foundationModelsDataSourceModel]
}

func (d *foundationModelsDataSource) Schema(ctx context.Context, request datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"by_customization_type": schema.StringAttribute{
				CustomType: fwtypes.StringEnumType[awstypes.ModelCustomization](),
				Optional:   true,
			},
			"by_inference_type": schema.StringAttribute{
				CustomType: fwtypes.StringEnumType[awstypes.InferenceType](),
				Optional:   true,
			},
			"by_output_modality": schema.StringAttribute{
				CustomType: fwtypes.StringEnumType[awstypes.ModelModality](),
				Optional:   true,
			},
			"by_provider": schema.StringAttribute{
				Optional: true,
				Validators: []validator.String{
					stringvalidator.RegexMatches(regexache.MustCompile(`^[A-Za-z0-9- ]{1,63}$`), ""),
				},
			},
			names.AttrID:      framework.IDAttribute(),
			"model_summaries": framework.DataSourceComputedListOfObjectAttribute[foundationModelSummaryModel](ctx),
		},
	}
}

func (d *foundationModelsDataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	var data foundationModelsDataSourceModel
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	conn := d.Meta().BedrockClient(ctx)

	input := &bedrock.ListFoundationModelsInput{}
	response.Diagnostics.Append(fwflex.Expand(ctx, data, input)...)
	if response.Diagnostics.HasError() {
		return
	}

	output, err := conn.ListFoundationModels(ctx, input)

	if err != nil {
		response.Diagnostics.AddError("listing Bedrock Foundation Models", err.Error())

		return
	}

	response.Diagnostics.Append(fwflex.Flatten(ctx, output, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	data.ID = types.StringValue(d.Meta().Region(ctx))

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}

type foundationModelsDataSourceModel struct {
	framework.WithRegionModel
	ByCustomizationType fwtypes.StringEnum[awstypes.ModelCustomization]              `tfsdk:"by_customization_type"`
	ByInferenceType     fwtypes.StringEnum[awstypes.InferenceType]                   `tfsdk:"by_inference_type"`
	ByOutputModality    fwtypes.StringEnum[awstypes.ModelModality]                   `tfsdk:"by_output_modality"`
	ByProvider          types.String                                                 `tfsdk:"by_provider"`
	ID                  types.String                                                 `tfsdk:"id"`
	ModelSummaries      fwtypes.ListNestedObjectValueOf[foundationModelSummaryModel] `tfsdk:"model_summaries"`
}

type foundationModelSummaryModel struct {
	CustomizationsSupported    fwtypes.SetOfString `tfsdk:"customizations_supported"`
	InferenceTypesSupported    fwtypes.SetOfString `tfsdk:"inference_types_supported"`
	InputModalities            fwtypes.SetOfString `tfsdk:"input_modalities"`
	ModelARN                   fwtypes.ARN         `tfsdk:"model_arn"`
	ModelID                    types.String        `tfsdk:"model_id"`
	ModelName                  types.String        `tfsdk:"model_name"`
	OutputModalities           fwtypes.SetOfString `tfsdk:"output_modalities"`
	ProviderName               types.String        `tfsdk:"provider_name"`
	ResponseStreamingSupported types.Bool          `tfsdk:"response_streaming_supported"`
}
