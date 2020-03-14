# donatetohospitals-web
The main webapp project for [donatetohospitals.com](https://donatetohospitals.com)

DonateToHospitals.com will be where normal people who have stocked excess medical supplies far beyond their needs will be able to donate those supplies to hospitals that need them to save lives while facing covid19

Architecture
-
- Will be a single server-rendered golang webapp to start
- Users will submit information about what they have
- We will manually match users to hospitals on our end

Collaboration
- 
- Github for code
- Github project board and issues for project management
- Discord for chatter

Stack
- 
- golang for language
- http/template for templates. Potentially `fasttemplate` if we need it
- `chi`
- postgres for db
- home-rolled auth/sessions
- aws s3 for media
- digitalocean for hosting
- nginx
- cloudflare
- fastmail

Application
-

Schema
-
- `USERS` - a user
- `HOSPITALS` - These will be manually entered by us after corresponding with the hospital via email. 
We must identify that they're legit, what supplies they need, and where/how they'd like it to be sent or dropped off 
- `STOCKPILES` - These will be batches of `SUPPLIES` a user has
```
ID
USER_ID: foreign key
IMAGE_URL: string
```
- `ITEMS` - This is an individual item in a stockpile (gloves, masks, etc)
```
COUNT: int
TITLE: string
STOCKPILE_ID: foreign key
```

Landing page
 -
- Jumbotron explains what it is. Name and single sentence should explain it

- Two big buttons
1. I HAVE STOCKPILED SUPPLIES - This goes to flow (A)
2. I AM HOSPITAL STAFF - this goes to flow (B)

- A list of the most recent 5-10 submitted `STOCKPILES` (looks like a twitter feed)

A1 - User signup. username, Email, password, state. If international go to (A2), if not go to (A3)

A2 - Form where the user enters their email and country to be notified if donatetohospitals.com is 
added to their country

A3 - After signup the logged in user goes to a UI for submitting a `STOCKPILE` of `ITEMS`

B1 - The hospital is directed to the contact email and guided on what information we need to start getting them supplies 