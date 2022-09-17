import requests

authresult = requests.post("https://stratosphaere.codelix.de/api/v1/auth", data={
    "username": "Felix Weglehner",
    "password": "KyP[vxhoRs64"
}).json()

print(authresult)

articlesresult = requests.post("https://stratosphaere.codelix.de/api/v1/blog/articles", headers={"Authorization":"Bearer " + authresult["data"]["token"]}).json()

print(articlesresult)

articleresult = requests.get("https://stratosphaere.codelix.de/api/v1/blog/articles/"+str(articlesresult["data"]["id"])).json()

print(articleresult)

wrongarticleresult = requests.get("https://stratosphaere.codelix.de/api/v1/blog/articles/3123").json()

print(wrongarticleresult)

deletearticleresult = requests.delete("https://stratosphaere.codelix.de/api/v1/blog/articles/"+str(articlesresult["data"]["id"])).json()

print(deletearticleresult)


articleresult = requests.get("https://stratosphaere.codelix.de/api/v1/blog/articles/"+str(articlesresult["data"]["id"])).json()

print(articleresult)