# CDK Demo App of How to Do S3 Notifications Triggering Lambda Functions

## I. Setup
1. Install AWS CDK: `npm install -g aws-cdk@latest`
2. Install Golang 1.16 if you have not

## II. How to run
1. Before running any command for CDK app, build the binary files first: `make build-all`
2. Run `cdk synth` to test if the CDK app generate the templates succesfully
3. For fully deployment stacks to AWS, run `cdk deploy --all`

