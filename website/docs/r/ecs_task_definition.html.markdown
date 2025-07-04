---
subcategory: "ECS (Elastic Container)"
layout: "aws"
page_title: "AWS: aws_ecs_task_definition"
description: |-
  Manages a revision of an ECS task definition.
---

# Resource: aws_ecs_task_definition

Manages a revision of an ECS task definition to be used in `aws_ecs_service`.

## Example Usage

### Basic Example

```terraform
resource "aws_ecs_task_definition" "service" {
  family = "service"
  container_definitions = jsonencode([
    {
      name      = "first"
      image     = "service-first"
      cpu       = 10
      memory    = 512
      essential = true
      portMappings = [
        {
          containerPort = 80
          hostPort      = 80
        }
      ]
    },
    {
      name      = "second"
      image     = "service-second"
      cpu       = 10
      memory    = 256
      essential = true
      portMappings = [
        {
          containerPort = 443
          hostPort      = 443
        }
      ]
    }
  ])

  volume {
    name      = "service-storage"
    host_path = "/ecs/service-storage"
  }

  placement_constraints {
    type       = "memberOf"
    expression = "attribute:ecs.availability-zone in [us-west-2a, us-west-2b]"
  }
}
```

### With AppMesh Proxy

```terraform
resource "aws_ecs_task_definition" "service" {
  family                = "service"
  container_definitions = file("task-definitions/service.json")

  proxy_configuration {
    type           = "APPMESH"
    container_name = "applicationContainerName"
    properties = {
      AppPorts         = "8080"
      EgressIgnoredIPs = "169.254.170.2,169.254.169.254"
      IgnoredUID       = "1337"
      ProxyEgressPort  = 15001
      ProxyIngressPort = 15000
    }
  }
}
```

### Example Using `docker_volume_configuration`

```terraform
resource "aws_ecs_task_definition" "service" {
  family                = "service"
  container_definitions = file("task-definitions/service.json")

  volume {
    name = "service-storage"

    docker_volume_configuration {
      scope         = "shared"
      autoprovision = true
      driver        = "local"

      driver_opts = {
        "type"   = "nfs"
        "device" = "${aws_efs_file_system.fs.dns_name}:/"
        "o"      = "addr=${aws_efs_file_system.fs.dns_name},rsize=1048576,wsize=1048576,hard,timeo=600,retrans=2,noresvport"
      }
    }
  }
}
```

### Example Using `efs_volume_configuration`

```terraform
resource "aws_ecs_task_definition" "service" {
  family                = "service"
  container_definitions = file("task-definitions/service.json")

  volume {
    name = "service-storage"

    efs_volume_configuration {
      file_system_id          = aws_efs_file_system.fs.id
      root_directory          = "/opt/data"
      transit_encryption      = "ENABLED"
      transit_encryption_port = 2999
      authorization_config {
        access_point_id = aws_efs_access_point.test.id
        iam             = "ENABLED"
      }
    }
  }
}
```

### Example Using `fsx_windows_file_server_volume_configuration`

```terraform
resource "aws_ecs_task_definition" "service" {
  family                = "service"
  container_definitions = file("task-definitions/service.json")

  volume {
    name = "service-storage"

    fsx_windows_file_server_volume_configuration {
      file_system_id = aws_fsx_windows_file_system.test.id
      root_directory = "\\data"

      authorization_config {
        credentials_parameter = aws_secretsmanager_secret_version.test.arn
        domain                = aws_directory_service_directory.test.name
      }
    }
  }
}

resource "aws_secretsmanager_secret_version" "test" {
  secret_id     = aws_secretsmanager_secret.test.id
  secret_string = jsonencode({ username : "admin", password : aws_directory_service_directory.test.password })
}
```

### Example Using `container_definitions`

```terraform
resource "aws_ecs_task_definition" "test" {
  family                = "test"
  container_definitions = <<TASK_DEFINITION
[
  {
    "cpu": 10,
    "command": ["sleep", "10"],
    "entryPoint": ["/"],
    "environment": [
      {"name": "VARNAME", "value": "VARVAL"}
    ],
    "essential": true,
    "image": "jenkins",
    "memory": 128,
    "name": "jenkins",
    "portMappings": [
      {
        "containerPort": 80,
        "hostPort": 8080
      }
    ]
  }
]
TASK_DEFINITION
}
```

### Example Using `runtime_platform` and `fargate`

```terraform
resource "aws_ecs_task_definition" "test" {
  family                   = "test"
  requires_compatibilities = ["FARGATE"]
  network_mode             = "awsvpc"
  cpu                      = 1024
  memory                   = 2048
  container_definitions    = <<TASK_DEFINITION
[
  {
    "name": "iis",
    "image": "mcr.microsoft.com/windows/servercore/iis",
    "cpu": 1024,
    "memory": 2048,
    "essential": true
  }
]
TASK_DEFINITION

  runtime_platform {
    operating_system_family = "WINDOWS_SERVER_2019_CORE"
    cpu_architecture        = "X86_64"
  }
}
```

## Argument Reference

The following arguments are required:

