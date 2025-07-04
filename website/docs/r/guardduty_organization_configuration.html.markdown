---
subcategory: "GuardDuty"
layout: "aws"
page_title: "AWS: aws_guardduty_organization_configuration"
description: |-
  Manages the GuardDuty Organization Configuration
---

# Resource: aws_guardduty_organization_configuration

Manages the GuardDuty Organization Configuration in the current AWS Region. The AWS account utilizing this resource must have been assigned as a delegated Organization administrator account, e.g., via the [`aws_guardduty_organization_admin_account` resource](/docs/providers/aws/r/guardduty_organization_admin_account.html). More information about Organizations support in GuardDuty can be found in the [GuardDuty User Guide](https://docs.aws.amazon.com/guardduty/latest/ug/guardduty_organizations.html).

~> **NOTE:** This is an advanced Terraform resource. Terraform will automatically assume management of the GuardDuty Organization Configuration without import and perform no actions on removal from the Terraform configuration.

## Example Usage

```terraform
resource "aws_guardduty_detector" "example" {
  enable = true
}

resource "aws_guardduty_organization_configuration" "example" {
  auto_enable_organization_members = "ALL"

  detector_id = aws_guardduty_detector.example.id

  datasources {
    s3_logs {
      auto_enable = true
    }
    kubernetes {
      audit_logs {
        enable = true
      }
    }
    malware_protection {
      scan_ec2_instance_with_findings {
        ebs_volumes {
          auto_enable = true
        }
      }
    }
  }
}
```

## Argument Reference

This resource supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `auto_enable_organization_members` - (Required) Indicates the auto-enablement configuration of GuardDuty for the member accounts in the organization.
  Valid values are `ALL`, `NEW`, `NONE`.
* `detector_id` - (Required) The detector ID of the GuardDuty account.
* `datasources` - (Optional) Configuration for the collected datasources. [Deprecated](https://docs.aws.amazon.com/guardduty/latest/ug/guardduty-feature-object-api-changes-march2023.html) in favor of [`aws_guardduty_organization_configuration_feature` resources](guardduty_organization_configuration_feature.html).

~> **NOTE:** One of `auto_enable` or `auto_enable_organization_members` must be specified.

`datasources` supports the following:

* `s3_logs` - (Optional) Enable S3 Protection automatically for new member accounts.
* `kubernetes` - (Optional) Enable Kubernetes Audit Logs Monitoring automatically for new member accounts.
* `malware_protection` - (Optional) Enable Malware Protection automatically for new member accounts.

### S3 Logs

`s3_logs` block supports the following:

* `auto_enable` - (Optional) Set to `true` if you want S3 data event logs to be automatically enabled for new members of the organization. Default: `false`

### Kubernetes

`kubernetes` block supports the following:

* `audit_logs` - (Required) Enable Kubernetes Audit Logs Monitoring automatically for new member accounts. [Kubernetes protection](https://docs.aws.amazon.com/guardduty/latest/ug/kubernetes-protection.html).
  See [Kubernetes Audit Logs](#kubernetes-audit-logs) below for more details.

#### Kubernetes Audit Logs

The `audit_logs` block supports the following:

* `enable` - (Required) If true, enables Kubernetes audit logs as a data source for [Kubernetes protection](https://docs.aws.amazon.com/guardduty/latest/ug/kubernetes-protection.html).
  Defaults to `true`.

### Malware Protection

`malware_protection` block supports the following:

* `scan_ec2_instance_with_findings` - (Required) Configure whether [Malware Protection](https://docs.aws.amazon.com/guardduty/latest/ug/malware-protection.html) for EC2 instances with findings should be auto-enabled for new members joining the organization.
   See [Scan EC2 instance with findings](#scan-ec2-instance-with-findings) below for more details.

#### Scan EC2 instance with findings

The `scan_ec2_instance_with_findings` block supports the following:

* `ebs_volumes` - (Required) Configure whether scanning EBS volumes should be auto-enabled for new members joining the organization
  See [EBS volumes](#ebs-volumes) below for more details.

#### EBS volumes

The `ebs_volumes` block supports the following:

* `auto_enable` - (Required) If true, enables [Malware Protection](https://docs.aws.amazon.com/guardduty/latest/ug/malware-protection.html) for all new accounts joining the organization.
  Defaults to `true`.

## Attribute Reference

This resource exports no additional attributes.

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import GuardDuty Organization Configurations using the GuardDuty Detector ID. For example:

```terraform
import {
  to = aws_guardduty_organization_configuration.example
  id = "00b00fd5aecc0ab60a708659477e9617"
}
```

Using `terraform import`, import GuardDuty Organization Configurations using the GuardDuty Detector ID. For example:

```console
% terraform import aws_guardduty_organization_configuration.example 00b00fd5aecc0ab60a708659477e9617
```
