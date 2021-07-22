<template lang="">
    <div class="page404">
        <router-link to="/">
            <h3>&ensp;Shop&ensp;</h3>
        </router-link>
        <img src="@/assets/img/404.png" alt="" srcset="">
    </div>
</template>
<script>
    export default {
        name: "404",
    }
</script>
<style scoped>
    .page404 {
        position: relative;
        text-align: center;
    }

    .page404 h3 {
        position: absolute;
        z-index: 2;
        /* width: 50%; */
        transform: translate(50%, 50%);
        border: 1px solid #ee4d2d;
        border-radius: 6px;
        display: inline-block;
        padding: 8px;
    }

    .page404 h3:hover {
        cursor: pointer;
    }

    @media (max-width:1023px) {
        .page404 {
            margin-top: 10rem;
        }
    }
</style>