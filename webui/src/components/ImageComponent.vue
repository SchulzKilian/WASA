<template>
    <div class="image-container">
      <img :src="imageSrc" />
      <div class="image-info">
        <p>{{ photoDetails.username }} - Likes: {{ photoDetails.LikesCount }}, Comments: {{ photoDetails.CommentsCount }}</p>
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
        console.log(typeof this.photoDetails.imageData);
        return `data:image/jpeg;base64,${this.photoDetails.imageData}`;
        // return `data:image/jpeg;base64,${btoa(String.fromCharCode(...new Uint8Array(this.photoDetails.imageData)))}`;
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
            await api.delete(url,{headers: {
                        Authorization: localStorage.getItem("token")}
                    });
                    this.photoDetails.LikesCount = this.photoDetails.LikesCount -1
          } else {
            await api.post(url,{},{headers: {
                        Authorization: localStorage.getItem("token")}
                    });
                    this.photoDetails.LikesCount = this.photoDetails.LikesCount +1
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
  function base64ToUint8Array(base64) {
    var binaryString = window.atob(base64);
    var len = binaryString.length;
    var bytes = new Uint8Array(len);
    for (var i = 0; i < len; i++) {
        bytes[i] = binaryString.charCodeAt(i);
    }
    return bytes;
}
  </script>
  
  <style>
  .image-container img {
  width: 100%; /* or a specific pixel value */
  height: auto; /* maintain aspect ratio */
}

  </style>
  