"""
An example program to implement Mapping abc
"""

from __future__ import annotations

import json
from collections import defaultdict
from collections.abc import Mapping
from dataclasses import dataclass, field
from functools import lru_cache
from operator import itemgetter
from typing import Any, ClassVar, Iterator, KeysView, Optional, Self

import asteval
import inflect
import yaml


class InternalError(Exception):
    """Represents an inernal error. This shouldn't happen basically."""


@dataclass
class NoChildError(Exception):
    """aaa"""

    child_name: str


@dataclass
class Scope:
    """aaa"""

    raw_data: dict[str, Any] | list[dict[str, Any]]
    descriptor: Descriptor

    def keys(self) -> KeysView[str]:
        "aaa"
        if isinstance(self.raw_data, list):
            return self.raw_data[0].keys()
        return self.raw_data.keys()

    def here(self, field_name: str) -> bool:
        "aaa"
        if isinstance(self.raw_data, list):
            return field_name in self.raw_data[0]
        return field_name in self.raw_data

    def value(self, field_name: str) -> Any:
        "aaa"
        if isinstance(self.raw_data, list):
            return [x[field_name] for x in self.raw_data]
        return self.raw_data[field_name]

    def in_parent(self, field_name: str) -> bool:
        "aaa"
        if self.descriptor.parent is None:
            return False
        return self.descriptor.parent[1].name == field_name

    def in_children(self, field_name: str) -> bool:
        "aaa"
        return field_name in self.descriptor.children

    def in_custom_fields(self, field_name: str) -> bool:
        "aaa"
        return field_name in self.descriptor.custom_fields


@dataclass
class Data(Mapping[str, Any]):
    "aaa"

    store: DataStore
    scope: Scope
    descriptors: dict[str, Descriptor]

    def __iter__(self) -> Iterator[str]:
        return iter(self.scope.keys())

    def __len__(self) -> int:
        return len(self.scope.keys())

    def __getitem__(self, name: str) -> Any:
        if self.scope.here(name):
            return self.scope.value(name)
        if self.scope.in_parent(name):
            return self.parent()
        if self.scope.in_children(name):
            return self.child(name)
        if self.scope.in_custom_fields(name):
            return self.custom_field_value(name)
        raise KeyError(name)

    def __getattr__(self, name: str) -> Any:
        return self.__getitem__(name)

    def parent(self) -> Data:
        "aa"
        if self.scope.descriptor.parent is None:
            raise InternalError
        field_name_to_parent, parent_descriptor = itemgetter(0, 1)(
            self.scope.descriptor.parent
        )
        parent_id = self[field_name_to_parent]
        parent_raw_data = self.store.get_raw_data(
            parent_descriptor.name, parent_descriptor.id_field, parent_id
        )
        parent_scope = Scope(parent_raw_data, parent_descriptor)
        return self.store.get_expandable_data(parent_scope, self.descriptors)

    def child(self, name: str) -> Any:
        "aa"
        child_descriptor = self.scope.descriptor.children[name]
        if child_descriptor.parent is None:
            raise InternalError
        filed_name_to_me = child_descriptor.parent[0]
        if isinstance(self.scope.raw_data, list):
            raise InternalError
        my_id = self.scope.raw_data[self.scope.descriptor.id_field]
        try:
            child_raw_data_list = self.store.get_raw_data_list(
                child_descriptor.name, filed_name_to_me, my_id
            )
        except KeyError as exc:
            if exc.args[0] == my_id:
                raise NoChildError(child_name=name) from exc
        else:
            child_scope = Scope(child_raw_data_list, child_descriptor)
        return self.store.get_expandable_data(child_scope, self.descriptors)

    def custom_field_value(self, field_name: str) -> Any:
        "a"
        expression = self.scope.descriptor.custom_fields[field_name]
        symtable = asteval.make_symbol_table(**self, ref=self)
        aeval = asteval.Interpreter(symtable=symtable)
        return aeval.eval(expression)


