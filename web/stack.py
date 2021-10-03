#!/usr/local/bin/python3

#from stackapi import StackAPI
#from datetime import datetime
#SITE = StackAPI('askubuntu')
#
#
## Parameters for search
#SITE.page_size = 10
#SITE.max_pages = 1
#
#
## Get search 
#query = SITE.fetch('questions', fromdate=datetime(2020, 1, 1), todate=datetime(2021, 11, 3), min=10, sort='votes')
#comments = SITE.fetch('comments')
#
#print(query)


import requests
from bs4 import BeautifulSoup

res = requests.get("https://stackoverflow.com/questions")

soup = BeautifulSoup(res.text, "html.parser")


questions = soup.select(".question-summary")

print(questions[0].attrs)
