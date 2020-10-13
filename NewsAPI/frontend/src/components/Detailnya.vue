<template >
  <div>
    <ul v-for="data in berita" :key="data.id" style="margin-right:2%;">
      <div>
        <v-img
          class="white--text align-end"
          height="100%"
          width="100%"
          :src="data.urlToImage"
          gradient="to top right, rgba(100,115,201,.33), rgba(25,32,72,.7)"
        >
          <v-row>
            <v-col class="display-2" v-text="data.title"></v-col>
            <v-col></v-col>
          </v-row>
          <v-card-title ></v-card-title>
        </v-img>
        
        <v-card-text class="text--primary">
          <div class="pb-0 green--text" v-text="data.sourcename"></div>
          <div class="pb-0 grey--text" v-text="data.author"></div>
          <div class="pb-10 grey--text" v-text="data.publishedAt"></div>
          <div v-text="data.description" size="30"></div>
          <div v-text="data.content"></div>
        </v-card-text>

        <v-card-actions>
          
          <a :href="data.url" >
            <v-btn color="orange" >
              Sumber
            </v-btn>
          </a>

        </v-card-actions>
      </div>
    </ul>
  </div>
</template>
<script>
import axios from 'axios'
export default {
  data(){
    return{
        berita: [],
    }
  },
  created() {
    this.load()
  },
  methods: {
    load(){
      axios.get('http://localhost:3000/articles')
      .then(res =>{
          //console.log(res.data)
          this.berita = res.data //respon dari rest api dimasukan ke users
          console.log(this.berita)
        }).catch ((err) => {
        console.log(err)
      })
    },
  }
}
</script>