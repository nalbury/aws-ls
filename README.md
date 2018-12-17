# aws-ls

A quick easy alternative to parsing aws cli output, or the AWS web console for listing all resources of a particular type (e.g. ec2 instances). 

## Usage

```
➜  ~ aws-ls --help
ls your AWS resources

Usage:
  aws-ls [command]

Available Commands:
  ec2         ls your ec2 instances
  help        Help about any command

Flags:
  -h, --help             help for aws-ls
  -n, --no-headers       Supress column names
  -p, --profile string   AWS profile to use (default "default")

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
INDEX    NAME             INSTANCE_ID            PRIVATE_IP    STATUS
0        my-instance-1    i-0c3b898n45039f7a9    1.2.3.4       running
1        my-instance-2    i-0e639e5ceb5cb0637    1.2.3.4       stopped
2        my-instance-3    i-0a638fd2cxdedbdb5    1.2.3.4       terminated
```

## Installing

[Downloads for each release can be found here.](https://gitlab.pizza/nalbury/aws-ls/tags)

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

