# walnut db design | designed with:
# https://app.quickdatabasediagrams.com/#/

##
User
-
ID PK int
Username varchar(45) INDEX
Password varchar(45) INDEX
Role int FK - Role.ID INDEX
LastLoggedin NULL date
Token int FK - Token.ID

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
Title varchar(45)
Token int FK - Token.ID

##
Token
-
ID PK int
Title varchar(45)
ExpiresAt NULL date

##
Role
-
ID PK int
Title varchar(45)

##
Module
-
ID PK int
Url NULL varchar(200)
Path NULL varchar(200)
Kind int FK - ModuleKind.ID
Visibility int FK - ModuleVisibility.ID
Creator int FK - User.ID
CreatedOn NULL date
LastFetched NULL date
LastUsaged NULL date

##
ModuleKind
-
ID PK int
Title varchar(45)

##
ModuleVisibility
-
ID PK int
Title varchar(45)




