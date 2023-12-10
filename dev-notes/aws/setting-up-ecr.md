# Setting up ECR

https://docs.aws.amazon.com/AmazonECR/latest/userguide/getting-started-console.html

https://docs.aws.amazon.com/AmazonECR/latest/userguide/getting-started-cli.html

[ECR Registry](https://us-east-1.console.aws.amazon.com/ecr/private-registry/repositories?region=us-east-1)

## Commands

```sh
# Tag docker image
> docker tag docker-gs-ping:latest 038417966468.dkr.ecr.us-east-1.amazonaws.com/hello-repository

# Login
$ aws ecr get-login-password --region us-east-1 | docker login --username AWS --password-stdin 038417966468.dkr.ecr.us-east-1.amazonaws.com

# Push docker image
$ docker push 038417966468.dkr.ecr.us-east-1.amazonaws.com/hello-repository
```
