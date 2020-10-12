<template>
  <v-sheet
    class="mx-auto"
    elevation="10"
    max-width="1200"
  >
    <v-card-text>
      <h2 style="color:green; font-weight: bold; font-size:200%;">| Business</h2>
    </v-card-text>
    
    <v-slide-group
      v-model="model"
      class="pa-4"
      show-arrows
    >
    
      <v-slide-item
        v-for="data in artikels" :key="data.id"
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
              color="green"
              class="white--text"
              fab
              large
              right
              top
            >
              <v-icon>mdi-heart</v-icon>
            </v-btn>
            
            <div class="font-weight-light green--text" v-text="data.title" size="5"></div>
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
                <router-link :to="'/detail/'+data.id">
                  <v-btn
                    v-if="active"
                    color="green lighten-2"
                    @click="overlay = false"
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
    data(){
      return{
        model: null,
        absolute: true,
        overlay: false,
        artikels: [],
      }
    },
    mounted(){
      this.load()
    },
    methods: {
      load(){
        axios.get('http://localhost:8090')
        .then(res => {
          this.artikels = res.data
          console.log(this.artikels)
        }).catch( err=> {
          console.log(err)
        })
      },
    },
  }
</script>