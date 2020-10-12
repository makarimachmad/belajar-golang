<template>
  <v-sheet
    class="mx-auto"
    elevation="10"
    max-width="1200"
  >
    <v-card-text>
      <h2 style="color:orange; font-weight: bold; font-size:200%;">| Entertainment</h2>
    </v-card-text>
    
    <v-slide-group
      v-model="model"
      class="pa-4"
      show-arrows
    >
    
      <v-slide-item
        v-for="data in artikels.articles"
        :key="data.id"
        v-slot:default="{ active, toggle, overlay }"
      >
        <v-card
          :color="active ? undefined : 'grey lighten-4'"
          class="ma-4"
          height="300"
          width="300"
          @click="toggle"
        >
         <v-img
            :aspect-ratio="16/9"
            :src="data.urlToImage"
          >
            <!-- <v-expand-transition>
              <div
                v-if="hover"
                class="d-flex transition-fast-in-fast-out orange darken-2 v-card--reveal display-3 white--text"
                style="height: 100%;"
              >
                $14.99
              </div>
            </v-expand-transition> -->
          </v-img>
          <v-card-text
            class="pt-6"
            style="position: relative;"
          >
            <v-btn
              absolute
              color="orange"
              class="white--text"
              fab
              large
              right
              top
            >
              <v-icon>mdi-heart</v-icon>
            </v-btn>
            
            <div class="font-weight-light orange--text" v-text="data.title" size="5"></div>
            <div class="display-0.75 font-weight-light grey--text mb-2" v-text="data.author" size="3"></div>
            <!-- <div class="font-weight-light title mb-2">
              Our Vintage kitchen utensils delight any chef.<br>
              Made of bamboo by hand
            </div> -->
          </v-card-text>
          <v-overlay
              v-if="active"
              :absolute="absolute"
              :value="overlay"
              :sementara="data"
          >
              
                <router-link :to="{ name: 'detailnya'}">
                  <v-btn
                    v-if="active"
                    color="orange lighten-2"
                    @click="overlay = false, sementara(data)"
                  >
                  Selengkapnya
                  </v-btn>
                </router-link>
                
              
          </v-overlay>
        </v-card>
      </v-slide-item>
    </v-slide-group>
  </v-sheet>
</template>
<script>
  import axios from 'axios'
  export default {
    data: () => ({
      model: null,
      absolute: true,
      overlay: false,
      artikels: [],
      tampung: {
        id: 1,
      }
    }),
    mounted(){
      this.load()
    },
    methods: {
      load(){
        axios.get('https://newsapi.org/v2/top-headlines?apikey=6bc3cbc8dcf3473fb2527028734aedee&country=id&category=entertainment')
        .then(res => {
          this.artikels = res.data
          console.log(this.artikels)
        }).catch( err=> {
          console.log(err)
        })
      },
      sementara(data){
        console.log(data)
        
        return axios.put('http://localhost:3000/articles/' + this.tampung.id , 
        { sourceid: data.source.id,
          sourcename: data.source.name,
          title: data.title,
          url: data.url,
          urlToImage: data.urlToImage,
          description: data.description,
          author: data.author,
          publishedAt: data.publishedAt,
          content: data.content
        })
        .then(res => {
          console.log(res.data)
        }).catch((err) => {
          console.log(err);
        })
      },
      // update(tampung){
      //   return axios.put('http://localhost:3000/users/' + form.id , {name: this.form.name})
      //   .then(res => {
      //     this.form.id = ''
      //     this.form.name = ''
      //     this.updateSubmit = false
      //     console.log(res)
      //   }).catch((err) => {
      //     console.log(err);
          
      //   })
      // },
    },
  }
</script>