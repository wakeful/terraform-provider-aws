---
subcategory: "ACM PCA (Certificate Manager Private Certificate Authority)"
layout: "aws"
page_title: "AWS: aws_acmpca_policy"
description: |-
  Attaches a resource based policy to an AWS Certificate Manager Private Certificate Authority (ACM PCA)
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_acmpca_policy

Attaches a resource based policy to a private CA.

## Example Usage

### Basic

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { Token, TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { AcmpcaPolicy } from "./.gen/providers/aws/acmpca-policy";
import { DataAwsIamPolicyDocument } from "./.gen/providers/aws/data-aws-iam-policy-document";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    const example = new DataAwsIamPolicyDocument(this, "example", {
      statement: [
        {
          actions: [
            "acm-pca:DescribeCertificateAuthority",
            "acm-pca:GetCertificate",
            "acm-pca:GetCertificateAuthorityCertificate",
            "acm-pca:ListPermissions",
            "acm-pca:ListTags",
          ],
          effect: "Allow",
          principals: [
            {
              identifiers: [Token.asString(current.accountId)],
              type: "AWS",
            },
          ],
          resources: [Token.asString(awsAcmpcaCertificateAuthorityExample.arn)],
          sid: "1",
        },
        {
          actions: ["acm-pca:IssueCertificate"],
          condition: [
            {
              test: "StringEquals",
              values: ["arn:aws:acm-pca:::template/EndEntityCertificate/V1"],
              variable: "acm-pca:TemplateArn",
            },
          ],
          effect: allow,
          principals: [
            {
              identifiers: [Token.asString(current.accountId)],
              type: "AWS",
            },
          ],
          resources: [Token.asString(awsAcmpcaCertificateAuthorityExample.arn)],
          sid: "2",
        },
      ],
    });
    const awsAcmpcaPolicyExample = new AcmpcaPolicy(this, "example_1", {
      policy: Token.asString(example.json),
      resourceArn: Token.asString(awsAcmpcaCertificateAuthorityExample.arn),
    });
    /*This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.*/
    awsAcmpcaPolicyExample.overrideLogicalId("example");
  }
}

```

## Argument Reference

This resource supports the following arguments:

* `resourceArn` - (Required) ARN of the private CA to associate with the policy.
* `policy` - (Required) JSON-formatted IAM policy to attach to the specified private CA resource.

## Attribute Reference

This resource exports no additional attributes.

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import `aws_acmpca_policy` using the `resourceArn` value. For example:

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { AcmpcaPolicy } from "./.gen/providers/aws/acmpca-policy";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    AcmpcaPolicy.generateConfigForImport(
      this,
      "example",
      "arn:aws:acm-pca:us-east-1:123456789012:certificate-authority/12345678-1234-1234-1234-123456789012"
    );
  }
}

```

Using `terraform import`, import `aws_acmpca_policy` using the `resourceArn` value. For example:

```console
% terraform import aws_acmpca_policy.example arn:aws:acm-pca:us-east-1:123456789012:certificate-authority/12345678-1234-1234-1234-123456789012
```

<!-- cache-key: cdktf-0.20.8 input-3897a945f24b08cf8623246e958727ed8480522a140850c9c8a5a835f4ee024d -->