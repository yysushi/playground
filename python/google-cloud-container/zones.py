from google.cloud import compute
import google.auth


credentials, project_id = google.auth.default()
print(project_id)

zc = compute.ZonesClient(credentials=credentials)
response = zc.list(project=project_id)
for z in response.items:
    print(z.name)
