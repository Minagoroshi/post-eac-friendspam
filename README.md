# VRChat FriendSpam
Yet another simple friendspam script

## Preface
I wish to preface this by saying, that I am making this --> SIMPLE <-- script public because it seems that immature children on VRChat have been collecting ego's by making other simple scripts similar to this. I encourage any of these children reading this to go learn about actual software engineering, and pursue it as a legit career and learning opportunity instead of sitting on VRChat collecting ego's over HTTP requests.
(Kuro can be excused he actually knows more outside of simple shit like this)

## Compiling
1. Install [Go](https://go.dev/dl/)
2. Clone or Download this repo, extracting the contents if necessary 
3. Open a terminal and change directory into the project
4. `go build .`
5. 😁

## Proxies
- Proxies are optional but reccomended as they will prevent your main account from being banned, and you can send more than 4 requests at a time

## Accounts
- Accounts should be structured as a wordlist, meaning each account string should be on a new line
- Format them as username:password:twofactorcode
- Getting twofactorcode is as easy as going to vrchat.com, logging in, opening dev-tools (Shift + F12), going to network, refresh the page, and look at the cookies under any VRC bound request (EX: https://vrchat.com/home)

## Files 
- Make sure you create a file named "accounts" and populate it with your accounts
- proxies are optional but the file should be named "proxies"