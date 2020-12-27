<template>
  <CFlex
    as="footer"
    h="50px"
    align="center"
    :width="['100%', '100%', '100%', '1200px']"
    m="0 auto"
    borderTopColor="gray.50"
    borderTopWidth="2px"
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
    await new Promise(r => setTimeout(r, 2000));
    try {
      await this.setSysInfo();
    } catch(err) {
      console.log(err); 
    }
    console.log("done")
  },
  methods: {
    ...mapActions([
      'setSysInfo'
    ])
  }
}
</script>
