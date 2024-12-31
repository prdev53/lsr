# lsr

lsr is a simple tool for listing files and filtering unwanted ones. It is similar to `fd --type file`, but with its own defaults and configuration file.

### Getting started

From the root directory of the project, run:

```console
$ cd src
$ go build
```

This will generate `lsr` in the `src` folder.

Alternatively, for non-windows users, if you have Perl installed, you can run:

```console
$ perl make.pl
```

The `lsr` executable will be in the root directory of the project if you do so.

### Configuring filters

Create a [.lsrignore](/.lsrignore) file in the directory where `lsr` is meant to be run. This file is composed of two sections:
1. Name suffixes
2. Full names

Both sections separated by the first empty line in the file.

**Note:** 
- `lsr` does not distinguish between directories and files when filtering. 
- It also does not support paths, only file names.
