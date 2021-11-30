package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/dynamodb"
	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/lambda"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		tableName := "money-log-table"
		_, err := dynamodb.NewTable(ctx, tableName, &dynamodb.TableArgs{
			Name: pulumi.String(tableName),
			Attributes: dynamodb.TableAttributeArray{
				&dynamodb.TableAttributeArgs{
					Name: pulumi.String("Title"),
					Type: pulumi.String("S"),
				},
				&dynamodb.TableAttributeArgs{
					Name: pulumi.String("Date"),
					Type: pulumi.String("S"),
				},
			},
			BillingMode:   pulumi.String("PROVISIONED"),
			HashKey:       pulumi.String("Title"),
			RangeKey:      pulumi.String("Date"),
			ReadCapacity:  pulumi.Int(1),
			WriteCapacity: pulumi.Int(1),
			Tags: pulumi.StringMap{
				"Environment": pulumi.String("production"),
				"Name":        pulumi.String("money-log-table"),
			},
		})
		if err != nil {
			return err
		}

		// Create the lambda using the args.
		_, err := lambda.NewFunction(ctx, "basicLambda", &lambda.FunctionArgs{
			Handler: pulumi.String("handler"),
			Role:    role.Arn,
			Runtime: pulumi.String("go1.x"),
			Code:    pulumi.NewFileArchive("./lambda/handler.zip"),
		},
		)
		if err != nil {
			return err
		}

		return nil
	})
}
