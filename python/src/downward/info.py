from __future__ import annotations

import os
from dataclasses import dataclass
from typing import Mapping, Optional

from .vars import (
    ENV_NODE_ID,
    ENV_NODE_NAME,
    ENV_SERVICE_ID,
    ENV_SERVICE_NAME,
    ENV_STACK_NAME,
    ENV_TASK_ID,
    ENV_TASK_NAME,
    ENV_TASK_SLOT,
)


@dataclass(frozen=True)
class Stack:
    name: str

    def to_dict(self) -> dict[str, object]:
        return {ENV_STACK_NAME: self.name}


@dataclass(frozen=True)
class Service:
    id: str
    name: str

    def short_name(self, stack_name: str) -> str:
        if not stack_name:
            return self.name

        return self.name.removeprefix(f"{stack_name}_")

    def to_dict(self) -> dict[str, object]:
        return {
            ENV_SERVICE_ID: self.id,
            ENV_SERVICE_NAME: self.name,
        }


@dataclass(frozen=True)
class Task:
    id: str
    name: str
    slot: Optional[int]

    def to_dict(self) -> dict[str, object]:
        return {
            ENV_TASK_ID: self.id,
            ENV_TASK_NAME: self.name,
            ENV_TASK_SLOT: self.slot,
        }


@dataclass(frozen=True)
class Node:
    id: str
    name: str

    def to_dict(self) -> dict[str, object]:
        return {
            ENV_NODE_ID: self.id,
            ENV_NODE_NAME: self.name,
        }


@dataclass(frozen=True)
class Info:
    stack: Stack
    service: Service
    task: Task
    node: Node

    def to_dict(self) -> dict[str, object]:
        return {
            **self.stack.to_dict(),
            **self.service.to_dict(),
            **self.task.to_dict(),
            **self.node.to_dict(),
        }


def load(environ: Optional[Mapping[str, str]] = None) -> Info:
    env = os.environ if environ is None else environ

    raw_slot = env.get(ENV_TASK_SLOT)
    slot = None
    if raw_slot is not None:
        try:
            slot = int(raw_slot)
        except ValueError as exc:
            raise ValueError(f"parse {ENV_TASK_SLOT}: {exc}") from exc

    return Info(
        stack=Stack(name=env.get(ENV_STACK_NAME, "")),
        service=Service(
            id=env.get(ENV_SERVICE_ID, ""),
            name=env.get(ENV_SERVICE_NAME, ""),
        ),
        task=Task(
            id=env.get(ENV_TASK_ID, ""),
            name=env.get(ENV_TASK_NAME, ""),
            slot=slot,
        ),
        node=Node(
            id=env.get(ENV_NODE_ID, ""),
            name=env.get(ENV_NODE_NAME, ""),
        ),
    )
