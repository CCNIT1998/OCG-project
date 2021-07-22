<template>
  <div class="container">
    <p>home > Furniture</p>
    <div class="row">
      <div class="col">Furniture</div>
      <div class="col">
        <div class="form-group">
          <label class="text-end" for="exampleFormControlSelect1"
            >Sort by</label
          >
          <select class="form-control" id="exampleFormControlSelect1">
            <option value="name:asc">Alphabet, A -&gt; Z</option>
            <option value="name:desc">Alphabet, Z -&gt; A</option>
            <option value="price:asc">Price, Low to high</option>
            <option value="price:desc">Price, High to low</option>
          </select>
        </div>
      </div>
    </div>
    <!-- product -->
    <div class="row">
        <div class="col-sm-12 col-md-6 col-lg-3">
        <img src="https://via.placeholder.com/250x250" class="card-img-top"/>
            <div class="card-body">
                <h5 class="card-title">product 1</h5>
                <div class="d-flex justify-content-between">
                <p class="card-text">123456<sup>vnđ</sup></p>
                <a href="#" class="btn btn-outline-primary" >Add to Cart</a>
                </div>
            </div>
        </div>

        <div class="col-sm-12 col-md-6 col-lg-3">
        <img src="https://via.placeholder.com/250x250" class="card-img-top"/>
            <div class="card-body">
                <h5 class="card-title">product 2</h5>
                <div class="d-flex justify-content-between">
                <p class="card-text">123456<sup>vnđ</sup></p>
                <a href="#" class="btn btn-outline-primary" >Add to Cart</a>
                </div>
            </div>
        </div>

        <div class="col-sm-12 col-md-6 col-lg-3">
        <img src="https://via.placeholder.com/250x250" class="card-img-top"/>
            <div class="card-body">
                <h5 class="card-title">product 3</h5>
                <div class="d-flex justify-content-between">
                <p class="card-text">123456<sup>vnđ</sup></p>
                <a href="#" class="btn btn-outline-primary" >Add to Cart</a>
                </div>
            </div>
        </div>

        <div class="col-sm-12 col-md-6 col-lg-3">
        <img src="https://via.placeholder.com/250x250" class="card-img-top"/>
            <div class="card-body">
                <h5 class="card-title">product 4</h5>
                <div class="d-flex justify-content-between">
                <p class="card-text">123456<sup>vnđ</sup></p>
                <a href="#" class="btn btn-outline-primary" >Add to Cart</a>
                </div>
            </div>
        </div>

        <div class="col-sm-12 col-md-6 col-lg-3">
        <img src="https://via.placeholder.com/250x250" class="card-img-top"/>
            <div class="card-body">
                <h5 class="card-title">product 1</h5>
                <div class="d-flex justify-content-between">
                <p class="card-text">123456<sup>vnđ</sup></p>
                <a href="#" class="btn btn-outline-primary" >Add to Cart</a>
                </div>
            </div>
        </div>

        <div class="col-sm-12 col-md-6 col-lg-3">
        <img src="https://via.placeholder.com/250x250" class="card-img-top"/>
            <div class="card-body">
                <h5 class="card-title">product 2</h5>
                <div class="d-flex justify-content-between">
                <p class="card-text">123456<sup>vnđ</sup></p>
                <a href="#" class="btn btn-outline-primary" >Add to Cart</a>
                </div>
            </div>
        </div>

        <div class="col-sm-12 col-md-6 col-lg-3">
        <img src="https://via.placeholder.com/250x250" class="card-img-top"/>
            <div class="card-body">
                <h5 class="card-title">product 3</h5>
                <div class="d-flex justify-content-between">
                <p class="card-text">123456<sup>vnđ</sup></p>
                <a href="#" class="btn btn-outline-primary" >Add to Cart</a>
                </div>
            </div>
        </div>

        <div class="col-sm-12 col-md-6 col-lg-3">
        <img src="https://via.placeholder.com/250x250" class="card-img-top"/>
            <div class="card-body">
                <h5 class="card-title">product 4</h5>
                <div class="d-flex justify-content-between">
                <p class="card-text">123456<sup>vnđ</sup></p>
                <a href="#" class="btn btn-outline-primary" >Add to Cart</a>
                </div>
            </div>
        </div>
    </div>


  </div>
</template>
<script>
export default {
  name: "collectionDetail",
};
</script>
<style scoped>
.form-group {
  display: flex;
  padding-bottom: 20px;

}
label {
  width: 100%;
  padding-right: 10px;
  line-height: 38px;
}

</style>