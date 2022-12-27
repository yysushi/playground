import json
import uuid
from dataclasses import dataclass, field
from typing import Any, Mapping

import yaml
from jsonschema.exceptions import ValidationError
from openapi_core import Spec
from openapi_core.unmarshalling.schemas.exceptions import InvalidSchemaValue
from openapi_core.validation.exceptions import MissingRequiredParameter
from openapi_core.validation.request import openapi_request_validator
from openapi_core.validation.request.datatypes import RequestParameters
from openapi_core.validation.response import openapi_response_validator


@dataclass
class MyRequest:
    method: str
    host_url: str
    path: str
    mimetype: str
    parameters: RequestParameters
    body: None = field(init=False)


@dataclass
class MyResponse:
    status_code: int
    mimetype: str
    headers: Mapping[str, Any]
    data: str


with open("openapi.yaml", "r", encoding="utf-8") as spec_file:
    spec_dict = yaml.safe_load(spec_file)
spec = Spec.create(spec_dict)

request = MyRequest(
    "get", "http://example.com", "/hoge", "application/json", RequestParameters()
)
result = openapi_request_validator.validate(spec, request)
try:
    result.raise_for_errors()
except MissingRequiredParameter as e:
    assert e.name == "X-Request-ID", "unexpected openaapi validation error"
else:
    raise AssertionError("unexpected no error")

param = RequestParameters(header={"X-Request-ID": str(uuid.uuid4())})
request2 = MyRequest("get", "http://example.com", "/hoge", "application/json", param)
result = openapi_request_validator.validate(spec, request2)
result.raise_for_errors()


response = MyResponse(
    200,
    "application/json",
    {},
    json.dumps({}),
)
result2 = openapi_response_validator.validate(spec, request2, response)
try:
    result2.raise_for_errors()
except InvalidSchemaValue as e:
    for idx, schema_error in enumerate(e.schema_errors):
        assert idx == 0, "unexpected length of schema errors"
        assert type(schema_error) == ValidationError, "unexpected exception"
        assert (
            schema_error.validator_value[0] == "color"
        ), "unexpected json schema validation"
        break
    else:
        raise AssertionError("unexpected length of schema errors")
else:
    raise AssertionError("unexpected no error")


response = MyResponse(
    200,
    "application/json",
    {},
    json.dumps({"color": "cyan", "added": True}),
)
result2 = openapi_response_validator.validate(spec, request2, response)
try:
    result2.raise_for_errors()
except InvalidSchemaValue as e:
    for idx, schema_error in enumerate(e.schema_errors):
        assert idx == 0, "unexpected length of schema errors"
        assert type(schema_error) == ValidationError, "unexpected exception"
        assert "Additional properties are not allowed ('added' was unexpected)" in schema_error.message, "no additional error"
        break
    else:
        raise AssertionError("unexpected length of schema errors")
else:
    raise AssertionError("unexpected no error")
