<template>
    <div>
        <TopHeadlines/>
        <br><br><br>
        <Business/>
        <br><br><br>
        <Health/>
        <br><br><br>
        <Entertainment/>
    </div>
</template>
<script>
// import Hover from './components/Hover.vue'
import TopHeadlines from './TopHeadlines'
import Health from './Health.vue'
import Business from './Business.vue'
import Entertainment from './Entertainment.vue'

export default {
  components: {
    // HelloWorld,
    // Hover,
    TopHeadlines,
    Health,
    Business,
    Entertainment,
  },

  data: () => ({
    //
  }),
};
</script>