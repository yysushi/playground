from ruamel.yaml import YAML

inp = """\
- &CENTER {x: 1, y: 2}
- &LEFT {x: 0, y: 2}
- &BIG {r: 10}
- &SMALL {r: 1}
# All the following maps are equal:
# Explicit keys
- x: 1
  y: 2
  r: 10
  label: center/big
# Merge one map
- <<: *CENTER
  r: 10
  label: center/big
# Merge multiple maps
- <<: [*CENTER, *BIG]
  label: center/big
# Override
- <<: [*BIG, *LEFT, *SMALL]
  x: 1
  label: center/big
  fuga: *CENTER
"""

yaml = YAML()
data = yaml.load(inp)
for d in data:
    print(d.anchor.value)
    print(dir(d))
    # print(d.alias.value)
print(data)
