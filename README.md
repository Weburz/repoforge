# Terox: Scaffold Your Project With Ease and Speed

Terox is a CLI tool to scaffold projects from pre-built templates. Invoke
Terox with a path to a template (either local or online) and watch your
project get scaffolded automatically. Setting up projects from scratch and in a
reproducible manner is time-consuming and difficult and that is where Terox
will prove most helpful.

## Why Terox Exists

Development on Terox is heavily inspired from existing tools like
[Create-React-App (CRA)](https://create-react-app.dev),
[Cookiecutter](https://cookiecutter.readthedocs.io) and the likes. While the
existing tools have their use cases and they fulfill their roles quite aptly,
Terox aims to resolve the following pain points:

1. The need for a cross-platform and distributable binary executable unlike
   Cookiecutter which requires [Python](https://www.python.org) (or
   [Node.js](https://nodejs.org) for CRA) to be installed on a system where they
   are expected to be used.

2. The inbuilt templating engine in Go is amazing and very capable of lots of
   things! We want to utilise these capabilities and allow the users to
   dynamically scaffold their projects according to the preconfigured template.

3. Tools like Cookiecutter while extremely mature, are not well maintained
   anymore (see
   [one such discussion](https://github.com/cookiecutter/cookiecutter/issues/1642)
   among many others) and CRA is only used to scaffold React.js projects. We
   wanted a more universal tool to get a similar job done quickly and
   efficiently.

## Functionalities of Terox

**NOTE**: Terox is still a WIP project in its very initial phase and is not stable. Hence,
its UI/UX elements are subject to change without prior notice. We request you to
try out the pre-release versions of Terox and provide us with feedback. Any
feedback, no matter how trivial will help us improve the experience of the tool significantly!

That said, Terox provides the following core functionalities:

1. The ability to scaffold a project based on a pre-existing template (hosted
   either on GitHub or some archived file elsewhere). Terox will be able to
   fetch Git repositories from anywhere as long as the remote repository
   supports the Git protocol.

2. Help the user create templates which will be used in the future for
   scaffolding more projects. Do note though, Terox is not intelligent
   enough to decide what the template is supposed to be like, some manual
   tinkering will still be required from the user's end!

3. Provide the necessary utilities to manage and maintain all the locally
   available templates, like keeping them updated, listing local templates,
   deleting them when necessary and such.

## Contribution and Development Guide

<!-- TODO: Add a contributing and development guide later on -->

## Licensing and Distribution Rights

Terox is developed and open-sourced under a public open-source license (the
MIT license). Hence, you are free to use, copy and distribute the code for the
project under the terms of conditions of the aforementioned license. For more
information on the licensing details, refer to the [LICENSE](./LICENSE)
document.
