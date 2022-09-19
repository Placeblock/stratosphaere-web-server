import requests

authresult = requests.post("https://stratosphaere.codelix.de/api/v1/auth", data={
    "username": "Felix Weglehner",
    "password": "KyP[vxhoRs64"
}).json()

print(authresult)

articlesresult = requests.post("https://stratosphaere.codelix.de/api/v1/blog/articles", headers={"Authorization":"Bearer " + authresult["data"]["token"]}).json()

print(articlesresult)


articledata = requests.get("https://stratosphaere.codelix.de/api/v1/blog/articles/"+str(articlesresult["data"]["id"])).json()["data"]

print(articledata)

articledata["description"] = "Das ist eine Beschreibung"


""" articlesresult = requests.put("https://stratosphaere.codelix.de/api/v1/blog/articles/"+str(articlesresult["data"]["id"]), headers={"Authorization":"Bearer " + authresult["data"]["token"]}, data={
    "id": articledata["id"],
    "title": articledata["title"],
    "description": articledata["description"],
    "content": articledata["content"],
    "cover_image_url": articledata["cover_image_url"],
    "published": articledata["published"]
}).json()

print(articlesresult) """