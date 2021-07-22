<template>
  <div class="row order">
    <div class="col-sm-0 col-md-3 col-12 container-left">
      <Category />
    </div>

    <div class="col-sm-12 col-md-9 col-12">
      <p style="color: #0d6efd">User info</p>
      <p></p>
      <div class="row">
        <div class="col-4">
          <div class="card" style="width: 18rem">
            <img
              src="https://via.placeholder.com/250x250"
              class="card-img-top"
              alt="..."
            />
            <div class="card-body">
              <h5 class="card-title">Dolar Trump</h5>
              <p class="card-text">staff</p>
              <p class="card-text">Trump@gmail.com</p>
              <p class="card-text">0392182333</p>
            </div>
          </div>
        </div>

        <div class="col-6">
          <div class="input-group mb-3">
            <span class="input-group-text" id="inputGroup-sizing-default">Name</span>
            <input class="form-control"/>
          </div>

          <div class="input-group mb-3">
            <span class="input-group-text" id="inputGroup-sizing-default"
              >Email</span
            >

            <input
              type="text"
              class="form-control"
              aria-label="Recipient's username"
              aria-describedby="basic-addon2"
            />
            <span class="input-group-text" id="basic-addon2">@gmail.com</span>
          </div>

          <div class="input-group mb-3">
            <span class="input-group-text" id="inputGroup-sizing-default">Password</span>
            <input class="form-control"/>
          </div>

         <div class="input-group mb-3">
            <span class="input-group-text" id="inputGroup-sizing-default">Phone</span>
            <input class="form-control"/>
          </div>

          <div class="input-group mb-3">
            <span class="input-group-text" id="inputGroup-sizing-default">Address</span>
            <input class="form-control"/>
          </div>

        <button type="button" class="btn btn-outline-success">Save</button>

        </div>
      </div>
    </div>
  </div>
</template>
<script>
import Category from "@/admin/product/category.vue";
export default {
  name: "CategoryAdd",
  components: {
    Category,
  },
};
</script>
<style scoped>
.order {
  margin: 0 !important;
  position: absolute;
  height: 100vh;
  width: 100vw;
  top: 0px;
  left: 0px !important;
  bottom: 0px;
  z-index: 100;
  background-color: white;
}
#inputGroup-sizing-default{
    width: 20%;
}

</style>