# go-webview
create desktop apps using webview in go


#installation

`go get github.com/zserge/webview`

-> Create a file and import this webview and there are lot of functionalities available like Open to open a webview with url


Open Terminal and enter

# MacOS uses app bundles for GUI apps
$ mkdir -p learning.app/Contents/MacOS
$ go build -o learning.app/Contents/MacOS/learning
$ open learning.app # Or click on the app in Finder
