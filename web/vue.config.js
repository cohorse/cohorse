// Copyright 2020 The Cohorse Author
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
