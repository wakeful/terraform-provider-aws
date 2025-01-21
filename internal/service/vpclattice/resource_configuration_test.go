// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vpclattice_test

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/YakDriver/regexache"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/vpclattice"
	sdkacctest "github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/create"
	tfvpclattice "github.com/hashicorp/terraform-provider-aws/internal/service/vpclattice"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
	"github.com/hashicorp/terraform-provider-aws/names"
)

func TestAccVPCLatticeResourceConfiguration_basic(t *testing.T) {
	ctx := acctest.Context(t)
	if testing.Short() {
		t.Skip("skipping long-running test in short mode")
	}

	var resourceconfiguration vpclattice.GetResourceConfigurationOutput
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_vpclattice_resource_configuration.test"
	resourceGatewayName := "aws_vpclattice_resource_gateway.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acctest.PreCheck(ctx, t)
			acctest.PreCheckPartitionHasService(t, names.VPCLatticeEndpointID)
			testAccPreCheck(ctx, t)
		},
		ErrorCheck:               acctest.ErrorCheck(t, names.VPCLatticeServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckResourceConfigurationDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceConfigurationConfig_basic(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckResourceConfigurationExists(ctx, resourceName, &resourceconfiguration),
					resource.TestCheckResourceAttr(resourceName, names.AttrName, rName),
					resource.TestCheckResourceAttrPair(resourceName, "resource_gateway_identifier", resourceGatewayName, names.AttrID),
					resource.TestCheckResourceAttr(resourceName, "port_ranges.0", "80"),
					resource.TestCheckResourceAttr(resourceName, "resource_configuration_definition.0.dns_resource.0.domain_name", "example.com"),
					resource.TestCheckResourceAttr(resourceName, "resource_configuration_definition.0.dns_resource.0.ip_address_type", "IPV4"),
					acctest.MatchResourceAttrRegionalARN(ctx, resourceName, names.AttrARN, "vpc-lattice", regexache.MustCompile(`resourceconfiguration/+.`)),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccVPCLatticeResourceConfiguration_update(t *testing.T) {
	ctx := acctest.Context(t)
	if testing.Short() {
		t.Skip("skipping long-running test in short mode")
	}

	var v1, v2 vpclattice.GetResourceConfigurationOutput
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_vpclattice_resource_configuration.test"
	resourceGatewayName := "aws_vpclattice_resource_gateway.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acctest.PreCheck(ctx, t)
			acctest.PreCheckPartitionHasService(t, names.VPCLatticeEndpointID)
			testAccPreCheck(ctx, t)
		},
		ErrorCheck:               acctest.ErrorCheck(t, names.VPCLatticeServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckResourceConfigurationDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceConfigurationConfig_basic(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckResourceConfigurationExists(ctx, resourceName, &v1),
					resource.TestCheckResourceAttr(resourceName, names.AttrName, rName),
					resource.TestCheckResourceAttr(resourceName, "allow_association_to_shareable_service_network", acctest.CtTrue),
					resource.TestCheckResourceAttrPair(resourceName, "resource_gateway_identifier", resourceGatewayName, names.AttrID),
					resource.TestCheckResourceAttr(resourceName, "port_ranges.0", "80"),
					resource.TestCheckResourceAttr(resourceName, "resource_configuration_definition.0.dns_resource.0.domain_name", "example.com"),
					resource.TestCheckResourceAttr(resourceName, "resource_configuration_definition.0.dns_resource.0.ip_address_type", "IPV4"),
					acctest.MatchResourceAttrRegionalARN(ctx, resourceName, names.AttrARN, "vpc-lattice", regexache.MustCompile(`resourceconfiguration/+.`)),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccResourceConfigurationConfig_update(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckResourceConfigurationExists(ctx, resourceName, &v2),
					testAccCheckResourceConfigurationNotRecreated(&v1, &v2),
					resource.TestCheckResourceAttr(resourceName, names.AttrName, rName),
					resource.TestCheckResourceAttr(resourceName, "allow_association_to_shareable_service_network", acctest.CtTrue),
					resource.TestCheckResourceAttrPair(resourceName, "resource_gateway_identifier", resourceGatewayName, names.AttrID),
					resource.TestCheckResourceAttr(resourceName, "port_ranges.0", "80"),
					resource.TestCheckResourceAttr(resourceName, "port_ranges.1", "8080"),
					resource.TestCheckResourceAttr(resourceName, "resource_configuration_definition.0.dns_resource.0.domain_name", "example.com"),
					resource.TestCheckResourceAttr(resourceName, "resource_configuration_definition.0.dns_resource.0.ip_address_type", "IPV4"),
					acctest.MatchResourceAttrRegionalARN(ctx, resourceName, names.AttrARN, "vpc-lattice", regexache.MustCompile(`resourceconfiguration/+.`)),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccVPCLatticeResourceConfiguration_disappears(t *testing.T) {
	ctx := acctest.Context(t)
	if testing.Short() {
		t.Skip("skipping long-running test in short mode")
	}

	var resourceconfiguration vpclattice.GetResourceConfigurationOutput
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_vpclattice_resource_configuration.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acctest.PreCheck(ctx, t)
			acctest.PreCheckPartitionHasService(t, names.VPCLatticeEndpointID)
			testAccPreCheck(ctx, t)
		},
		ErrorCheck:               acctest.ErrorCheck(t, names.VPCLatticeServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckResourceConfigurationDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceConfigurationConfig_basic(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckResourceConfigurationExists(ctx, resourceName, &resourceconfiguration),
					acctest.CheckFrameworkResourceDisappears(ctx, acctest.Provider, tfvpclattice.ResourceConfiguration, resourceName),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func testAccCheckResourceConfigurationDestroy(ctx context.Context) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		conn := acctest.Provider.Meta().(*conns.AWSClient).VPCLatticeClient(ctx)

		for _, rs := range s.RootModule().Resources {
			if rs.Type != "aws_vpclattice_resource_configuration" {
				continue
			}

			_, err := tfvpclattice.FindResourceConfigurationByID(ctx, conn, rs.Primary.ID)
			if tfresource.NotFound(err) {
				return nil
			}
			if err != nil {
				return create.Error(names.VPCLattice, create.ErrActionCheckingDestroyed, tfvpclattice.ResNameResourceConfiguration, rs.Primary.ID, err)
			}

			return create.Error(names.VPCLattice, create.ErrActionCheckingDestroyed, tfvpclattice.ResNameResourceConfiguration, rs.Primary.ID, errors.New("not destroyed"))
		}

		return nil
	}
}

func testAccCheckResourceConfigurationExists(ctx context.Context, name string, resourceconfiguration *vpclattice.GetResourceConfigurationOutput) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return create.Error(names.VPCLattice, create.ErrActionCheckingExistence, tfvpclattice.ResNameResourceConfiguration, name, errors.New("not found"))
		}

		if rs.Primary.ID == "" {
			return create.Error(names.VPCLattice, create.ErrActionCheckingExistence, tfvpclattice.ResNameResourceConfiguration, name, errors.New("not set"))
		}

		conn := acctest.Provider.Meta().(*conns.AWSClient).VPCLatticeClient(ctx)

		resp, err := tfvpclattice.FindResourceConfigurationByID(ctx, conn, rs.Primary.ID)
		if err != nil {
			return create.Error(names.VPCLattice, create.ErrActionCheckingExistence, tfvpclattice.ResNameResourceConfiguration, rs.Primary.ID, err)
		}

		*resourceconfiguration = *resp

		return nil
	}
}

func testAccCheckResourceConfigurationNotRecreated(before, after *vpclattice.GetResourceConfigurationOutput) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if before, after := aws.ToString(before.Id), aws.ToString(after.Id); before != after {
			return create.Error(names.VPCLattice, create.ErrActionCheckingNotRecreated, tfvpclattice.ResNameResourceConfiguration, before, errors.New("recreated"))
		}

		return nil
	}
}

func testAccResourceConfigurationConfig_basic(rName string) string {
	return acctest.ConfigCompose(testAccResourceGatewayConfig_basic(rName),
		fmt.Sprintf(`
resource "aws_vpclattice_resource_configuration" "test" {
  name = %[1]q

  resource_gateway_identifier = aws_vpclattice_resource_gateway.test.id

  port_ranges = ["80"]

  resource_configuration_definition {
    dns_resource {
      domain_name     = "example.com"
      ip_address_type = "IPV4"
    }
  }
}

`, rName))
}

func testAccResourceConfigurationConfig_update(rName string) string {
	return acctest.ConfigCompose(testAccResourceGatewayConfig_basic(rName),
		fmt.Sprintf(`
resource "aws_vpclattice_resource_configuration" "test" {
  name = %[1]q

  resource_gateway_identifier = aws_vpclattice_resource_gateway.test.id

  port_ranges = ["80", "8080"]

  resource_configuration_definition {
    dns_resource {
      domain_name     = "example.com"
      ip_address_type = "IPV4"
    }
  }
}

`, rName))
}
