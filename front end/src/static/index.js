// Import stylesheets
require('./styles/main.sass')

// Inject bundled Elm app
var Elm = require('../elm/Main')
var app = Elm.Main.fullscreen()

app.ports.download.subscribe(function(filePath) {
    downloadFile(filePath)
})

function getFileName(str) {
    return str.substring(str.lastIndexOf("/") + 1, str.length)
}

function downloadFile(filePath) {
    var req = new XMLHttpRequest()
    var url = "https://s3-us-west-2.amazonaws.com/czrciwipidotrtoq/" + filePath    
    req.open("GET", url, true)
    req.responseType = "blob"
    req.onload = function (event) {
        var blob = req.response
        var fileName = getFileName(filePath)
        var link = document.createElement('a')
        link.href = window.URL.createObjectURL(blob)
        link.download = fileName
        link.click()
    }
    req.send()
}
