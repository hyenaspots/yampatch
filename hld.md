# Yampatch CLI-tool syntax
yampatch -t target.yml -o op1.yml -o op2.yml

--target, -t
  sets the following file as the original YAML to be patched

--ops-file, -o
  flags the following file as containing operations for the patcher to perform

No other flags planned; current flags are purely for usability in their actual
use context.
