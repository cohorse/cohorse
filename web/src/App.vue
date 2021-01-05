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

<template>
  <CThemeProvider>
    <CReset />

    <Loading v-if="!appLoaded" />

    <Nav v-if="appLoaded" />
    <Content v-if="appLoaded" />
    <Footer v-if="appLoaded" />
  </CThemeProvider>
</template>

<script>
import Nav from './components/base/Nav'
import Footer from './components/base/Footer'
import Content from './components/base/Content'
import Loading from './components/layouts/Loading'

import { mapState, mapActions } from 'vuex'

export default {
  name: "App",
  components: {
    Nav,
    Footer,
    Content,
    Loading
  },
  computed: mapState([
    'appLoaded'
  ]),
  methods: {
    ...mapActions([
      'setSysInfo'
    ])
  },
  async mounted() {
    // Show the loading animation for at least 2 seconds, which is half of the
    // animation's timing (one full painting without disassembly).
    await new Promise(r => setTimeout(r, 2000))
    try {
      await this.setSysInfo()
    } catch(err) {
      console.log(err) 
    }
  }
}
</script>
