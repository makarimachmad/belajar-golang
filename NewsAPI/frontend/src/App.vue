<template>
  <v-app>
    <v-app-bar
      app
      color="light-grey"
    >
      <div class="d-flex align-center">
        <v-img
          alt="Vuetify Logo"
          class="shrink mr-2"
          contain
          src="https://spesolution.com/images/layouts/main/logo.png"
          transition="scale-transition"
          width="80"
        />
        News
      </div>

      <v-spacer></v-spacer>

      <!-- <v-btn
        href="https://github.com/vuetifyjs/vuetify/releases/latest"
        target="_blank"
        text
      >
        <span class="mr-2">Latest Release</span>
        <v-icon>mdi-open-in-new</v-icon>
      </v-btn> -->
    </v-app-bar>
    <v-main>
      <!-- <HelloWorld/> -->
      <router-view></router-view>
      <br><br><br>
      <Footer/>
    </v-main>
  </v-app>
</template>

<script>
// import HelloWorld from './components/HelloWorld';
import Footer from './components/Footer.vue'

export default {
  name: 'App',

  components: {
    // HelloWorld,
    Footer,
  },

  data: () => ({
    //
  }),
};
</script>