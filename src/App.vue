<script setup>
 const handleClick = () => {
    ;
 };
</script>

<template>
  <div>
    <img src="/BroncoLogo.svg" class="logo" alt="Vite logo" />
  </div>

  <div>
    <input v-model="username" placeholder="Enter User ID" />
  </div>
  <div> 
    <button @click="handleClick">Login</button>
  </div>
</template>

<style scoped>

</style>
