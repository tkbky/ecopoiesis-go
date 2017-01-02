# Trf

Export AWS resources to Terraform resources.

## Pre-requisite

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
trf aocl
```
