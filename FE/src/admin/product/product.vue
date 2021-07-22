<template>
  <div class=" row main">
    <div class="col-sm-0 col-md-3">
      <Category />
    </div>

    <div class="col-sm-12 col-md-9">
      <ProductList />
    </div>
  </div>
</template>
<script>
import ProductList from "@/admin/product/product-list.vue";
import Category from "@/admin/product/category.vue";
export default {
  name: "AdminProduct",
  components: {
    ProductList,
    Category,
  },
};
</script>
<style scoped>
footer.footer {
    display: none;
}
#carouselExampleDark{
  display: none;
}
.main {
  height: 100vh;
  padding-right: 20px;
  margin: 0 !important;
  position: absolute;
  top: 0px;
  left: 0px !important;
  right: 0px !important;
  bottom: 0px;
  width: 100vw;
  background-color: yellow;
  z-index: 10;
  background-color:white;
}

</style>