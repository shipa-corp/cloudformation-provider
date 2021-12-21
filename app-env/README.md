# Shipa::AppEnv::Item

Add and remove environment variables to applications.

Applications should already exist for this to work properly.

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
$ aws cloudformation list-type-versions --type "RESOURCE" --type-name "Shipa::AppEnv::Item"
```

Set default version
```bash
$ aws cloudformation set-type-default-version \
  --type "RESOURCE" \
  --type-name "Shipa::AppEnv::Item" \
  --version-id "00000003"
```

Deregister version
```bash
$ aws cloudformation deregister-type --type "RESOURCE" \
  --type-name "Shipa::AppEnv::Item" \
  --version-id "00000001"
```

Describe registration type
```bash
$ aws cloudformation describe-type-registration \
--registration-token a8ecadb8-3dcb-4e68-be9c-dbc85f375e57
```

Provision the resource in a CloudFormation stack (FAILING)
```bash
$ aws cloudformation create-stack --region eu-west-1 \
  --template-body "file://./examples/stack-template.yaml" \
  --stack-name "shipa-app-env"
```

Describe type
```bash
$ aws cloudformation describe-type --type-name "Shipa::AppEnv::Item" --type RESOURCE
```

Describe stack
```bash
$ aws cloudformation describe-stacks --region eu-west-1
```

List stack events
```bash
$ aws cloudformation describe-stack-events --stack-name shipa-app-env > events.log
```

Create secret
```bash
$ aws secretsmanager create-secret --name ShipaToken \
--description "Shipa token" \
--secret-string <token>

```
