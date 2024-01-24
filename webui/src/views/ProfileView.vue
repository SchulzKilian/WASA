<template>
  <div>
    <input v-model="username" placeholder="Search Username" />
    <button @click="fetchUserProfile">Search</button>
    <div v-if="userProfile">
      <p>Followers: {{ userProfile.Followers }}</p>
      <p>Following: {{ userProfile.Following }}</p>
      <p>Posts: {{ userProfile.PhotosCount }}</p>
      <div v-if="images.length">
    <ImageComponent
      v-for="image in images"
      :key="image.photoId"
      :photoDetails="image"
    />
  </div>
    </div>
      <button v-if="userProfile && !isOwnProfile" @click="toggleFollow">
        {{ userProfile.IsFollowing ? 'Unfollow' : 'Follow' }}
      </button>
    </div>

</template>

<script>
import ImageComponent from '@/webui/src/components/ImageComponent.vue'; 
import api from "@/services/axios"; 

export default {
  components: {
    ImageComponent
  },

  data() {
    return {
      username: '', // Username to search
      userProfile: null,
      images: []
    };
  },
  computed: {
    isOwnProfile() {
      return this.username === localStorage.getItem("username");
    }
  },
  methods: {

    async toggleFollow() {
      if (this.userProfile.IsFollowing) {
        await this.unfollowUser();
        
      } else {
        await this.followUser();
      }
      this.userProfile.IsFollowing= !this.userProfile.IsFollowing
    },
    async followUser() {
      // the API call to follow the user
      try {
        await api.post(`/users/${this.username}/followers/`, {}, {
          headers: { Authorization: localStorage.getItem("token") }
        });

      } catch (error) {
        console.error("Failed to follow user:", error);
      }
    },
    async unfollowUser() {
      // the API call to unfollow the user
      try {
        await api.delete(`/users/${this.username}/followers/`, {
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
    this.images = response.data.Photos && Array.isArray(response.data.Photos)
      ? response.data.Photos.map(photo => ({
          ...photo,
          ImageData: this.arrayBufferToBase64(photo.ImageData)
        }))
      : [];

    this.userProfile = {
      ...response.data,
      photos: this.images
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
