#!/usr/local/bin/python3
from stackapi import StackAPI
from datetime import datetime




query = "How to turn on computer"
SITE = StackAPI('stackoverflow')

# Parameters for search
SITE.page_size = 10
SITE.max_pages = 1


# Get search 
questions = SITE.fetch('search/advanced', title=query)

print(questions)

for q in questions:
    print(q)


#
#import requests
#from bs4 import BeautifulSoup
#
#res = requests.get("https://stackoverflow.com/search?q=How+to+turn+off+the+computer")
#soup = BeautifulSoup(res.text, "html.parser")
#questions = soup.select(".question-summary")
#
## Print out questions
#for que in questions:
#    q = que.select_one('.question-hyperlink').getText()
#    print(q)
#
