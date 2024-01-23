<template>
  <div>
    <input v-model="username" placeholder="Search Username" />
    <button @click="fetchUserProfile">Search</button>
    <div v-if="userProfile">
      <div v-for="photo in userProfile.Photos" :key="photo.PhotoID">
        <!-- Convert binary image data to a data URL for display -->
        <img :src="`data:image/jpeg;base64,${arrayBufferToBase64(photo.imageData)}`" />
      </div>
      <p>Followers: {{ userProfile.Followers }}</p>
      <p>Following: {{ userProfile.Following }}</p>
      <p>Posts: {{ userProfile.PhotosCount }}</p>
    </div>
      <button v-if="!isOwnProfile" @click="toggleFollow">
        {{ isFollowing ? 'Unfollow' : 'Follow' }}
      </button>
    </div>

</template>

<script>
import api from "@/services/axios"; 

export default {
  data() {
    return {
      username: '', // Username to search
      userProfile: null
    };
  },
  methods: {

    async toggleFollow() {
      if (this.isFollowing) {
        await this.unfollowUser();
      } else {
        await this.followUser();
      }
    },
    async followUser() {
      // Implement the API call to follow the user
      try {
        await api.post(`/users/${this.username}/followers/`, {}, {
          headers: { Authorization: localStorage.getItem("token") }
        });

      } catch (error) {
        console.error("Failed to follow user:", error);
      }
    },
    async unfollowUser() {
      // Implement the API call to unfollow the user
      try {
        await api.delete(`/users/${this.username}/followers/`, {}, {
          headers: { Authorization: localStorage.getItem("token") }
        });

      } catch (error) {
        console.error("Failed to unfollow user:", error);
      }
    }
      ,
    
  

    async fetchUserProfile() {
  try {
    const response = await api.get(`/users/${this.username}`, {
      headers: { Authorization: localStorage.getItem("token") }
    });

    // Check if the response contains 'photos' and it is an array
    const photos = response.data.Photos && Array.isArray(response.data.Photos)
      ? response.data.Photos.map(photo => ({
          ...photo,
          ImageData: this.arrayBufferToBase64(photo.ImageData)
        }))
      : [];

    this.userProfile = {
      ...response.data,
      photos: photos
    };
  } catch (error) {
    console.error("Failed to fetch user profile:", error);
  }
},
arrayBufferToBase64(buffer) {
      let binary = '';
      let bytes = new Uint8Array(buffer);
      let len = bytes.byteLength;
      for (let i = 0; i < len; i++) {
        binary += String.fromCharCode(bytes[i]);
      }
      return window.btoa(binary);
    }
  }
}


</script>
