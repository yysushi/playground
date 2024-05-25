import abc
from typing import ClassVar, Literal

from pydantic import BaseModel, Field


class Resource(BaseModel, abc.ABC):
    url_version: ClassVar[str] = "v2"

    @property
    @classmethod
    @abc.abstractmethod
    def type_name(self):
        pass

    @property
    @classmethod
    @abc.abstractmethod
    def type_plural_name(self):
        pass

    resource_id: str = ""

    def show_url(self):
        if not self.resource_id:
            raise ValueError("not ready to serve show url")
        return "/".join([self.url_version,
                         self.type_plural_name,
                         self.resource_id])


class A(Resource):
    type_name: ClassVar[str] = "a"
    type_plural_name: ClassVar[str] = "as"

    field_a: str


class B(Resource):
    type_name: ClassVar[str] = "b"
    type_plural_name: ClassVar[str] = "bs"
    url_prefix: ClassVar[str] = "v1"

    field_b: str


b = B(resource_id="1", field_b="hoge")
print(b.model_dump())
print(b.show_url())
a = A(resource_id="1", field_a="fuga")
print(a.show_url())
aa = A(field_a="fuga")
try:
    aa.show_url()
except ValueError:
    pass
