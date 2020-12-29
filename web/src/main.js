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

import Vue from 'vue'
import App from './App.vue'
import Chakra from '@chakra-ui/vue'
import cohorseTheme from './misc/theme.js'
import store from './store'

// Expose the Vuex store in the global namespace so that Go can easily find and
// access it.
window.global.store = store

// Try to load the WASM binary and instantiate it.
const go = new window.global.Go()
WebAssembly.instantiateStreaming(fetch("cohorse.wasm"), go.importObject).then((result) => {
  go.run(result.instance)
})

Vue.use(Chakra, {
  extendTheme: cohorseTheme
})

Vue.config.productionTip = false

new Vue({
  render: h => h(App),
  store
}).$mount('#app')
