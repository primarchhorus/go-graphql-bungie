from dataclasses import field
import json
import os
import sys
import logging
import errno

logger = logging.getLogger()
logger.setLevel(logging.DEBUG)

handler = logging.StreamHandler(sys.stdout)
handler.setLevel(logging.DEBUG)
formatter = logging.Formatter('%(asctime)s - %(name)s - %(levelname)s - %(message)s')
handler.setFormatter(formatter)
logger.addHandler(handler)

lang = "go"

def create_dir(dir):
    try:
        os.mkdir(dir)
    except OSError as exc:
        if exc.errno != errno.EEXIST:
            raise
        pass

def peek(iterable):
    try:
        first = next(iterable)
        return first
    except StopIteration:
        return None

def open_api_v2(data_file):
    try:
        schema_file = open("bungie_schema.go", "w")
        schema_file.write("package main\n\n")
        with open ("bungie_data.go", "w") as data_file:
            data_file.write("package main\n\n")
            for v in oapi["definitions"]:
                object = ""
                enum = ""
                schema = ""

                if "properties" in oapi["definitions"][v]:
                    comma_counter = 0
                    object_name = v.split(".")[-1]
                    object+=f"type {object_name} struct {{\n"
                    schema+= f"var {object_name} = graphql.NewObject(\n"
                    schema+= "    graphql.ObjectConfig{\n" 
                    schema+= f"        Name: \"{object_name}\",\n"
                    schema+= "        Fields: graphql.Fields{\n"
                    for p in oapi["definitions"][v]["properties"]:
                        # Extract the property name to be used as Struct variable, sett first letter to Upper case
                        property_name = list(p)
                        quoted_json_arg = f'"{p}"'
                        property_name[0] = property_name[0].upper()
                        property_type = ""
                        property_array_object_name = ""

                        if len(oapi["definitions"][v]["properties"][p]) == 1 and "$ref" in oapi["definitions"][v]["properties"][p]:
                            property_type = oapi["definitions"][v]["properties"][p]["$ref"].split("/")[-1].split(".")[-1]
                        else:
                            if "type" in oapi["definitions"][v]["properties"][p]:
                                property_type = go_types[oapi["definitions"][v]["properties"][p]['type']]
                            else:
                                print("no type", oapi["definitions"][v]["properties"][p])
                            type_of_item = ""
                            if property_type == "[]":
                                if "x-enum-reference" in oapi["definitions"][v]["properties"][p]["items"]:
                                    property_array_object_name = go_types[oapi["definitions"][v]["properties"][p]["items"]["type"]]
                                    property_type += property_array_object_name
                                elif "$ref" in oapi["definitions"][v]["properties"][p]["items"]:
                                    property_array_object_name = oapi["definitions"][v]["properties"][p]["items"]["$ref"].split("/")[-1].split(".")[-1]
                                    property_type += property_array_object_name
                                elif "type" in oapi["definitions"][v]["properties"][p]["items"]:
                                    property_array_object_name = go_types[oapi["definitions"][v]["properties"][p]["items"]['type']]
                                    property_type += property_array_object_name

                            if property_type == "struct":
                                if "additionalProperties" in oapi["definitions"][v]["properties"][p]:
                                    if "$ref" in oapi["definitions"][v]["properties"][p]["additionalProperties"]:
                                        property_type = oapi["definitions"][v]["properties"][p]["additionalProperties"]["$ref"].split("/")[-1].split(".")[-1]
                                    elif "type" in oapi["definitions"][v]["properties"][p]["additionalProperties"]:
                                        if oapi["definitions"][v]["properties"][p]["additionalProperties"]["type"] == "array":
                                            property_array_object_name = oapi["definitions"][v]["properties"][p]["additionalProperties"]["items"]["$ref"].split("/")[-1].split(".")[-1]
                                            property_type = "[]"+property_array_object_name
                                        else:
                                            property_type = go_types[oapi["definitions"][v]["properties"][p]["additionalProperties"]["type"]]
                                        if property_type == "struct":
                                            property_type = "interface{}"
                                else:
                                    if "items" in oapi["definitions"][v]["properties"][p]:
                                        if "$ref" in oapi["definitions"][v]["properties"][p]["items"]:
                                            property_type = oapi["definitions"][v]["properties"][p]["items"]["$ref"].split("/")[-1].split(".")[-1]
                                    if "allOf" in oapi["definitions"][v]["properties"][p]:
                                        for entry in oapi["definitions"][v]["properties"][p]["allOf"]:
                                            property_type = "*"+entry["$ref"].split("/")[-1].split(".")[-1]
                        if "[]" in property_type:
                            graphql_type = graph_ql_types["[]"]+f"({property_array_object_name})"
                        elif property_type in graph_ql_types:
                            graphql_type = graph_ql_types[property_type]
                        else:
                            graphql_type = property_type
                        comma = ""
                        if comma_counter > 0:
                            comma = ",\n"
                        schema+=f"       \"{''.join(property_name)}\": &graphql.Field{{\n"
                        schema+=f"            Type: {graphql_type},\n"
                        schema+="        },\n"
                        object+=f"    {''.join(property_name)} {property_type} `json:{quoted_json_arg}`\n"
                        comma_counter=comma_counter+1
                    object+="}\n"
                    schema+="     },\n"
                    schema+="  },\n"
                    schema+=")\n"
                    data_file.write(object)
                    schema_file.write(schema)
                    print(schema)
        data_file.close()
        schema_file.close()
    except KeyError as e:
        logger.exception(e)
    except IOError as e:
        logger.exception(e)
    finally:
        pass

