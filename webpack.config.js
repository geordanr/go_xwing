/* global __dirname, module, require */
var path = require('path');

module.exports = {
    entry: path.resolve(__dirname, './web/src/app.js'),
    output: {
        path: path.resolve(__dirname, './web/public/js'),
        filename: 'bundle.js'
    },
    module: {
        loaders: [
            {
                test: /\.jsx?$/,
                exclude: /node_modules/,
                loader: 'babel-loader',
            },
            {
                test: /\.sass$/,
                loaders: ['style', 'css', 'sass'],
            }
        ]
    },
    devtool: '#inline-source-map',
};
