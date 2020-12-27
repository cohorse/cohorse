<template>
  <c-theme-provider>
    <c-reset />
    <Nav />
    <Content />
    <Footer />
  </c-theme-provider>
</template>

<script>
import Nav from './components/base/Nav';
import Footer from './components/base/Footer';
import Content from './components/base/Content';

export default {
  name: "App",
  components: {
    Nav,
    Footer,
    Content
  }
};
</script>
