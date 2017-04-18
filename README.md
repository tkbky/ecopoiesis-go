# Trf

[![Build Status](https://travis-ci.org/tkbky/trf.svg?branch=master)](https://travis-ci.org/tkbky/trf)

Export AWS resources to Terraform resources.

## Disclaimer

This is built for fun only, not for production use, use [this](https://github.com/dtan4/terraforming) instead.

## Prerequisites

1. AWS credentials

  Either,

  ```bash
  export AWS_ACCESS_KEY_ID=MY-ACCESS-KEY
  export AWS_SECRET_ACCESS_KEY=MY-SECRET-KEY
  export AWS_REGION=my-region
  ```

  Or,

  ```bash
  # ~/.aws/credentials

  [default]
  aws_access_key_id = MY-ACCESS-KEY
  aws_secret_access_key = MY-SECRET-KEY
  ```

## Usage

```bash
# Usage:
# trf [global options] command [command options] [arguments...]

trf aos # Get OpsWorks Stack resource
```

## Available Resources

| Command | Resource Name          |
|---------|------------------------|
| aos     | OpsWorks Stack         |
| aocl    | OpsWorks Custom Layer  |
| aoa     | OpsWorks Application   |

## Contributing

The normal workflow of adding a new resource

1. Refer to [AWS Cli reference](http://docs.aws.amazon.com/cli/latest/reference/) for available commands & their response. e.g. for OpsWorks, http://docs.aws.amazon.com/cli/latest/reference/opsworks
2. Refer to [AWS SDK for Go API documentation](http://docs.aws.amazon.com/sdk-for-go/api/service/) for how to issue a command via API.
3. Refer to [Terraform AWS providers documentation](https://www.terraform.io/docs/providers/aws/) when defining a struct for a particular resource.

## License

[MIT License](https://opensource.org/licenses/MIT)
