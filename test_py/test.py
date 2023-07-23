import argparse
import sys
import requests

parser = argparse.ArgumentParser()
parser.add_argument("-w", "--word")
args = parser.parse_args()
def main():
    word = args.word.split()
    api_url = 'https://api.dictionaryapi.dev/api/v2/entries/en/'
    api_url = api_url + str(word[0])
    response = requests.get(api_url)
    response = response.json()
    definition = response[0]
    definition = definition["meanings"]
    definition = definition[0]
    definition = definition["definitions"]
    definition = definition[0]
    definition = definition["definition"]
    sys.stdout.write(str(definition))

if __name__ == "__main__":
    main()