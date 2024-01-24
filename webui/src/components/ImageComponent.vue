<template>
    <div class="image-container">
      <img :src="imageSrc" />
      <div class="image-info">
        <p>{{ photoDetails.username }} - Likes: {{ photoDetails.LikesCount }}, Comments: {{ photoDetails.CommentsCount }}</p>
        <button @click="toggleLike">{{ liked ? 'Unlike' : 'Like' }}</button>
        <button @click="toggleComment">{{ commented ? 'Remove Comment' : 'Comment' }}</button>
        <button v-if="isCurrentUser" @click="deletePhoto">Delete</button>
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
      },
      isCurrentUser() {
      const currentUsername = localStorage.getItem("username");
      console.log(currentUsername)
      console.log(this.photoDetails.username)
      return this.photoDetails.username == currentUsername;
    }
    },
    created() {
    if (this.photoDetails && this.photoDetails.liked !== undefined) {
      this.liked = this.photoDetails.liked;
    }
  },
    methods: {
        async deletePhoto() {
      try {
        const url = `/photos/${this.photoDetails.photoId}`;
        await api.delete(url, {
          headers: {
            Authorization: localStorage.getItem("token")
          }
        });
        location.reload()
        // Handle the UI update or redirection after successful deletion
      } catch (error) {
        console.error('Error deleting photo:', error);
      }
    },
  
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
  width: 10%; /* or a specific pixel value */
  height: auto; /* maintain aspect ratio */
}

  </style>
  