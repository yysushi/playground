import requests

res = requests.get('httpbin.org/ip')
print(res.json())
