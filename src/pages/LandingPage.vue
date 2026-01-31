<script setup>
import { ref } from 'vue';
import { useRouter } from 'vue-router';

const router = useRouter();
const userID = ref('');

 

const handleClick = async () => {
  if(userID.value.trim() !== "") {
    try {
      const response = await fetch('http://localhost:8080/checl-user', {
        method: 'POST',
        body: JSON.stringify({ userID: username.value }),
        headers: { 'Content-type': 'application/json'}
      });

      const data = await response.json();
      if(data.exists) {
        router.push('/UserPage')
      } else {
        alert("Incorrect ID: " + userID.value.trim());
      }
    } catch (error) {
      console.error("Go server did not respond", error);        
    }
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
