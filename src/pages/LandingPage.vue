<script setup>
import { ref } from 'vue';
import { useRouter } from 'vue-router';

const router = useRouter();
const username = ref('');

 const handleClick = () => {
    router.push('/UserPage');
 };
</script>

<template>
  <div class ="landingPageContainer">
    <div class = "loginSlice">
        <img src="/BroncoLogo.svg" class="logo" alt="Bronco Logo" />
      
      <div class="loginContainer">
        <input v-model="username" placeholder="Enter User ID" />
        <button @click="handleClick">Login</button>
      </div>
    </div>
  </div>
</template>

<style scoped>
  
</style>
