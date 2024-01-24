<template>
    <div class="image-container">
      <img :src="imageSrc" />
      <div class="image-info">
        <p>{{ photoDetails.username }} - Likes: {{ photoDetails.likesCount }}, Comments: {{ photoDetails.commentsCount }}</p>
        <button @click="toggleLike">{{ liked ? 'Unlike' : 'Like' }}</button>
        <button @click="toggleComment">{{ commented ? 'Remove Comment' : 'Comment' }}</button>
      </div>
    </div>
  </template>
  
  <script>
  import api from "@/services/axios"; 
  export default {
    props: {
      photoDetails: {
        type: Object,
        required: true
      }
    },
    data() {
      return {
        liked: false,
        commented: false
      }
    },
    computed: {
      imageSrc() {
        return `data:image/jpeg;base64,${btoa(String.fromCharCode(...new Uint8Array(this.photoDetails.imageData)))}`;
      }
    },
    created() {
    if (this.photoDetails && this.photoDetails.liked !== undefined) {
      this.liked = this.photoDetails.liked;
    }
  },
    methods: {
      async toggleLike() {
        try {
          const url = `/photos/${this.photoDetails.photoId}/likes/`;
          if (this.liked) {
            await api.delete(url,{},{headers: {
                        Authorization: localStorage.getItem("token")}
                    });
          } else {
            await api.post(url,{},{headers: {
                        Authorization: localStorage.getItem("token")}
                    });
          }
          this.liked = !this.liked;
        } catch (error) {
          console.error('Error toggling like:', error);
        }
      },
      async toggleComment() {
        try {
          const url = `/photos/${this.photoDetails.photoId}/comments/`;
          if (this.commented) {
            await api.delete(url,{},{headers: {
                        Authorization: localStorage.getItem("token")}
                    });
          } else {
            await api.post(url,{},{headers: {
                        Authorization: localStorage.getItem("token")}
                    });
          }
          this.commented = !this.commented;
        } catch (error) {
          console.error('Error toggling comment:', error);
        }
      }
    }
  }
  </script>
  
  <style>
  .image-container {
    /* Add your styling here */
  }
  .image-info {
    /* Add your styling here */
  }
  </style>
  