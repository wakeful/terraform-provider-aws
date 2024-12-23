---
subcategory: "Timestream Write"
layout: "aws"
page_title: "AWS: aws_timestreamwrite_database"
description: |-
  Terraform data source for managing an AWS Timestream Write Database.
---


<!-- Please do not edit this file, it is generated. -->
# Data Source: aws_timestreamwrite_database

Terraform data source for managing an AWS Timestream Write Database.

## Example Usage

### Basic Usage

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.data_aws_timestreamwrite_database import DataAwsTimestreamwriteDatabase
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        DataAwsTimestreamwriteDatabase(self, "test",
            name="database-example"
        )
```

## Argument Reference

The following arguments are required:

* `database_name` – (Required) The name of the Timestream database. Minimum length of 3. Maximum length of 256.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `arn` - The ARN that uniquely identifies this database.
* `created_time` - Creation time of database.
* `database_name` – (Required) The name of the Timestream database. Minimum length of 3. Maximum length of 256.
* `kms_key_id` - The ARN of the KMS key used to encrypt the data stored in the database.
* `last_updated_time` - Last time database was updated.
* `table_count` -  Total number of tables in the Timestream database.

<!-- cache-key: cdktf-0.20.8 input-09f00a2fd4d9fce1e8e04daeb035e6aa811517fe8fcd6ee09f5600b143473d33 -->