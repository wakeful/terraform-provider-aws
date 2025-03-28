---
subcategory: "WAF Classic Regional"
layout: "aws"
page_title: "AWS: aws_wafregional_geo_match_set"
description: |-
  Provides a AWS WAF Regional Geo Match Set resource.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_wafregional_geo_match_set

Provides a WAF Regional Geo Match Set Resource

## Example Usage

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.wafregional_geo_match_set import WafregionalGeoMatchSet
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        WafregionalGeoMatchSet(self, "geo_match_set",
            geo_match_constraint=[WafregionalGeoMatchSetGeoMatchConstraint(
                type="Country",
                value="US"
            ), WafregionalGeoMatchSetGeoMatchConstraint(
                type="Country",
                value="CA"
            )
            ],
            name="geo_match_set"
        )
```

## Argument Reference

This resource supports the following arguments:

* `name` - (Required) The name or description of the Geo Match Set.
* `geo_match_constraint` - (Optional) The Geo Match Constraint objects which contain the country that you want AWS WAF to search for.

## Nested Blocks

### `geo_match_constraint`

#### Arguments

* `type` - (Required) The type of geographical area you want AWS WAF to search for. Currently Country is the only valid value.
* `value` - (Required) The country that you want AWS WAF to search for.
  This is the two-letter country code, e.g., `US`, `CA`, `RU`, `CN`, etc.
  See [docs](https://docs.aws.amazon.com/waf/latest/APIReference/API_GeoMatchConstraint.html) for all supported values.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `id` - The ID of the WAF Regional Geo Match Set.

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import WAF Regional Geo Match Set using the id. For example:

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.wafregional_geo_match_set import WafregionalGeoMatchSet
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        WafregionalGeoMatchSet.generate_config_for_import(self, "geoMatchSet", "a1b2c3d4-d5f6-7777-8888-9999aaaabbbbcccc")
```

Using `terraform import`, import WAF Regional Geo Match Set using the id. For example:

```console
% terraform import aws_wafregional_geo_match_set.geo_match_set a1b2c3d4-d5f6-7777-8888-9999aaaabbbbcccc
```

<!-- cache-key: cdktf-0.20.8 input-076708092d61d2d1e6c278cff3a1595b6f32efbba2e3ac89150e24303f965c61 -->