# create-dcos-app
A tool to create DC/OS apps built using the dcos-commons SDK

*NOTE:* This is still a work in progress

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

### There is no single place to make Marathon configuration changes affecting multiple frameworks

### There is no single place to make svc.yml changes affecting multiple frameworks

### There is no consistent directory structure for an SDK framework


## The goal

`create-dcos-app` is a "compiler" for DC/OS applications.



## The folder structure
*


## Using `create-dcos-app`:
*TODO:* implement the build command as in this section

Running the command:
```bash
$ create-dcos-app build application.yml
```
(where `applcition.yml` is the YAML definition of the application)
will result in the following:

The binaries (e.g. scheduler, CLI) will be built. The files from each of the sections of the YAML file (see `template.yml`)
will be generated.



TODO:
* Transfer contents of https://docs.google.com/document/d/1652Wxbq9sPme82irN29-rsxhz9WXBtrpWHS6QmPRPZA/edit to here
* Add support for compiling jars
