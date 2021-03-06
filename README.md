# aws-ls

A quick easy alternative to parsing aws cli output, or the AWS web console for listing all resources of a particular type (e.g. ec2 instances). 

## Installation(macos)

To install to `/usr/local/bin`

```
curl -o /usr/local/bin/aws-ls https://s3.amazonaws.com/aws-ls/latest/macos/aws-ls && chmod +x /usr/local/bin/aws-ls
```

## Usage

```
➜  ~ aws-ls --help
ls your AWS resources

Usage:
  aws-ls [command]

Available Commands:
  asg         ls your auto scaling groups
  ec2         ls your ec2 instances
  ecr         ls your ecr repos
  elb         ls your ec2 load balancers
  help        Help about any command

Flags:
  -h, --help             help for aws-ls
  -n, --no-headers       Supress column names
  -p, --profile string   AWS profile to use (default "default")
  -t, --tags strings     Tag filters in the format Key:Value,Key2:Value2

Use "aws-ls [command] --help" for more information about a command.
```

## Available Resource Types

### ec2

Lists ec2 instaces.

**Default Attributes**:
- Name Tag
- Instance ID
- Private IP Address
- Status (running, stopped, terminated etc.)

**Example**:
```
➜  ~ aws-ls ec2
INDEX    NAME             INSTANCE_ID            PRIVATE_IP    INSTANCE_TYPE    STATUS
0        my-instance-1    i-0c3b898n45039f7a9    1.2.3.4       t2.medium        running
1        my-instance-2    i-0e639e5ceb5cb0637    1.2.3.4       t2.medium        stopped
2        my-instance-3    i-0a638fd2cxdedbdb5    1.2.3.4       c5.xlarge        terminated
```

### elb

Lists elastic load balancers

**Default Attributes**:
- DNS Name
- Instance Count
- Health Check Endpoint
- Listener Configuration

**Example**:
```
➜  ~ aws-ls elb
INDEX    DNS_NAME                                                     INSTANCE_COUNT    HEALTH_CHECK          LISTENERS
0        my-elb-123456789.us-east-1.elb.amazonaws.com                 0                 HTTP:80/index.html    HTTPS:443 -> HTTP:80
1        internal-other-elb-1234566789.us-east-1.elb.amazonaws.com    4                 TCP:1234              SSL:443 -> TCP:1234
```

### asg

Lists auto scaling groups.

**Default Attributes**:
- Name
- Desired Instances
- Current Instances
- Min Instances
- Max Instances
- Launch Config Name

**Example**:
```
➜  ~ aws-ls asg
INDEX    NAME           DESIRED    CURRENT    MIN    MAX    LAUNCH_CONFIG
0        nginx-asg      3          3          1      3      nginx-launch-config
1        backend-asg    3          3          1      3      backend-launch-config
```


### ecr

Lists ecr repos
**Default Attributes**:
- Name
- Uri

**Example**
```
➜  ~ aws-ls ecr
INDEX    NAME       URI
0        nginx-image    1234566789.dkr.ecr.us-east-1.amazonaws.com/nginx-image
1        backend-image    1234566789.dkr.ecr.us-east-1.amazonaws.com/backend-image
```


## Installing

[Downloads for each release can be found here.](https://github.com/nalbury/aws-ls/releases)

You can also build from source (requires Go 1.11 for go.mod support):
```
git clone https://gitlab.pizza/nalbury/aws-ls
cd ./aws-ls
#If outside of $GOPATH
go build
#If inside of $GOPATH
GO111MODULE=on go build
```

[More info on Go Modules can be found here.](https://github.com/golang/go/wiki/Modules)

