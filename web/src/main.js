import Vue from 'vue'
import App from './App.vue'
import Chakra from '@chakra-ui/vue'
import cohorseTheme from './theme.js'
import store from './store'

const go = new window.global.Go()
WebAssembly.instantiateStreaming(fetch("cohorse.wasm"), go.importObject).then((result) => {
  go.run(result.instance);
  window.global.cohorse = result.instance;
});

Vue.use(Chakra, {
  extendTheme: cohorseTheme
})

window.global.store = store

Vue.config.productionTip = false

new Vue({
  render: h => h(App),
  store
}).$mount('#app')
