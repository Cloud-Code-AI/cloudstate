# CloudState
CloudState is an innovative open-source tool designed to empower developers, system administrators, and cloud architects. It provides a detailed, real-time snapshot of all resources deployed across your cloud environments, generating an accurate and up-to-date state file. This tool is your go-to solution for cloud resource management, offering clarity and control over your cloud infrastructure.

## Key Features:
- Resource Visualization: CloudState scans your cloud environment, presenting a clear and comprehensive view of all deployed resources.
- State File Generation: Automatically generates a state file that reflects your current cloud resource setup, facilitating better infrastructure management and planning.
- Multi-Cloud Support: Designed with versatility in mind, CloudState seamlessly integrates with various cloud providers, ensuring broad compatibility and utility.
- Real-Time Updates: Stay informed with real-time updates, ensuring that your state files always reflect the latest changes in your cloud environment.
- Easy Integration: CloudState is built to easily integrate into your existing cloud management workflow, enhancing your resource monitoring and decision-making processes.

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

```
./main -provider=aws -region=us-east-1
```

This will store all the meta data for AWS resource in `output/` directory

## Contributing:
We welcome contributions from the community! If you're interested in contributing to CloudState, please check out our contributing guidelines [here]. Whether it's adding new features, fixing bugs, or improving documentation, your help is greatly appreciated.

### License:
CloudState is released under the [MIT License](/LICENSE). Feel free to use, modify, and distribute it as part of your projects.

### Stay Connected:
For updates, follow us on [Github](https://github.com/Cloud-Code-AI) and join our community on [Discord](https://discord.gg/tEPMDTxX9K)

