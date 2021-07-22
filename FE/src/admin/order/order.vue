<template>
    <div class="row order">
        <div class="col-sm-0 col-md-3 col-12 container-left">
            <Category />
        </div>

        <div class="col-sm-12 col-md-9 col-12">
            <p style="color: #0d6efd">List order</p>
            <p></p>
            <div>
                <div class="row product-striped">
                    <div class="col-1">ID</div>
                    <div class="col">Name</div>
                    <div class="col">Price</div>
                    <div class="col">Category</div>
                    <div class="col">Trade Mark</div>
                    <!-- <div class="col-3">Description</div> -->
                    <div class="col"></div>
                </div>
                <div class="row product-content">
                    <div class="col-1">1</div>
                    <div class="col">Tai nghe M5</div>
                    <div class="col">125</div>
                    <div class="col">Tai nghe</div>
                    <div class="col">Viet Nam</div>
                    <!-- <div class="col-3 Description">DescriptionDescriptionDescriptionDescriptionDescriptionDescriptionDescription DescriptionDescriptionDescriptionDescriptionDescrDescription DescriptionDescriptionDescriptionDescriptionDescrDescription DescriptionDescriptionDescriptionDescriptionDescr</div> -->
                    <div class="col edit">
                        <div class="div"><i class="bi bi-pencil"></i>Edit</div>
                        <div class="div"><i class="bi bi-trash"></i>Remove</div>
                    </div>
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

    .edit :is(div){
    display: inline-block;
    margin-right: 2px;
    font-size:10px;
    border: 1px solid rgba(0, 0, 0, 0.212);
    border-radius: 5px;
    padding: 4px;
    color: white
  }
  .edit :is(div):nth-child(1){
    background-color: #00a8ff;
  }
  .edit :is(div):nth-child(2){
    background-color: #e84118;
  }
  
  .edit :is(div):hover{
    cursor: pointer;
  }

  .edit :is(div) i{
    padding-right: 2px;
    color: white
  }
  .product-striped{
    font-weight: bold;
    padding-bottom: 10px;
  }
  .product-content{
    border-top: 1px solid rgba(0, 0, 0, 0.247);
    padding: 5px 0;
  }
  .table-p{
    font-size:15px;
    opacity: 0.8;
  }
  .Description{
    overflow: hidden;
    display: -webkit-box;
    -webkit-box-orient: vertical;
    -webkit-line-clamp: 2;
  }
  .header{
    padding-left: 12px;
  }
  
  .header p{
    display: inline-block;
    font-size: 18 px;
    margin-right: 15px;
    opacity: 0.8;

  }
  .header p:hover{
    cursor: pointer;
    opacity: 1;
  }

  .header-icon i{
    margin: 5px 8px;
    opacity: 0.8;
  }

  .header-icon i:nth-child(2) sup{
    background-color: rgba(250, 2, 2, 0.952);
    color: white;
    padding: 0px 5px; 
    border-radius: 6px;
  }

  .header-icon i:nth-child(3) sup{
    background-color: rgba(255, 230, 0, 0.952);
    color: black;
    padding: 0px 5px; 
    border-radius: 6px;
  }


  .header-icon i:hover{
    opacity: 1;
    cursor: pointer;
  }
  .header-product{
    margin-top: 10px;
  }
</style>