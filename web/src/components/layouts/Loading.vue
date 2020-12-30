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
    direction="column"
    justify="center"
    align="center"
    width="100vw"
    height="100vh"
    opacity="0"
    animation="fade-in 2s linear forwards"
  >
    <CImage
      width="150px"
      :src="require('../../assets/logo_loading.svg')" />
  </CFlex>
</template>

<style>
  @keyframes fade-in {
    to {
      opacity: 1;
    }
  }
</style>

<script>
export default {
  name: 'Loading',
}
</script>
