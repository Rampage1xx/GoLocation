"use strict";
exports.__esModule = true;
var express = require('express');
var path = require('path');
exports.app = express();
exports.app.use(express.static('dist'));
exports.app.use('*', function (req, res) {
    res.sendFile(path.join(__dirname + '/dist/index.html'));
});
exports.server = exports.app.listen(8000, function () {
    console.log('server running on 8000');
});
