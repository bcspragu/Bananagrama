<template>
  <div id="notice"></div>
</template>

<script lang="ts">
import { Component, Prop, Vue, Watch } from 'vue-property-decorator';
import * as d3 from 'd3';
import {BaseType, Selection} from 'd3';
import {Letter} from '@/data';

@Component
export default class Notice extends Vue {
  @Prop() private notice!: string[];
  private board: Selection<BaseType, any, HTMLElement, any> = d3.select('#notice');
  private sizeX = 0;
  private sizeY = 0;

  private mounted(): void {
    this.board = d3.select('#notice').append('svg')
      .attr('width', 0)
      .attr('height', 0);
  }

  @Watch('notice')
  private renderNotice(): void {
    this.resizeSVG();

    const notice = this.board.selectAll('text').data(this.notice);

    notice.exit().remove();
    const newNotices = notice.enter().append('text');

    notice.merge(newNotices)
        .attr('x', this.sizeX / 2)
        .attr('y', (d: any, i: number) => {
          return (this.sizeY / this.notice.length) * (i + 1) - (this.sizeY / this.notice.length / 2);
        })
        .attr('text-anchor', 'middle')
        .attr('alignment-baseline', 'middle')
        .text((d: any) => d)
        .style('font-size', '30px');
  }

  private resizeSVG(): void {
    const board = document.getElementById('notice');
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
#notice {
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
