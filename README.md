# ReadItLater

(readitlater in PowerShell 7)[readitlater.png]

## Built With
Built in go. I chose it because it's fast enough to be snappy and simple enough to develop in one evening. This project is a part of my series 'built it in one evening' during the summer.

## How to use
* `readitlater add <name> <content>` - adds a note with a name and an invisible timestamp
* `readitlater get <patter>` - matches all the notes' names to the parretn as a regex and outputs the ones that match
* `readitlater search` - the same as `readitlater get` but interactive and fuzzy
