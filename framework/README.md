# Shipa::Framework::Item

Create and manage policy frameworks to be enforced to applications across multiple clusters

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
$ aws cloudformation list-type-versions --type "RESOURCE" --type-name "Shipa::Framework::Item"
```

Set default version
```bash
$ aws cloudformation set-type-default-version \
  --type "RESOURCE" \
  --type-name "Shipa::Framework::Item" \
  --version-id "00000007"
```

Deregister version
```bash
$ aws cloudformation deregister-type --type "RESOURCE" \
  --type-name "Shipa::Framework::Item" \
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
  --template-body "file://stack.json" \
  --stack-name "shipa-3"

aws cloudformation create-stack --region eu-west-1 \
  --template-body "file://stack-template.yaml" \
  --stack-name "shipa-4"
```

Describe type
```bash
$ aws cloudformation describe-type --type-name "Shipa::Framework::Item" --type RESOURCE
```

Describe stack
```bash
$ aws cloudformation describe-stacks --region eu-west-1
```

List stack events
```bash
$ aws cloudformation describe-stack-events --stack-name shipa > events.log
```

Create secret
```bash
$ aws secretsmanager create-secret --name ShipaToken \
--description "Shipa token" \
--secret-string <token>

```
