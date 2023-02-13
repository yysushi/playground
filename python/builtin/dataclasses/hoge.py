from collections import defaultdict
from dataclasses import InitVar, dataclass, field
from typing import Optional


@dataclass
class Hoge:
    # two init methods
    # i. __init__ (cls(**kwargs))
    # ii. __post_init__

    # how to define fields
    #
    # a1. normal field
    # => be arg in __init__
    # a2. normal field & default
    # => omittable in __init__ and set with default if omitted
    #
    # b1. `field(init=False)`
    # => no arg in __init__
    # => needs be instantiated anywhere (__post_init__ or ..)
    # b2. `field(init=False)` & default given
    # => set with default
    #
    # c1. `InitVar[T]`
    # => no arg in __init__ but be arg in __post_init__
    # c2. `InitVar[T]` & default value
    # => no arg in __init__ but be omittable arg in __post_init__

    # how to set default
    # d1: str | None = None
    # d2: Optional[str] = None
    # d3: Optional[str] = field(default=None)
    # e1: dict[str, list[str]] = field(default_factory=dict)
    # e2: dict[str, list[str]] = {} => NG!!!
    # e3: dict[str, list[str]] = dict
    # f1: dict[str, list[str]] = field(default_factory=lambda: defaultdict(list))
    # g1: Optional[InitVar[str]] = None
    # g2: InitVar[Optional[str]] = None => NG!!!

    field_a1: int
    # field_a2: int = -1
    field_c1: InitVar[int]

    field_b1: int = field(init=False)
    field_b2: int = field(init=False, default=-1)

    # field_c1: InitVar[int]
    field_c2: InitVar[int] = -1

    field_d1: str | None = None
    field_d2: Optional[str] = None
    field_d3: Optional[str] = field(default=None)

    field_e1: dict[str, list[str]] = field(default_factory=dict)
    # field_e2: dict[str, list[str]] = {} => NG!!!
    field_e3: dict[str, list[str]] = dict

    field_f1: dict[str, list[str]] = field(
        default_factory=lambda: defaultdict(list[str])
    )

    field_g1: Optional[InitVar[str]] = None
    # field_g2: InitVar[InitVar[str]] = None => NG!!!

    def __post_init__(self, field_c1, field_c2):
        print(f"{field_c1=} {field_c2=}")


print(vars(Hoge(field_a1=1, field_c1=2)))
