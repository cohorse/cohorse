<template>
  <CFlex
    align="center"
  >
    <CFlex
      direction="column"
      pr="6"
    >
      <CStack :spacing="5">
        <CHeading
          as="h1"
          size="xl"
          fontWeight="bold"
          color="primary.800"
        >
          Create perfect cohorts for your school
        </CHeading>
        <CText
          color="gray.800"
        >
          Cohorse helps you to form perfects cohorts of students for every class your school hosts while maintaining perfect secrecy of your student's personal details. The goal is to split school classes into two clusters that have minimum, at best no contact with one another.
          <CLink href="https://docs.cohorse.io" isExternal>Find out more about this approach.</CLink>
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
        <CButton>Onboard new school</CButton>
        <CText size="xs" color="gray.500">or</CText>
        <CButton>Load existing school</CButton>
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

    <CImage ml="auto" rounded="lg" transform="skew(0, 1deg)" :width="['80%', '60%', '50%']" boxShadow="1px 1px 7px 0px rgb(0 0 0 / 27%)" my="6" height="400px" :src="require('../../assets/students_social_distancing.jpg')" />
  </CFlex>
</template>

<script>
export default {
  name: 'Content',
}
</script>
