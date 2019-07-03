# spot-termination-handler

The tool is supposed to run on instances of a spot fleet backing an ECS cluster. Runs in a 5s loop polling the instance metadata for a spot termination notification (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/spot-interruptions.html#spot-instance-termination-notices). If there is a notification the instance is marked as "DRAINING" to evacuate the workload to other instances before termination.

## requirements

```
brew install go
brew install dep
```

## build

```
go build
```
