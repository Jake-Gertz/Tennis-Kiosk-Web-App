<script setup>
import { useRouter } from 'vue-router';

    const router = useRouter();
    const handleClick = () => {
    router.push('/');
 };
</script>
<template>
    <div>
        <p>
            This is a test for the page
            <button @click="handleClick">Logout</button>

        </p>
    </div>
</template>