* `container_definitions` - (Required) A list of valid [container definitions](http://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_ContainerDefinition.html) provided as a single valid JSON document. Please note that you should only provide values that are part of the container definition document. For a detailed description of what parameters are available, see the [Task Definition Parameters](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task_definition_parameters.html) section from the official [Developer Guide](https://docs.aws.amazon.com/AmazonECS/latest/developerguide).
* `family` - (Required) A unique name for your task definition.

The following arguments are optional:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `cpu` - (Optional) Number of cpu units used by the task. If the `requires_compatibilities` is `FARGATE` this field is required.
* `enable_fault_injection` - (Optional) Enables fault injection and allows for fault injection requests to be accepted from the task's containers. Default is `false`.
* `execution_role_arn` - (Optional) ARN of the task execution role that the Amazon ECS container agent and the Docker daemon can assume.
* `ipc_mode` - (Optional) IPC resource namespace to be used for the containers in the task The valid values are `host`, `task`, and `none`.
* `memory` - (Optional) Amount (in MiB) of memory used by the task. If the `requires_compatibilities` is `FARGATE` this field is required.
* `network_mode` - (Optional) Docker networking mode to use for the containers in the task. Valid values are `none`, `bridge`, `awsvpc`, and `host`.
* `runtime_platform` - (Optional) Configuration block for [runtime_platform](#runtime_platform) that containers in your task may use.
* `pid_mode` - (Optional) Process namespace to use for the containers in the task. The valid values are `host` and `task`.
* `placement_constraints` - (Optional) Configuration block for rules that are taken into consideration during task placement. Maximum number of `placement_constraints` is `10`. [Detailed below](#placement_constraints).
* `proxy_configuration` - (Optional) Configuration block for the App Mesh proxy. [Detailed below.](#proxy_configuration)
* `ephemeral_storage` - (Optional)  The amount of ephemeral storage to allocate for the task. This parameter is used to expand the total amount of ephemeral storage available, beyond the default amount, for tasks hosted on AWS Fargate. See [Ephemeral Storage](#ephemeral_storage).
* `requires_compatibilities` - (Optional) Set of launch types required by the task. The valid values are `EC2` and `FARGATE`.
* `skip_destroy` - (Optional) Whether to retain the old revision when the resource is destroyed or replacement is necessary. Default is `false`.
* `tags` - (Optional) Key-value map of resource tags. If configured with a provider [`default_tags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
* `task_role_arn` - (Optional) ARN of IAM role that allows your Amazon ECS container task to make calls to other AWS services.
* `track_latest` - (Optional) Whether should track latest `ACTIVE` task definition on AWS or the one created with the resource stored in state. Default is `false`. Useful in the event the task definition is modified outside of this resource.
* `volume` - (Optional) Configuration block for [volumes](#volume) that containers in your task may use. Detailed below.

~> **NOTE:** Proper escaping is required for JSON field values containing quotes (`"`) such as `environment` values. If directly setting the JSON, they should be escaped as `\"` in the JSON,  e.g., `"value": "I \"love\" escaped quotes"`. If using a Terraform variable value, they should be escaped as `\\\"` in the variable, e.g., `value = "I \\\"love\\\" escaped quotes"` in the variable and `"value": "${var.myvariable}"` in the JSON.

~> **Note:** Fault injection only works with tasks using the `awsvpc` or `host` network modes. Fault injection isn't available on Windows.

### volume

* `docker_volume_configuration` - (Optional) Configuration block to configure a [docker volume](#docker_volume_configuration). Detailed below.
* `efs_volume_configuration` - (Optional) Configuration block for an [EFS volume](#efs_volume_configuration). Detailed below.
* `fsx_windows_file_server_volume_configuration` - (Optional) Configuration block for an [FSX Windows File Server volume](#fsx_windows_file_server_volume_configuration). Detailed below.
* `host_path` - (Optional) Path on the host container instance that is presented to the container. If not set, ECS will create a nonpersistent data volume that starts empty and is deleted after the task has finished.
* `configure_at_launch` - (Optional) Whether the volume should be configured at launch time. This is used to create Amazon EBS volumes for standalone tasks or tasks created as part of a service. Each task definition revision may only have one volume configured at launch in the volume configuration.
* `name` - (Required) Name of the volume. This name is referenced in the `sourceVolume`
parameter of container definition in the `mountPoints` section.

### docker_volume_configuration

For more information, see [Specifying a Docker volume in your Task Definition Developer Guide](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/docker-volumes.html#specify-volume-config)

* `autoprovision` - (Optional) If this value is `true`, the Docker volume is created if it does not already exist. *Note*: This field is only used if the scope is `shared`.
* `driver_opts` - (Optional) Map of Docker driver specific options.
* `driver` - (Optional) Docker volume driver to use. The driver value must match the driver name provided by Docker because it is used for task placement.
* `labels` - (Optional) Map of custom metadata to add to your Docker volume.
* `scope` - (Optional) Scope for the Docker volume, which determines its lifecycle, either `task` or `shared`.  Docker volumes that are scoped to a `task` are automatically provisioned when the task starts and destroyed when the task stops. Docker volumes that are scoped as `shared` persist after the task stops.

### efs_volume_configuration

For more information, see [Specifying an EFS volume in your Task Definition Developer Guide](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/efs-volumes.html#specify-efs-config)

* `file_system_id` - (Required) ID of the EFS File System.
* `root_directory` - (Optional) Directory within the Amazon EFS file system to mount as the root directory inside the host. If this parameter is omitted, the root of the Amazon EFS volume will be used. Specifying / will have the same effect as omitting this parameter. This argument is ignored when using `authorization_config`.
* `transit_encryption` - (Optional) Whether or not to enable encryption for Amazon EFS data in transit between the Amazon ECS host and the Amazon EFS server. Transit encryption must be enabled if Amazon EFS IAM authorization is used. Valid values: `ENABLED`, `DISABLED`. If this parameter is omitted, the default value of `DISABLED` is used.
* `transit_encryption_port` - (Optional) Port to use for transit encryption. If you do not specify a transit encryption port, it will use the port selection strategy that the Amazon EFS mount helper uses.
* `authorization_config` - (Optional) Configuration block for [authorization](#authorization_config) for the Amazon EFS file system. Detailed below.

### runtime_platform

* `operating_system_family` - (Optional) If the `requires_compatibilities` is `FARGATE` this field is required; must be set to a valid option from the [operating system family in the runtime platform](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task_definition_parameters.html#runtime-platform) setting
* `cpu_architecture` - (Optional) Must be set to either `X86_64` or `ARM64`; see [cpu architecture](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task_definition_parameters.html#runtime-platform)

#### authorization_config

* `access_point_id` - (Optional) Access point ID to use. If an access point is specified, the root directory value will be relative to the directory set for the access point. If specified, transit encryption must be enabled in the EFSVolumeConfiguration.
* `iam` - (Optional) Whether or not to use the Amazon ECS task IAM role defined in a task definition when mounting the Amazon EFS file system. If enabled, transit encryption must be enabled in the EFSVolumeConfiguration. Valid values: `ENABLED`, `DISABLED`. If this parameter is omitted, the default value of `DISABLED` is used.

### fsx_windows_file_server_volume_configuration

For more information, see [Specifying an FSX Windows File Server volume in your Task Definition Developer Guide](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/tutorial-wfsx-volumes.html)

* `file_system_id` - (Required) The Amazon FSx for Windows File Server file system ID to use.
* `root_directory` - (Required) The directory within the Amazon FSx for Windows File Server file system to mount as the root directory inside the host.
* `authorization_config` - (Required) Configuration block for [authorization](#authorization_config) for the Amazon FSx for Windows File Server file system detailed below.

#### authorization_config

* `credentials_parameter` - (Required) The authorization credential option to use. The authorization credential options can be provided using either the Amazon Resource Name (ARN) of an AWS Secrets Manager secret or AWS Systems Manager Parameter Store parameter. The ARNs refer to the stored credentials.
* `domain` - (Required) A fully qualified domain name hosted by an AWS Directory Service Managed Microsoft AD (Active Directory) or self-hosted AD on Amazon EC2.

### placement_constraints

* `expression` -  (Optional) Cluster Query Language expression to apply to the constraint. For more information, see [Cluster Query Language in the Amazon EC2 Container Service Developer Guide](http://docs.aws.amazon.com/AmazonECS/latest/developerguide/cluster-query-language.html).
* `type` - (Required) Type of constraint. Use `memberOf` to restrict selection to a group of valid candidates. Note that `distinctInstance` is not supported in task definitions.

### proxy_configuration

* `container_name` - (Required) Name of the container that will serve as the App Mesh proxy.
* `properties` - (Required) Set of network configuration parameters to provide the Container Network Interface (CNI) plugin, specified a key-value mapping.
* `type` - (Optional) Proxy type. The default value is `APPMESH`. The only supported value is `APPMESH`.

### ephemeral_storage

* `size_in_gib` - (Required) The total amount, in GiB, of ephemeral storage to set for the task. The minimum supported value is `21` GiB and the maximum supported value is `200` GiB.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `arn` - Full ARN of the Task Definition (including both `family` and `revision`).
* `arn_without_revision` - ARN of the Task Definition with the trailing `revision` removed. This may be useful for situations where the latest task definition is always desired. If a revision isn't specified, the latest ACTIVE revision is used. See the [AWS documentation](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_StartTask.html#ECS-StartTask-request-taskDefinition) for details.
* `revision` - Revision of the task in a particular family.
* `tags_all` - Map of tags assigned to the resource, including those inherited from the provider [`default_tags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block).

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import ECS Task Definitions using their ARNs. For example:

```terraform
import {
  to = aws_ecs_task_definition.example
  id = "arn:aws:ecs:us-east-1:012345678910:task-definition/mytaskfamily:123"
}
```

Using `terraform import`, import ECS Task Definitions using their ARNs. For example:

```console
% terraform import aws_ecs_task_definition.example arn:aws:ecs:us-east-1:012345678910:task-definition/mytaskfamily:123
```