def open_api_v3(data_file):
    try:
        schema_file = open("bungie_schema.go", "w")
        schema_file.write("package main\n\n")
        with open ("bungie_data.go", "w") as data_file:
            data_file.write("package main\n\n")
            for schema_def in oapi["components"]:
                object = ""
                enum = ""
                schema = ""

                if "properties" in oapi["definitions"][schema_def]:
                    comma_counter = 0
                    object_name = schema_def.split(".")[-1]
                    object+=f"type {object_name} struct {{\n"
                    schema+= f"var {object_name} = graphql.NewObject(\n"
                    schema+= "    graphql.ObjectConfig{\n" 
                    schema+= f"        Name: \"{object_name}\",\n"
                    schema+= "        Fields: graphql.Fields{\n"
                    for p in oapi["definitions"][schema_def]["properties"]:
                        # Extract the property name to be used as Struct variable, sett first letter to Upper case
                        property_name = list(p)
                        quoted_json_arg = f'"{p}"'
                        property_name[0] = property_name[0].upper()
                        property_type = ""
                        property_array_object_name = ""

                        if len(oapi["definitions"][schema_def]["properties"][p]) == 1 and "$ref" in oapi["definitions"][schema_def]["properties"][p]:
                            property_type = oapi["definitions"][schema_def]["properties"][p]["$ref"].split("/")[-1].split(".")[-1]
                        else:
                            if "type" in oapi["definitions"][schema_def]["properties"][p]:
                                property_type = go_types[oapi["definitions"][schema_def]["properties"][p]['type']]
                            else:
                                print("no type", oapi["definitions"][schema_def]["properties"][p])
                            type_of_item = ""
                            if property_type == "[]":
                                if "x-enum-reference" in oapi["definitions"][schema_def]["properties"][p]["items"]:
                                    property_array_object_name = go_types[oapi["definitions"][schema_def]["properties"][p]["items"]["type"]]
                                    property_type += property_array_object_name
                                elif "$ref" in oapi["definitions"][schema_def]["properties"][p]["items"]:
                                    property_array_object_name = oapi["definitions"][schema_def]["properties"][p]["items"]["$ref"].split("/")[-1].split(".")[-1]
                                    property_type += property_array_object_name
                                elif "type" in oapi["definitions"][schema_def]["properties"][p]["items"]:
                                    property_array_object_name = go_types[oapi["definitions"][schema_def]["properties"][p]["items"]['type']]
                                    property_type += property_array_object_name

                            if property_type == "struct":
                                if "additionalProperties" in oapi["definitions"][schema_def]["properties"][p]:
                                    if "$ref" in oapi["definitions"][schema_def]["properties"][p]["additionalProperties"]:
                                        property_type = oapi["definitions"][schema_def]["properties"][p]["additionalProperties"]["$ref"].split("/")[-1].split(".")[-1]
                                    elif "type" in oapi["definitions"][schema_def]["properties"][p]["additionalProperties"]:
                                        if oapi["definitions"][schema_def]["properties"][p]["additionalProperties"]["type"] == "array":
                                            property_array_object_name = oapi["definitions"][schema_def]["properties"][p]["additionalProperties"]["items"]["$ref"].split("/")[-1].split(".")[-1]
                                            property_type = "[]"+property_array_object_name
                                        else:
                                            property_type = go_types[oapi["definitions"][schema_def]["properties"][p]["additionalProperties"]["type"]]
                                        if property_type == "struct":
                                            property_type = "interface{}"
                                else:
                                    if "items" in oapi["definitions"][schema_def]["properties"][p]:
                                        if "$ref" in oapi["definitions"][schema_def]["properties"][p]["items"]:
                                            property_type = oapi["definitions"][schema_def]["properties"][p]["items"]["$ref"].split("/")[-1].split(".")[-1]
                                    if "allOf" in oapi["definitions"][schema_def]["properties"][p]:
                                        for entry in oapi["definitions"][schema_def]["properties"][p]["allOf"]:
                                            property_type = "*"+entry["$ref"].split("/")[-1].split(".")[-1]
                        if "[]" in property_type:
                            graphql_type = graph_ql_types["[]"]+f"({property_array_object_name})"
                        elif property_type in graph_ql_types:
                            graphql_type = graph_ql_types[property_type]
                        else:
                            graphql_type = property_type
                        comma = ""
                        if comma_counter > 0:
                            comma = ",\n"
                        schema+=f"       \"{''.join(property_name)}\": &graphql.Field{{\n"
                        schema+=f"            Type: {graphql_type},\n"
                        schema+="        },\n"
                        object+=f"    {''.join(property_name)} {property_type} `json:{quoted_json_arg}`\n"
                        comma_counter=comma_counter+1
                    object+="}\n"
                    schema+="     },\n"
                    schema+="  },\n"
                    schema+=")\n"
                    data_file.write(object)
                    schema_file.write(schema)
                    print(schema)
        data_file.close()
        schema_file.close()
    except KeyError as e:
        logger.exception(e)
    except IOError as e:
        logger.exception(e)
    finally:
        pass    

