<template>
  <div>
    <form @submit.prevent="add">
        <input type="hidden" v-model="form.id">
        <input type="text" v-model="form.nama">
        <input type="text" v-model="form.nim">
        <input type="text" v-model="form.email">
        <button type="submit" v-show="!updateSubmit">Tambah</button> <!-- jika tidak update maka tombol update tidak muncul --> 
        <button type="button" v-show="updateSubmit" @click="update(form)">Update</button> <!-- jika tombol edit di klik maka tombol add akan berubah menjadi update -->
    </form>
    <ul v-for="user in users" :key="user.id">
      <li>
        <span>{{user.nama}}</span> &#160;
        <button @click="edit(user)">Edit</button> ||  <router-link :to="'/detail/'+user.id">Detail</router-link> || <button @click.prevent="del(user)">Delete</button>
      </li>
    </ul>
  </div>
</template>
<script>
/* eslint-disable */ 
import axios from 'axios'
export default {
  data(){
    return{
        users: [],
        form: {
          id: '',
          nama: '',
          nim: '',
          email: '',
        },
        updateSubmit: false
    }
  },
  mounted() {
    this.load()
  },
  methods: {
    load(){
        axios.get('http://localhost:8090/biodata')
        .then(res => {
          this.users = res.data.Data //respon dari rest api dimasukan ke users
        }).catch ((err) => {
          console.log(err)
        })
    },
    add(){
        axios.post('http://localhost:8090/tambah', this.form)
        .then(res => {
            this.load()
            this.form.nama = ''
            this.form.nim = ''
            this.form.email = ''
        })
    },
    del(user){
      axios.delete('http://localhost:8090/hapus/'+user.id)
      .then(res =>{
          this.load()
          let index = this.users.indexOf(form.nama)
          this.users.splice(index,1)
      }).catch(error => {
          console.log(error.res) 
      })
    },
    edit(user){ 
        this.updateSubmit = true
        this.form.id = user.id 
        this.form.nama = user.nama
        this.form.nim = user.nim
        this.form.email = user.email 
    },
    update(form){
        return axios.put('http://localhost:8090/update', {nama: this.form.nama, nim: this.form.nim, email: this.form.email})
        .then(res => {
        this.load()
        this.form.id = ''
        nama : this.form.nama
        nim : this.form.nim
        email : this.form.email
        this.updateSubmit = false
      }).catch((err) => {
        console.log(err)
      })
    },
  }
}
</script>