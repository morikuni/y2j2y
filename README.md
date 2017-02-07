# y2j2y
yaml to json / json to yaml converter

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