def extract_struct_name(input):
    try:
        split_res = input.split(".")
        data_blob_name = split_res[-1]
        if "[]" in data_blob_name:
            data_blob_name = data_blob_name.strip("[]") + "List"
        return data_blob_name
    except KeyError and ValueError as e:
        logger.exception(e)

go_types = {
    "integer": "int",
    "string": "string",
    "boolean": "bool",
    "array": "[]",
    "number": "float64",
    "object": "interface{}",
    "dictionary": "interface{}"
}
graph_ql_types = {
    "integer": "graphql.Int",
    "number": "graphql.Float",
    "string": "graphql.String",
    "array": "graphql.NewList",
    "boolean": "graphql.Boolean"
}

json_types = {
    "integer": "json.Number"
}

def convert_types(field_type, language="go", graphql=False, json=False):
    if json:
        if field_type in json_types:
            return json_types[field_type]
        else:
            return go_types[field_type]
    if graphql:
        if field_type in graph_ql_types:
            return graph_ql_types[field_type]
    else:
        return go_types[field_type]

def create_struct(object, object_name):
    object+=f"type {object_name} struct {{\n"
    return object
    
def create_new_graphql_object(schema, object_name):
    schema+= f"var {object_name} = graphql.NewObject(\n"
    schema+= "    graphql.ObjectConfig{\n" 
    schema+= f"        Name: \"{object_name}\",\n"
    schema+= "        Fields: graphql.Fields{\n"
    return schema

def create_aggregate_struct_type(object, property_name, property_type, items):
    field_type = ""
    if "$ref" in items:
        field_type = convert_types(property_type)+items["$ref"].split("/")[-1].split(".")[-1]
    else:
        field_type = convert_types(property_type)+convert_types(items["type"])
    field_name = list(property_name)
    quoted_json_arg = f'"{property_name.lower()}"'
    field_name[0] = field_name[0].upper()
    object+=f"    {''.join(field_name)} {field_type} `json:{quoted_json_arg}`\n"
    return object

def create_basic_struct_type(object, property_name, property_type):
    field_name = list(property_name)
    quoted_json_arg = f'"{property_name.lower()}"'
    field_name[0] = field_name[0].upper()
    object+=f"    {''.join(field_name)} {convert_types(property_type, json=True)} `json:{quoted_json_arg}`\n"
    return object

