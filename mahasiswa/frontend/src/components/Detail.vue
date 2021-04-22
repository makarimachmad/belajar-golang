<template>
    <div>
        <ul v-for="user in users" :key="user.id">
            <li>
                <span>{{user.nama}}</span> &#160;
                <span>{{user.nim}}</span> &#160;
                <span>{{user.email}}</span> &#160;
            </li>
        </ul>
    </div>
</template>
<script>
import axios from 'axios'
export default {
  data(){
    return{
        users: [],
    }
  },
  created() {
    this.det()
  },
  methods: {
    det(){
      axios.get(`http://localhost:8090/detail?id=${this.$route.params.id}`)
      .then(res =>{
        //console.log(res.data.Data)
          this.users = res.data //respon dari rest api dimasukan ke users
          console.log(this.users)
        }).catch ((err) => {
        console.log(err)
      })
    }
  }
}
</script>