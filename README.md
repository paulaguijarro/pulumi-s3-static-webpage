# Pulumi S3 static-webpage

Host a static website on AWS S3 with Pulumi and Golang

## How to use it

### Requirements

* Install Pulumi: `brew install pulumi/tap/pulumi`
* [Install](https://go.dev/doc/install) Go
* [Install](https://docs.aws.amazon.com/cli/latest/userguide/cli-chap-configure.html) AWS CLI and configure it to access your account

### Configuration

* Create a new [stack](https://www.pulumi.com/docs/intro/concepts/stack/): `pulumi stack init <stack-name>`
* Configure the stack: `pulumi config set aws:region <aws-region>`

### Deployment

Deploy the stack: `pulumi up`

```text
$ pulumi up
Previewing update (<stack-name>)

     Type                    Name                           Plan       
 +   pulumi:pulumi:Stack     pulumi-s3-static-webpage-dev   create     
 +   ├─ aws:s3:Bucket        website-bucket                 create     
 +   └─ aws:s3:BucketObject  index.html                     create     
 
Outputs:
    bucketEndpoint: output<string>
    bucketName    : output<string>

Resources:
    + 3 to create

Do you want to perform this update? yes
Updating (<stack-name>)

     Type                    Name                           Status      
 +   pulumi:pulumi:Stack     pulumi-s3-static-webpage-dev   created     
 +   ├─ aws:s3:Bucket        website-bucket                 created     
 +   └─ aws:s3:BucketObject  index.html                     created     
 
Outputs:
    bucketEndpoint: "http://website-bucket-f7900c5.s3-website-eu-west-1.amazonaws.com"
    bucketName    : "website-bucket-f7900c5"

Resources:
    + 3 created

Duration: 7s
```

Check out your static website at the URL in the `outputs`

`curl $(pulumi stack output bucketEndpoint)`

### Remove reources

Destroy the stack: `pulumi destroy`
Delete the stack itself: `pulumi stack rm`
