# Week 2 Scripting like a Developer

The code found in this repository is to help you learn how to script

## WIP

Its a work in progress (WIP), and will be updated when ready

## Powershell Code
THe code is for anyone that wants to create a Resource Group

## How to use the Powershell Code


## Python Code


## How to use the Python code

```
function New-ResourceGroup {
    [cmdletbinding(SupportsShouldProcess)]
    param (
        [parameter(Mandatory)]
        [string]$rgName,

        [parameter(Mandatory)]
        [string]$location

    )

    $params =@{
        'Name' = $rgName
        'Location' = $location
    }
    if ($PSCmdlet.ShouldProcess('location')){
        New-AzResourcegroup @params
    }
}
```
```
import sys
import boto3


try:
    def main():
        create_s3bucket(bucket_name)

except Exception as e:
    print(e)

def create_s3bucket(bucket_name):
    s3_bucket=boto3.client(
        's3',
    )

    bucket = s3_bucket.create_bucket(
        Bucket=bucket_name,
        ACL='private',
    )
    print(bucket)

bucket_name = sys.argv[1]

if __name__ == '__main__':
    main()

```

## Testing
There are tests available on the `powershell` an `python` directories