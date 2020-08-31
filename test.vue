<template>
  <div class="pay" :style="{width: width, maxWidth: max}">
    <div :class="['pay-button',{'pay-button-disable':!enable}]" 
            @click="checkout" 
            :style="{
                width: '100%', 
                height: height,
                margin: margin,
                fontFamily: mono ? 
                `'Roboto Mono','Droid Sans Mono', Menlo, monospace;`: 
                `'Avenir Next', -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell, 'Open Sans', 'Helvetica Neue', sans-serif;`,
            }">
      <h1>{{title}}</h1>
    </div>

    <p class="powered">Powered by
      <a href="https://stripe.com" target="_blank">
        <span style="color:#7A78F1"><b>Stripe</b></span>
      </a>
    </p>
    <script src="https://checkout.stripe.com/checkout.js"></script>
  </div>
    
</template>
<script>
export default {
    data() {
        return {}
      },
      props: {
        mono: {
          type: Boolean,
          default: false,
        },
        enable: {
          type: Boolean,
          default: true,
        },
        currency: {
          type: String,
          default: 'usd',
        },
        name: String,
        desc: String,
        price: Number,

        width: {
          type: String,
          default: '400px',
        },
        margin: {
          type: String,
          default: '20px 0',
        },
        max: {
          type: String,
          default: '90vw',
        },
        height: {
          type: String,
          default: '70px',
        },
        title: String,
      },

      methods: {
        checkout(e) {
          var vm = this
          var handler = StripeCheckout.configure({
            key: 'pk_test_FVwib8bCxlGuikUTuYD8Qi90',
            //image: '/images/villass_round.png',
            locale: 'auto',
            allowRememberMe: false,
            token: function (token) {
              this.$emit('token', token)
            }
          });
          handler.open({
            currency: vm.currency,
            amount: vm.price * 100,
            name: vm.name,
            description: vm.desc,
          });
        },
        formatNumber: function (value) {
          console.log(value)
          if (!value) return ''
          return value.toString().replace(/\B(?=(\d{3})+(?!\d))/g, ",") || '';
        },



      },
}
</script>


<style>
@media only screen and (max-width: 500px) {
    body {
        background-color: lightblue;margin: 4px;color: #913BAF;

    }
    .nice:hover{
      color: #4E60D2
    }

    .nice::-moz-list-bullet{
      color: #4E60D2
    }
    .nice.james::-moz-focus-outer{
      color: #E6366D;
    }
}

@keyframes fly {
    from {
        left: -50vw;
        bottom: -50vh;
    }
    to {
       left: 100vw;
        bottom: 100vh;
    }
}

.plane{
  position: absolute;
    /* right: 13vw;
    top: 7vh; */
    width: 40vw;
    height: auto;
    pointer-events: none;
    animation-name: fly;
    animation-duration: 30s;
}

.pay{
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
}
.pay-button {
  display: flex;
  align-items: center;
  justify-content: center;
  background-image: linear-gradient(-235deg, #4E60D2 69%, #913BAF 70%, #D52D88 79%, #D72E85 76%, #E6366D 86%, #F26D4F 93%);
  border-radius: 10px;
  cursor: pointer; 
  user-select: none;
  max-width: 100%;
}

.pay-button-disable{
  background: rgba(255, 255, 255, 0.034);
  cursor: default;
}



.pay-button h1 {
  font-weight: 600;
  text-transform: uppercase;
  font-size: 24px;
  color: #FFFFFF;
  letter-spacing: 0.1em;
  margin: 0;
  text-align: center;
}

.pay-button-disable h1{
  color: rgba(255, 255, 255, 0.281);
}



.pay-button:hover {
  /* transition: all 0.1s ease; */
  filter: brightness(1.1)
}

.pay-button:active {
  transform: translate(0px,2px);
}

.pay-button-disable:active {
  transform: translate(0px,0px);
}

@media screen and (max-width: 600px){
.pay-button {
  border-radius: 8px;
  margin: 5vw 0;
  max-width: 90vw;
}

.pay-button h1 {
  /* font-family: 'Avenir Next', -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell, 'Open Sans', 'Helvetica Neue', sans-serif; */
}

}
</style>

