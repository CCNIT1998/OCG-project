<template>
  <div>
    <Slide />
    <header class="contener">
      <nav class="navbar navbar-expand-lg navbar-light bg-light">
        <div class="container-fluid">
          <router-link to="/">
            <a class="navbar-brand" href="#"
              ><img src="https://via.placeholder.com/50x50" alt="logo" />
            </a>
          </router-link>
          <button
            class="navbar-toggler"
            type="button"
            data-bs-toggle="collapse"
            data-bs-target="#navbarSupportedContent"
            aria-controls="navbarSupportedContent"
            aria-expanded="false"
            aria-label="Toggle navigation"
          >
            <span class="navbar-toggler-icon"></span>
          </button>
          <div class="collapse navbar-collapse" id="navbarSupportedContent">
            <ul class="navbar-nav me-auto mb-2 mb-lg-0">
              <router-link to="/">
                <li class="nav-item">
                  <a class="nav-link active" aria-current="page" href="#"
                    >Home</a
                  >
                </li>
              </router-link>
              <li class="nav-item">
                <a class="nav-link" href="#"></a>
              </li>
              <li class="nav-item dropdown">
                <a
                  class="nav-link dropdown-toggle"
                  href="#"
                  id="navbarDropdown"
                  role="button"
                  data-bs-toggle="dropdown"
                  aria-expanded="false"
                >
                  Products
                </a>
                <ul class="dropdown-menu" aria-labelledby="navbarDropdown">
                  <li><a class="dropdown-item" href="#">Loa</a></li>
                  <li><hr class="dropdown-divider" /></li>
                  <li>
                    <a class="dropdown-item" href="#">Tai nghe</a>
                  </li>
                </ul>
              </li>
              <li class="nav-item">
                <a class="nav-link" href="#" tabindex="-1" aria-disabled="true"
                  >Contact</a
                >
              </li>
            </ul>
            <form class="d-flex" @submit.prevent="searchProducts()">
              <input
                class="form-control position-relative"
                type="search"
                placeholder="Search"
                aria-label="Search"
                v-model="textSearch"
              />
              <i class="bi bi-search header-icon-search"></i>
              <!-- <button class="btn btn-outline-success position-relative" type="submit">Search</button> -->
            </form>
            <div class="header-cart">
              <router-link to="/cart">
                <i class="bi bi-cart-plus">
                  <sup>{{ totalItems }}</sup>
                </i>
              </router-link>
            </div>

            <router-link to="/admin/product">
              <a class="navbar-brand" href="#">Admin </a>
            </router-link>
            <router-link to="/login">
              <a class="navbar-brand" href="#">Login </a>
            </router-link>
            <router-link to="/register">
              <a class="navbar-brand" href="#">Register </a>
            </router-link>
          </div>
        </div>
      </nav>
    </header>
  </div>
</template>
￼￼

<script>
import Slide from "@/components/Slide.vue";
import { mapGetters } from "vuex";
export default {
  name: "Header",
  components: {
    Slide,
  },
  data() {
    return { textSearch: "" };
  },
  created() {
    this.$store.dispatch("cart/getProductsInCart");
  },
  computed: {
    ...mapGetters("cart", ["totalItems"]),
  },
  methods: {
    async searchProducts() {
      console.log("ok");
      await this.$store.dispatch("products/getProducts", {
        search: this.textSearch,
        pageIndex: 1,
      });
    },
  },
};
</script>
<style scoped>
a {
  text-decoration: none;
  font-size: 15px;
}

header {
  position: fixed;
  top: 0;
  right: 0;
  left: 0;
  z-index: 5;
}

form input {
  border: 1px solid transparent;
  box-shadow: 0 2px 5px 1px rgb(64 60 67 / 16%);
  border-radius: 24px;
}

.header-cart {
  padding: 5px 20px 0 10px;
  font-size: 25px;
}
.header-cart a {
  padding: 5px 20px 0 10px;
  font-size: 20px;
  color: black;
}
sup {
  border: 1px solid rgba(0, 0, 0, 0.144);
  padding: 0 2px;
  background-color: #ffb243;
  border-radius: 50%;
  text-align: center;
}
.header-icon-search {
  right: 26px;
  top: 9px;
  position: relative;
  cursor: pointer;
}
</style>
