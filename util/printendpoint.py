import requests
# import prettyprints
import json
import sys

try:
    config = {}
    with open("bungieAPI_config.json", "r") as jfile:
        config = json.load(jfile)
    jfile.close()

    headers_dict = {"Content-Type": "application/json", "X-API-KEY": config["x-api-key"]}
    response = requests.get(f'{config["base-path"]}{sys.argv[1]}', headers=headers_dict)
    # json_object = json.loads()

    json_formatted_str = json.dumps(response.json(), indent=2)

    print(json_formatted_str)

except Exception as e:
    print(e)