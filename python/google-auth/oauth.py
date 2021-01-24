from google_auth_oauthlib.flow import Flow


flow = Flow.from_client_secrets_file(
        'client_secret_141815874697-ihsj4n0t7ncrql39r5qqfj1coc2u5cdv.apps.googleusercontent.com.json',
        scopes=['openid', 'https://www.googleapis.com/auth/userinfo.email', 'https://www.googleapis.com/auth/userinfo.profile'],
        redirect_uri='urn:ietf:wg:oauth:2.0:oob')
auth_url, _ = flow.authorization_url(prompt='consent')
print('Please go to this URL: {}'.format(auth_url))
code = input('Enter the authorization code: ')
token = flow.fetch_token(code=code)
print(token)
