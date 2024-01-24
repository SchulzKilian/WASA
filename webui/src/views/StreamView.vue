<template>
  <div v-if="images.length">
    <ImageComponent
      v-for="image in images"
      :key="image.photoId"
      :photoDetails="image"
    />
    </div>
</template>

<script>
import ImageComponent from '@/webui/src/components/ImageComponent.vue'; 
import axios from 'axios';
import api from "@/services/axios"; 

export default {
  components: {
    ImageComponent
  },
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
        
        const response = await api.get('/stream',{headers: {
                        Authorization: localStorage.getItem("token")}
                    }); // Replace with the full API URL if necessary
        if (response.data == null){
          return {
            images: [] // This array will hold the processed photo objects
    }
        }
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
