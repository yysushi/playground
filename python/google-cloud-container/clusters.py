from google.cloud import container
from google.cloud import compute
import google.auth


credentials, project_id = google.auth.default()
location = 'asia-northeast1'

parent = f'projects/{project_id}/locations/{location}'
client = container.ClusterManagerClient(credentials=credentials)
response = client.list_clusters(
    parent=parent,
)
for c in response.clusters:
    print(c)
