<template>
  <div>
    <div v-for="image in images" :key="image.id">
      <img :src="image.url" />
      <p>Likes: {{ image.likes }}, Comments: {{ image.comments }}</p>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      images: [] // This should be fetched from your /stream endpoint
    }
  },
  mounted() {
    // Fetch images from /stream endpoint and populate `images` array
  }
}
</script>
