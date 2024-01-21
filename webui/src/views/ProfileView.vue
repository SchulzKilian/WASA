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
      searchUsername: '',
      userProfile: null // This will hold the user profile data
    }
  },
  methods: {
    searchProfile() {
      // Fetch user profile based on searchUsername and set to userProfile
    }
  }
}
</script>