@dataclass
class DataStore:
    "aa"
    raw_data_dict: dict[str, list[dict[str, Any]]]

    def get_expandable_data(
        self, scope: Scope, descriptors: dict[str, Descriptor]
    ) -> Data:
        "aaa"
        return Data(store=self, scope=scope, descriptors=descriptors)

    def get_raw_data(
        self, name: str, field_name_for_id: str, _id: str
    ) -> dict[str, Any]:
        "aaa"
        return self.get_raw_data_list(name, field_name_for_id, _id)[0]

    def get_raw_data_list(
        self, name: str, field_name: str, field_value: str
    ) -> list[dict[str, Any]]:
        "aaa"
        return self._get_table(name, field_name)[field_value]

    @lru_cache
    def _get_table(self, name: str, field_name: str) -> dict[str, list[dict[str, Any]]]:
        table: defaultdict[str, list[dict[str, Any]]] = defaultdict(
            list[dict[str, Any]]
        )
        for raw_data in self.raw_data_dict[name]:
            table[raw_data[field_name]].append(raw_data)
        return dict(table)

    def __hash__(self) -> int:
        return hash(json.dumps(self.raw_data_dict, sort_keys=True))


@dataclass
class Manager:
    "aaa"

    data_store: DataStore
    descriptors: dict[str, Descriptor]

    @classmethod
    def from_dict(
        cls,
        descriptor_dict: dict[str, Any],
        raw_data_dict: dict[str, list[dict[str, Any]]],
    ) -> Self:
        "a"

        data_store = DataStore(raw_data_dict)
        parent_dict: dict[str, tuple[str, str]] = {}
        descriptors = {}
        for name, raw_descriptor in descriptor_dict.items():
            try:
                raw_parent = raw_descriptor.pop("parent")
            except KeyError:
                pass
            else:
                parent_dict[name] = itemgetter("name", "field")(raw_parent)
            descriptors[name] = Descriptor(name=name, **raw_descriptor)
        for name, (parent_name, field_name_to_parent) in parent_dict.items():
            descriptors[name].associate(field_name_to_parent, descriptors[parent_name])
        return cls(data_store=data_store, descriptors=descriptors)

    def eval_expression(self, name: str, expression: str) -> None:
        "a"
        print(f"evaluating {expression} for {name}")
        descriptor = self.descriptors[name]
        for idx, raw_data in enumerate(self.data_store.raw_data_dict[name]):
            print(f"evaluating {idx}")
            data = self.data_store.get_expandable_data(
                Scope(raw_data, descriptor), self.descriptors
            )
            symtable = asteval.make_symbol_table(**data, ref=data)
            aeval = asteval.Interpreter(symtable=symtable)
            try:
                print(aeval.eval(expression, raise_errors=True))
            except NoChildError:
                pass


@dataclass
class Descriptor:
    "a"
    name: str
    id_field: str = "id"
    custom_fields: dict[str, str] = field(default_factory=dict)
    parent: Optional[tuple[str, Descriptor]] = None
    children: dict[str, Descriptor] = field(default_factory=dict)

    inflect_engine: ClassVar[inflect.engine] = inflect.engine()

    @property
    def name_plural(self) -> str:
        "a"
        return self.inflect_engine.plural(self.name)

    def associate(
        self, field_name_to_parent: str, parent_descriptor: Descriptor
    ) -> None:
        "a"
        self.parent = (field_name_to_parent, parent_descriptor)
        parent_descriptor.children[self.name_plural] = self


DESCRIPTOR_YAML = """
---
Aa:
  id_field: id
Bb:
  id_field: id
  parent:
    name: Aa
    field: aa_id
  custom_fields:
    afield1: ref.Aa.field1
"""

DATA_YAML = """
---
Aa:
  - id: aaa1
    field1: 123
    field2: "234"
    field4: true
  - id: aaa2
    field1: 123
    field2: "234"
    field4: false
Bb:
  - id: bbb1
    field3: 345
    aa_id: aaa2
"""


if __name__ == "__main__":
    m = Manager.from_dict(
        descriptor_dict=yaml.safe_load(DESCRIPTOR_YAML),
        raw_data_dict=yaml.safe_load(DATA_YAML),
    )
    m.eval_expression("Aa", "123")
    m.eval_expression("Aa", "id")
    m.eval_expression("Bb", "ref.Aa.field1")
    m.eval_expression("Bb", "ref.Aa.field2 + '123'")
    m.eval_expression("Aa", "ref.Bbs.field3")
    m.eval_expression("Bb", "ref.afield1")
    m.eval_expression("Bb", "ref.afield1 // 3")
    m.eval_expression("Aa", "field1 if field4 else field2")