def create_object_struct_type(object, property_name, ref):
    field_name = list(property_name)
    quoted_json_arg = f'"{property_name.lower()}"'
    field_name[0] = field_name[0].upper()
    if "additionalProperties" in ref:
        if "$ref" in ref["additionalProperties"]:
            property_type = ref["additionalProperties"]["$ref"].split("/")[-1].split(".")[-1]
        elif "type" in ref["additionalProperties"]:
            if ref["additionalProperties"]["type"] == "array":
                if "$ref" in ref["additionalProperties"]:
                    if ref["additionalProperties"]["type"] == "object":
                        ref_type = "dictionary"
                    else:
                        ref_type = ref["additionalProperties"]["type"]
                    property_type = convert_types(ref_type)+ref["additionalProperties"]["$ref"].split("/")[-1].split(".")[-1]
                elif "items" in ref["additionalProperties"]:
                    if ref["additionalProperties"]["type"] == "object":
                        ref_type = "dictionary"
                    else:
                        ref_type = ref["additionalProperties"]["type"]
                    property_type = convert_types(ref_type)+ref["additionalProperties"]["items"]["$ref"].split("/")[-1].split(".")[-1]               
            else:
                property_type = convert_types(ref["additionalProperties"]["type"])
        object+=f"    {''.join(field_name)} {property_type} `json:{quoted_json_arg}`\n"
    elif "allOf" in ref:
        property_type = ref["allOf"][-1]["$ref"].split("/")[-1].split(".")[-1]

        object+=f"    {''.join(field_name)} {property_type} `json:{quoted_json_arg}`\n"
    else:
        property_type = ref.split("/")[-1].split(".")[-1]
        object+=f"    {''.join(field_name)} {property_type} `json:{quoted_json_arg}`\n"
    return object

def close_struct(object):
    object+="}\n"
    return object

def close_schema(object):
    object+="     },\n  },\n)\n\n"
    return object

def create_basic_graphql_schema_type(schema, property_name, property_type):
    field_type = convert_types(property_type, graphql=True)
    field_name = list(property_name)
    field_name[0] = field_name[0].upper()
    if field_type == None:
        field_type = list(property_name)
        field_type[0] = field_name[0].upper()
    schema+=f"       \"{''.join(field_name)}\": &graphql.Field{{\n"
    schema+=f"            Type: {''.join(field_type)},\n"
    schema+="        },\n"
    return schema

def create_object_graphql_schema_type(object, property_name, ref):
    field_name = list(property_name)
    quoted_json_arg = f'"{property_name}"'
    field_name[0] = field_name[0].upper()
    if "additionalProperties" in ref:
        if "$ref" in ref["additionalProperties"]:
            field_type = ref["additionalProperties"]["$ref"].split("/")[-1].split(".")[-1]
        elif "type" in ref["additionalProperties"]:
            if ref["additionalProperties"]["type"] == "array":
                if "$ref" in ref["additionalProperties"]:
                    field_type = convert_types(ref["additionalProperties"]["type"])+ref["additionalProperties"]["$ref"].split("/")[-1].split(".")[-1]
                elif "items" in ref["additionalProperties"]:
                    field_type = convert_types(ref["additionalProperties"]["type"])+ref["additionalProperties"]["items"]["$ref"].split("/")[-1].split(".")[-1]               
            else:
                field_type = convert_types(ref["additionalProperties"]["type"])
    elif "allOf" in ref:
        field_type = ref["allOf"][-1]["$ref"].split("/")[-1].split(".")[-1]
    else:
        field_type = ref.split("/")[-1].split(".")[-1]
    object+=f"       \"{''.join(field_name)}\": &graphql.Field{{\n"
    object+=f"            Type: {''.join(field_type)},\n"
    object+="        },\n"
    return object

def create_aggregate_graphql_schema_type(schema, property_name, property_type, items):
    field_type = ""
    field_name = list(property_name)
    field_name[0] = field_name[0].upper()
    if "$ref" in items:
        field_type = items["$ref"].split("/")[-1].split(".")[-1]
    elif "type" in items:
        field_type = convert_types(items["type"], graphql=True)
    else:
        field_type = convert_types(items["type"], graphql=True)
    schema+=f"       \"{''.join(field_name)}\": &graphql.Field{{\n"
    schema+=f"            Type: graphql.NewList({field_type}),\n"
    schema+="        },\n"
    return schema

