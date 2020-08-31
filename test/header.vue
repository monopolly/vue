<template>
<div class="header">
  <div class="header-inside">
    <!-- left -->
    <div class="left">
      <nuxt-link class="menu" active-class="active" to="/" exact>home</nuxt-link>
      <!-- <nuxt-link class="menu" active-class="active" to="/invest" exact>invest</nuxt-link> -->
      <!-- <nuxt-link class="menu" active-class="active" to="/sale" exact>sale</nuxt-link>
      <nuxt-link class="menu" active-class="active" to="/read" exact>read</nuxt-link>
      <nuxt-link class="menu" active-class="active" to="/read" exact>shopping</nuxt-link>
      <nuxt-link class="menu" active-class="active" to="/read" exact>services</nuxt-link> -->
    </div>

    <!-- right -->
    <div class="right">
      <nuxt-link class="menu" active-class="active" to="/list" exact>+list a villa</nuxt-link>
      <nuxt-link v-show="!$store.state.logged" class="menu" style="margin:0px 5px 0" active-class="active" to="/login" exact>👤</nuxt-link>
      <nuxt-link v-show="$store.state.logged" class="menu" active-class="active" to="/profile" exact>{{$store.state.user.name}}</nuxt-link>
      <p v-show="$store.state.logged" class="menu" active-class="active" @click="logout()">➥</p>
    </div>
  </div>

</div>
</template>
<script>
export default {
  //name:'header',
  methods:{
  
    async logout(){
          console.log("try logout")
          var req = {name:'logout'}
          const ok = await this.$axios.post('/api', req)
          console.log("api logout", ok)
          
          if (!ok) return
          console.log("commit logout")
          this.$store.commit('logout')
          if (process.browser) {
            window.location = '/'
            console.log("logout!")
          }

      },
  },
 
  data() {
    return {
      
    }
  }
}
</script>



<style scoped>
.header {
  position: fixed;
  z-index: 10000;
  top: 15px;
}

.header-inside {
  width: 100vw;
  position: relative;
  display: flex;
  justify-content: space-between;
}

.left {
  position: absolute;
  display: flex;
  left: 20px;
  
}

.right {
  position: absolute;
  display: flex;
  justify-content: center;
  align-items: center;
  text-align: right;
  right: 20px;
  
}

a.logo {
  /* font-family: 'Avernir Next',  -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell, 'Open Sans', 'Helvetica Neue', sans-serif; */
  text-decoration: none;
  transition: all 0.1s ease;
  margin: 0 12px;
  /* text-transform: uppercase; */

  font-weight: 100;
  font-size: 0.83em;
  color: rgb(255, 255, 255);
  /* letter-spacing: 0.02em; */
}

a.menu, .menu {
  font-family:  -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell, 'Open Sans', 'Helvetica Neue', sans-serif;
  text-decoration: none;
  transition: all 0.1s ease;
  margin: 0 12px;
  /* text-transform: uppercase; */
  text-transform: lowercase;

  font-weight: 100;
  font-size: 0.83em;
  color: rgb(255, 255, 255);
  /* letter-spacing: 0.02em; */
  transition: all 0.2s ease;
  cursor: pointer;
}

a.menu:hover {
  opacity: 0.6;
  text-decoration: dotted;
}

a.menu.active {
  opacity: 0.4;
}

a.menu.active:hover {
  text-decoration: none;
}
/* 
@media screen and (max-width: 600px) {
  .header {
    position: relative;
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
  }
  .logo {
    display: none;
  }
  .menu {
    display: none;
  }
  
} */

.magiclogo{
    background: -webkit-linear-gradient(180deg, #4E60D3, #913BAF, #D52D88, #D72E85, #E6366D, #F26D4F);
    -webkit-background-clip: text;
    -webkit-text-fill-color: transparent;
}

</style>

