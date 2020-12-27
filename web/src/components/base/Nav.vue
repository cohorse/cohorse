<template>
  <CBox as="nav" boxShadow="0px 1px 3px rgb(16 16 16 / 10%)" h="50px" px="20px" py="5px" d="flex" justifyContent="left" align-items="center">

    <CFlex
      :width="['100%', '100%', '100%', '1200px']"
      m="0 auto"
      align="center"
    >
      <CImage :src="require('../../assets/logo.svg')" h="35px" pr="5px" />
      <CText font-size="xl" font-weight="bold" color="brand.100">Cohorse</CText>
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
