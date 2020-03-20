# donatetohospitals-web
The main webapp project for [donatetohospitals.com](https://donatetohospitals.com)

DonateToHospitals.com will be where normal people who have stocked excess medical supplies beyond their needs will be able to donate those supplies to hospitals that need them to save lives while facing covid19

Architecture
-
- Will be a single server-rendered golang webapp to start
- `SUPPLIERS` (everyday people probably) will submit information about what supplies they have
- No authentication or sessions. People simply input their email address along with what items they have and we email them when there is a match
- We will manually match users to hospitals on our end for now

Collaboration
- 
- Github for code
- [Github project board](https://github.com/donatetohospitals/donatetohospitals-web/projects/1) and issues for project management
- [Discord for chatter](https://discord.gg/tbAmwZR) This is the best starting place to see how you can chip in
- Volunteer project page at [https://helpwithcovid.com/projects/56](https://helpwithcovid.com/projects/56)

Stack
- 
- golang for language
- http/template for templates
- `chi`
- postgres (gorm) for db
- aws s3 for media
- digitalocean for hosting
- nginx
- cloudflare
- fastmail for email UI

Application
-

Schema
-
See the structs for `SUPPLIERS` and `ITEMS` in the code

Hospitals won't even be in the db for now. We will correspond with them via email and do matching with `SUPPLIERS` by state manually. 
We must identify that they're legit, what supplies they need, and where/how they'd like it to be sent or dropped off 

To Create a test supplier with items in the DB
-
Note that you'll have to change the email each time if you want to post multiple
```
curl -i -X POST \
   -H "Content-Type:application/json" \
   -d \
'{
	"email": "awesome8@gmail.com",
	"geo": "MN",
	"ImageUrl": "https://donatetohospitalsdev.s3-us-west-1.amazonaws.com/supply-placeholder.png",
	"items" : [
		{"Name": "asdf3", "Count": 10}, {"Name": "asdf4", "Count": 20}	
	]
}' \
 'http://localhost:9990/suppliers'

```

Landing page
 -
- Jumbotron explains what it is. Name and single sentence should explain it

- Two big buttons
1. I HAVE SUPPLIES - This goes to flow (A)
2. I AM HOSPITAL STAFF - this goes to flow (B)

- A list of the most recent 5-10 submitted `SUPPLIERS` (looks like a twitter feed)

A1 - supplier submission form. Includes email, state. If international go to (A2)

A2 - Form where the user enters their email and country to be notified if donatetohospitals.com is 
added to their country

B1 - The hospital is directed to the contact email and guided on what information we need to start getting them supplies 