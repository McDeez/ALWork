# Simple SQL Injection Filter
### By Nathan Giardina


### High Level

...This super simple SQL Injection Filter uses compiled Regex read in from a JSON file to provide a very simplistic two stage inspection system for SQL Vulnerabilities..
Currently the only vulnerabilities designed and loaded are focused around SQL injection in the Query string, however it would be possible to get any other vulnerability 
dectection for the QueryString as well. 

### Detection 

The two stages of the filter are focused around:
1. QueryStrings that cannot contain malicious code are allowed to pass
... Zero Query String length
... URIs that contain just alpha numeric and forward slashes
2. Deeper inspection of all other packets for detecting:
... Single quote text block based SQL Injection attacks (name' OR 'a'='a')
... SQL Statement detection (ALTER, CREATE, DELETE, etc)
... Statement Breaks (a simple semi-colon)

The main assumption here is that firstly the quick regex's will take a crack at the Query String. If it matches one of those Regexes it will pass on the query as 
those are marked as "safe" queries, ones that could not contain a SQL Injection vulnerability. Depending on the amount and type of traffic the webserver would see
this might save quite some processing time, however it would need to be adjusted depending on the quantities of traffic seen. Next, it will try the more complex regexes 
for the vulnerabilities.  Matching one of those will result in a HTTP Error code being sent back to the request. In a more robust solution, I would provide a bit more
low level inspection of the request, as well as a more robust body and payload checking part that would provide security for POST, and PUT requests as well. The 
modifications there are slight and could be done quite easily.

### Considerations for Longer Term

In Making this a true solution and not one that has just been developed for 4 hours, I would load in the regexes from a database (perhaps Mongo, or MySQL,), and use a 
notification system to ensure that regex vulnerability signatures added to the database could be added and compiled into the running product without any downtime. 
Considerations for a multi-server proxy would be taken into account with this database as well. Auditing attacks and providing firewalls with possible compromised IPs 
and hosts would also be considered. Currently the logging is fairly simple as it just outputs all logs as the same level. Some unit testing might be required on the
filter.go file to provide some level of assurance for the main filtering algorythims. Currently I'm relying on integration level testing broken as the project is so 
small I cannot justify a true unit level test setup, especially before an API is setup. As such a real API would have to be designed and layed out for this proxy 
to allow log extraction, reporting, and auditing information.


### Directory Structure

ROOT
  integrationTest.sh        //Integration tests for positive and negative cases
  src                       //All the Go source code is in here
    config
      config.go             //Congfiguration 
    direct
      direct.go             //Directs requests
    filter
      filter.go             //Filters requests
    filters.json            //The signatures
    config.json             //The config
    main.go                 //Main code to run it
  webServer                 //Simple NodeJS WebServer
    package.json
    server.js



### How to Run 

Running this project is simple. Prereqs:

1. GoLang is installed
2. NodeJS is installed (if you are running the build in webserver

Steps:

1. cd into ROOT/webServer
2. Run 'npm install' (no quotes)
3. Run 'npm start'
...This will start the built in webServer that allow for fairly simple testing
4. In a new window, cd into ROOT/src
5. Run 'go build'
6. Execute the compiled binary (by default it's listed as "src")

At this point the Proxy and the WebServer are up and running. Next you can run the integrationTests to be able to see it accept and reject requests:
7. In a new window cd into ROOT/
8. Edit the integrationTest.sh file to change the host, and port to correct values (or leave them alone if you are running the default setup)
9. Run './integrationTest.sh
