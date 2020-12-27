const { ChakraLoaderPlugin } = require('chakra-loader')
const AddAssetHtmlPlugin = require('add-asset-html-webpack-plugin')
const express = require('express')

module.exports = {
  configureWebpack: {
    plugins: [
      new ChakraLoaderPlugin(),
      new AddAssetHtmlPlugin({ filepath: require.resolve('./src/misc/wasm_exec.js') })
    ],
    // We have to make sure that the development server is also serving the
    // built WASM binary.
    devServer: {
      before: function(app, server, compiler) {
        app.all('/cohorse.wasm', function (req, res, next) {
          res.type("application/wasm");
          next()
        })
        app.use(express.static('./public/wasm'))
      }
    }
  }
}
