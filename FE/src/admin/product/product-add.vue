<template>
  <div class="row add-product">
    <div class="col-sm-0 col-md-3 col-12 container-left">
      <Category />
    </div>

    <div class="col-sm-12 col-md-9 col-12">
      <p style="color: #0d6efd">Add new Product</p>
      <p>Product > Category</p>
      <form method="POST" class="product-form">
        <!-- name -->
        <div class="">
          <label for="exampleFormControlInput1" class="form-label">Name</label>
          <input
            type="email"
            class="form-control"
            id="exampleFormControlInput1"
            v-model="name"
          />
        </div>

        <!-- price -->
        <div class="">
          <label for="exampleFormControlInput1" class="form-label">Price</label>
          <input
            type="email"
            class="form-control"
            id="exampleFormControlInput1"
            v-model="price"
          />
        </div>

        <!-- Description -->
        <div class="">
          <label for="exampleFormControlTextarea1" class="form-label"
            >Description</label
          >
          <textarea
            class="form-control"
            id="exampleFormControlTextarea1"
            rows="3"
            v-model="description"
          ></textarea>
        </div>

        <!-- images -->
        <!-- <div class="">
          <label for="formFileSm" class="form-label">Images</label>
          <input
            class="form-control form-control-sm"
            id="formFileSm"
            type="file"
            @change="onFileUpload"
          />
        </div> -->

        
        <div class="row">
          <!-- Category -->
          <div class="col-3">
            <label for="exampleDataList" class="form-label">Category</label>
            <input
              class="form-control"
              list="datalistOptions"
              id="exampleDataList"
              value="Tai nghe"
            />
            <datalist id="datalistOptions">
              <option value="Loa"></option>
              <option value="Tai nghe"></option>
              
            </datalist><br />
          </div>
          <!-- Trade Mark -->
          <div class="col-3">
            <label for="TradeMark" class="form-label">Trade Mark</label>
            <select class="form-select" id="TradeMark">
              <option value="Việt Nam">Việt Nam</option>
              <option value="Chine">Chine</option>
              <option value="US">US</option>
            </select>
          </div>

          <div class="col-3">
            <button
              @click="onSubmit"
              type="button"
              class="btn btn-outline-success"
            >
              Submit
            </button>
          </div>
        </div>
      </form>
    </div>
  </div>
</template>
<script>
import Category from "@/admin/product/category.vue";
import axios from "axios";
export default {
  name: "ProductAdd",
  components: {
    Category,
  },
  data() {
    return {
      name: "",
      description: "",
      price: "",
      category_id: "",
      trade_mark: "",
      image: null,
    };
  },

  methods: {
    onSubmit() {
      let TradeMark = document.querySelector('#TradeMark')['value']
      axios
        .post("http://localhost:3000/products", {
          name: this.name,
          description: this.description,
          price: Number(this.price),
          category_id: 1,
          trade_mark: TradeMark,
          // image: this.FILE
        })
        .then((res) => {
          console.log(res);
        });

      alert("added to cart");
      location.reload();
    },
  },
};
</script>
<style scoped>
button {
  margin-top: 20px;
}
.add-product {
  margin: 0 !important;
  position: absolute;
  height: 100vh;
  right: 10px !important;
  width: 100vw;
  top: 0px;
  left: 0px !important;
  bottom: 0px;
  z-index: 100;
  background-color: white;
}

label {
  color: black;
  position: relative;
  top: 10px;
  background: white;
  margin: 0px 4px;
  padding: 0px 5px;
  font-size: 13px;
  font-weight: 297px !important;
  font-weight: 300;
}

.product-form {
  background-color: white;
  height: 100vh;
  padding-right: 20px;
  /* width: 100%; */
  position: relative;
}
</style>