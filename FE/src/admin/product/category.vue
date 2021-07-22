<template>
    <div class="container-left">
        <p>Admin</p>
        <hr />
        <router-link to="/">
            <p>Shop</p>
        </router-link>
        <hr />
        <router-link to="/admin/username">
            <p class="username">@Trump</p>
        </router-link>
        <hr />
        <router-link to="/admin/list-user">
            <p class="category">User <i class="bi bi-caret-right"></i></p>
        </router-link>
        <router-link to="/admin/order">
            <p class="category">Order <i class="bi bi-caret-right"></i></p>
        </router-link>
        <!-- products -->
        <router-link to="/admin/product">
        <p @click="showText = !showText" class="category">Product <i class="bi bi-caret-down"></i></p>
        </router-link>
        <div class="btn-group dropend category-btn">
            <button v-show="!showText" type="button" class="btn" data-bs-toggle="dropdown" aria-expanded="false">
                <p>Tai nghe</p>
                <i class="bi bi-plus"></i>
            </button>

            <ul class="dropdown-menu">
                <li><a class="dropdown-item" href="#">Products</a></li>
                <router-link to="/admin/product/add">
                    <li><a class="dropdown-item" href="#">Add Product</a></li>
                </router-link>
            </ul>
        </div>

        <div class="btn-group dropend category-btn">
            <button v-show="!showText" type="button" class="btn" data-bs-toggle="dropdown" aria-expanded="false">
                <p>Loa</p>
                <i class="bi bi-plus"></i>
            </button>

            <ul class="dropdown-menu">
                <li><a class="dropdown-item" href="#">Product</a></li>
                <router-link to="/admin/product/add">
                    <li><a class="dropdown-item" href="#">Add Product</a></li>
                </router-link>
            </ul>
        </div>
        <!-- end products -->

        <!-- category -->
        <!-- <p class="category">Category <i class="bi bi-caret-down"></i></p> -->

        <div class=" list-category">
            <button type="button" class="btn" data-bs-toggle="dropdown" aria-expanded="false">
                <p>Category</p>
                <i class="bi bi-caret-down"></i>
            </button>

            <ul class="dropdown-menu display-category">
                <li><a class="dropdown-item" href="#">Tai nghe</a></li>
                <li><a class="dropdown-item" href="#">Loa</a></li>
                <li><a class="dropdown-item" href="#">Sac</a></li>
                <router-link to="/admin/category/add">
                    <li><a class="dropdown-item" href="#" style="border:1px solid rgba(255, 255, 255, 0.226)">Add
                            Product</a></li>
                </router-link>
            </ul>
        </div>
        <!-- end category -->
    </div>
</template>
<script>
    export default {
        name: "Category",
        data() {
            return {
                showText: true,
            }
        },
    }

</script>
<style scoped>
    .container-left p:hover {
        color: rgb(9, 184, 9);
    }

    .username:hover {
        cursor: pointer;
    }

    .display-category li a {
        border-bottom: 1px solid rgba(255, 255, 255, 0.226)
    }

    a {
        color: white;
    }

    .container-left {
        background-color: #212529;
        color: white;
        height: 100vh;
        padding: 10px;
        top: 0px;
        bottom: 0px;
    }

    button.btn p {
        padding-top: 6px;
        left: -15px;
        position: relative;
    }

    button:hover {
        cursor: pointer;
    }

    .list-category {
        /* margin-left: -10px; */
        border-top: 1px solid rgba(255, 255, 255, 0.233);
        display: block;
    }

    .list-category button {
        width: 100%;
        display: flex;
        justify-content: space-between;
    }

    .category {
        display: flex;
        justify-content: space-between;
        padding-right: 13px;
    }

    .category:hover {
        cursor: pointer;
    }

    .category-btn {
        background-color: #212529;
        color: white;
        display: block;
    }

    button {
        color: white !important;
        border-top: 1px solid white !important;
    }

    .category-btn button {
        width: 100%;
        display: flex;
        justify-content: space-between;
    }

    .dropdown-menu {
        background-color: #212529;
        color: white;
    }

    .dropdown-item {
        color: white;
        opacity: 0.8;
    }

    .dropdown-item:hover {
        background-color: #212529;
        opacity: 1;
    }

    a {
        text-decoration: none;
    }
</style>