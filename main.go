package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/dynamodb"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := dynamodb.NewTable(ctx, "money-log-table", &dynamodb.TableArgs{
			Attributes: dynamodb.TableAttributeArray{
				&dynamodb.TableAttributeArgs{
					Name: pulumi.String("Title"),
					Type: pulumi.String("S"),
				},
				&dynamodb.TableAttributeArgs{
					Name: pulumi.String("JapaneseTitle"),
					Type: pulumi.String("S"),
				},
				&dynamodb.TableAttributeArgs{
					Name: pulumi.String("Price"),
					Type: pulumi.String("N"),
				},
				&dynamodb.TableAttributeArgs{
					Name: pulumi.String("PrimaryCategory"),
					Type: pulumi.String("S"),
				},
				&dynamodb.TableAttributeArgs{
					Name: pulumi.String("SecondaryCategory"),
					Type: pulumi.String("S"),
				},
				&dynamodb.TableAttributeArgs{
					Name: pulumi.String("Date"),
					Type: pulumi.String("S"),
				},
			},
			BillingMode: pulumi.String("PROVISIONED"),
			HashKey:      pulumi.String("Title"),
			RangeKey:     pulumi.String("Date"),
			ReadCapacity: pulumi.Int(1),
			WriteCapacity: pulumi.Int(1),
			Tags: pulumi.StringMap{
				"Environment": pulumi.String("production"),
				"Name":        pulumi.String("money-log-table"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}