<template>
    <box class="test1 test2"></box>
</template>

<script>
export default {
  props: {
    mono:Boolean,
    enabled:Boolean,
    width:{
      type: String,
      default: '400px',
    },
    skin:'',
     margin:{
      type: String,
      default: '20px 1px 20px 0',
    },
    max:{
      type: String,
      default: '90vw',
    },
    height:{
      type: String,
      default: '70px',
    },
    title: String,
  },
  methods:{
  },
  data() {
    return {}
  },
}
</script>

<style>

.test1 {
  display: flex;
  align-items: center;
  justify-content: center;
  background-image: linear-gradient(-235deg, #4E60D2 69%, #913BAF 70%, #D52D88 79%, #D72E85 76%, #E6366D 86%, #F26D4F 93%);
  border-radius: 10px;
  cursor: pointer; 
  user-select: none;
  max-width: 100%;
}


.test2{
  font-weight: 400;
  text-transform: none;
  font-size: 18px;
  color: rgba(255, 255, 255, .6);
  letter-spacing: 0.04em;
  margin: 0;
  text-align: center;
}

</style>
