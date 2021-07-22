<template>
    <div class="row add-product">
        <div class="col-sm-0 col-md-3 col-12 container-left">
            <Category />
        </div>

        <div class="col-sm-12 col-md-9 col-12">
            <p style="color: #0d6efd">Add new Category</p>
            <p>Product > Category</p>
            <form method="" class="product-form">
                <div class="mb-3">
                    <label for="exampleFormControlInput1" class="form-label">Category</label>
                    <input type="email" class="form-control" id="exampleFormControlInput1" placeholder="name category">
                </div>


                <button type="button" class="btn btn-outline-success">Add</button>

            </form>


        </div>
    </div>

</template>
<script>
    import Category from "@/admin/product/category.vue"
    export default {
        name: "CategoryAdd",
        components: {
            Category,
        }
    }
</script>
<style scoped>
    .add-product {
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

    /* .product-form {

        background-color: white;
        height: 100vh;
        width: 100vw;
        margin: 0px;
        margin-top: 20px;
    } */
</style>