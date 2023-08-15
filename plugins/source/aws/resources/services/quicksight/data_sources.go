package quicksight

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/quicksight"
	"github.com/aws/aws-sdk-go-v2/service/quicksight/types"
	"github.com/aws/smithy-go"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/pkg/errors"
)

func DataSources() *schema.Table {
	tableName := "aws_quicksight_data_sources"
	return &schema.Table{
		Name:        tableName,
		Description: "https://docs.aws.amazon.com/quicksight/latest/APIReference/API_DataSource.html",
		Resolver:    fetchQuicksightDataSources,
		Transform: transformers.TransformWithStruct(&types.DataSource{},
			transformers.WithPrimaryKeys("Arn"),
			transformers.WithSkipFields("AlternateDataSourceParameters"),
		),
		Multiplex: client.ServiceAccountRegionMultiplexer(tableName, "quicksight"),
		Columns:   []schema.Column{client.DefaultAccountIDColumn(true), client.DefaultRegionColumn(true), tagsCol},
	}
}

func fetchQuicksightDataSources(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceQuicksight).Quicksight
	input := quicksight.ListDataSourcesInput{
		AwsAccountId: aws.String(cl.AccountID),
	}
	var ae smithy.APIError

	paginator := quicksight.NewListDataSourcesPaginator(svc, &input)
	for paginator.HasMorePages() {
		result, err := paginator.NextPage(ctx, func(options *quicksight.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			if errors.As(err, &ae) && ae.ErrorCode() == "UnsupportedUserEditionException" {
				return nil
			}

			return err
		}
		res <- result.DataSources
	}
	return nil
}
