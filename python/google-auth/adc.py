import google.auth
from google.auth.transport.requests import AuthorizedSession


credentials, _ = google.auth.default()

authed_session = AuthorizedSession(credentials)
response = authed_session.get(
    'https://www.googleapis.com/storage/v1/b')
response.raise_for_status()
print(response)
