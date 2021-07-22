<template>
    <footer class="footer">
        <div class="container">
            <div class="row">

                <div class="col-sm-6 col-lg-4">
                    <div>
                        <i class="bi bi-facebook"></i>
                        <p>Fanpage</p>
                    </div>

                    <div>
                        <i class="bi bi-geo-alt"></i>
                        <p>Thanh Xuân, Hà Nội</p>
                    </div>

                    <div>
                        <i class="bi bi-envelope"></i>
                        <p>mail-your-shop@gmail.com</p>
                    </div>

                    <div><i class="bi bi-telephone-forward"></i>
                        <p>0987 654 321</p>
                    </div>

                    <div>
                        <i class="bi bi-link"></i>
                        <p>shopOCG.com.vn</p>
                    </div>
                </div>

                <div class="col-sm-6 col-lg-4">
                    <div>
                        <i></i>
                        <p>Blog</p>
                    </div>
                    <div>
                        <i></i>
                        <p>Contact</p>
                    </div>
                    <div>
                        <i></i>
                        <p>Products</p>
                    </div>
                    <div>
                        <i></i>
                        <p>Location</p>
                    </div>
                    <div>
                        <i></i>
                        <p></p>
                    </div>
                </div>

                <div class="col-12 col-lg-4 ">
                    <iframe
                        src="https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d3724.8961847807204!2d105.82321841376394!3d20.996798194247177!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x3135ac88f2b2a107%3A0x4da0e40f92fe6ea7!2zMTMwYiBIb8OgbmcgVsSDbiBUaMOhaSwgS2jGsMahbmcgVHJ1bmcsIFRoYW5oIFh1w6JuLCBIw6AgTuG7mWksIFZp4buHdCBOYW0!5e0!3m2!1svi!2s!4v1591427153216!5m2!1svi!2s"
                        width="450" height="200" frameborder="0" style="border:0;" allowfullscreen=""
                        aria-hidden="false" tabindex="0"></iframe>
                </div>
            </div>
        </div>
        <div class="copyright"></div>
    </footer>
</template>
<script>
    export default {
        name: "Footer",
    };
</script>
<style scoped>
    footer {
        padding-top: 15px;
        background-color: #00000015;
    }

    .footer p {
        display: inline-block;
        margin-left: 10px;
    }

    .footer i {
        font-size: 20px;
        color: #f58551;
    }

    .footer div iframe {
        width: 100%;
        right: 0px;
    }

    .copyright {
        height: 50px;
        border: 1px solid;
        background: gray;
    }
</style>