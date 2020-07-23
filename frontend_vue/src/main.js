import Vue from 'vue'
import App from './App.vue'
import store from './components/store.js'


Vue.config.productionTip = false

new Vue({
  render: h => h(App),
  store
}).$mount('#app')

// new Vue({
//   el: '#app',
//   store,
//   components: { App },
//   template: '<App/>'
// })
