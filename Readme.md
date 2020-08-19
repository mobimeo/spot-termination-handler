# Spot Termination Handler

Welcome to Spot Termination Handler !

Please note

- This is not an official Reach-Now product.
- This functionality has since been integrated into the official [ECS Agent](https://github.com/aws/amazon-ecs-agent/blob/master/CHANGELOG.md#1320) with [PR 2205](https://github.com/aws/amazon-ecs-agent/pull/2205).

---

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