def gen_fields(oapi):
    field = "var (\n    Fields = graphql.Fields{\n"
    resolvers = ""
    for endpoints in oapi["paths"]:
        array = False
        if "get" in oapi["paths"][endpoints]:
            args = oapi["paths"][endpoints]["get"]["responses"]["200"]["$ref"].split("/")
            response = oapi[args[1]][args[2]][args[3]]["content"]["application/json"]["schema"]["properties"]["Response"]
            config_name = oapi["paths"][endpoints]["get"]["operationId"]
            field_id = oapi["paths"][endpoints]["get"]["operationId"].split(".")[-1]
            field += f"        \"{field_id.lower()}\": &graphql.Field{{\n"
            if "$ref" in response:
                go_type =  "schema."+response["$ref"].split("/")[-1].split(".")[-1]
                field_type = go_type
            elif "type" in response:
                if response["type"] == "object":
                    if "additionalProperties" in response:
                        additional = response["additionalProperties"]
                        if "additionalProperties" in additional:
                            nested_additional = additional["additionalProperties"]
                            go_type = "schema."+nested_additional["$ref"].split("/")[-1].split(".")[-1]
                            field_type = go_type
                        elif "$ref" in additional:
                            go_type = "schema."+additional["$ref"].split("/")[-1].split(".")[-1]
                            field_type = go_type
                        else:
                            go_type = convert_types(additional["type"])
                            field_type =  convert_types(additional["type"], graphql=True)
                    else:
                        go_type = convert_types(response["type"])
                        field_type =  convert_types(response["type"], graphql=True)
                elif response["type"] == "array":
                    array = True
                    go_type = "schema."+response["items"]["$ref"].split("/")[-1].split(".")[-1]
                    field_type = "graphql.NewList("+go_type+")"
                else:
                    go_type = convert_types(response["type"])
                    field_type =  convert_types(response["type"], graphql=True)
            if field_type == None:
                field_type = "nil"
            field += f"            Type: {field_type},\n"
        elif "post" in oapi["paths"][endpoints]:
            args = oapi["paths"][endpoints]["post"]["responses"]["200"]["$ref"].split("/")
            response = oapi[args[1]][args[2]][args[3]]["content"]["application/json"]["schema"]["properties"]["Response"]
            config_name = oapi["paths"][endpoints]["post"]["operationId"]
            field_id = oapi["paths"][endpoints]["post"]["operationId"].split(".")[-1]
            field += f"        \"{field_id.lower()}\": &graphql.Field{{\n"
            if "$ref" in response:
                go_type = "schema."+response["$ref"].split("/")[-1].split(".")[-1]
                field_type = go_type
            elif "type" in response:
                if response["type"] == "object":
                    if "additionalProperties" in response:
                        additional = response["additionalProperties"]
                        if "additionalProperties" in additional:
                            nested_additional = additional["additionalProperties"]
                            go_type = "schema."+nested_additional["$ref"].split("/")[-1].split(".")[-1]
                        elif "$ref" in additional:
                            go_type = "schema."+additional["$ref"].split("/")[-1].split(".")[-1]
                            field_type = go_type
                        else:
                            go_type = convert_types(additional["type"])
                            field_type =  convert_types(additional["type"], graphql=True)
                    else:
                        go_type = convert_types(response["type"])
                        field_type =  convert_types(response["type"], graphql=True)
                elif response["type"] == "array":
                    array = True
                    go_type = "schema."+response["items"]["$ref"].split("/")[-1].split(".")[-1]
                    field_type =  "graphql.NewList("+go_type+")"
                else:
                    go_type = convert_types(response["type"])
                    field_type =  convert_types(response["type"], graphql=True)
            field += f"            Type: {field_type},\n"
   
        matches = list((x for x in endpoints.split("/") if "{" in x))
        input_args = []
        if len(matches):
            field += "            Args: graphql.FieldConfigArgument{\n"
            for i in matches:
                field += "                \""+i.strip("{").strip("}")+"\"" +  ": &graphql.ArgumentConfig{\n" + "                        Type: graphql.String,\n" + "                },\n"
                input_args.append(i.strip("{").strip("}"))

            field += "            },\n"
        func_return = "(interface{}, error)"
        field += f"            Resolve: func(p graphql.ResolveParams) {func_return} {{\n"
        field += f"                return resolvers.{field_id}(p)\n            }},\n"
        field += "        },\n\n"
        resolvers += gen_resolver_function(resolvers, input_args, field_id, config_name, field_type, go_type, array=array)
    field += "    }\n"
    field += ")\n"
    write_fields(field)
    write_resolvers(resolvers)

