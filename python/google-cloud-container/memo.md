# google-cloud-container

## Log for local use case

- to allow application to login

```
koketani: ~/g/g/k/p/p/google-cloud-container (master ?)$ gcloud auth application-default login --no-launch-browser
Go to the following link in your browser:

    https://accounts.google.com/o/oauth2/auth?response_type=code&client_id=...

Enter verification code: ...

Credentials saved to file: [/Users/koketani/.config/gcloud/application_default_credentials.json]

Quota project "..." was added to ADC which can be used by Google client libraries for billing and quota. Note that some services may still bill the project owning the resource.
```

- revoke

```
koketani: ~/g/g/k/p/p/google-cloud-container (master ?)$ gcloud auth application-default revoke
```

## Log for container

```
$ docker build -t google-cloud-container .
$ docker run --name google-cloud-container -itd google-cloud-container bash
$ docker exec -it google-cloud-container gcloud config set account $EMAIL
$ docker exec -it google-cloud-container gcloud config set project $PROJECT_ID
$ docker exec -it google-cloud-container gcloud auth login
$ docker exec -it google-cloud-container gcloud auth application-default login
$ docker exec google-cloud-container python clusters.py
```
