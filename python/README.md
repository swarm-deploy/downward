# downward for Python

Python SDK for the [downward](https://github.com/swarm-deploy/downward) Docker Swarm metadata contract.

Install directly from the repository:

```bash
pip install "downward @ git+https://github.com/swarm-deploy/downward.git#subdirectory=python"
```

Load metadata from the environment:

```python
import downward

info = downward.load()

print(info.stack.name)
print(info.service.name)
print(info.task.slot)
print(info.node.name)
```
