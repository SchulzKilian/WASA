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
    login() {
      // Call login function from your backend with this.username
      // Redirect to stream view on successful login
      this.$router.push('/stream');
    }
  }
}
</script>
