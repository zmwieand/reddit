from optparse import OptionParser
from lxml import html
import requests

REDDIT_URL = "https://reddit.com/"
SUFFIXES = {}


def print_headlines(one, two, three, four):
    url = REDDIT_URL + SUFFIXES[str(one)]
    
    page = requests.get(url)
    tree = html.fromstring(page.content)
    
    elems = tree.find_class('outbound')
    for i, e in enumerate(elems):
        print("[" + str(i)+ "]:", e.text_content())



def main():
    parser = OptionParser()

    with open('reddit.conf') as conf_file:
        for line in conf_file:
            options = line.strip('\n').split(' ')
            flag_name = options[0]

            parser.add_option("-" + flag_name,
                              action="callback",
                              callback=print_headlines,
                              help="don't print status messages to stdout")

            SUFFIXES["-" + flag_name] = options[1]

    args = parser.parse_args()

if __name__ == "__main__":
    main()
