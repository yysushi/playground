import uuid
from dataclasses import dataclass, field

import yaml
from openapi_core import Spec
from openapi_core.validation.exceptions import MissingRequiredParameter
from openapi_core.validation.request import openapi_request_validator
from openapi_core.validation.request.datatypes import RequestParameters


@dataclass
class MyRequest:
    method: str
    host_url: str
    path: str
    mimetype: str
    parameters: RequestParameters
    body: None = field(init=False)


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

param = RequestParameters(header={"X-Request-ID": str(uuid.uuid4())})
request2 = MyRequest("get", "http://example.com", "/hoge", "application/json", param)
result = openapi_request_validator.validate(spec, request2)
result.raise_for_errors()
