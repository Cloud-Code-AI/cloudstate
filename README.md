  [![Go Report Card](https://goreportcard.com/badge/github.com/Cloud-Code-AI/cloudstate)](https://goreportcard.com/report/github.com/Cloud-Code-AI/cloudstate) 

# CloudState
CloudState is an innovative open-source tool designed to empower developers, system administrators, and cloud architects. It provides a detailed, real-time snapshot of all resources deployed across your cloud environments, generating an accurate and up-to-date state file. This tool is your go-to solution for cloud resource management, offering clarity and control over your cloud infrastructure.

## Key Features:
- State File Generation: Automatically generates a state file that reflects your current cloud resource setup, facilitating better infrastructure management and planning.
- Multi-Cloud Support: Designed with versatility in mind, CloudState seamlessly integrates with various cloud providers, ensuring broad compatibility and utility.
- Real-Time Updates: Stay informed with real-time updates, ensuring that your state files always reflect the latest changes in your cloud environment.

## Getting Started:
To get started with CloudState, please refer to our comprehensive documentation [here]. This includes installation instructions, usage guides, and best practices for leveraging CloudState in your cloud infrastructure.

## Installation

Instructions on how to install your project. This section usually starts with cloning the repository and then proceeding with specific steps.

### Prerequisites

- Ensure Go 1.18 is installed

### Build Project

```
git clone https:/github.com/Cloud-Code-AI/cloudstate.git
cd cloudstate
go build cmd/main.go
```

### Usage

To Gather all resources
```
./main gather -provider=aws -region=us-east-1
```
This will store all the meta data for AWS resource in `output/` directory


To Generate report:

Note: Make sure to run gather command first as report runs on local meta data gathered from cloud resources
```
./main report --provider=aws
```

You should see an output like this:
```
us-east-1
    kms                       
        keys                   2                    
    rds                       
        instances              0                    
    route53                   
        domains                0                    
    lambda                    
        functions              4                    
        layers                 0                    
    s3                        
        buckets                9                    
    cloudfront                
        websites               2                    
    ec2                       
        amis                   0                    
        instances              0                    
        snapshots              0                    
    elasticache               
        cache_clusters         0                    
        snapshots              0                    
    codebuild                 
        builds                 0                    
        enviroment_images      5                    
        projects               0                    
    dynamodb                  
        tables                 2                    
    ess                       
        domains                0                    
    ecr                       
        repositories           0                    
    iam                       
        policies               29                   
        roles                  37                   
        unused_policies        3                    
        users                  5                    
    apigateway                
        stacks                 1                    
    cloudformation            
        stacks                 0                    
    cloudwatch                
        metrics                272                  
        dashboards             0                    
        metric_streams         0                    
us-east-2
    cloudfront                
        websites               2                    
    dynamodb                  
        tables                 3                    
    ec2                       
        amis                   0                    
        instances              1                    
        snapshots              0                    
    iam                       
        policies               29                   
        roles                  38                   
        unused_policies        3                    
        users                  5                    
    lambda                    
        functions              5                    
        layers                 1                    
    rds                       
        instances              1                    
    s3                        
        buckets                9
```


And you can also find your results at `output/aws_report.json`


## Contributing:
We welcome contributions from the community! If you're interested in contributing to CloudState, please check out our contributing guidelines [here]. Whether it's adding new features, fixing bugs, or improving documentation, your help is greatly appreciated.

### License:
CloudState is released under the [MIT License](/LICENSE). Feel free to use, modify, and distribute it as part of your projects.

### Stay Connected:
For updates, follow us on [Github](https://github.com/Cloud-Code-AI) and join our community on [Discord](https://discord.gg/tEPMDTxX9K)

