<template>
  <div class="banner">
    <p v-if="storedUsername">Welcome, {{ storedUsername }}</p>
    <p v-else>Welcome, Guest</p>
  </div>
  <div>
    <input v-model="username" placeholder="Enter Username" />
    <button @click="login">Login</button>
  </div>
</template>

<script>
import axios from 'axios';
import api from "@/services/axios"; 
import router from '@/router';

export default {

  data() {
    return {
      username: '', // Tracks login status
      storedUsername: localStorage.getItem("username") || '' 
    }
  },

  methods: {
    async login() {
      try {
        const response = await api.post('/session', { name: this.username });
        console.log(response.data)
        localStorage.setItem("token", response.data);
        localStorage.setItem("username", this.username);
        axios.defaults.headers.common['Authorization'] = response.data;
        this.$router.push('/'); // Redirect to home after successful login
        location.reload()
      } catch (error) {
        console.error("Login failed:", error);
      }
    }
  }
}
</script>
