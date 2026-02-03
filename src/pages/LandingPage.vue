<script setup>
import { ref } from 'vue';
import { useRouter } from 'vue-router';

const router = useRouter();
const userID = ref('');

 

const handleClick = async () => {
  if(userID.value.trim() !== "") {
    try {
      const response = await fetch('http://localhost:8080/check-user', {
        method: 'POST',
        body: JSON.stringify({ userId: userID.value }),
        headers: { 'Content-type': 'application/json'}
      });

      const data = await response.json();
      if(data.exists && userID.value == "9999") {
        router.push('/AdminPage')
      } else if (data.exists && userID.value == "8888") {
          router.push('/StringerPage')
      } else if (data.exists && userID.value == "7777"){
          router.push('/UserPage')
      }  else {
        alert("Incorrect ID: " + userID.value.trim());
      }
    } catch (error) {
      console.error("Go server did not respond", error);        
    }
  } else {
    alert("Please enter a user ID:");
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
