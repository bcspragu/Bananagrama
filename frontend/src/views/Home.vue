<template>
  <Grid/>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator';
import Grid from '@/components/Grid.vue'; // @ is an alias to /src

@Component({
  components: {
    Grid,
  },
})
export default class Home extends Vue {}
</script>