def gen_resolver_function(resolver, arguments, func_name, config_name, return_type, go_type, array=False):
    resolver = f"func {func_name}(p graphql.ResolveParams) (interface{{}}, error) {{\n"
    if config_name.split(".")[0]:
        resolver += f"    url := fmt.Sprintf(configuration.API.{config_name}.Endpoint"
    else:
        single_name = config_name.split(".")[-1]
        resolver += f"    url := fmt.Sprintf(configuration.API.{single_name}.Endpoint"

    if arguments:
        for arg in arguments:
            resolver +=  f", p.Args[\"{arg}\"].(string)"
    resolver += ")\n"
    # if return_type == "nil":
    #     go_type
    return_val = go_type.replace("schema.", "data.")
    resolver += f"    var return_val {return_val}\n"
    if config_name.split(".")[0]:
        resolver += f"    ret, err := base_resolver(p, url, configuration.API.{config_name}.Method, &return_val)\n"
    else:
        single_name = config_name.split(".")[-1]
        resolver += f"    ret, err := base_resolver(p, url, configuration.API.{single_name}.Method, &return_val)\n"
    resolver += "    if err != nil {\n        log.Println(err)\n    }\n"
    resolver += "    return ret, nil\n}\n"
    return resolver
    # print(arguments, func_name, config_name, return_type, go_type, array)

def parse_responses(oapi):
    pass

def gen_response_structs(oapi):
    try:
        object = ""
        for schema_def in oapi["components"]["schemas"]:
            schema_object_name = extract_struct_name(schema_def)
            schema_object = oapi["components"]["schemas"][schema_def]
            schema_type = schema_object["type"]

            object = create_struct(object, schema_object_name)
            if schema_type == "object" and "properties" in schema_object:
                for property in schema_object["properties"]:
                    if "type" in schema_object["properties"][property]:
                        property_type = schema_object["properties"][property]["type"]
                        if property_type == "array":
                            items = schema_object["properties"][property]["items"]
                            object = create_aggregate_struct_type(object, property, property_type, items)
                        if property_type == "object":
                            object = create_object_struct_type(object, property, schema_object["properties"][property])
                        elif property_type != "array" and property_type != "object":
                            object = create_basic_struct_type(object, property, property_type)
                    elif "$ref" in schema_object["properties"][property]:
                        object = create_object_struct_type(object, property, schema_object["properties"][property]["$ref"])
                    else:
                        print(schema_object["properties"][property])
                object = close_struct(object)
            elif "enum" in schema_object:
                property_type = schema_object["type"]
                object = create_basic_struct_type(object, schema_object_name, property_type)
                object = close_struct(object)
            else:
                if "items" in schema_object:
                    property_type = schema_object["items"]["type"]
                else:
                    property_type = "dictionary"
                object = create_basic_struct_type(object, schema_object_name, property_type)
                object = close_struct(object)
        write_structs(object)
    except KeyError as e:
        logger.exception(e)
    except IOError as e:
        logger.exception(e)
    finally:
        pass  

