import json
import sys

try:
    config_output = {}
    open_api_file = sys.argv[1]
    with open(open_api_file, "r") as jfile:
        oapi = json.load(jfile)
    jfile.close()
    for tag in oapi["tags"]:
        # print(tag["name"])
        config_output.update({tag["name"]: {}})
    for endpoints in oapi["paths"]:
        method = ""
        try:
            tval = oapi["paths"][endpoints]["get"]
            method = "GET"
        except Exception as e:   
            pass
        try:
            tval = oapi["paths"][endpoints]["post"]
            method = "POST"
        except Exception as e:   
            pass 
        epoint = []
        params = []
        for part in endpoints.split("/"):
            if "{" in part:
                # print(part)
                params.append(part.strip("{").strip("}"))
                epoint.append("%s")
            else:
                epoint.append(part)
        parsed_endpoint = '/'.join(epoint)
        oid = tval["operationId"].split(".")[1]
        if params:
            config_output[tval["tags"][0]].update({oid: {"method": method, "endpoint": parsed_endpoint, "parameters": params}})
        else:
            config_output[tval["tags"][0]].update({oid: {"method": method, "endpoint": parsed_endpoint}})

        
    json_formatted_str = json.dumps(config_output, indent=2)
    with open("endpoint_configuration.json", "w") as njfile:
        json.dump(config_output, njfile, indent=2)
    njfile.close()
    # print(json_formatted_str)
except Exception as e:
    print(e)

