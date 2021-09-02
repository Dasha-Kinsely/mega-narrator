## Please Adhere to the Following Naming Convention for Any Context & Env & Form Variables:

authenticated current user inside gin.Context: "current"
jwt inside gin.Context: "claim"

any env variables format: [ CAPS_LOCK_WITH_UNDERSCORES ]

password passed via forms: "password"
username passed via forms: "username"
files passed via forms: "attachment"

## The Following Conventions are Optional:

directory names: [ onesingleword ] (no cap on first word, no camels, no underscore, no dashes)
file names: [ no cap on first word ] & [ CamelOnNextWords ]
variables: [ no cap on first word ] & [ CamelOnNextWords ]
function and methods: [ cap on first word ] & [ CamelOnNextWords ]