# Google Auth

## Ref

[google-auth](https://googleapis.dev/python/google-auth/latest/index.html)
[google-auth-oauthlib](https://google-auth-oauthlib.readthedocs.io/en/latest/index.html)

## Account and identitication

there are two types of ways of identifying an applicatino or user. two types of accounts are service account and user account.

Credentials from service accounts identify a particular application. These types of credentials are used in server-to-server use cases, such as accessing a database. This library primarily focuses on service account credentials.

Credentials from user accounts are obtained by asking the user to authorize access to their data. These types of credentials are used in cases where your application needs access to a user’s data in another service, such as accessing a user’s documents in Google Drive. This library provides no support for obtaining user credentials, but does provide limited support for using user credentials.

## How to obtain credentials

1. ADC (application default credentials)

    1. Enable locally by gcloud command
    ```
    gcloud auth application-default login
    ```

    2. Install default credentials
    ```
    import google.auth
    credentials, project = google.auth.default()
    ```

2. Service account private key files
3. Compute engine, container engine, and app engine
4. User credentials

    1. Make an oauth client in [Google API Console](https://console.developers.google.com/apis/credentials)
    2. Download client secret file
    3. Get access token by using google-auth-oauthlib
    4. Obtain credentials

5. Impersonated credentials
6. Identity tokens

## ADC Log

```
docker build -t google-auth .
docker run --name google-auth -itd google-auth bash
docker exec -it google-auth gcloud auth application-default login
docker exec google-auth python adc.py
```

```
docker stop google-auth
docker rm google-auth
```
