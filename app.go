package main

import (
	"os"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambdaeventsources"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3assets"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3notifications"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type StorageStackProps struct {
	awscdk.StackProps
}

type TriggerFunc2StackProps struct {
	awscdk.StackProps
}

func NewStorageStack(scope constructs.Construct, id string, props *StorageStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	s3Bucket := awss3.NewBucket(stack,
		jsii.String("ExampleBucket"), &awss3.BucketProps{
			AccessControl:     awss3.BucketAccessControl_BUCKET_OWNER_FULL_CONTROL,
			BlockPublicAccess: awss3.BlockPublicAccess_BLOCK_ALL(),
			BucketName:        jsii.String("cdk-example-bucket"),
			RemovalPolicy:     awscdk.RemovalPolicy_DESTROY,
			AutoDeleteObjects: jsii.Bool(true),
		})

	triggerLambdaFunc1 := awslambda.NewFunction(stack, jsii.String("TriggerFunc1"), &awslambda.FunctionProps{
		FunctionName: jsii.String("trigger-func-1"),
		MemorySize:   jsii.Number(128),
		Timeout:      awscdk.Duration_Seconds(jsii.Number(10)),
		VpcSubnets:   &awsec2.SubnetSelection{},
		Code:         awslambda.AssetCode_FromAsset(jsii.String("build/trigger-func1/"), &awss3assets.AssetOptions{}),
		Handler:      jsii.String("main"),
		Runtime:      awslambda.Runtime_GO_1_X(),
	})

	triggerLambdaFunc1.AddEventSource(awslambdaeventsources.NewS3EventSource(s3Bucket, &awslambdaeventsources.S3EventSourceProps{
		Events: &[]awss3.EventType{
			awss3.EventType_OBJECT_CREATED,
		},
	}))

	s3Bucket.GrantReadWrite(triggerLambdaFunc1, nil)

	return stack
}

func NewTriggerFunc2Stack(scope constructs.Construct, id string, props *TriggerFunc2StackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	triggerLambdaFunc2 := awslambda.NewFunction(stack, jsii.String("TriggerFunc2"), &awslambda.FunctionProps{
		FunctionName: jsii.String("trigger-func-2"),
		MemorySize:   jsii.Number(128),
		Timeout:      awscdk.Duration_Seconds(jsii.Number(10)),
		VpcSubnets:   &awsec2.SubnetSelection{},
		Code:         awslambda.AssetCode_FromAsset(jsii.String("build/trigger-func2/"), &awss3assets.AssetOptions{}),
		Handler:      jsii.String("main"),
		Runtime:      awslambda.Runtime_GO_1_X(),
	})

	s3Bucket := awss3.Bucket_FromBucketName(stack, jsii.String("ExampleBucket"), jsii.String("cdk-example-bucket"))
	s3Bucket.AddEventNotification(awss3.EventType_OBJECT_CREATED, awss3notifications.NewLambdaDestination(triggerLambdaFunc2))

	s3Bucket.GrantReadWrite(triggerLambdaFunc2, nil)

	return stack
}

func main() {
	app := awscdk.NewApp(nil)

	NewStorageStack(app, "StorageStack", &StorageStackProps{
		awscdk.StackProps{
			Env: env(),
		},
	})

	NewTriggerFunc2Stack(app, "TriggerFunc2Stack", &TriggerFunc2StackProps{
		awscdk.StackProps{
			Env: env(),
		},
	})

	app.Synth(nil)
}

func env() *awscdk.Environment {
	var account = os.Getenv("CDK_DEFAULT_ACCOUNT")
	var region = os.Getenv("CDK_DEFAULT_REGION")
	if region == "" {
		region = "us-west-2"
	}

	return &awscdk.Environment{
		Account: jsii.String(account),
		Region:  jsii.String(region),
	}
}
