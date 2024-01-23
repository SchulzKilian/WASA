<template>
      <div class="banner">
      <p>Welcome, {{ storedUsername }} </p>
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
      storedUsername: localStorage.getItem("username") // To store the username after login
    }
  },
methods: {
  async login() {
    try {
      const response = await api.post('/session', { name: this.username });
      console.log(response.data)
      localStorage.setItem("token", response.data);
      localStorage.setItem("username", this.username);
      this.isLoggedIn = true;
      this.storedUsername = this.username;
      axios.defaults.headers.common['Authorization'] = response.data;


    } catch (error) {
      console.error("Login failed:", error);
    }
  }
}
}
</script>
