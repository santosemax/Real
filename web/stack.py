#!/usr/local/bin/python3
from stackapi import StackAPI
from datetime import datetime

query = "Black screen after install ubuntu"

SITE = StackAPI('stackoverflow')

# Parameters for search
#SITE.page_size = 10
#SITE.max_pages = 10

# Get search 
questions = SITE.fetch('search/advanced', order='desc', sort='relevance', q=query, site='stackoverflow')

print(questions)
