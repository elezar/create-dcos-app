# create-dcos-app
A tool to create DC/OS apps built using the dcos-commons SDK

## What problem does this solve?

Currently, an SDK-based DC/OS framework is made up of multiple parts:
* The Java source code defining:
    * A scheduler
    * An (optional) executor
* The service specification `svc.yml`
* The universe package definition:
    * A `package.json` file
    * A `config.json` file
    * A `resource.json` file
    * A `marathon.json.mustache` file
* The artifacts defined on the `resource.json` file

This results in the following problems with the SDK or Framework development process:

### Parameters need to specified in at least 2 (sometimes 3) places



## The folder structure



## Using `create-dcos-app`:

```bash
$ create-dcos-app compile service.yml
```





TODO:
* Transfer contents of https://docs.google.com/document/d/1652Wxbq9sPme82irN29-rsxhz9WXBtrpWHS6QmPRPZA/edit to here
* Add support for compiling jars