def gen_response_graphql_objects(oapi):
    try:
        object = ""
        for schema_def in oapi["components"]["schemas"]:
            schema_object_name = extract_struct_name(schema_def)
            schema_object = oapi["components"]["schemas"][schema_def]
            schema_type = schema_object["type"]
            object = create_new_graphql_object(object, schema_object_name)
            if schema_type == "object" and "properties" in schema_object:
                for property in schema_object["properties"]:
                    if "x-dictionary-key" in schema_object["properties"][property]:
                        property_type = schema_object["properties"][property]["x-dictionary-key"]["type"]
                        object = create_basic_graphql_schema_type(object, property, property_type) 
                    elif "type" in schema_object["properties"][property]:
                        property_type = schema_object["properties"][property]["type"]
                        if property_type == "array":
                            items = schema_object["properties"][property]["items"]
                            object = create_aggregate_graphql_schema_type(object, property, property_type, items)
                        if property_type == "object":
                            object = create_object_graphql_schema_type(object, property, schema_object["properties"][property])
                        elif property_type != "array" and property_type != "object":
                            object = create_basic_graphql_schema_type(object, property, property_type)
                    elif "$ref" in schema_object["properties"][property]:
                        object = create_object_graphql_schema_type(object, property, schema_object["properties"][property]["$ref"])
                    elif "$ref" in schema_object:
                        object = create_object_graphql_schema_type(object, property, schema_object["properties"][property]["$ref"])
                       
                    else:
                        object = create_basic_graphql_schema_type(object, property, property_type)
                object = close_schema(object)
            elif "enum" in schema_object:
                property_type = schema_object["type"]
                object = create_basic_graphql_schema_type(object, schema_object_name, property_type)
                object = close_schema(object)
            elif schema_type == "array":
                items = schema_object["items"]
                property = schema_object["items"]["x-enum-reference"]["$ref"].split("/")[-1]
                object = create_aggregate_graphql_schema_type(object, property, property_type, items)
                object = close_schema(object)
            else:
                print(schema_object_name)
                object = close_schema(object)
        write_schema(object)
    except KeyError as e:
        logger.exception(e)
    except IOError as e:
        logger.exception(e)
    finally:
        pass  

def gen_reponse_graphql_fields():
    pass

def write_schema(object):
    create_dir("schema")
    schema = "package schema\n"
    schema += "import (\n\"github.com/graphql-go/graphql\"\n)\n\n"
    schema += object
    write_out_code_file(schema, "schema/bungie_schema.go")

def write_structs(object):
    create_dir("data")
    structs = "package data\n\n"
    structs += "import (\n\"encoding/json\"\n)\n\n"
    structs += object
    write_out_code_file(structs, "data/bungie_data.go")

def write_fields(object):
    create_dir("fields")
    fields = "package fields\n\n"
    fields += "import (\n\"github.com/graphql-go/graphql\"\n\"github.com/primarchhorus/go-graphql-bungie/schema\"\n\"github.com/primarchhorus/go-graphql-bungie/resolvers\"\n)\n\n"
    fields += object
    write_out_code_file(fields, "fields/bungie_fields.go")

def write_resolvers(object):
    create_dir("resolvers")
    fields = "package resolvers\n\n"
    fields += "import (\n\"github.com/graphql-go/graphql\"\n\"github.com/primarchhorus/go-graphql-bungie/data\"\n\"github.com/primarchhorus/go-graphql-bungie/configuration\"\n\"fmt\"\n\"log\"\n)\n\n"
    fields += object
    write_out_code_file(fields, "resolvers/bungie_resolvers.go")

def write_out_code_file(object, destination):
    try:
        output = destination
        with open(output, "w") as ofile:
            ofile.write(object)
        ofile.close()
    except IOError as e:
        logger.exception(e)

def gen_response_from_schema(object, schema_def, name):
    schema_object_name = name
    schema_object = schema_def
    schema_type = schema_object["type"]

    object = create_struct(object, schema_object_name)
    if schema_type == "object" and "properties" in schema_object:
        for property in schema_object["properties"]:
            if "type" in schema_object["properties"][property]:
                property_type = schema_object["properties"][property]["type"]
                if property_type == "array":
                    items = schema_object["properties"][property]["items"]
                    object = create_aggregate_struct_type(object, property, property_type, items)
                if property_type == "object":
                    object = create_object_struct_type(object, property, schema_object["properties"][property])
                elif property_type != "array" and property_type != "object":
                    object = create_basic_struct_type(object, property, property_type)
            elif "$ref" in schema_object["properties"][property]:
                object = create_object_struct_type(object, property, schema_object["properties"][property]["$ref"])
            else:
                print(schema_object["properties"][property])
        object = close_struct(object)
    elif "enum" in schema_object:
        property_type = schema_object["type"]
        object = create_basic_struct_type(object, schema_object_name, property_type)
        object = close_struct(object)
    print(object)

try:
    open_api_file = sys.argv[1]
    with open(open_api_file, "r") as jfile:
        oapi = json.load(jfile)
    jfile.close()
except IOError as e:
    logger.exception(e)

gen_response_graphql_objects(oapi)
gen_response_structs(oapi)
gen_fields(oapi)
