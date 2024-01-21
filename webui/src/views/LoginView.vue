<template>
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
      username: ''
    }
  },
methods: {
  async login() {
    try {
      const response = await api.post('/session', { name: this.username });
      console.log(response.data)
      localStorage.setItem("token", response.data);
      axios.defaults.headers.common['Authorization'] = response.data;
      this.$router.push('/stream');

    } catch (error) {
      console.error("Login failed:", error);
    }
  }
}
}
</script>
