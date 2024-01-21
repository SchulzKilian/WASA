<template>
  <div>
    <input v-model="searchUsername" placeholder="Search Username" />
    <button @click="searchProfile">Search</button>
    <div v-if="userProfile">
      <div v-for="photo in userProfile.photos" :key="photo.id">
        <img :src="photo.url" />
      </div>
      <p>Followers: {{ userProfile.followers }}</p>
      <p>Following: {{ userProfile.following }}</p>
      <p>Posts: {{ userProfile.photos.length }}</p>
    </div>
  </div>
</template>

<script>
export default {
 data() {
  return {
    username: '', // Username to search
    userProfile: null
  };
},
methods: {
  async fetchUserProfile() {
    try {
      const response = await axios.get(`/users/${this.username}`);
      this.userProfile = response.data;
    } catch (error) {
      console.error("Failed to fetch user profile:", error);
    }
  }
}
}
</script>
