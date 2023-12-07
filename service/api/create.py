import os

listee = ["unfollowUser",
"banUser",
"unbanUser",
"getUserProfile",
"getMyStream",
"likePhoto",
"unlikePhoto",
"commentPhoto",
"uncommentPhoto",
"deletePhoto",
"doLogin", 
"setMyUserName",
"uploadPhoto",
"followUser"
]


for item in listee:
    name = item+".go"
    with open(name, "w") as fl:
        print(name)