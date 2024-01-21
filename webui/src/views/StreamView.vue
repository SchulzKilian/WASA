<template>
  <div>
    <div v-for="image in images" :key="image.photoId">
      <img :src="image.src" />
      <p>{{ image.username }} - Likes: {{ image.likesCount }}, Comments: {{ image.commentsCount }}</p>
    </div>
  </div>
</template>

<script>
import axios from 'axios';
import api from "@/services/axios"; 

export default {
  data() {
    return {
      images: [] // This array will hold the processed photo objects
    }
  },
  mounted() {
    this.fetchImages();
  },
  methods: {
    async fetchImages() {
      try {
        const response = await api.get('/stream'); // Replace with the full API URL if necessary
        this.images = response.data.map(photo => ({
          ...photo,
          src: `data:image/jpeg;base64,${btoa(String.fromCharCode(...new Uint8Array(photo.imageData)))}`,
          // Assuming the image data is JPEG. Change the MIME type if it's different.
        }));
      } catch (error) {
        console.error('Error fetching images:', error);
      }
    }
  }
}
</script>
