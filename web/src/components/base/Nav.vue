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
  <CBox
    as="nav"
    h="50px"
    px="20px"
    py="5px"
    d="flex"
    justifyContent="left"
    align-items="center"
    boxShadow="0px 1px 3px rgb(16 16 16 / 10%)"
  >
    <CFlex
      :width="['100%', '100%', '100%', '1200px']"
      m="0 auto"
      align="center"
    >
      <CImage :src="require('../../assets/logo.svg')" h="35px" pr="5px" />
      <CText font-size="xl" font-weight="bold" color="brand.100">
        Cohorse
      </CText>

      <CFlex ml="auto">
        <a href="https://docs.cohorse.io" target="_blank" marginLeft="20">
          <CButton left-icon="question" variant="outline" size="xs">
            Documentation
          </CButton>
        </a>
      </CFlex>
    </CFlex>
  </CBox>
</template>

<script>
export default {
  name: 'Nav',
}
</script>
