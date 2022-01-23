import requests
# import prettyprints
import json
import sys
import os

try:
    config = {}
    with open("../config/bungieAPI_config.json", "r") as jfile:
        config = json.load(jfile)
    jfile.close()

    with open("../config/temp_tokens.json", "r") as jfile:
        tokens = json.load(jfile)
    jfile.close()

    headers_dict = {"Content-Type": "application/json", "X-API-KEY": os.getenv("X_API_KEY"), "Authorization": "Bearer " + tokens["access_token"]}
    print(tokens["access_token"])
    response = requests.get(f'{config["base-path"]}{sys.argv[1]}', headers=headers_dict)
    # json_object = json.loads()

    json_formatted_str = json.dumps(response.json(), indent=2)

    print(json_formatted_str)

except Exception as e:
    print(e)