# service_templated module template

## Motivation

A [templatectl](https://github.com/4thel00z/templatectl) template to generate modules for [service_templated](https://github.com/4thel00z/service_templated)

## Prerequesites

This template uses [templatectl](https://github.com/4thel00z/templatectl/releases/latest).

## Usage

```
templatectl template download 4thel00z/service-templated-module service
```

Inside the root of your [service_templated](https://github.com/4thel00z/service_templated) project you can do:

```
templatectl template use service ./pkg/libyourproject/
```

## License

This project is licensed under the GPL-3 license.
