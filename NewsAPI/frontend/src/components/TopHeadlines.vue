<template>
  <v-carousel
    cycle
    height="400"
    hide-delimiter-background
    show-arrows-on-hover
  >
    <v-carousel-item
      v-for="data in artikels.articles"
      :key="data.id"
    >
      <v-img
        width="100%"
        height="100%"
        :src="data.urlToImage"
        gradient="to top right, rgba(100,115,201,.33), rgba(25,32,72,.7)"
      >
        <v-row
          class="fill-height"
          align="center"
          justify="center"
          style="margin-left:2%;"
        >
          <div class="display-2" v-text="data.title">
          </div>
        </v-row>
        
      </v-img>
    </v-carousel-item>
  </v-carousel>
</template>
<script>
import axios from 'axios'
  export default {
    data () {
      return {
        colors: [
          'indigo',
          'warning',
          'pink darken-2',
          'red lighten-1',
          'deep-purple accent-4',
        ],
        slides: [
          'First',
          'Second',
          'Third',
          'Fourth',
          'Fifth',
        ],
        artikels:[]
      }
    },
    mounted() {
        this.load()
    },
    methods: {
        load(){
            axios.get('https://newsapi.org/v2/top-headlines?apikey=6bc3cbc8dcf3473fb2527028734aedee&country=id')
            .then(res => {
              this.artikels = res.data
              console.log(this.artikels)
            }).catch( err=> {
              console.log(err)
            })
        }
    }
  }
</script>
