<script setup>
/* TODO Update admin page */
import { useRouter } from 'vue-router';

    const router = useRouter();

    const handleLogoutClick = () => {
        router.push('/');
    };
</script>

<template>
    <p>
        Admin page under construction
    </p>

    <div class = "logoutButton">
            <button @click="handleLogoutClick">Logout</button>
    </div>

</template>