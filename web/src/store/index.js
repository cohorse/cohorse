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
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
  state: () => ({
    appLoaded: false,
    sysInfo: {
      version: null,
      build: null,
      commit: null
    }
  }),
  actions: {
    async setSysInfo() {
      let resp = window.global.CohorseVuexDispatch("global", "setSysInfo")
      if(resp.err != null) {
        throw new Error(resp.err)
      }
    } 
  },
  mutations: {
    finishAppLoading(state) {
      state.appLoaded = true
    },
    setSysInfoVersion(state, version) {
      state.sysInfo.version = version
    },
    setSysInfoBuild(state, build) {
      state.sysInfo.build = build
    },
    setSysInfoCommit(state, commit) {
      state.sysInfo.commit = commit
    }
  },
  getters: {
    sysInfo: state => {
      return state.sysInfo
    }
  }
})
