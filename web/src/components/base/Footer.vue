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
  <CFlex
    as="footer"
    px="6"
    align="center"
    :maxWidth="['100%', '100%', '100%', '1200px']"
    m="0 auto"
  >
    <CFlex
      borderTopColor="gray.50"
      borderTopWidth="2px"
      line-height="48px"
      style="flex-grow: 1;"
    >
      <CFlex mr="4">
        <CText color="gray.500" fontWeight="bold" mr="1">Version:</CText>
        <CText color="gray.500">{{ sysInfo.version }}</CText>
      </CFlex>

      <CFlex mr="4">
        <CText color="gray.500" fontWeight="bold" mr="1">Build:</CText>
        <CText color="gray.500">{{ sysInfo.build }}</CText>
      </CFlex>

      <CFlex>
        <CText color="gray.500" fontWeight="bold" mr="1">Commit:</CText>
        <CText color="gray.500">{{ sysInfo.commit }}</CText>
      </CFlex>
    </CFlex>
  </CFlex>
</template>

<script>
import { mapGetters, mapActions } from 'vuex'

export default {
  name: 'Footer',
  computed: {
    ...mapGetters([
      'sysInfo'
    ])
  },
  async mounted() {
    // TODO: For now, we need to delay the action dispatch. This should of
    // course be no longer needed as soon as we have the global application
    // loading state.
    await new Promise(r => setTimeout(r, 2000))
    try {
      await this.setSysInfo()
    } catch(err) {
      console.log(err) 
    }
  },
  methods: {
    ...mapActions([
      'setSysInfo'
    ])
  }
}
</script>
