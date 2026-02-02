<script setup>
/* TODO Figure out how to make all the buttons the same size */
import { useRouter } from 'vue-router';

    const router = useRouter();

    const handleLogoutClick = () => {
        router.push('/');
    };

    const handleEditPlayerClick = () => {
        /* Update the logic here*/
        router.push();
    };

    const handleEditStringsClick = () => {
        /* Update the logic here*/
        router.push();
    };

</script>
<template>
    <div class ="stringerButtons">

        <div class = "editPlayersButton">
            <button @click="handleEditPlayerClick">Edit Players</button>
        </div>

        <div class = "racketListButton">
            <button @click="handleRacketListClick">Racket List</button>
        </div>

        <div class = "editStringsButton">
            <button @click="handleEditStringsClick">Edit Strings</button>
        </div>

        <div class = "logoutButton">
            <button @click="handleLogoutClick">Logout</button>
        </div> 

    </div>
    
</template>

<style scoped>

    .stringerButtons {
        display: grid;
        grid-template-columns: repeat(2, 1fr);
        gap: 20px;
        padding: 25px;
    }

    button {
        font-size: 200%;
    }
</style>