<script setup lang="ts">
import { lang } from "../../utils/language";
</script>

<template>
  <div class="mx-20 mt-4 flex justify-start border p-3">
    <p>
      {{ lang("pageAbout.privacy1") }}
    </p>

    <ol>
      <li>
        <strong> {{ lang("pageAbout.privacy2") }}</strong
        >{{ lang("pageAbout.privacy3") }}
      </li>

      <li>
        <strong>{{ lang("pageAbout.privacy4") }}</strong
        >{{ lang("pageAbout.privacy5") }}
      </li>
    </ol>

    <p>
      {{ lang("pageAbout.privacy6") }}
    </p>

    <p>
      {{ lang("pageAbout.privacy7") }}
    </p>

    <p>
      {{ lang("pageAbout.privacy8") }}
    </p>
  </div>
</template>

<style scoped>
/* Add any desired styling here for readability and presentation */
body {
  font-family: Arial, sans-serif;
  line-height: 1.6;
  margin: 0 auto;
  max-width: 800px;
  padding: 20px;
}

h1,
h2,
h3 {
  color: #333;
  font-weight: bold;
  margin: 2em 0 1em;
}

h1 {
  font-size: 24px;
}

h2 {
  font-size: 20px;
}

h3 {
  font-size: 18px;
}

ul {
  list-style: disc;
  margin-left: 2em;
  margin-bottom: 1em;
}

li {
  margin-bottom: 0.5em;
}

a {
  color: #007bff;
  text-decoration: none;
}

a:hover {
  text-decoration: underline;
}
</style>
