import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
  state: () => ({
    sysInfo: {
      version: null,
      build: null,
      commit: null
    }
  }),
  actions: {
    async setSysInfo() {
      let resp = window.global.CohorseVuexDispatch("global", "setSysInfo")
      console.log(resp)
    }
  },
  mutations: {
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
