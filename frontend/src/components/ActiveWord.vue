<template>
  <div id="word"></div>
</template>

<script lang="ts">
import { Component, Prop, Vue, Watch } from 'vue-property-decorator';
import * as d3 from 'd3';
import {BaseType, Selection} from 'd3';
import {Letter} from '@/data';

@Component
export default class ActiveWord extends Vue {
  @Prop() private word!: string;
  private board: Selection<BaseType, any, HTMLElement, any> = d3.select('#word');
  private sizeX = 0;
  private sizeY = 0;

  private mounted(): void {
    this.board = d3.select('#word').append('svg')
      .attr('width', 0)
      .attr('height', 0);
  }

  @Watch('word')
  private renderWord(): void {
    this.resizeSVG();

    const word = this.board.selectAll('text').data([this.word]);

    word.exit().remove();
    const newWords = word.enter().append('text');

    word.merge(newWords)
        .attr('x', this.sizeX / 2)
        .attr('y', this.sizeY / 2)
        .attr('text-anchor', 'middle')
        .attr('alignment-baseline', 'middle')
        .text((d: any) => d)
        .style('font-size', '30px');
  }

  private resizeSVG(): void {
    const board = document.getElementById('word');
    if (board) {
      this.sizeX = board.offsetWidth;
      this.sizeY = board.offsetHeight;
    }

    this.board
      .attr('width', this.sizeX + 'px')
      .attr('height', this.sizeY + 'px');
  }
}
</script>

<style>
#word {
  height: 100%;

  -webkit-touch-callout: none; /* iOS Safari */
    -webkit-user-select: none; /* Safari */
     -khtml-user-select: none; /* Konqueror HTML */
       -moz-user-select: none; /* Firefox */
        -ms-user-select: none; /* Internet Explorer/Edge */
            user-select: none; /* Non-prefixed version, currently
                                  supported by Chrome and Opera */
}
</style>
