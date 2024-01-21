<template>
  <div>
    <input v-model="username" placeholder="Enter Username" />
    <button @click="login">Login</button>
  </div>
</template>

<script>
export default {
  data() {
    return {
      username: ''
    }
  },
methods: {
  async login() {
    try {
      const response = await axios.post('/session', { username: this.username });
      axios.defaults.headers.common['Authorization'] = response.data;
      // Handle the response, such as saving the token
    } catch (error) {
      console.error("Login failed:", error);
    }
  }
}
}
</script>
