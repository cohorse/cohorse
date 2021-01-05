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
  <CStack align="center" p="6" is-inline>
    <CFlex direction="column" pr="6">
      <CStack :spacing="5">
        <CHeading as="h1" size="xl" fontWeight="bold" color="primary.800">
          Create perfect cohorts for your school
        </CHeading>
        <CText color="gray.800">
          Cohorse helps you to form perfects cohorts of students for every
          class your school hosts while maintaining perfect secrecy of your
          student's personal details. The goal is to split school classes into
          two clusters that have minimum, at best no contact with one another.

          <CLink href="https://docs.cohorse.io" isExternal>
            Find out more about this approach.
          </CLink>
        </CText>
      </CStack>

      <CHeading
        as="h5"
        size="s"
        textTransform="uppercase"
        fontWeight="bold"
        mt="4"
        mb="2"
      >
        For instituion leaders
      </CHeading>
      <CStack :spacing="2" alignItems="center" is-inline>
        <CButton>Onboard a new school</CButton>
        <CText size="xs" color="gray.500">or</CText>
        <CButton>Load an existing school</CButton>
      </CStack>

      <CHeading
        as="h5"
        size="s"
        textTransform="uppercase"
        fontWeight="bold"
        mt="4"
        mb="2"
      >
        For students
      </CHeading>
      <CStack :spacing="2" is-inline>
        <CButton>Fill out survey</CButton>
      </CStack>
    </CFlex>

    <CImage
      ml="auto"
      rounded="lg"
      transform="skew(0, 1deg)"
      :height="['0', '0', '300px', '400px']"
      :display="['hidden', 'hidden', 'block', 'block']"
      boxShadow="1px 1px 7px 0px rgb(0 0 0 / 27%)"
      my="6"
      :src="require('../../assets/students_social_distancing.jpg')" />
  </CStack>
</template>

<script>
export default {
  name: 'Content',
}
</script>
