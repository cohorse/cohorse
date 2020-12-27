<template>
  <CFlex
    :width="['100%', '100%', '100%', '1200px']"
    m="0 auto"
    minH="calc(100vh - 50px - 50px)"
  >
    <Landing />
  </CFlex>
</template>

<script>
import Landing from '../layouts/Landing';

export default {
  name: 'Content',
  components: {
    Landing
  }
}
</script>
