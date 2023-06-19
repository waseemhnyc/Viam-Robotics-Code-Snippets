import os
import json


def get_json_from_file(file_path):
    with open(file_path, 'r') as file:
        return json.load(file)


def get_body_from_file(file_path):
    with open(file_path, 'r') as file:
        lines = file.readlines()

    body_list = []
    for line in lines:
        body_list.append(line.rstrip('\n'))

    return body_list


def write_output_to_file(output, output_file_name):
    with open(output_file_name, 'w') as file:
        json.dump(output, file, indent=4)


def create_snippets():
    combined_data = {}

    for dirpath, dirnames, filenames in os.walk('./python'):
        if 'details.json' in filenames and 'snippet.py' in filenames:
            # prefix = dirpath.strip('./')
            json_data = get_json_from_file(
                os.path.join(dirpath, 'details.json'))
            body = get_body_from_file(os.path.join(dirpath, 'snippet.py'))

            combined_data[json_data.get('prefix')] = {
                'prefix': json_data.get('prefix'),
                'body': body,
                'description': json_data.get('description')
            }

    write_output_to_file(combined_data, 'snippets/snippets.code-snippets')


if __name__ == '__main__':
    create_snippets()
