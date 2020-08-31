<template>
<div class="gallery">
  <div v-show="images && images.length" class="gallery-preview">
      <div @click="previewFile = p" v-for="p in images" :key="p" class="gallery-images">
            <img  class="gallery-image" :class="`gallery-${size}`" :src="p">
    </div>
  </div>

      <transition name="fade">
        
          <img v-show="previewFile" class="preview-image" @click="previewFile = ''" :src="previewFile">
        
        
      </transition>
  
        
  
</div>
</template>

<script>

export default {
    props:{ 
        size:{
            type: String,
            default: "1",
        },
        alt:'',
        images: Array,
        preview: Boolean,
    },
    data() {
        return {
          previewFile:'',
        }
    },
    mounted() {
      
    },
    methods: {

    }
    
};
</script>
<style scope>
.gallery {
  position: relative;
  display: flex;
  flex-direction: column;
  align-items: center;
  max-width: 100vw;
}

.gallery-menu {
  display: flex;
  justify-content: center;
}

.gallery-preview {
  display: flex;
  flex-wrap: wrap;
  flex-direction: row;
  width: 100vw;
  justify-content: center;
}

.gallery-images {
    position: relative;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: flex-start;
}

.preview-image{
  height: 96vh;
  width: auto;
  margin: 0;
  position: fixed;
  top: 40%;
  left: 50%;
  transform: translate(-50%, -40%);
  vertical-align: middle;
  z-index: 1000;
  box-shadow: 14px 18px 300px 0 rgba(0, 0, 0, 1);
  cursor: pointer;
  
  border-radius: 20px;
}


.fade-enter-active, .fade-leave-active {
  transition: opacity .4s;
}
.fade-enter, .fade-leave-to /* .fade-leave-active below version 2.1.8 */ {
  opacity: 0;
}

.gallery-info {
  margin: 0;
  padding: 0;
  width: 180px;
}

.cover-wrap{
    position: relative;
    display: flex;
    flex-direction: column;
    align-items: center;
}

.cover-image {
    position: relative;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: flex-start;
}

.gallery-image-wrap{
    position: relative;
    display: flex;
    flex-direction: column;
}

.gallery-image {
  height: auto;
  border-radius: 10px;
  box-shadow: 14px 18px 30px 0 rgba(0, 0, 0, 0.3);
  margin: 20px 20px 8px;
  cursor: pointer;
  transition: all 0.2s ease;
  z-index: 5;
  max-width: 100vw;
}

.gallery-1 {
  /* width: auto;
  max-width: 90vw; */
  max-height: 70vw;
  margin: 20px;
}

.gallery-2 {
  width: 650px;
  height: auto;
  margin: 15px 0 8px;
}

.gallery-3 {
  width: 400px;
  height: auto;
  margin: 15px;
}

.gallery-4 {
  width: 300px;
  height: auto;
  margin: 10px;
}

.gallery-5 {
  width: 230px;
  height: auto;
}

.gallery-6 {
  width: 180px;
  height: auto;
}


.gallery-image-preview{
  height: 100vh;
  width: 100vw;
  z-index: 1000;
  margin: 0;
    position: absolute;
    top: 40%;
    left: 50%;
    transform: translate(-50%, -40%);
    vertical-align: middle;

}

.gallery-image.main {
  height: 500px;
  width: auto;
  border-radius: 10px;
  margin: 10px 20px 50px;
  cursor: default;
}

.gallery-image:hover {
  filter: brightness(1.1)
}

.image-delete {
  position: absolute;
  margin-left: 108px;
  margin-top: 10px;
  display: flex;
  width: 22px;
  height: 22px;
  cursor: pointer;
  z-index: 100;
  box-shadow: 1px 4px 7px 0 rgba(0, 0, 0, 0.15);
}

.image-delete:active {
  transform: translate(0px, 1px);
}

.image-delete:hover {
  filter: brightness(0.9)
}

.image-delete.main {
  position: absolute;
  margin-left: 347px;
  margin-top: 0px;
  display: flex;
  width: 26px;
  height: 26px;
  cursor: pointer;
  z-index: 100;
  box-shadow: 1px 4px 7px 0 rgba(0, 0, 0, 0.15);
}

input.inline {
  -webkit-appearance: none;
  width: 100%;
  padding: 5px 0px 15px 13px;
  margin: 0px 0px 0px 0px;
  font-size: 1em;
  font-weight: 400;
  color: black;
  background: none;
  box-shadow: none;
  box-sizing: border-box;
  text-align: center;
}

.drop {
  margin: 20px;
  width: 600px;
  height: 200px;
  border-radius: 15px;
  border: 2px dotted rgba(70, 189, 235, 0.71);
  background: rgba(199, 239, 255, 0.19);
  display: flex;
  justify-content: center;
  align-items: center;
  cursor: pointer;
}

.drop:hover {
  border: 2px dotted rgba(60, 196, 250, 0.856);
  background: rgba(190, 235, 253, 0.521);
}

.drop.dark{
    border: 2px dotted rgba(63, 63, 63, 0.71);
    background: rgba(44, 44, 44, 0.19);
}

.drop.dark:hover {
  border: 2px dotted rgba(63, 63, 63, 1);
  /*   background: rgba(44, 44, 44, 0.4); */
}


.dragLabel {
  font-weight: 500;
  font-size: 18px;
  color: rgba(53, 129, 170, 0.29);
  text-align: center;
}
.dragLabel.dark {
  color: rgba(82, 82, 82, 0.29);
}

p.gallery-close {
    position: absolute;
    right: 10px;
    top: 10px;
    cursor: pointer;
    font-size: 12px;
    width: 10px;
    margin: 0;
    line-height: 7px;
    color: #0000001a;
}
p.gallery-close.dark {
    color: #e0e0e0;
}
p.gallery-close:hover {
    color: #e58f8fce;
}
.gallery-close:active{
    transform: translate(0px,2px);
}
p.gallery-close::before{
    content: "\2715";
}

/* iphones */
@media screen and (max-width: 700px) {
  .gallery-image {
    width: 90vw;
    margin: 5vh 0px 3vh;
  }
  .gallery-image.main {
    width: 90vw;
    height: auto;
    margin: 4vh 0px 3vh;
    /* display: none; */
  }
  .preview-image{
    display: none;
  }
  p.gallery-close {
    position: absolute;
    right: 20px;
    top: 40px;
    z-index: 10;
    /* width: 10px;
    margin: 0; */
    line-height: 7px;
    font-size: 31px;
    text-shadow: 2px 5px 9px rgba(0, 0, 0, .78);
}
.covert-title{
    padding: 16px 0 7px;;
}

  .drop {
    /* display: none; */
    width: 90vw;
  }
  input.inline {
    padding: 10px 0px 15px 23px;
    font-size: 1.2em;
    font-weight: 400;
  }
}

</style>
