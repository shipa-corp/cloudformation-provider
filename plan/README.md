# Shipa::Plan::Item

Define resource plan limits to be applied and enforced when applications are deployed.

# Development

First need to activate python virtual env
```bash
   source ../venv/bin/activate
```

Build
```bash
$ make build
```

Submit
```bash
$ cfn submit -v --region eu-west-1
```

# Helpful 

List versions
```bash
$ aws cloudformation list-type-versions --type "RESOURCE" --type-name "Shipa::Plan::Item"
```

Set default version
```bash
$ aws cloudformation set-type-default-version \
  --type "RESOURCE" \
  --type-name "Shipa::Plan::Item" \
  --version-id "00000007"
```

Deregister version
```bash
$ aws cloudformation deregister-type --type "RESOURCE" \
  --type-name "Shipa::Plan::Item" \
  --version-id "00000001"
```

Describe registration type
```bash
$ aws cloudformation describe-type-registration \
--registration-token a8ecadb8-3dcb-4e68-be9c-dbc85f375e57
```

Provision the resource in a CloudFormation stack
```bash
$ aws cloudformation create-stack --region eu-west-1 \
  --template-body "file://examples/stack-template.yaml" \
  --stack-name "shipa-plan"
```

Delete stack
```bash
$ aws cloudformation delete-stack --stack-name shipa-plan
```

Describe type
```bash
$ aws cloudformation describe-type --type-name "Shipa::Plan::Item" --type RESOURCE
```

Describe stack
```bash
$ aws cloudformation describe-stacks --region eu-west-1
```

List stack events
```bash
$ aws cloudformation describe-stack-events --stack-name shipa-plan > events.log
```

Create secret
```bash
$ aws secretsmanager create-secret --name ShipaToken \
--description "Shipa token" \
--secret-string <token>

```
