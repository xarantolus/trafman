<template>
  <nav class="navbar">
    <div class="navbar-brand">
      <router-link to="/" class="navbar-item">
        <img class="logo-image invert-dm" src="/img/icons/safari-pinned-tab.svg" alt="Repositories">
      </router-link>
      <a class="navbar-burger" :class="{ 'is-active': menuOpen }" v-on:click="menuOpen = !menuOpen">
        <span aria-hidden="true"></span>
        <span aria-hidden="true"></span>
        <span aria-hidden="true"></span>
      </a>
    </div>
    <div class="navbar-menu" :class="{ 'is-active': menuOpen }">
      <div class="navbar-start">
        <router-link class="navbar-item" to="/">
          Repositories
        </router-link>
      </div>
    </div>
  </nav>
  <router-view />
</template>

<script lang="ts">
import { defineComponent } from 'vue';

export default defineComponent({
  name: 'App',
  data() {
    return {
      menuOpen: false,
    }
  }
});
</script>

<style>
@import "../node_modules/bulma/css/bulma.min.css";
@import "../node_modules/bulma-prefers-dark/css/bulma-prefers-dark.css";

/* ALL color definitions go here */
:root {
  --font-color: #1f2933;
  --border-color: #ccc;
  --green: #00ff55;
  --yellow: #f7ae40;
  --font-color-on-yellow: var(--font-color);
  --blue: #6fcaff;
  --card-color: #fff;
  --button-color: #bbb;
}

@media (prefers-color-scheme: dark) {
  :root {
    --font-color: #eed7c0;
    --border-color: #999;
    --green: #158c11;
    --yellow: #df8b1d;
    --font-color-on-yellow: #333;
    --blue: #004770;
    --card-color: #222;
    --button-color: #333;
  }

  /* Some additional color fixes */
  .help {
    color: #aaa;
  }

  .invert-dm {
    filter: invert(1) !important;
  }


  /* Some style fixes for bulma */
  .navbar-link.is-active,
  .navbar-link:focus,
  .navbar-link:focus-within,
  .navbar-link:hover,
  a.navbar-item.is-active,
  a.navbar-item:focus,
  a.navbar-item:focus-within,
  a.navbar-item:hover {
    background-color: var(--button-color);
    color: var(--green);
  }
}

html,
body {
  height: 100vh;
}

.navbar {
  position: sticky;
  width: 100%;
  height: 3%;
  top: 0;
  padding: 0;
  filter: drop-shadow(0 0 .25rem var(--font-color));
  margin-bottom: 10px;
}

#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: var(--font-color);
}

nav {
  padding: 30px;
}

nav a {
  font-weight: bold;
  color: #2c3e50;
}

nav a.router-link-exact-active {
  color: var(--green);
}
</style>
