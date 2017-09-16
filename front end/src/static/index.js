// Import stylesheets
require('./styles/main.sass')

// Inject bundled Elm app
var Elm = require('../elm/Main')
Elm.Main.embed(document.getElementById('main'))
