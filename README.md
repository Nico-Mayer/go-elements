# ⚛️ Go-Elements API
Go-Elements is a simple and efficient REST API built with Golang, providing comprehensive information about the periodic table's elements through a variety of endpoints.

# Features
- Retrieve detailed information about elements based on atomic number or symbol
- Access elements by period, group, block, or category
- Search elements by name or symbol
- Easy-to-use RESTful design

# API Usage
The API provides the following endpoints:

- GET `/elements` - Retrieve a list of all elements with their basic information
- GET `/elements/atomic_number/{atomic_number}` - Retrieve detailed information about an element based on its atomic number
- GET `/elements/symbol/{symbol}` - Retrieve detailed information about an element based on its symbol
- GET `/elements/period/{period}` - Retrieve a list of elements in a specific period
- GET `/elements/group/{group}` - Retrieve a list of elements in a specific group
- GET `/elements/block/{block}` - Retrieve a list of elements in a specific block (s, p, d, or f block)
- GET `/elements/category/{category}` - Retrieve a list of elements in a specific category (e.g., alkali metal, noble gas, etc.)
- GET `/elements/search?query={query}` - Search for elements by name or symbol and return a list of matching elements
