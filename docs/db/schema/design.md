# walnut db design | designed with:
# https://app.quickdatabasediagrams.com/#/

##
User
-
ID PK int
Username string INDEX
Password string
Role int FK >- Role.ID
LastLoggedin NULL date

##
UserTeam
-
ID PK int
Team int FK >- Team.ID
User int FK >- User.ID

##
Team
-
ID PK int
Title string

##
Role
-
ID PK int
Title string

##
Module
-
ID PK int
Kind int FK >- Kind.ID
Creator int FK >- User.ID

##
Kind
-
ID PK int
Title string


