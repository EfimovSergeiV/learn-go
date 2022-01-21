import json

with open('upload-data.json', 'r') as file:
    data = json.load(file)


    count = 0

    for cursor in data:
        count += 1
        print(count, cursor)
