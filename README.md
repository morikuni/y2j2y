# y2j2y
[![Build Status](https://travis-ci.org/morikuni/y2j2y.svg?branch=master)](https://travis-ci.org/morikuni/y2j2y)

yaml to json / json to yaml converter.

## Install

Download from [release page](https://github.com/morikuni/y2j2y/releases).

## Example

```sh
$ cat yaml.yaml
yaml:
  hello: world
  array:
    - 1
    - 2
    - 3
    
$ cat yaml.yaml | yaml2json
{"yaml":{"hello":"world","array":[1,2,3]}}

% cat yaml.yaml | yaml2json | json2yaml
yaml:
  array:
  - 1
  - 2
  - 3
  hello: world
```
