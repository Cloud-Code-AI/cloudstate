---
name: Add new service monitoring
about: Describe this issue template's purpose here.
title: 'feat:'
labels: needs approval
assignees: ''

---

## Description
The purpose of this task is to provide a list or description of <SERVICE NAME> related services and resources. 

## How-to Guide
To complete this task, follow these steps:
1. In awshandler, copy one of the go files (e.g. iam.go) and rename it to `<service>.go`.
2. Add all the services and resources you can find in the code to the newly created .go file.
3. Once the .go file is ready, import the function and add it to initialize.go.

## Reference
You can find the GoDoc at this URL: <SERVICE REFERENCE URL>
