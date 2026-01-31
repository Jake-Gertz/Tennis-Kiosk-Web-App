<script setup>
import { ref } from 'vue';
import { useRouter } from 'vue-router';

const router = useRouter();
const userID = ref('');

 const handleClick = () => {
  if(userID.value.trim() === '9999') {
    router.push('/UserPage');
  } else {
    /* Update this alert message*/
    alert('Wrong UserID')
  }
 };
</script>

<template>
  <div class ="landingPageContainer">
    <div class = "loginSlice">
        <img src="/BroncoLogo.svg" class="logo" alt="Bronco Logo" />
      
      <div class="loginContainer">
        <input v-model="userID" placeholder="Enter User ID" />
        <button @click="handleClick">Login</button>
      </div>
    </div>
  </div>
</template>

<style scoped>
  
</style>
